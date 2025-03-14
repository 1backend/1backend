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

// checks if the PromptSvcEngineParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptSvcEngineParameters{}

// PromptSvcEngineParameters struct for PromptSvcEngineParameters
type PromptSvcEngineParameters struct {
	LlamaCppParameters *PromptSvcLlamaCppParameters `json:"llamaCppParameters,omitempty"`
	StableDiffusion *PromptSvcStableDiffusionParameters `json:"stableDiffusion,omitempty"`
}

// NewPromptSvcEngineParameters instantiates a new PromptSvcEngineParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptSvcEngineParameters() *PromptSvcEngineParameters {
	this := PromptSvcEngineParameters{}
	return &this
}

// NewPromptSvcEngineParametersWithDefaults instantiates a new PromptSvcEngineParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptSvcEngineParametersWithDefaults() *PromptSvcEngineParameters {
	this := PromptSvcEngineParameters{}
	return &this
}

// GetLlamaCppParameters returns the LlamaCppParameters field value if set, zero value otherwise.
func (o *PromptSvcEngineParameters) GetLlamaCppParameters() PromptSvcLlamaCppParameters {
	if o == nil || IsNil(o.LlamaCppParameters) {
		var ret PromptSvcLlamaCppParameters
		return ret
	}
	return *o.LlamaCppParameters
}

// GetLlamaCppParametersOk returns a tuple with the LlamaCppParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcEngineParameters) GetLlamaCppParametersOk() (*PromptSvcLlamaCppParameters, bool) {
	if o == nil || IsNil(o.LlamaCppParameters) {
		return nil, false
	}
	return o.LlamaCppParameters, true
}

// HasLlamaCppParameters returns a boolean if a field has been set.
func (o *PromptSvcEngineParameters) HasLlamaCppParameters() bool {
	if o != nil && !IsNil(o.LlamaCppParameters) {
		return true
	}

	return false
}

// SetLlamaCppParameters gets a reference to the given PromptSvcLlamaCppParameters and assigns it to the LlamaCppParameters field.
func (o *PromptSvcEngineParameters) SetLlamaCppParameters(v PromptSvcLlamaCppParameters) {
	o.LlamaCppParameters = &v
}

// GetStableDiffusion returns the StableDiffusion field value if set, zero value otherwise.
func (o *PromptSvcEngineParameters) GetStableDiffusion() PromptSvcStableDiffusionParameters {
	if o == nil || IsNil(o.StableDiffusion) {
		var ret PromptSvcStableDiffusionParameters
		return ret
	}
	return *o.StableDiffusion
}

// GetStableDiffusionOk returns a tuple with the StableDiffusion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcEngineParameters) GetStableDiffusionOk() (*PromptSvcStableDiffusionParameters, bool) {
	if o == nil || IsNil(o.StableDiffusion) {
		return nil, false
	}
	return o.StableDiffusion, true
}

// HasStableDiffusion returns a boolean if a field has been set.
func (o *PromptSvcEngineParameters) HasStableDiffusion() bool {
	if o != nil && !IsNil(o.StableDiffusion) {
		return true
	}

	return false
}

// SetStableDiffusion gets a reference to the given PromptSvcStableDiffusionParameters and assigns it to the StableDiffusion field.
func (o *PromptSvcEngineParameters) SetStableDiffusion(v PromptSvcStableDiffusionParameters) {
	o.StableDiffusion = &v
}

func (o PromptSvcEngineParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptSvcEngineParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LlamaCppParameters) {
		toSerialize["llamaCppParameters"] = o.LlamaCppParameters
	}
	if !IsNil(o.StableDiffusion) {
		toSerialize["stableDiffusion"] = o.StableDiffusion
	}
	return toSerialize, nil
}

type NullablePromptSvcEngineParameters struct {
	value *PromptSvcEngineParameters
	isSet bool
}

func (v NullablePromptSvcEngineParameters) Get() *PromptSvcEngineParameters {
	return v.value
}

func (v *NullablePromptSvcEngineParameters) Set(val *PromptSvcEngineParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcEngineParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcEngineParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcEngineParameters(val *PromptSvcEngineParameters) *NullablePromptSvcEngineParameters {
	return &NullablePromptSvcEngineParameters{value: val, isSet: true}
}

func (v NullablePromptSvcEngineParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcEngineParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


