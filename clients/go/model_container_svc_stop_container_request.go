/*
1Backend

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ContainerSvcStopContainerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcStopContainerRequest{}

// ContainerSvcStopContainerRequest struct for ContainerSvcStopContainerRequest
type ContainerSvcStopContainerRequest struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewContainerSvcStopContainerRequest instantiates a new ContainerSvcStopContainerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcStopContainerRequest() *ContainerSvcStopContainerRequest {
	this := ContainerSvcStopContainerRequest{}
	return &this
}

// NewContainerSvcStopContainerRequestWithDefaults instantiates a new ContainerSvcStopContainerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcStopContainerRequestWithDefaults() *ContainerSvcStopContainerRequest {
	this := ContainerSvcStopContainerRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerSvcStopContainerRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcStopContainerRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerSvcStopContainerRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContainerSvcStopContainerRequest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerSvcStopContainerRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcStopContainerRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerSvcStopContainerRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerSvcStopContainerRequest) SetName(v string) {
	o.Name = &v
}

func (o ContainerSvcStopContainerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcStopContainerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableContainerSvcStopContainerRequest struct {
	value *ContainerSvcStopContainerRequest
	isSet bool
}

func (v NullableContainerSvcStopContainerRequest) Get() *ContainerSvcStopContainerRequest {
	return v.value
}

func (v *NullableContainerSvcStopContainerRequest) Set(val *ContainerSvcStopContainerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcStopContainerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcStopContainerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcStopContainerRequest(val *ContainerSvcStopContainerRequest) *NullableContainerSvcStopContainerRequest {
	return &NullableContainerSvcStopContainerRequest{value: val, isSet: true}
}

func (v NullableContainerSvcStopContainerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcStopContainerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


