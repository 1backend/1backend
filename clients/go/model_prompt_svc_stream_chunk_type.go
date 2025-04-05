/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.34
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PromptSvcStreamChunkType the model 'PromptSvcStreamChunkType'
type PromptSvcStreamChunkType string

// List of prompt_svc_stream.ChunkType
const (
	ChunkTypeProgress PromptSvcStreamChunkType = "progress"
	ChunkTypeDone PromptSvcStreamChunkType = "done"
)

// All allowed values of PromptSvcStreamChunkType enum
var AllowedPromptSvcStreamChunkTypeEnumValues = []PromptSvcStreamChunkType{
	"progress",
	"done",
}

func (v *PromptSvcStreamChunkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PromptSvcStreamChunkType(value)
	for _, existing := range AllowedPromptSvcStreamChunkTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PromptSvcStreamChunkType", value)
}

// NewPromptSvcStreamChunkTypeFromValue returns a pointer to a valid PromptSvcStreamChunkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromptSvcStreamChunkTypeFromValue(v string) (*PromptSvcStreamChunkType, error) {
	ev := PromptSvcStreamChunkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromptSvcStreamChunkType: valid values are %v", v, AllowedPromptSvcStreamChunkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromptSvcStreamChunkType) IsValid() bool {
	for _, existing := range AllowedPromptSvcStreamChunkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to prompt_svc_stream.ChunkType value
func (v PromptSvcStreamChunkType) Ptr() *PromptSvcStreamChunkType {
	return &v
}

type NullablePromptSvcStreamChunkType struct {
	value *PromptSvcStreamChunkType
	isSet bool
}

func (v NullablePromptSvcStreamChunkType) Get() *PromptSvcStreamChunkType {
	return v.value
}

func (v *NullablePromptSvcStreamChunkType) Set(val *PromptSvcStreamChunkType) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcStreamChunkType) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcStreamChunkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcStreamChunkType(val *PromptSvcStreamChunkType) *NullablePromptSvcStreamChunkType {
	return &NullablePromptSvcStreamChunkType{value: val, isSet: true}
}

func (v NullablePromptSvcStreamChunkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcStreamChunkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

