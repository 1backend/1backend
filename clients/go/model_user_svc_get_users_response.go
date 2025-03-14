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

// checks if the UserSvcGetUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcGetUsersResponse{}

// UserSvcGetUsersResponse struct for UserSvcGetUsersResponse
type UserSvcGetUsersResponse struct {
	After *string `json:"after,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Users []UserSvcUser `json:"users,omitempty"`
}

// NewUserSvcGetUsersResponse instantiates a new UserSvcGetUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcGetUsersResponse() *UserSvcGetUsersResponse {
	this := UserSvcGetUsersResponse{}
	return &this
}

// NewUserSvcGetUsersResponseWithDefaults instantiates a new UserSvcGetUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcGetUsersResponseWithDefaults() *UserSvcGetUsersResponse {
	this := UserSvcGetUsersResponse{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *UserSvcGetUsersResponse) GetAfter() string {
	if o == nil || IsNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcGetUsersResponse) GetAfterOk() (*string, bool) {
	if o == nil || IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *UserSvcGetUsersResponse) HasAfter() bool {
	if o != nil && !IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *UserSvcGetUsersResponse) SetAfter(v string) {
	o.After = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UserSvcGetUsersResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcGetUsersResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UserSvcGetUsersResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UserSvcGetUsersResponse) SetCount(v int32) {
	o.Count = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *UserSvcGetUsersResponse) GetUsers() []UserSvcUser {
	if o == nil || IsNil(o.Users) {
		var ret []UserSvcUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcGetUsersResponse) GetUsersOk() ([]UserSvcUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UserSvcGetUsersResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserSvcUser and assigns it to the Users field.
func (o *UserSvcGetUsersResponse) SetUsers(v []UserSvcUser) {
	o.Users = v
}

func (o UserSvcGetUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcGetUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.After) {
		toSerialize["after"] = o.After
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableUserSvcGetUsersResponse struct {
	value *UserSvcGetUsersResponse
	isSet bool
}

func (v NullableUserSvcGetUsersResponse) Get() *UserSvcGetUsersResponse {
	return v.value
}

func (v *NullableUserSvcGetUsersResponse) Set(val *UserSvcGetUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcGetUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcGetUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcGetUsersResponse(val *UserSvcGetUsersResponse) *NullableUserSvcGetUsersResponse {
	return &NullableUserSvcGetUsersResponse{value: val, isSet: true}
}

func (v NullableUserSvcGetUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcGetUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


