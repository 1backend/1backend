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

// checks if the ContainerSvcContainerIsRunningResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcContainerIsRunningResponse{}

// ContainerSvcContainerIsRunningResponse struct for ContainerSvcContainerIsRunningResponse
type ContainerSvcContainerIsRunningResponse struct {
	IsRunning bool `json:"isRunning"`
}

type _ContainerSvcContainerIsRunningResponse ContainerSvcContainerIsRunningResponse

// NewContainerSvcContainerIsRunningResponse instantiates a new ContainerSvcContainerIsRunningResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcContainerIsRunningResponse(isRunning bool) *ContainerSvcContainerIsRunningResponse {
	this := ContainerSvcContainerIsRunningResponse{}
	this.IsRunning = isRunning
	return &this
}

// NewContainerSvcContainerIsRunningResponseWithDefaults instantiates a new ContainerSvcContainerIsRunningResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcContainerIsRunningResponseWithDefaults() *ContainerSvcContainerIsRunningResponse {
	this := ContainerSvcContainerIsRunningResponse{}
	return &this
}

// GetIsRunning returns the IsRunning field value
func (o *ContainerSvcContainerIsRunningResponse) GetIsRunning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRunning
}

// GetIsRunningOk returns a tuple with the IsRunning field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainerIsRunningResponse) GetIsRunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRunning, true
}

// SetIsRunning sets field value
func (o *ContainerSvcContainerIsRunningResponse) SetIsRunning(v bool) {
	o.IsRunning = v
}

func (o ContainerSvcContainerIsRunningResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcContainerIsRunningResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isRunning"] = o.IsRunning
	return toSerialize, nil
}

func (o *ContainerSvcContainerIsRunningResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isRunning",
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

	varContainerSvcContainerIsRunningResponse := _ContainerSvcContainerIsRunningResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerSvcContainerIsRunningResponse)

	if err != nil {
		return err
	}

	*o = ContainerSvcContainerIsRunningResponse(varContainerSvcContainerIsRunningResponse)

	return err
}

type NullableContainerSvcContainerIsRunningResponse struct {
	value *ContainerSvcContainerIsRunningResponse
	isSet bool
}

func (v NullableContainerSvcContainerIsRunningResponse) Get() *ContainerSvcContainerIsRunningResponse {
	return v.value
}

func (v *NullableContainerSvcContainerIsRunningResponse) Set(val *ContainerSvcContainerIsRunningResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcContainerIsRunningResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcContainerIsRunningResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcContainerIsRunningResponse(val *ContainerSvcContainerIsRunningResponse) *NullableContainerSvcContainerIsRunningResponse {
	return &NullableContainerSvcContainerIsRunningResponse{value: val, isSet: true}
}

func (v NullableContainerSvcContainerIsRunningResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcContainerIsRunningResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


