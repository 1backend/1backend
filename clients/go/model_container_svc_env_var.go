/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.31
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ContainerSvcEnvVar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcEnvVar{}

// ContainerSvcEnvVar struct for ContainerSvcEnvVar
type ContainerSvcEnvVar struct {
	// Key is the environment variable name.
	Key string `json:"key"`
	// Value is the environment variable value.
	Value string `json:"value"`
}

type _ContainerSvcEnvVar ContainerSvcEnvVar

// NewContainerSvcEnvVar instantiates a new ContainerSvcEnvVar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcEnvVar(key string, value string) *ContainerSvcEnvVar {
	this := ContainerSvcEnvVar{}
	this.Key = key
	this.Value = value
	return &this
}

// NewContainerSvcEnvVarWithDefaults instantiates a new ContainerSvcEnvVar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcEnvVarWithDefaults() *ContainerSvcEnvVar {
	this := ContainerSvcEnvVar{}
	return &this
}

// GetKey returns the Key field value
func (o *ContainerSvcEnvVar) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcEnvVar) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ContainerSvcEnvVar) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *ContainerSvcEnvVar) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcEnvVar) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ContainerSvcEnvVar) SetValue(v string) {
	o.Value = v
}

func (o ContainerSvcEnvVar) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcEnvVar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ContainerSvcEnvVar) UnmarshalJSON(data []byte) (err error) {
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

	varContainerSvcEnvVar := _ContainerSvcEnvVar{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerSvcEnvVar)

	if err != nil {
		return err
	}

	*o = ContainerSvcEnvVar(varContainerSvcEnvVar)

	return err
}

type NullableContainerSvcEnvVar struct {
	value *ContainerSvcEnvVar
	isSet bool
}

func (v NullableContainerSvcEnvVar) Get() *ContainerSvcEnvVar {
	return v.value
}

func (v *NullableContainerSvcEnvVar) Set(val *ContainerSvcEnvVar) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcEnvVar) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcEnvVar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcEnvVar(val *ContainerSvcEnvVar) *NullableContainerSvcEnvVar {
	return &NullableContainerSvcEnvVar{value: val, isSet: true}
}

func (v NullableContainerSvcEnvVar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcEnvVar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


