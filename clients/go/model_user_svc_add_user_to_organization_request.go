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

// checks if the UserSvcAddUserToOrganizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcAddUserToOrganizationRequest{}

// UserSvcAddUserToOrganizationRequest struct for UserSvcAddUserToOrganizationRequest
type UserSvcAddUserToOrganizationRequest struct {
	UserId *string `json:"userId,omitempty"`
}

// NewUserSvcAddUserToOrganizationRequest instantiates a new UserSvcAddUserToOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcAddUserToOrganizationRequest() *UserSvcAddUserToOrganizationRequest {
	this := UserSvcAddUserToOrganizationRequest{}
	return &this
}

// NewUserSvcAddUserToOrganizationRequestWithDefaults instantiates a new UserSvcAddUserToOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcAddUserToOrganizationRequestWithDefaults() *UserSvcAddUserToOrganizationRequest {
	this := UserSvcAddUserToOrganizationRequest{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserSvcAddUserToOrganizationRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcAddUserToOrganizationRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserSvcAddUserToOrganizationRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UserSvcAddUserToOrganizationRequest) SetUserId(v string) {
	o.UserId = &v
}

func (o UserSvcAddUserToOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcAddUserToOrganizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUserSvcAddUserToOrganizationRequest struct {
	value *UserSvcAddUserToOrganizationRequest
	isSet bool
}

func (v NullableUserSvcAddUserToOrganizationRequest) Get() *UserSvcAddUserToOrganizationRequest {
	return v.value
}

func (v *NullableUserSvcAddUserToOrganizationRequest) Set(val *UserSvcAddUserToOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcAddUserToOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcAddUserToOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcAddUserToOrganizationRequest(val *UserSvcAddUserToOrganizationRequest) *NullableUserSvcAddUserToOrganizationRequest {
	return &NullableUserSvcAddUserToOrganizationRequest{value: val, isSet: true}
}

func (v NullableUserSvcAddUserToOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcAddUserToOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


