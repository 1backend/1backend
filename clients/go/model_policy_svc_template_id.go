/*
1Backend

A unified backend development platform for your AI applications—microservices-based and built to scale.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PolicySvcTemplateId the model 'PolicySvcTemplateId'
type PolicySvcTemplateId string

// List of policy_svc.TemplateId
const (
	TemplateIdRateLimit PolicySvcTemplateId = "rate-limit"
	TemplateIdBlocklist PolicySvcTemplateId = "blocklist"
)

// All allowed values of PolicySvcTemplateId enum
var AllowedPolicySvcTemplateIdEnumValues = []PolicySvcTemplateId{
	"rate-limit",
	"blocklist",
}

func (v *PolicySvcTemplateId) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicySvcTemplateId(value)
	for _, existing := range AllowedPolicySvcTemplateIdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicySvcTemplateId", value)
}

// NewPolicySvcTemplateIdFromValue returns a pointer to a valid PolicySvcTemplateId
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicySvcTemplateIdFromValue(v string) (*PolicySvcTemplateId, error) {
	ev := PolicySvcTemplateId(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicySvcTemplateId: valid values are %v", v, AllowedPolicySvcTemplateIdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicySvcTemplateId) IsValid() bool {
	for _, existing := range AllowedPolicySvcTemplateIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to policy_svc.TemplateId value
func (v PolicySvcTemplateId) Ptr() *PolicySvcTemplateId {
	return &v
}

type NullablePolicySvcTemplateId struct {
	value *PolicySvcTemplateId
	isSet bool
}

func (v NullablePolicySvcTemplateId) Get() *PolicySvcTemplateId {
	return v.value
}

func (v *NullablePolicySvcTemplateId) Set(val *PolicySvcTemplateId) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySvcTemplateId) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySvcTemplateId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySvcTemplateId(val *PolicySvcTemplateId) *NullablePolicySvcTemplateId {
	return &NullablePolicySvcTemplateId{value: val, isSet: true}
}

func (v NullablePolicySvcTemplateId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySvcTemplateId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

