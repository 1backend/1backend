/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcErrorResponse{}

// UserSvcErrorResponse struct for UserSvcErrorResponse
type UserSvcErrorResponse struct {
	Error *string `json:"error,omitempty"`
}

// NewUserSvcErrorResponse instantiates a new UserSvcErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcErrorResponse() *UserSvcErrorResponse {
	this := UserSvcErrorResponse{}
	return &this
}

// NewUserSvcErrorResponseWithDefaults instantiates a new UserSvcErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcErrorResponseWithDefaults() *UserSvcErrorResponse {
	this := UserSvcErrorResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *UserSvcErrorResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcErrorResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *UserSvcErrorResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *UserSvcErrorResponse) SetError(v string) {
	o.Error = &v
}

func (o UserSvcErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableUserSvcErrorResponse struct {
	value *UserSvcErrorResponse
	isSet bool
}

func (v NullableUserSvcErrorResponse) Get() *UserSvcErrorResponse {
	return v.value
}

func (v *NullableUserSvcErrorResponse) Set(val *UserSvcErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcErrorResponse(val *UserSvcErrorResponse) *NullableUserSvcErrorResponse {
	return &NullableUserSvcErrorResponse{value: val, isSet: true}
}

func (v NullableUserSvcErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


