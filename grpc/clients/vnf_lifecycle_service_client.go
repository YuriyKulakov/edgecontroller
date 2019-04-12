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

package clients

import (
	"context"

	"github.com/pkg/errors"
	"github.com/smartedgemec/controller-ce/grpc"
	"github.com/smartedgemec/controller-ce/pb"
)

// VNFLifecycleServiceClient wraps the PB client.
type VNFLifecycleServiceClient struct {
	pbCli pb.VNFLifecycleServiceClient
}

// NewVNFLifecycleServiceClient creates a new client.
func NewVNFLifecycleServiceClient(
	conn *grpc.ClientConn,
) *VNFLifecycleServiceClient {
	return &VNFLifecycleServiceClient{
		conn.NewVNFLifecycleServiceClient(),
	}
}

// Start starts a stopped VNF.
func (c *VNFLifecycleServiceClient) Start(
	ctx context.Context,
	id string,
) error {
	_, err := c.pbCli.Start(
		ctx,
		&pb.LifecycleCommand{
			Id:  id,
			Cmd: pb.LifecycleCommand_START,
		})

	if err != nil {
		return errors.Wrap(err, "error starting vnf")
	}

	return nil
}

// Stop stops a running VNF.
func (c *VNFLifecycleServiceClient) Stop(
	ctx context.Context,
	id string,
) error {
	_, err := c.pbCli.Stop(
		ctx,
		&pb.LifecycleCommand{
			Id:  id,
			Cmd: pb.LifecycleCommand_STOP,
		})

	if err != nil {
		return errors.Wrap(err, "error stopping vnf")
	}

	return nil
}

// Restart restarts a running VNF.
func (c *VNFLifecycleServiceClient) Restart(
	ctx context.Context,
	id string,
) error {
	_, err := c.pbCli.Restart(
		ctx,
		&pb.LifecycleCommand{
			Id:  id,
			Cmd: pb.LifecycleCommand_RESTART,
		})

	if err != nil {
		return errors.Wrap(err, "error restarting vnf")
	}

	return nil
}