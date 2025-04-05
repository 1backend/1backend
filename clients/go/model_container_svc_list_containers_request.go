/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.33
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ContainerSvcListContainersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcListContainersRequest{}

// ContainerSvcListContainersRequest struct for ContainerSvcListContainersRequest
type ContainerSvcListContainersRequest struct {
	ContainerId *string `json:"containerId,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	NodeId *string `json:"nodeId,omitempty"`
}

// NewContainerSvcListContainersRequest instantiates a new ContainerSvcListContainersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcListContainersRequest() *ContainerSvcListContainersRequest {
	this := ContainerSvcListContainersRequest{}
	return &this
}

// NewContainerSvcListContainersRequestWithDefaults instantiates a new ContainerSvcListContainersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcListContainersRequestWithDefaults() *ContainerSvcListContainersRequest {
	this := ContainerSvcListContainersRequest{}
	return &this
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *ContainerSvcListContainersRequest) GetContainerId() string {
	if o == nil || IsNil(o.ContainerId) {
		var ret string
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcListContainersRequest) GetContainerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerId) {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *ContainerSvcListContainersRequest) HasContainerId() bool {
	if o != nil && !IsNil(o.ContainerId) {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given string and assigns it to the ContainerId field.
func (o *ContainerSvcListContainersRequest) SetContainerId(v string) {
	o.ContainerId = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ContainerSvcListContainersRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcListContainersRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ContainerSvcListContainersRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ContainerSvcListContainersRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ContainerSvcListContainersRequest) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcListContainersRequest) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ContainerSvcListContainersRequest) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *ContainerSvcListContainersRequest) SetNodeId(v string) {
	o.NodeId = &v
}

func (o ContainerSvcListContainersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcListContainersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerId) {
		toSerialize["containerId"] = o.ContainerId
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.NodeId) {
		toSerialize["nodeId"] = o.NodeId
	}
	return toSerialize, nil
}

type NullableContainerSvcListContainersRequest struct {
	value *ContainerSvcListContainersRequest
	isSet bool
}

func (v NullableContainerSvcListContainersRequest) Get() *ContainerSvcListContainersRequest {
	return v.value
}

func (v *NullableContainerSvcListContainersRequest) Set(val *ContainerSvcListContainersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcListContainersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcListContainersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcListContainersRequest(val *ContainerSvcListContainersRequest) *NullableContainerSvcListContainersRequest {
	return &NullableContainerSvcListContainersRequest{value: val, isSet: true}
}

func (v NullableContainerSvcListContainersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcListContainersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


