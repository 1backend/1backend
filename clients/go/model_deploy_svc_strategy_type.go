/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.35
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// DeploySvcStrategyType the model 'DeploySvcStrategyType'
type DeploySvcStrategyType string

// List of deploy_svc.StrategyType
const (
	StrategyRollingUpdate DeploySvcStrategyType = "RollingUpdate"
	StrategyRecreate DeploySvcStrategyType = "Recreate"
)

// All allowed values of DeploySvcStrategyType enum
var AllowedDeploySvcStrategyTypeEnumValues = []DeploySvcStrategyType{
	"RollingUpdate",
	"Recreate",
}

func (v *DeploySvcStrategyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeploySvcStrategyType(value)
	for _, existing := range AllowedDeploySvcStrategyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeploySvcStrategyType", value)
}

// NewDeploySvcStrategyTypeFromValue returns a pointer to a valid DeploySvcStrategyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeploySvcStrategyTypeFromValue(v string) (*DeploySvcStrategyType, error) {
	ev := DeploySvcStrategyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeploySvcStrategyType: valid values are %v", v, AllowedDeploySvcStrategyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeploySvcStrategyType) IsValid() bool {
	for _, existing := range AllowedDeploySvcStrategyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to deploy_svc.StrategyType value
func (v DeploySvcStrategyType) Ptr() *DeploySvcStrategyType {
	return &v
}

type NullableDeploySvcStrategyType struct {
	value *DeploySvcStrategyType
	isSet bool
}

func (v NullableDeploySvcStrategyType) Get() *DeploySvcStrategyType {
	return v.value
}

func (v *NullableDeploySvcStrategyType) Set(val *DeploySvcStrategyType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploySvcStrategyType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySvcStrategyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySvcStrategyType(val *DeploySvcStrategyType) *NullableDeploySvcStrategyType {
	return &NullableDeploySvcStrategyType{value: val, isSet: true}
}

func (v NullableDeploySvcStrategyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySvcStrategyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

