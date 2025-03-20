/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// DeploySvcDeploymentStatus the model 'DeploySvcDeploymentStatus'
type DeploySvcDeploymentStatus string

// List of deploy_svc.DeploymentStatus
const (
	DeploymentStatusOK DeploySvcDeploymentStatus = "OK"
	DeploymentStatusError DeploySvcDeploymentStatus = "Error"
	DeploymentStatusPending DeploySvcDeploymentStatus = "Pending"
	DeploymentStatusFailed DeploySvcDeploymentStatus = "Failed"
	DeploymentStatusDeploying DeploySvcDeploymentStatus = "Deploying"
)

// All allowed values of DeploySvcDeploymentStatus enum
var AllowedDeploySvcDeploymentStatusEnumValues = []DeploySvcDeploymentStatus{
	"OK",
	"Error",
	"Pending",
	"Failed",
	"Deploying",
}

func (v *DeploySvcDeploymentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeploySvcDeploymentStatus(value)
	for _, existing := range AllowedDeploySvcDeploymentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeploySvcDeploymentStatus", value)
}

// NewDeploySvcDeploymentStatusFromValue returns a pointer to a valid DeploySvcDeploymentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeploySvcDeploymentStatusFromValue(v string) (*DeploySvcDeploymentStatus, error) {
	ev := DeploySvcDeploymentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeploySvcDeploymentStatus: valid values are %v", v, AllowedDeploySvcDeploymentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeploySvcDeploymentStatus) IsValid() bool {
	for _, existing := range AllowedDeploySvcDeploymentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to deploy_svc.DeploymentStatus value
func (v DeploySvcDeploymentStatus) Ptr() *DeploySvcDeploymentStatus {
	return &v
}

type NullableDeploySvcDeploymentStatus struct {
	value *DeploySvcDeploymentStatus
	isSet bool
}

func (v NullableDeploySvcDeploymentStatus) Get() *DeploySvcDeploymentStatus {
	return v.value
}

func (v *NullableDeploySvcDeploymentStatus) Set(val *DeploySvcDeploymentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploySvcDeploymentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySvcDeploymentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySvcDeploymentStatus(val *DeploySvcDeploymentStatus) *NullableDeploySvcDeploymentStatus {
	return &NullableDeploySvcDeploymentStatus{value: val, isSet: true}
}

func (v NullableDeploySvcDeploymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySvcDeploymentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

