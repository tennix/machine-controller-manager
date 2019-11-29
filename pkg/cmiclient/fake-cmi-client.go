/*
Copyright (c) 2017 SAP SE or an SAP affiliate company. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package cmiclient contains the cloud provider specific implementations to manage machines
package cmiclient

import (
	corev1 "k8s.io/api/core/v1"
)

// FakeCMIClient is a fake driver returned when none of the actual drivers match
type FakeCMIClient struct {
	MachineID string
	NodeName  string
	Err       error
}

// NewFakeCMIClient returns a new fakedriver object
func NewFakeCMIClient(machineID string, nodeName string, err error) *FakeCMIClient {
	return &FakeCMIClient{
		MachineID: machineID,
		NodeName:  nodeName,
		Err:       err,
	}
}

// CreateMachine makes a gRPC call to the driver to create the machine.
func (c *FakeCMIClient) CreateMachine() (string, string, string, error) {
	return c.MachineID, c.NodeName, "", c.Err
}

// DeleteMachine make a grpc call to the driver to delete the machine.
func (c *FakeCMIClient) DeleteMachine() (string, error) {
	return "", c.Err
}

// GetMachineStatus makes a gRPC call to the driver to check existance of machine
func (c *FakeCMIClient) GetMachineStatus() (string, string, string, error) {
	if c.Err == nil {
		return "", "", "", nil
	}

	return "", "", "", c.Err
}

// ListMachines have to list machines
func (c *FakeCMIClient) ListMachines() (map[string]string, error) {
	var mapOfMachines map[string]string
	return mapOfMachines, c.Err
}

// ShutDownMachine implements shutdownmachine
func (c *FakeCMIClient) ShutDownMachine() error {
	return c.Err
}

// GetProviderID returns the GetProviderID
func (c *FakeCMIClient) GetProviderID() string {
	return c.MachineID
}

// GetVolumeIDs returns a list of VolumeIDs for the PV spec list supplied
func (c *FakeCMIClient) GetVolumeIDs(pvSpecs []*corev1.PersistentVolumeSpec) ([]string, error) {
	return []string{}, c.Err
}
