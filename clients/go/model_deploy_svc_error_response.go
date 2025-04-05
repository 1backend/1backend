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

// checks if the DeploySvcErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploySvcErrorResponse{}

// DeploySvcErrorResponse struct for DeploySvcErrorResponse
type DeploySvcErrorResponse struct {
	Error *string `json:"error,omitempty"`
}

// NewDeploySvcErrorResponse instantiates a new DeploySvcErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploySvcErrorResponse() *DeploySvcErrorResponse {
	this := DeploySvcErrorResponse{}
	return &this
}

// NewDeploySvcErrorResponseWithDefaults instantiates a new DeploySvcErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploySvcErrorResponseWithDefaults() *DeploySvcErrorResponse {
	this := DeploySvcErrorResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DeploySvcErrorResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcErrorResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DeploySvcErrorResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *DeploySvcErrorResponse) SetError(v string) {
	o.Error = &v
}

func (o DeploySvcErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploySvcErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableDeploySvcErrorResponse struct {
	value *DeploySvcErrorResponse
	isSet bool
}

func (v NullableDeploySvcErrorResponse) Get() *DeploySvcErrorResponse {
	return v.value
}

func (v *NullableDeploySvcErrorResponse) Set(val *DeploySvcErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploySvcErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySvcErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySvcErrorResponse(val *DeploySvcErrorResponse) *NullableDeploySvcErrorResponse {
	return &NullableDeploySvcErrorResponse{value: val, isSet: true}
}

func (v NullableDeploySvcErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySvcErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


