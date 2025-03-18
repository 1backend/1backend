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
)

// checks if the RegistrySvcSaveDefinitionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcSaveDefinitionRequest{}

// RegistrySvcSaveDefinitionRequest struct for RegistrySvcSaveDefinitionRequest
type RegistrySvcSaveDefinitionRequest struct {
	Definition *RegistrySvcDefinition `json:"definition,omitempty"`
}

// NewRegistrySvcSaveDefinitionRequest instantiates a new RegistrySvcSaveDefinitionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcSaveDefinitionRequest() *RegistrySvcSaveDefinitionRequest {
	this := RegistrySvcSaveDefinitionRequest{}
	return &this
}

// NewRegistrySvcSaveDefinitionRequestWithDefaults instantiates a new RegistrySvcSaveDefinitionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcSaveDefinitionRequestWithDefaults() *RegistrySvcSaveDefinitionRequest {
	this := RegistrySvcSaveDefinitionRequest{}
	return &this
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *RegistrySvcSaveDefinitionRequest) GetDefinition() RegistrySvcDefinition {
	if o == nil || IsNil(o.Definition) {
		var ret RegistrySvcDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcSaveDefinitionRequest) GetDefinitionOk() (*RegistrySvcDefinition, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *RegistrySvcSaveDefinitionRequest) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given RegistrySvcDefinition and assigns it to the Definition field.
func (o *RegistrySvcSaveDefinitionRequest) SetDefinition(v RegistrySvcDefinition) {
	o.Definition = &v
}

func (o RegistrySvcSaveDefinitionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcSaveDefinitionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	return toSerialize, nil
}

type NullableRegistrySvcSaveDefinitionRequest struct {
	value *RegistrySvcSaveDefinitionRequest
	isSet bool
}

func (v NullableRegistrySvcSaveDefinitionRequest) Get() *RegistrySvcSaveDefinitionRequest {
	return v.value
}

func (v *NullableRegistrySvcSaveDefinitionRequest) Set(val *RegistrySvcSaveDefinitionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcSaveDefinitionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcSaveDefinitionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcSaveDefinitionRequest(val *RegistrySvcSaveDefinitionRequest) *NullableRegistrySvcSaveDefinitionRequest {
	return &NullableRegistrySvcSaveDefinitionRequest{value: val, isSet: true}
}

func (v NullableRegistrySvcSaveDefinitionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcSaveDefinitionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


