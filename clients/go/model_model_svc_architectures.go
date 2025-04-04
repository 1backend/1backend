/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.32
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ModelSvcArchitectures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSvcArchitectures{}

// ModelSvcArchitectures struct for ModelSvcArchitectures
type ModelSvcArchitectures struct {
	// CUDA-specific container parameters, if applicable.
	Cuda *ModelSvcCudaParameters `json:"cuda,omitempty"`
	// Default container configuration for non-GPU environments.
	Default *ModelSvcDefaultParameters `json:"default,omitempty"`
}

// NewModelSvcArchitectures instantiates a new ModelSvcArchitectures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSvcArchitectures() *ModelSvcArchitectures {
	this := ModelSvcArchitectures{}
	return &this
}

// NewModelSvcArchitecturesWithDefaults instantiates a new ModelSvcArchitectures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSvcArchitecturesWithDefaults() *ModelSvcArchitectures {
	this := ModelSvcArchitectures{}
	return &this
}

// GetCuda returns the Cuda field value if set, zero value otherwise.
func (o *ModelSvcArchitectures) GetCuda() ModelSvcCudaParameters {
	if o == nil || IsNil(o.Cuda) {
		var ret ModelSvcCudaParameters
		return ret
	}
	return *o.Cuda
}

// GetCudaOk returns a tuple with the Cuda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcArchitectures) GetCudaOk() (*ModelSvcCudaParameters, bool) {
	if o == nil || IsNil(o.Cuda) {
		return nil, false
	}
	return o.Cuda, true
}

// HasCuda returns a boolean if a field has been set.
func (o *ModelSvcArchitectures) HasCuda() bool {
	if o != nil && !IsNil(o.Cuda) {
		return true
	}

	return false
}

// SetCuda gets a reference to the given ModelSvcCudaParameters and assigns it to the Cuda field.
func (o *ModelSvcArchitectures) SetCuda(v ModelSvcCudaParameters) {
	o.Cuda = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ModelSvcArchitectures) GetDefault() ModelSvcDefaultParameters {
	if o == nil || IsNil(o.Default) {
		var ret ModelSvcDefaultParameters
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcArchitectures) GetDefaultOk() (*ModelSvcDefaultParameters, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ModelSvcArchitectures) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given ModelSvcDefaultParameters and assigns it to the Default field.
func (o *ModelSvcArchitectures) SetDefault(v ModelSvcDefaultParameters) {
	o.Default = &v
}

func (o ModelSvcArchitectures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSvcArchitectures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cuda) {
		toSerialize["cuda"] = o.Cuda
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return toSerialize, nil
}

type NullableModelSvcArchitectures struct {
	value *ModelSvcArchitectures
	isSet bool
}

func (v NullableModelSvcArchitectures) Get() *ModelSvcArchitectures {
	return v.value
}

func (v *NullableModelSvcArchitectures) Set(val *ModelSvcArchitectures) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSvcArchitectures) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSvcArchitectures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSvcArchitectures(val *ModelSvcArchitectures) *NullableModelSvcArchitectures {
	return &NullableModelSvcArchitectures{value: val, isSet: true}
}

func (v NullableModelSvcArchitectures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSvcArchitectures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


