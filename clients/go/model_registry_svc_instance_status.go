/*
1Backend

AI-native microservices platform.

API version: 0.7.6
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// RegistrySvcInstanceStatus the model 'RegistrySvcInstanceStatus'
type RegistrySvcInstanceStatus string

// List of registry_svc.InstanceStatus
const (
	InstanceStatusUnknown RegistrySvcInstanceStatus = "Unknown"
	InstanceStatusHealthy RegistrySvcInstanceStatus = "Healthy"
	InstanceStatusDegraded RegistrySvcInstanceStatus = "Degraded"
	InstanceStatusUnreachable RegistrySvcInstanceStatus = "Unreachable"
	InstanceStatusError RegistrySvcInstanceStatus = "Error"
)

// All allowed values of RegistrySvcInstanceStatus enum
var AllowedRegistrySvcInstanceStatusEnumValues = []RegistrySvcInstanceStatus{
	"Unknown",
	"Healthy",
	"Degraded",
	"Unreachable",
	"Error",
}

func (v *RegistrySvcInstanceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegistrySvcInstanceStatus(value)
	for _, existing := range AllowedRegistrySvcInstanceStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegistrySvcInstanceStatus", value)
}

// NewRegistrySvcInstanceStatusFromValue returns a pointer to a valid RegistrySvcInstanceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegistrySvcInstanceStatusFromValue(v string) (*RegistrySvcInstanceStatus, error) {
	ev := RegistrySvcInstanceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegistrySvcInstanceStatus: valid values are %v", v, AllowedRegistrySvcInstanceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegistrySvcInstanceStatus) IsValid() bool {
	for _, existing := range AllowedRegistrySvcInstanceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to registry_svc.InstanceStatus value
func (v RegistrySvcInstanceStatus) Ptr() *RegistrySvcInstanceStatus {
	return &v
}

type NullableRegistrySvcInstanceStatus struct {
	value *RegistrySvcInstanceStatus
	isSet bool
}

func (v NullableRegistrySvcInstanceStatus) Get() *RegistrySvcInstanceStatus {
	return v.value
}

func (v *NullableRegistrySvcInstanceStatus) Set(val *RegistrySvcInstanceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcInstanceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcInstanceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcInstanceStatus(val *RegistrySvcInstanceStatus) *NullableRegistrySvcInstanceStatus {
	return &NullableRegistrySvcInstanceStatus{value: val, isSet: true}
}

func (v NullableRegistrySvcInstanceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcInstanceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

