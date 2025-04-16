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
)

// checks if the ContainerSvcListContainersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcListContainersResponse{}

// ContainerSvcListContainersResponse struct for ContainerSvcListContainersResponse
type ContainerSvcListContainersResponse struct {
	Containers []ContainerSvcContainer `json:"containers,omitempty"`
}

// NewContainerSvcListContainersResponse instantiates a new ContainerSvcListContainersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcListContainersResponse() *ContainerSvcListContainersResponse {
	this := ContainerSvcListContainersResponse{}
	return &this
}

// NewContainerSvcListContainersResponseWithDefaults instantiates a new ContainerSvcListContainersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcListContainersResponseWithDefaults() *ContainerSvcListContainersResponse {
	this := ContainerSvcListContainersResponse{}
	return &this
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *ContainerSvcListContainersResponse) GetContainers() []ContainerSvcContainer {
	if o == nil || IsNil(o.Containers) {
		var ret []ContainerSvcContainer
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcListContainersResponse) GetContainersOk() ([]ContainerSvcContainer, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *ContainerSvcListContainersResponse) HasContainers() bool {
	if o != nil && !IsNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []ContainerSvcContainer and assigns it to the Containers field.
func (o *ContainerSvcListContainersResponse) SetContainers(v []ContainerSvcContainer) {
	o.Containers = v
}

func (o ContainerSvcListContainersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcListContainersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	return toSerialize, nil
}

type NullableContainerSvcListContainersResponse struct {
	value *ContainerSvcListContainersResponse
	isSet bool
}

func (v NullableContainerSvcListContainersResponse) Get() *ContainerSvcListContainersResponse {
	return v.value
}

func (v *NullableContainerSvcListContainersResponse) Set(val *ContainerSvcListContainersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcListContainersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcListContainersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcListContainersResponse(val *ContainerSvcListContainersResponse) *NullableContainerSvcListContainersResponse {
	return &NullableContainerSvcListContainersResponse{value: val, isSet: true}
}

func (v NullableContainerSvcListContainersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcListContainersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


