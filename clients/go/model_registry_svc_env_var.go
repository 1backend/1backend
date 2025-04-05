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
	"bytes"
	"fmt"
)

// checks if the RegistrySvcEnvVar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcEnvVar{}

// RegistrySvcEnvVar struct for RegistrySvcEnvVar
type RegistrySvcEnvVar struct {
	// Key is the environment variable name.
	Key string `json:"key"`
	// Value is the environment variable value.
	Value string `json:"value"`
}

type _RegistrySvcEnvVar RegistrySvcEnvVar

// NewRegistrySvcEnvVar instantiates a new RegistrySvcEnvVar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcEnvVar(key string, value string) *RegistrySvcEnvVar {
	this := RegistrySvcEnvVar{}
	this.Key = key
	this.Value = value
	return &this
}

// NewRegistrySvcEnvVarWithDefaults instantiates a new RegistrySvcEnvVar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcEnvVarWithDefaults() *RegistrySvcEnvVar {
	this := RegistrySvcEnvVar{}
	return &this
}

// GetKey returns the Key field value
func (o *RegistrySvcEnvVar) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcEnvVar) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *RegistrySvcEnvVar) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *RegistrySvcEnvVar) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcEnvVar) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RegistrySvcEnvVar) SetValue(v string) {
	o.Value = v
}

func (o RegistrySvcEnvVar) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcEnvVar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *RegistrySvcEnvVar) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRegistrySvcEnvVar := _RegistrySvcEnvVar{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrySvcEnvVar)

	if err != nil {
		return err
	}

	*o = RegistrySvcEnvVar(varRegistrySvcEnvVar)

	return err
}

type NullableRegistrySvcEnvVar struct {
	value *RegistrySvcEnvVar
	isSet bool
}

func (v NullableRegistrySvcEnvVar) Get() *RegistrySvcEnvVar {
	return v.value
}

func (v *NullableRegistrySvcEnvVar) Set(val *RegistrySvcEnvVar) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcEnvVar) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcEnvVar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcEnvVar(val *RegistrySvcEnvVar) *NullableRegistrySvcEnvVar {
	return &NullableRegistrySvcEnvVar{value: val, isSet: true}
}

func (v NullableRegistrySvcEnvVar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcEnvVar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


