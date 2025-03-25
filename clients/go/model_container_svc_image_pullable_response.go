/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.30
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ContainerSvcImagePullableResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcImagePullableResponse{}

// ContainerSvcImagePullableResponse struct for ContainerSvcImagePullableResponse
type ContainerSvcImagePullableResponse struct {
	Pullable bool `json:"pullable"`
}

type _ContainerSvcImagePullableResponse ContainerSvcImagePullableResponse

// NewContainerSvcImagePullableResponse instantiates a new ContainerSvcImagePullableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcImagePullableResponse(pullable bool) *ContainerSvcImagePullableResponse {
	this := ContainerSvcImagePullableResponse{}
	this.Pullable = pullable
	return &this
}

// NewContainerSvcImagePullableResponseWithDefaults instantiates a new ContainerSvcImagePullableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcImagePullableResponseWithDefaults() *ContainerSvcImagePullableResponse {
	this := ContainerSvcImagePullableResponse{}
	return &this
}

// GetPullable returns the Pullable field value
func (o *ContainerSvcImagePullableResponse) GetPullable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pullable
}

// GetPullableOk returns a tuple with the Pullable field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcImagePullableResponse) GetPullableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pullable, true
}

// SetPullable sets field value
func (o *ContainerSvcImagePullableResponse) SetPullable(v bool) {
	o.Pullable = v
}

func (o ContainerSvcImagePullableResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcImagePullableResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pullable"] = o.Pullable
	return toSerialize, nil
}

func (o *ContainerSvcImagePullableResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pullable",
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

	varContainerSvcImagePullableResponse := _ContainerSvcImagePullableResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerSvcImagePullableResponse)

	if err != nil {
		return err
	}

	*o = ContainerSvcImagePullableResponse(varContainerSvcImagePullableResponse)

	return err
}

type NullableContainerSvcImagePullableResponse struct {
	value *ContainerSvcImagePullableResponse
	isSet bool
}

func (v NullableContainerSvcImagePullableResponse) Get() *ContainerSvcImagePullableResponse {
	return v.value
}

func (v *NullableContainerSvcImagePullableResponse) Set(val *ContainerSvcImagePullableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcImagePullableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcImagePullableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcImagePullableResponse(val *ContainerSvcImagePullableResponse) *NullableContainerSvcImagePullableResponse {
	return &NullableContainerSvcImagePullableResponse{value: val, isSet: true}
}

func (v NullableContainerSvcImagePullableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcImagePullableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


