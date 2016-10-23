/*
Copyright 2016 The Kubernetes Authors.

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

// Code generated by ""fitask" -type=natgateway"; DO NOT EDIT

package awstasks

import (
	"encoding/json"

	"k8s.io/kops/upup/pkg/fi"
)

// VPC

// JSON marshalling boilerplate
type realNatGateway NATGateway

func (o *NATGateway) UnmarshalJSON(data []byte) error {
	var jsonName string
	if err := json.Unmarshal(data, &jsonName); err == nil {
		o.AllocationID = &jsonName
		return nil
	}

	var r realNatGateway
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*o = NATGateway(r)
	return nil
}

var _ fi.HasName = &NATGateway{}

func (e *NATGateway) GetName() *string {
	return e.AllocationID
}

func (e *NATGateway) SetName(name string) {
	e.AllocationID = &name
}

func (e *NATGateway) String() string {
	return fi.TaskAsString(e)
}