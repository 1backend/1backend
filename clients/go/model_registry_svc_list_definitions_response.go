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
)

// checks if the RegistrySvcListDefinitionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcListDefinitionsResponse{}

// RegistrySvcListDefinitionsResponse struct for RegistrySvcListDefinitionsResponse
type RegistrySvcListDefinitionsResponse struct {
	Definitions []RegistrySvcDefinition `json:"definitions,omitempty"`
}

// NewRegistrySvcListDefinitionsResponse instantiates a new RegistrySvcListDefinitionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcListDefinitionsResponse() *RegistrySvcListDefinitionsResponse {
	this := RegistrySvcListDefinitionsResponse{}
	return &this
}

// NewRegistrySvcListDefinitionsResponseWithDefaults instantiates a new RegistrySvcListDefinitionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcListDefinitionsResponseWithDefaults() *RegistrySvcListDefinitionsResponse {
	this := RegistrySvcListDefinitionsResponse{}
	return &this
}

// GetDefinitions returns the Definitions field value if set, zero value otherwise.
func (o *RegistrySvcListDefinitionsResponse) GetDefinitions() []RegistrySvcDefinition {
	if o == nil || IsNil(o.Definitions) {
		var ret []RegistrySvcDefinition
		return ret
	}
	return o.Definitions
}

// GetDefinitionsOk returns a tuple with the Definitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcListDefinitionsResponse) GetDefinitionsOk() ([]RegistrySvcDefinition, bool) {
	if o == nil || IsNil(o.Definitions) {
		return nil, false
	}
	return o.Definitions, true
}

// HasDefinitions returns a boolean if a field has been set.
func (o *RegistrySvcListDefinitionsResponse) HasDefinitions() bool {
	if o != nil && !IsNil(o.Definitions) {
		return true
	}

	return false
}

// SetDefinitions gets a reference to the given []RegistrySvcDefinition and assigns it to the Definitions field.
func (o *RegistrySvcListDefinitionsResponse) SetDefinitions(v []RegistrySvcDefinition) {
	o.Definitions = v
}

func (o RegistrySvcListDefinitionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcListDefinitionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Definitions) {
		toSerialize["definitions"] = o.Definitions
	}
	return toSerialize, nil
}

type NullableRegistrySvcListDefinitionsResponse struct {
	value *RegistrySvcListDefinitionsResponse
	isSet bool
}

func (v NullableRegistrySvcListDefinitionsResponse) Get() *RegistrySvcListDefinitionsResponse {
	return v.value
}

func (v *NullableRegistrySvcListDefinitionsResponse) Set(val *RegistrySvcListDefinitionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcListDefinitionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcListDefinitionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcListDefinitionsResponse(val *RegistrySvcListDefinitionsResponse) *NullableRegistrySvcListDefinitionsResponse {
	return &NullableRegistrySvcListDefinitionsResponse{value: val, isSet: true}
}

func (v NullableRegistrySvcListDefinitionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcListDefinitionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


