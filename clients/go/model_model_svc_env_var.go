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

// checks if the ModelSvcEnvVar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSvcEnvVar{}

// ModelSvcEnvVar struct for ModelSvcEnvVar
type ModelSvcEnvVar struct {
	// Key is the environment variable name.
	Key *string `json:"key,omitempty"`
	// Value is the environment variable value.
	Value *string `json:"value,omitempty"`
}

// NewModelSvcEnvVar instantiates a new ModelSvcEnvVar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSvcEnvVar() *ModelSvcEnvVar {
	this := ModelSvcEnvVar{}
	return &this
}

// NewModelSvcEnvVarWithDefaults instantiates a new ModelSvcEnvVar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSvcEnvVarWithDefaults() *ModelSvcEnvVar {
	this := ModelSvcEnvVar{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ModelSvcEnvVar) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcEnvVar) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ModelSvcEnvVar) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ModelSvcEnvVar) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ModelSvcEnvVar) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcEnvVar) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ModelSvcEnvVar) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ModelSvcEnvVar) SetValue(v string) {
	o.Value = &v
}

func (o ModelSvcEnvVar) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSvcEnvVar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableModelSvcEnvVar struct {
	value *ModelSvcEnvVar
	isSet bool
}

func (v NullableModelSvcEnvVar) Get() *ModelSvcEnvVar {
	return v.value
}

func (v *NullableModelSvcEnvVar) Set(val *ModelSvcEnvVar) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSvcEnvVar) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSvcEnvVar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSvcEnvVar(val *ModelSvcEnvVar) *NullableModelSvcEnvVar {
	return &NullableModelSvcEnvVar{value: val, isSet: true}
}

func (v NullableModelSvcEnvVar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSvcEnvVar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


