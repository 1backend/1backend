/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.38
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PromptSvcPromptStatus the model 'PromptSvcPromptStatus'
type PromptSvcPromptStatus string

// List of prompt_svc.PromptStatus
const (
	PromptStatusScheduled PromptSvcPromptStatus = "scheduled"
	PromptStatusRunning PromptSvcPromptStatus = "running"
	PromptStatusCompleted PromptSvcPromptStatus = "completed"
	PromptStatusErrored PromptSvcPromptStatus = "errored"
	PromptStatusAbandoned PromptSvcPromptStatus = "abandoned"
	PromptStatusCanceled PromptSvcPromptStatus = "canceled"
)

// All allowed values of PromptSvcPromptStatus enum
var AllowedPromptSvcPromptStatusEnumValues = []PromptSvcPromptStatus{
	"scheduled",
	"running",
	"completed",
	"errored",
	"abandoned",
	"canceled",
}

func (v *PromptSvcPromptStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PromptSvcPromptStatus(value)
	for _, existing := range AllowedPromptSvcPromptStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PromptSvcPromptStatus", value)
}

// NewPromptSvcPromptStatusFromValue returns a pointer to a valid PromptSvcPromptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromptSvcPromptStatusFromValue(v string) (*PromptSvcPromptStatus, error) {
	ev := PromptSvcPromptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromptSvcPromptStatus: valid values are %v", v, AllowedPromptSvcPromptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromptSvcPromptStatus) IsValid() bool {
	for _, existing := range AllowedPromptSvcPromptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to prompt_svc.PromptStatus value
func (v PromptSvcPromptStatus) Ptr() *PromptSvcPromptStatus {
	return &v
}

type NullablePromptSvcPromptStatus struct {
	value *PromptSvcPromptStatus
	isSet bool
}

func (v NullablePromptSvcPromptStatus) Get() *PromptSvcPromptStatus {
	return v.value
}

func (v *NullablePromptSvcPromptStatus) Set(val *PromptSvcPromptStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcPromptStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcPromptStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcPromptStatus(val *PromptSvcPromptStatus) *NullablePromptSvcPromptStatus {
	return &NullablePromptSvcPromptStatus{value: val, isSet: true}
}

func (v NullablePromptSvcPromptStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcPromptStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

