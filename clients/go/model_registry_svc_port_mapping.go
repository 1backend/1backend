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
	"bytes"
	"fmt"
)

// checks if the RegistrySvcPortMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcPortMapping{}

// RegistrySvcPortMapping struct for RegistrySvcPortMapping
type RegistrySvcPortMapping struct {
	Host int32 `json:"host"`
	Internal int32 `json:"internal"`
}

type _RegistrySvcPortMapping RegistrySvcPortMapping

// NewRegistrySvcPortMapping instantiates a new RegistrySvcPortMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcPortMapping(host int32, internal int32) *RegistrySvcPortMapping {
	this := RegistrySvcPortMapping{}
	this.Host = host
	this.Internal = internal
	return &this
}

// NewRegistrySvcPortMappingWithDefaults instantiates a new RegistrySvcPortMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcPortMappingWithDefaults() *RegistrySvcPortMapping {
	this := RegistrySvcPortMapping{}
	return &this
}

// GetHost returns the Host field value
func (o *RegistrySvcPortMapping) GetHost() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcPortMapping) GetHostOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *RegistrySvcPortMapping) SetHost(v int32) {
	o.Host = v
}

// GetInternal returns the Internal field value
func (o *RegistrySvcPortMapping) GetInternal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcPortMapping) GetInternalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Internal, true
}

// SetInternal sets field value
func (o *RegistrySvcPortMapping) SetInternal(v int32) {
	o.Internal = v
}

func (o RegistrySvcPortMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcPortMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["internal"] = o.Internal
	return toSerialize, nil
}

func (o *RegistrySvcPortMapping) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"internal",
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

	varRegistrySvcPortMapping := _RegistrySvcPortMapping{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrySvcPortMapping)

	if err != nil {
		return err
	}

	*o = RegistrySvcPortMapping(varRegistrySvcPortMapping)

	return err
}

type NullableRegistrySvcPortMapping struct {
	value *RegistrySvcPortMapping
	isSet bool
}

func (v NullableRegistrySvcPortMapping) Get() *RegistrySvcPortMapping {
	return v.value
}

func (v *NullableRegistrySvcPortMapping) Set(val *RegistrySvcPortMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcPortMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcPortMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcPortMapping(val *RegistrySvcPortMapping) *NullableRegistrySvcPortMapping {
	return &NullableRegistrySvcPortMapping{value: val, isSet: true}
}

func (v NullableRegistrySvcPortMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcPortMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


