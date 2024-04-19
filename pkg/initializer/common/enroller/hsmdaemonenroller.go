/*
 * Copyright contributors to the Hyperledger Fabric Operator project
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 * 	  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package enroller

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	current "github.com/IBM-Blockchain/fabric-operator/api/v1beta1"
	"github.com/IBM-Blockchain/fabric-operator/pkg/initializer/common/config"
	k8sclient "github.com/IBM-Blockchain/fabric-operator/pkg/k8s/controllerclient"
	jobv1 "github.com/IBM-Blockchain/fabric-operator/pkg/manager/resources/job"
	"github.com/IBM-Blockchain/fabric-operator/pkg/util"
	"github.com/pkg/errors"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
)

// HSMDaemonEnroller is responsible for enrolling with CAs to generate cryptographic materical
// for fabric nodes
type HSMDaemonEnroller struct {
	CAClient HSMCAClient
	Client   k8sclient.Client
	Instance Instance
	Timeouts HSMEnrollJobTimeouts
	Scheme   *runtime.Scheme
	Config   *config.HSMConfig
}

// NewHSMDaemonEnroller initializes and returns a pointer to HSMDaemonEnroller
func NewHSMDaemonEnroller(cfg *current.Enrollment, instance Instance, caclient HSMCAClient, client k8sclient.Client, scheme *runtime.Scheme, timeouts HSMEnrollJobTimeouts, hsmConfig *config.HSMConfig) *HSMDaemonEnroller {
	return &HSMDaemonEnroller{
		CAClient: caclient,
		Client:   client,
		Instance: instance,
		Scheme:   scheme,
		Timeouts: timeouts,
		Config:   hsmConfig,
	}
}

// GetEnrollmentRequest returns the enrollment request defined on the ca client
func (e *HSMDaemonEnroller) GetEnrollmentRequest() *current.Enrollment {
	return e.CAClient.GetEnrollmentRequest()
}

// ReadKey is no-op method on HSM
func (e *HSMDaemonEnroller) ReadKey() ([]byte, error) {
	return nil, nil
}

// PingCA uses the ca client do ping the CA
func (e *HSMDaemonEnroller) PingCA(timeout time.Duration) error {
	return e.CAClient.PingCA(timeout)
}

// Enroll reaches out the CA to get back a signed certificate
func (e *HSMDaemonEnroller) Enroll() (*config.Response, error) {
	log.Info(fmt.Sprintf("Enrolling using HSM Daemon"))
	// Deleting CA client config is an unfortunate requirement since the ca client
	// config map was not properly deleted after a successfull reenrollment request.
	// This is problematic when recreating a resource with same name, as it will
	// try to use old settings in the config map, which might no longer apply, thus
	// it must be removed if found before proceeding.
	if err := deleteCAClientConfig(e.Client, e.Instance); err != nil {
		return nil, err
	}

	e.CAClient.SetHSMLibrary(filepath.Join("/hsm/lib", filepath.Base(e.Config.Library.FilePath)))
	if err := createRootTLSSecret(e.Client, e.CAClient, e.Scheme, e.Instance); err != nil {
		return nil, err
	}

	if err := createCAClientConfig(e.Client, e.CAClient, e.Scheme, e.Instance); err != nil {
		return nil, err
	}

	job := e.initHSMJob(e.Instance, e.Timeouts)
	if err := e.Client.Create(context.TODO(), job.Job, k8sclient.CreateOption{
		Owner:  e.Instance,
		Scheme: e.Scheme,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to create HSM ca initialization job")
	}
	log.Info(fmt.Sprintf("Job '%s' created", job.GetName()))

	if err := job.WaitUntilActive(e.Client); err != nil {
		return nil, err
	}
	log.Info(fmt.Sprintf("Job '%s' active", job.GetName()))

	if err := job.WaitUntilContainerFinished(e.Client, CertGen); err != nil {
		return nil, err
	}
	log.Info(fmt.Sprintf("Job '%s' finished", job.GetName()))

	status, err := job.ContainerStatus(e.Client, CertGen)
	if err != nil {
		return nil, err
	}

	log.Info(fmt.Sprintf("Job status at finish '%s'", status))

	switch status {
	case jobv1.FAILED:
		return nil, fmt.Errorf("Job '%s' finished unsuccessfully, not cleaning up pods to allow for error evaluation", job.GetName())
	case jobv1.COMPLETED:
		if err := job.Delete(e.Client); err != nil {
			return nil, err
		}

		if err := deleteRootTLSSecret(e.Client, e.Instance); err != nil {
			return nil, err
		}

		if err := deleteCAClientConfig(e.Client, e.Instance); err != nil {
			return nil, err
		}
	}

	name := fmt.Sprintf("ecert-%s-signcert", e.Instance.GetName())
	err = wait.Poll(2*time.Second, 30*time.Second, func() (bool, error) {
		sec := &corev1.Secret{}
		log.Info(fmt.Sprintf("Waiting for secret '%s' to be created", name))
		err = e.Client.Get(context.TODO(), types.NamespacedName{
			Name:      name,
			Namespace: e.Instance.GetNamespace(),
		}, sec)
		if err != nil {
			return false, nil
		}

		return true, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create secret '%s'", name)
	}

	if err := setControllerReferences(e.Client, e.Scheme, e.Instance); err != nil {
		return nil, err
	}

	return &config.Response{}, nil
}

const (
	// HSMClient is the name of container that contain the HSM client library
	HSMClient = "hsm-client"
	// CertGen is the name of container that runs the command to generate the certificate for the CA
	CertGen = "certgen"
)

func (e *HSMDaemonEnroller) initHSMJob(instance Instance, timeouts HSMEnrollJobTimeouts) *jobv1.Job {
	hsmConfig := e.Config
	req := e.CAClient.GetEnrollmentRequest()

	hsmLibraryPath := hsmConfig.Library.FilePath
	hsmLibraryName := filepath.Base(hsmLibraryPath)

	jobName := fmt.Sprintf("%s-enroll", instance.GetName())

	f := false
	t := true
	user := int64(0)
	backoffLimit := int32(0)
	mountPath := "/shared"
	pvcVolumeName := fmt.Sprintf("%s-pvc-volume", instance.GetName())

	k8sJob := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobName,
			Namespace: instance.GetNamespace(),
			Labels: map[string]string{
				"name":  jobName,
				"owner": instance.GetName(),
			},
		},
		Spec: batchv1.JobSpec{
			BackoffLimit: &backoffLimit,
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					ServiceAccountName: instance.GetName(),
					ImagePullSecrets:   util.AppendImagePullSecretIfMissing(instance.GetPullSecrets(), hsmConfig.BuildPullSecret()),
					RestartPolicy:      corev1.RestartPolicyNever,
					InitContainers: []corev1.Container{
						{
							Name:            HSMClient,
							Image:           hsmConfig.Library.Image,
							ImagePullPolicy: corev1.PullAlways,
							Command: []string{
								"sh",
								"-c",
								fmt.Sprintf("mkdir -p %s/hsm && dst=\"%s/hsm/%s\" && echo \"Copying %s to ${dst}\" && mkdir -p $(dirname $dst) && cp -r %s $dst", mountPath, mountPath, hsmLibraryName, hsmLibraryPath, hsmLibraryPath),
							},
							SecurityContext: &corev1.SecurityContext{
								RunAsUser:                &user,
								RunAsNonRoot:             &f,
								AllowPrivilegeEscalation: &t,
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "shared",
									MountPath: mountPath,
								},
							},
							Resources: instance.GetResource(current.INIT),
						},
					},
					Containers: []corev1.Container{
						{
							Name:            CertGen,
							Image:           instance.EnrollerImage(),
							ImagePullPolicy: corev1.PullAlways,
							SecurityContext: &corev1.SecurityContext{
								RunAsUser:  &user,
								Privileged: &t,
							},
							Env: hsmConfig.GetEnvs(),
							Command: []string{
								"sh",
								"-c",
							},
							Args: []string{
								fmt.Sprintf(config.DAEMON_CHECK_CMD+" && /usr/local/bin/enroller node enroll %s %s %s %s %s %s %s %s %s", e.CAClient.GetHomeDir(), "/tmp/fabric-ca-client-config.yaml", req.CAHost, req.CAPort, req.CAName, instance.GetName(), instance.GetNamespace(), req.EnrollID, req.EnrollSecret),
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "tlscertfile",
									MountPath: fmt.Sprintf("%s/tlsCert.pem", e.CAClient.GetHomeDir()),
									SubPath:   "tlsCert.pem",
								},
								{
									Name:      "clientconfig",
									MountPath: fmt.Sprintf("/tmp/%s", "fabric-ca-client-config.yaml"),
									SubPath:   "fabric-ca-client-config.yaml",
								},
								{
									Name:      "shared",
									MountPath: "/hsm/lib",
									SubPath:   "hsm",
								},
								{
									Name:      "shared",
									MountPath: "/shared",
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "shared",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{
									Medium: corev1.StorageMediumMemory,
								},
							},
						},
						{
							Name: "tlscertfile",
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName: fmt.Sprintf("%s-init-roottls", instance.GetName()),
								},
							},
						},
						{
							Name: "clientconfig",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: fmt.Sprintf("%s-init-config", instance.GetName()),
									},
								},
							},
						},
						{
							Name: pvcVolumeName,
							VolumeSource: corev1.VolumeSource{
								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: instance.PVCName(),
								},
							},
						},
					},
				},
			},
		},
	}

	job := jobv1.New(k8sJob, &jobv1.Timeouts{
		WaitUntilActive:   timeouts.JobStart.Get(),
		WaitUntilFinished: timeouts.JobCompletion.Get(),
	})

	job.Spec.Template.Spec.Volumes = append(job.Spec.Template.Spec.Volumes, hsmConfig.GetVolumes()...)
	job.Spec.Template.Spec.Containers[0].VolumeMounts = append(job.Spec.Template.Spec.Containers[0].VolumeMounts, hsmConfig.GetVolumeMounts()...)

	// If daemon settings are configured in HSM config, create a sidecar that is running the daemon image
	if e.Config.Daemon != nil {
		// Certain token information requires to be stored in persistent store, the administrator
		// responsible for configuring HSM sets the HSM config to point to the path where the PVC
		// needs to be mounted.
		var pvcMount *corev1.VolumeMount
		for _, vm := range e.Config.MountPaths {
			if vm.UsePVC {
				pvcMount = &corev1.VolumeMount{
					Name:      pvcVolumeName,
					MountPath: vm.MountPath,
				}
			}
		}

		// Add daemon container to the deployment
		config.AddDaemonContainer(e.Config, job, instance.GetResource(current.HSMDAEMON), pvcMount)

		// If a pvc mount has been configured in HSM config, set the volume mount on the CertGen container
		if pvcMount != nil {
			job.Spec.Template.Spec.Containers[0].VolumeMounts = append(job.Spec.Template.Spec.Containers[0].VolumeMounts, *pvcMount)
		}
	}

	return job
}
