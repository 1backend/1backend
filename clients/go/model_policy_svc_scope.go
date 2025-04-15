/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.37
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PolicySvcScope the model 'PolicySvcScope'
type PolicySvcScope string

// List of policy_svc.Scope
const (
	ScopeEndpoint PolicySvcScope = "endpoint"
	ScopeGlobal PolicySvcScope = "global"
)

// All allowed values of PolicySvcScope enum
var AllowedPolicySvcScopeEnumValues = []PolicySvcScope{
	"endpoint",
	"global",
}

func (v *PolicySvcScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicySvcScope(value)
	for _, existing := range AllowedPolicySvcScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicySvcScope", value)
}

// NewPolicySvcScopeFromValue returns a pointer to a valid PolicySvcScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicySvcScopeFromValue(v string) (*PolicySvcScope, error) {
	ev := PolicySvcScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicySvcScope: valid values are %v", v, AllowedPolicySvcScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicySvcScope) IsValid() bool {
	for _, existing := range AllowedPolicySvcScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to policy_svc.Scope value
func (v PolicySvcScope) Ptr() *PolicySvcScope {
	return &v
}

type NullablePolicySvcScope struct {
	value *PolicySvcScope
	isSet bool
}

func (v NullablePolicySvcScope) Get() *PolicySvcScope {
	return v.value
}

func (v *NullablePolicySvcScope) Set(val *PolicySvcScope) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySvcScope) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySvcScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySvcScope(val *PolicySvcScope) *NullablePolicySvcScope {
	return &NullablePolicySvcScope{value: val, isSet: true}
}

func (v NullablePolicySvcScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySvcScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

