// Copyright 2019 Smart-Edge.com, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clients_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/smartedgemec/controller-ce/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("Application Lifecycle Service", func() {
	var (
		containerAppID *string
		vmAppID        *string
	)

	BeforeEach(func() {
		var err error

		By("Deploying a container application")
		containerAppID, err = appDeploySvcCli.DeployContainer(
			ctx,
			&pb.Application{
				Name:        "test_container_app",
				Vendor:      "test_vendor",
				Description: "test container app",
				Image:       "http://test.com/container123",
				Cores:       4,
				Memory:      4096,
			})
		Expect(err).ToNot(HaveOccurred())

		By("Deploying a VM application")
		vmAppID, err = appDeploySvcCli.DeployVM(
			ctx,
			&pb.Application{
				Name:        "test_vm_app",
				Vendor:      "test_vendor",
				Description: "test vm app",
				Image:       "http://test.com/vm123",
				Cores:       4,
				Memory:      4096,
			})
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("Start", func() {
		Describe("Success", func() {
			It("Should start container applications", func() {
				By("Starting the container application")
				err := appLifeSvcCli.Start(ctx, *containerAppID)

				By("Verifying a success response")
				Expect(err).ToNot(HaveOccurred())

				By("Verifying the container application is started")
				containerApp, err := appDeploySvcCli.Get(ctx, *containerAppID)
				Expect(err).ToNot(HaveOccurred())
				Expect(containerApp.Status).To(Equal(
					pb.LifecycleStatus_RUNNING))
			})

			It("Should start VM applications", func() {
				By("Starting the VM application")
				err := appLifeSvcCli.Start(ctx, *vmAppID)
				Expect(err).ToNot(HaveOccurred())

				By("Verifying a success response")
				Expect(err).ToNot(HaveOccurred())

				By("Verifying the VM application is started")
				vmApp, err := appDeploySvcCli.Get(ctx, *vmAppID)
				Expect(err).ToNot(HaveOccurred())
				Expect(vmApp.Status).To(Equal(pb.LifecycleStatus_RUNNING))
			})
		})

		Describe("Errors", func() {
			It("Should return an error if the application is already "+
				"running", func() {
				By("Starting the container application")
				err := appLifeSvcCli.Start(ctx, *containerAppID)
				Expect(err).ToNot(HaveOccurred())

				By("Attempting to start the container application again")
				err = appLifeSvcCli.Start(ctx, *containerAppID)

				By("Verifying a FailedPrecondition response")
				Expect(err).To(HaveOccurred())
				Expect(errors.Cause(err)).To(Equal(
					status.Errorf(codes.FailedPrecondition,
						"Application %s not stopped", *containerAppID)))
			})
		})
	})

	Describe("Restart", func() {
		Describe("Success", func() {
			It("Should restart container applications", func() {
				By("Starting the container application")
				err := appLifeSvcCli.Start(ctx, *containerAppID)
				Expect(err).ToNot(HaveOccurred())

				By("Restarting the container application")
				err = appLifeSvcCli.Restart(ctx, *containerAppID)

				By("Verifying a success response")
				Expect(err).ToNot(HaveOccurred())

				By("Verifying the container application is restarted")
				containerApp, err := appDeploySvcCli.Get(ctx, *containerAppID)
				Expect(err).ToNot(HaveOccurred())
				Expect(containerApp.Status).To(Equal(
					pb.LifecycleStatus_RUNNING))
			})

			It("Should restart VM applications", func() {
				By("Starting the VM application")
				err := appLifeSvcCli.Start(ctx, *vmAppID)
				Expect(err).ToNot(HaveOccurred())

				By("Restarting the VM application")
				err = appLifeSvcCli.Restart(ctx, *vmAppID)

				By("Verifying a success response")
				Expect(err).ToNot(HaveOccurred())

				By("Verifying the VM application is restarted")
				vmApp, err := appDeploySvcCli.Get(ctx, *vmAppID)
				Expect(err).ToNot(HaveOccurred())
				Expect(vmApp.Status).To(Equal(pb.LifecycleStatus_RUNNING))
			})
		})

		Describe("Errors", func() {
			It("Should return an error if the application is not "+
				"running", func() {
				By("Attempting to restart the container application")
				err := appLifeSvcCli.Restart(ctx, *containerAppID)

				By("Verifying a FailedPrecondition response")
				Expect(err).To(HaveOccurred())
				Expect(errors.Cause(err)).To(Equal(
					status.Errorf(codes.FailedPrecondition,
						"Application %s not running", *containerAppID)))
			})
		})
	})

	Describe("Stop", func() {
		Describe("Success", func() {
			It("Should stop container applications", func() {
				By("Starting the container application")
				err := appLifeSvcCli.Start(ctx, *containerAppID)
				Expect(err).ToNot(HaveOccurred())

				By("Stopping the container application")
				err = appLifeSvcCli.Stop(ctx, *containerAppID)

				By("Verifying a success response")
				Expect(err).ToNot(HaveOccurred())

				By("Verifying the container application is stopped")
				containerApp, err := appDeploySvcCli.Get(ctx, *containerAppID)
				Expect(err).ToNot(HaveOccurred())
				Expect(containerApp.Status).To(Equal(
					pb.LifecycleStatus_STOPPED))
			})

			It("Should stop VM applications", func() {
				By("Starting the VM application")
				err := appLifeSvcCli.Start(ctx, *vmAppID)
				Expect(err).ToNot(HaveOccurred())

				By("Stopping the VM application")
				err = appLifeSvcCli.Stop(ctx, *vmAppID)

				By("Verifying a success response")
				Expect(err).ToNot(HaveOccurred())

				By("Verifying the VM application is stopped")
				vmApp, err := appDeploySvcCli.Get(ctx, *vmAppID)
				Expect(err).ToNot(HaveOccurred())
				Expect(vmApp.Status).To(Equal(pb.LifecycleStatus_STOPPED))
			})
		})

		Describe("Errors", func() {
			It("Should return an error if the application is already "+
				"stopped", func() {
				By("Starting the container application")
				err := appLifeSvcCli.Start(ctx, *containerAppID)
				Expect(err).ToNot(HaveOccurred())

				By("Stopping the container application")
				err = appLifeSvcCli.Stop(ctx, *containerAppID)
				Expect(err).ToNot(HaveOccurred())

				By("Attempting to stop the container application again")
				err = appLifeSvcCli.Stop(ctx, *containerAppID)

				By("Verifying a FailedPrecondition response")
				Expect(err).To(HaveOccurred())
				Expect(errors.Cause(err)).To(Equal(
					status.Errorf(codes.FailedPrecondition,
						"Application %s not running", *containerAppID)))

			})
		})
	})
})