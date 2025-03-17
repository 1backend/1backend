/*
1Backend

A unified backend for your AI applications—microservices-based and built to scale.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ContainerSvcLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcLabel{}

// ContainerSvcLabel struct for ContainerSvcLabel
type ContainerSvcLabel struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type _ContainerSvcLabel ContainerSvcLabel

// NewContainerSvcLabel instantiates a new ContainerSvcLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcLabel(key string, value string) *ContainerSvcLabel {
	this := ContainerSvcLabel{}
	this.Key = key
	this.Value = value
	return &this
}

// NewContainerSvcLabelWithDefaults instantiates a new ContainerSvcLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcLabelWithDefaults() *ContainerSvcLabel {
	this := ContainerSvcLabel{}
	return &this
}

// GetKey returns the Key field value
func (o *ContainerSvcLabel) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcLabel) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ContainerSvcLabel) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *ContainerSvcLabel) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcLabel) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ContainerSvcLabel) SetValue(v string) {
	o.Value = v
}

func (o ContainerSvcLabel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ContainerSvcLabel) UnmarshalJSON(data []byte) (err error) {
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

	varContainerSvcLabel := _ContainerSvcLabel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerSvcLabel)

	if err != nil {
		return err
	}

	*o = ContainerSvcLabel(varContainerSvcLabel)

	return err
}

type NullableContainerSvcLabel struct {
	value *ContainerSvcLabel
	isSet bool
}

func (v NullableContainerSvcLabel) Get() *ContainerSvcLabel {
	return v.value
}

func (v *NullableContainerSvcLabel) Set(val *ContainerSvcLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcLabel(val *ContainerSvcLabel) *NullableContainerSvcLabel {
	return &NullableContainerSvcLabel{value: val, isSet: true}
}

func (v NullableContainerSvcLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


