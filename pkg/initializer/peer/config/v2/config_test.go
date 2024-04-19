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

package v2_test

import (
	"github.com/IBM-Blockchain/fabric-operator/pkg/apis/common"
	v2core "github.com/IBM-Blockchain/fabric-operator/pkg/apis/peer/v2"
	v2 "github.com/IBM-Blockchain/fabric-operator/pkg/initializer/peer/config/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Peer configuration", func() {
	It("merges current configuration with overrides values", func() {
		core, err := v2.ReadCoreFile("../../../../../testdata/init/peer/core.yaml")
		Expect(err).NotTo(HaveOccurred())
		Expect(core.Peer.ID).To(Equal("jdoe"))

		newConfig := &v2.Core{
			Core: v2core.Core{
				Peer: v2core.Peer{
					BCCSP: &common.BCCSP{
						Default: "PKCS11",
						PKCS11: &common.PKCS11Opts{
							Library:  "library2",
							Label:    "label2",
							Pin:      "2222",
							Hash:     "SHA3",
							Security: 512,
							FileKeyStore: &common.FileKeyStoreOpts{
								KeyStorePath: "keystore3",
							},
						},
					},
				},
			},
		}

		Expect(core.Peer.Keepalive.MinInterval).To(Equal(common.MustParseDuration("60s")))

		err = core.MergeWith(newConfig, true)
		Expect(err).NotTo(HaveOccurred())

		Expect(*core.Peer.BCCSP.PKCS11).To(Equal(common.PKCS11Opts{
			Library:        "/usr/local/lib/libpkcs11-proxy.so",
			Label:          "label2",
			Pin:            "2222",
			Hash:           "SHA3",
			Security:       512,
			SoftwareVerify: true,
			FileKeyStore: &common.FileKeyStoreOpts{
				KeyStorePath: "keystore3",
			},
		}))
	})

	Context("chaincode configuration", func() {
		It("merges v2 current configuration with overrides values", func() {
			core, err := v2.ReadCoreFile("../../../../../testdata/init/peer/core.yaml")
			Expect(err).NotTo(HaveOccurred())
			Expect(core.Peer.ID).To(Equal("jdoe"))

			startupTimeout, err := common.ParseDuration("200s")
			Expect(err).NotTo(HaveOccurred())
			executeTimeout, err := common.ParseDuration("20s")
			Expect(err).NotTo(HaveOccurred())

			newConfig := &v2.Core{
				Core: v2core.Core{
					Chaincode: v2core.Chaincode{
						StartupTimeout: startupTimeout,
						ExecuteTimeout: executeTimeout,
						ExternalBuilders: []v2core.ExternalBuilder{
							v2core.ExternalBuilder{
								Path:                 "/scripts",
								Name:                 "go-builder",
								EnvironmentWhiteList: []string{"ENV1=Value1"},
								PropogateEnvironment: []string{"ENV1=Value1"},
							},
						},
					},
				},
			}

			err = core.MergeWith(newConfig, false)
			Expect(err).NotTo(HaveOccurred())
			Expect(core.Chaincode.StartupTimeout).To(Equal(startupTimeout))
			Expect(core.Chaincode.ExecuteTimeout).To(Equal(executeTimeout))

			Expect(core.Chaincode.ExternalBuilders[0]).To(Equal(
				v2core.ExternalBuilder{
					Path:                 "/scripts",
					Name:                 "go-builder",
					EnvironmentWhiteList: []string{"ENV1=Value1"},
					PropogateEnvironment: []string{"ENV1=Value1"},
				},
			))
		})
	})

	Context("read in core file", func() {
		It("reads core and converts peer.gossip.bootstrap", func() {
			core, err := v2.ReadCoreFile("../../../../../testdata/init/peer/core_bootstrap_test.yaml")
			Expect(err).NotTo(HaveOccurred())
			Expect(core.Peer.Gossip.Bootstrap).To(Equal([]string{"127.0.0.1:7051"}))
		})

		It("returns error if invalid core (besides peer.gossip.boostrap field)", func() {
			_, err := v2.ReadCoreFile("../../../../../testdata/init/peer/core_invalid.yaml")
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
