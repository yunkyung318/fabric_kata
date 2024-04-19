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

package serviceaccount

import (
	"context"
	"fmt"

	k8sclient "github.com/IBM-Blockchain/fabric-operator/pkg/k8s/controllerclient"
	"github.com/IBM-Blockchain/fabric-operator/pkg/manager/resources"
	"github.com/IBM-Blockchain/fabric-operator/pkg/operatorerrors"
	"github.com/IBM-Blockchain/fabric-operator/pkg/util"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var log = logf.Log.WithName("serviceaccount_manager")

type Manager struct {
	Client             k8sclient.Client
	Scheme             *runtime.Scheme
	ServiceAccountFile string
	Name               string

	LabelsFunc   func(v1.Object) map[string]string
	OverrideFunc func(v1.Object, *corev1.ServiceAccount, resources.Action) error
}

func (m *Manager) GetName(instance v1.Object) string {
	return GetName(instance.GetName(), m.Name)
}

func (m *Manager) Reconcile(instance v1.Object, update bool) error {
	name := m.GetName(instance)
	err := m.Client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: instance.GetNamespace()}, &corev1.ServiceAccount{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			log.Info(fmt.Sprintf("Creating service account '%s'", name))
			serviceAccount, err := m.GetServiceAccountBasedOnCRFromFile(instance)
			if err != nil {
				return err
			}

			err = m.Client.Create(context.TODO(), serviceAccount, k8sclient.CreateOption{Owner: instance, Scheme: m.Scheme})
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	// TODO: If needed, update logic for servie goes here

	return nil
}

func (m *Manager) GetServiceAccountBasedOnCRFromFile(instance v1.Object) (*corev1.ServiceAccount, error) {
	serviceAccount, err := util.GetServiceAccountFromFile(m.ServiceAccountFile)
	if err != nil {
		log.Error(err, fmt.Sprintf("Error reading service account configuration file: %s", m.ServiceAccountFile))
		return nil, err
	}

	name := m.GetName(instance)
	serviceAccount.Name = name
	serviceAccount.Namespace = instance.GetNamespace()
	serviceAccount.Labels = m.LabelsFunc(instance)

	return m.BasedOnCR(instance, serviceAccount)
}

func (m *Manager) BasedOnCR(instance v1.Object, serviceAccount *corev1.ServiceAccount) (*corev1.ServiceAccount, error) {
	if m.OverrideFunc != nil {
		err := m.OverrideFunc(instance, serviceAccount, resources.Create)
		if err != nil {
			return nil, operatorerrors.New(operatorerrors.InvalidServiceAccountCreateRequest, err.Error())
		}
	}

	return serviceAccount, nil
}

func (m *Manager) Get(instance v1.Object) (client.Object, error) {
	if instance == nil {
		return nil, nil // Instance has not been reconciled yet
	}

	name := m.GetName(instance)
	sa := &corev1.ServiceAccount{}
	err := m.Client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: instance.GetNamespace()}, sa)
	if err != nil {
		return nil, err
	}

	return sa, nil
}

func (m *Manager) Exists(instance v1.Object) bool {
	_, err := m.Get(instance)
	if err != nil {
		return false
	}

	return true
}

func (m *Manager) Delete(instance v1.Object) error {
	sa, err := m.Get(instance)
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			return err
		}
	}

	if sa == nil {
		return nil
	}

	err = m.Client.Delete(context.TODO(), sa)
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			return err
		}
	}

	return nil
}

func (m *Manager) CheckState(instance v1.Object) error {
	// NO-OP
	return nil
}

func (m *Manager) RestoreState(instance v1.Object) error {
	// NO-OP
	return nil
}

func (m *Manager) SetCustomName(name string) {
	// NO-OP
}

func GetName(instanceName string, suffix ...string) string {
	if len(suffix) != 0 {
		if suffix[0] != "" {
			return fmt.Sprintf("%s-%s", instanceName, suffix[0])
		}
	}
	return fmt.Sprintf("%s", instanceName)
}
