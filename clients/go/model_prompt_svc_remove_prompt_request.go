/*
1Backend

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PromptSvcRemovePromptRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptSvcRemovePromptRequest{}

// PromptSvcRemovePromptRequest struct for PromptSvcRemovePromptRequest
type PromptSvcRemovePromptRequest struct {
	PromptId *string `json:"promptId,omitempty"`
}

// NewPromptSvcRemovePromptRequest instantiates a new PromptSvcRemovePromptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptSvcRemovePromptRequest() *PromptSvcRemovePromptRequest {
	this := PromptSvcRemovePromptRequest{}
	return &this
}

// NewPromptSvcRemovePromptRequestWithDefaults instantiates a new PromptSvcRemovePromptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptSvcRemovePromptRequestWithDefaults() *PromptSvcRemovePromptRequest {
	this := PromptSvcRemovePromptRequest{}
	return &this
}

// GetPromptId returns the PromptId field value if set, zero value otherwise.
func (o *PromptSvcRemovePromptRequest) GetPromptId() string {
	if o == nil || IsNil(o.PromptId) {
		var ret string
		return ret
	}
	return *o.PromptId
}

// GetPromptIdOk returns a tuple with the PromptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcRemovePromptRequest) GetPromptIdOk() (*string, bool) {
	if o == nil || IsNil(o.PromptId) {
		return nil, false
	}
	return o.PromptId, true
}

// HasPromptId returns a boolean if a field has been set.
func (o *PromptSvcRemovePromptRequest) HasPromptId() bool {
	if o != nil && !IsNil(o.PromptId) {
		return true
	}

	return false
}

// SetPromptId gets a reference to the given string and assigns it to the PromptId field.
func (o *PromptSvcRemovePromptRequest) SetPromptId(v string) {
	o.PromptId = &v
}

func (o PromptSvcRemovePromptRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptSvcRemovePromptRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PromptId) {
		toSerialize["promptId"] = o.PromptId
	}
	return toSerialize, nil
}

type NullablePromptSvcRemovePromptRequest struct {
	value *PromptSvcRemovePromptRequest
	isSet bool
}

func (v NullablePromptSvcRemovePromptRequest) Get() *PromptSvcRemovePromptRequest {
	return v.value
}

func (v *NullablePromptSvcRemovePromptRequest) Set(val *PromptSvcRemovePromptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcRemovePromptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcRemovePromptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcRemovePromptRequest(val *PromptSvcRemovePromptRequest) *NullablePromptSvcRemovePromptRequest {
	return &NullablePromptSvcRemovePromptRequest{value: val, isSet: true}
}

func (v NullablePromptSvcRemovePromptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcRemovePromptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


