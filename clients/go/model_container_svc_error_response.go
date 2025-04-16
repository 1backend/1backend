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

// checks if the ContainerSvcErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcErrorResponse{}

// ContainerSvcErrorResponse struct for ContainerSvcErrorResponse
type ContainerSvcErrorResponse struct {
	Error *string `json:"error,omitempty"`
}

// NewContainerSvcErrorResponse instantiates a new ContainerSvcErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcErrorResponse() *ContainerSvcErrorResponse {
	this := ContainerSvcErrorResponse{}
	return &this
}

// NewContainerSvcErrorResponseWithDefaults instantiates a new ContainerSvcErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcErrorResponseWithDefaults() *ContainerSvcErrorResponse {
	this := ContainerSvcErrorResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ContainerSvcErrorResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcErrorResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ContainerSvcErrorResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ContainerSvcErrorResponse) SetError(v string) {
	o.Error = &v
}

func (o ContainerSvcErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableContainerSvcErrorResponse struct {
	value *ContainerSvcErrorResponse
	isSet bool
}

func (v NullableContainerSvcErrorResponse) Get() *ContainerSvcErrorResponse {
	return v.value
}

func (v *NullableContainerSvcErrorResponse) Set(val *ContainerSvcErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcErrorResponse(val *ContainerSvcErrorResponse) *NullableContainerSvcErrorResponse {
	return &NullableContainerSvcErrorResponse{value: val, isSet: true}
}

func (v NullableContainerSvcErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


