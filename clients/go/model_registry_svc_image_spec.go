/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.38
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RegistrySvcImageSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcImageSpec{}

// RegistrySvcImageSpec struct for RegistrySvcImageSpec
type RegistrySvcImageSpec struct {
	// InternalPorts are the ports the container will listen on internally
	InternalPorts []int32 `json:"internalPorts"`
	// Name is the container image name/URL to use for the container
	Name string `json:"name"`
}

type _RegistrySvcImageSpec RegistrySvcImageSpec

// NewRegistrySvcImageSpec instantiates a new RegistrySvcImageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcImageSpec(internalPorts []int32, name string) *RegistrySvcImageSpec {
	this := RegistrySvcImageSpec{}
	this.InternalPorts = internalPorts
	this.Name = name
	return &this
}

// NewRegistrySvcImageSpecWithDefaults instantiates a new RegistrySvcImageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcImageSpecWithDefaults() *RegistrySvcImageSpec {
	this := RegistrySvcImageSpec{}
	return &this
}

// GetInternalPorts returns the InternalPorts field value
func (o *RegistrySvcImageSpec) GetInternalPorts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.InternalPorts
}

// GetInternalPortsOk returns a tuple with the InternalPorts field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcImageSpec) GetInternalPortsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternalPorts, true
}

// SetInternalPorts sets field value
func (o *RegistrySvcImageSpec) SetInternalPorts(v []int32) {
	o.InternalPorts = v
}

// GetName returns the Name field value
func (o *RegistrySvcImageSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcImageSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RegistrySvcImageSpec) SetName(v string) {
	o.Name = v
}

func (o RegistrySvcImageSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcImageSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["internalPorts"] = o.InternalPorts
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *RegistrySvcImageSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"internalPorts",
		"name",
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

	varRegistrySvcImageSpec := _RegistrySvcImageSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrySvcImageSpec)

	if err != nil {
		return err
	}

	*o = RegistrySvcImageSpec(varRegistrySvcImageSpec)

	return err
}

type NullableRegistrySvcImageSpec struct {
	value *RegistrySvcImageSpec
	isSet bool
}

func (v NullableRegistrySvcImageSpec) Get() *RegistrySvcImageSpec {
	return v.value
}

func (v *NullableRegistrySvcImageSpec) Set(val *RegistrySvcImageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcImageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcImageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcImageSpec(val *RegistrySvcImageSpec) *NullableRegistrySvcImageSpec {
	return &NullableRegistrySvcImageSpec{value: val, isSet: true}
}

func (v NullableRegistrySvcImageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcImageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


