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

// checks if the ContainerSvcPortMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcPortMapping{}

// ContainerSvcPortMapping struct for ContainerSvcPortMapping
type ContainerSvcPortMapping struct {
	Host int32 `json:"host"`
	Internal int32 `json:"internal"`
}

type _ContainerSvcPortMapping ContainerSvcPortMapping

// NewContainerSvcPortMapping instantiates a new ContainerSvcPortMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcPortMapping(host int32, internal int32) *ContainerSvcPortMapping {
	this := ContainerSvcPortMapping{}
	this.Host = host
	this.Internal = internal
	return &this
}

// NewContainerSvcPortMappingWithDefaults instantiates a new ContainerSvcPortMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcPortMappingWithDefaults() *ContainerSvcPortMapping {
	this := ContainerSvcPortMapping{}
	return &this
}

// GetHost returns the Host field value
func (o *ContainerSvcPortMapping) GetHost() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcPortMapping) GetHostOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *ContainerSvcPortMapping) SetHost(v int32) {
	o.Host = v
}

// GetInternal returns the Internal field value
func (o *ContainerSvcPortMapping) GetInternal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcPortMapping) GetInternalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Internal, true
}

// SetInternal sets field value
func (o *ContainerSvcPortMapping) SetInternal(v int32) {
	o.Internal = v
}

func (o ContainerSvcPortMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcPortMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["internal"] = o.Internal
	return toSerialize, nil
}

func (o *ContainerSvcPortMapping) UnmarshalJSON(data []byte) (err error) {
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

	varContainerSvcPortMapping := _ContainerSvcPortMapping{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerSvcPortMapping)

	if err != nil {
		return err
	}

	*o = ContainerSvcPortMapping(varContainerSvcPortMapping)

	return err
}

type NullableContainerSvcPortMapping struct {
	value *ContainerSvcPortMapping
	isSet bool
}

func (v NullableContainerSvcPortMapping) Get() *ContainerSvcPortMapping {
	return v.value
}

func (v *NullableContainerSvcPortMapping) Set(val *ContainerSvcPortMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcPortMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcPortMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcPortMapping(val *ContainerSvcPortMapping) *NullableContainerSvcPortMapping {
	return &NullableContainerSvcPortMapping{value: val, isSet: true}
}

func (v NullableContainerSvcPortMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcPortMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


