/*
OpenOrch

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ContainerSvcRunContainerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcRunContainerResponse{}

// ContainerSvcRunContainerResponse struct for ContainerSvcRunContainerResponse
type ContainerSvcRunContainerResponse struct {
	// Ports is returned here as host ports might get mapped dynamically.
	Ports []ContainerSvcPortMapping `json:"ports,omitempty"`
	Started *bool `json:"started,omitempty"`
}

// NewContainerSvcRunContainerResponse instantiates a new ContainerSvcRunContainerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcRunContainerResponse() *ContainerSvcRunContainerResponse {
	this := ContainerSvcRunContainerResponse{}
	return &this
}

// NewContainerSvcRunContainerResponseWithDefaults instantiates a new ContainerSvcRunContainerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcRunContainerResponseWithDefaults() *ContainerSvcRunContainerResponse {
	this := ContainerSvcRunContainerResponse{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ContainerSvcRunContainerResponse) GetPorts() []ContainerSvcPortMapping {
	if o == nil || IsNil(o.Ports) {
		var ret []ContainerSvcPortMapping
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcRunContainerResponse) GetPortsOk() ([]ContainerSvcPortMapping, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ContainerSvcRunContainerResponse) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ContainerSvcPortMapping and assigns it to the Ports field.
func (o *ContainerSvcRunContainerResponse) SetPorts(v []ContainerSvcPortMapping) {
	o.Ports = v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *ContainerSvcRunContainerResponse) GetStarted() bool {
	if o == nil || IsNil(o.Started) {
		var ret bool
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcRunContainerResponse) GetStartedOk() (*bool, bool) {
	if o == nil || IsNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *ContainerSvcRunContainerResponse) HasStarted() bool {
	if o != nil && !IsNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given bool and assigns it to the Started field.
func (o *ContainerSvcRunContainerResponse) SetStarted(v bool) {
	o.Started = &v
}

func (o ContainerSvcRunContainerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcRunContainerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.Started) {
		toSerialize["started"] = o.Started
	}
	return toSerialize, nil
}

type NullableContainerSvcRunContainerResponse struct {
	value *ContainerSvcRunContainerResponse
	isSet bool
}

func (v NullableContainerSvcRunContainerResponse) Get() *ContainerSvcRunContainerResponse {
	return v.value
}

func (v *NullableContainerSvcRunContainerResponse) Set(val *ContainerSvcRunContainerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcRunContainerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcRunContainerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcRunContainerResponse(val *ContainerSvcRunContainerResponse) *NullableContainerSvcRunContainerResponse {
	return &NullableContainerSvcRunContainerResponse{value: val, isSet: true}
}

func (v NullableContainerSvcRunContainerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcRunContainerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


