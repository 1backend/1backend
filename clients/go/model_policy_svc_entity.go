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

// PolicySvcEntity the model 'PolicySvcEntity'
type PolicySvcEntity string

// List of policy_svc.Entity
const (
	EntityUserID PolicySvcEntity = "userId"
	EntityIP PolicySvcEntity = "ip"
)

// All allowed values of PolicySvcEntity enum
var AllowedPolicySvcEntityEnumValues = []PolicySvcEntity{
	"userId",
	"ip",
}

func (v *PolicySvcEntity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicySvcEntity(value)
	for _, existing := range AllowedPolicySvcEntityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicySvcEntity", value)
}

// NewPolicySvcEntityFromValue returns a pointer to a valid PolicySvcEntity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicySvcEntityFromValue(v string) (*PolicySvcEntity, error) {
	ev := PolicySvcEntity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicySvcEntity: valid values are %v", v, AllowedPolicySvcEntityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicySvcEntity) IsValid() bool {
	for _, existing := range AllowedPolicySvcEntityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to policy_svc.Entity value
func (v PolicySvcEntity) Ptr() *PolicySvcEntity {
	return &v
}

type NullablePolicySvcEntity struct {
	value *PolicySvcEntity
	isSet bool
}

func (v NullablePolicySvcEntity) Get() *PolicySvcEntity {
	return v.value
}

func (v *NullablePolicySvcEntity) Set(val *PolicySvcEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySvcEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySvcEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySvcEntity(val *PolicySvcEntity) *NullablePolicySvcEntity {
	return &NullablePolicySvcEntity{value: val, isSet: true}
}

func (v NullablePolicySvcEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySvcEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

