/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.32
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcIsAuthorizedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcIsAuthorizedResponse{}

// UserSvcIsAuthorizedResponse struct for UserSvcIsAuthorizedResponse
type UserSvcIsAuthorizedResponse struct {
	Authorized *bool `json:"authorized,omitempty"`
	User *UserSvcUser `json:"user,omitempty"`
}

// NewUserSvcIsAuthorizedResponse instantiates a new UserSvcIsAuthorizedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcIsAuthorizedResponse() *UserSvcIsAuthorizedResponse {
	this := UserSvcIsAuthorizedResponse{}
	return &this
}

// NewUserSvcIsAuthorizedResponseWithDefaults instantiates a new UserSvcIsAuthorizedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcIsAuthorizedResponseWithDefaults() *UserSvcIsAuthorizedResponse {
	this := UserSvcIsAuthorizedResponse{}
	return &this
}

// GetAuthorized returns the Authorized field value if set, zero value otherwise.
func (o *UserSvcIsAuthorizedResponse) GetAuthorized() bool {
	if o == nil || IsNil(o.Authorized) {
		var ret bool
		return ret
	}
	return *o.Authorized
}

// GetAuthorizedOk returns a tuple with the Authorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcIsAuthorizedResponse) GetAuthorizedOk() (*bool, bool) {
	if o == nil || IsNil(o.Authorized) {
		return nil, false
	}
	return o.Authorized, true
}

// HasAuthorized returns a boolean if a field has been set.
func (o *UserSvcIsAuthorizedResponse) HasAuthorized() bool {
	if o != nil && !IsNil(o.Authorized) {
		return true
	}

	return false
}

// SetAuthorized gets a reference to the given bool and assigns it to the Authorized field.
func (o *UserSvcIsAuthorizedResponse) SetAuthorized(v bool) {
	o.Authorized = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserSvcIsAuthorizedResponse) GetUser() UserSvcUser {
	if o == nil || IsNil(o.User) {
		var ret UserSvcUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcIsAuthorizedResponse) GetUserOk() (*UserSvcUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserSvcIsAuthorizedResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserSvcUser and assigns it to the User field.
func (o *UserSvcIsAuthorizedResponse) SetUser(v UserSvcUser) {
	o.User = &v
}

func (o UserSvcIsAuthorizedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcIsAuthorizedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authorized) {
		toSerialize["authorized"] = o.Authorized
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUserSvcIsAuthorizedResponse struct {
	value *UserSvcIsAuthorizedResponse
	isSet bool
}

func (v NullableUserSvcIsAuthorizedResponse) Get() *UserSvcIsAuthorizedResponse {
	return v.value
}

func (v *NullableUserSvcIsAuthorizedResponse) Set(val *UserSvcIsAuthorizedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcIsAuthorizedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcIsAuthorizedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcIsAuthorizedResponse(val *UserSvcIsAuthorizedResponse) *NullableUserSvcIsAuthorizedResponse {
	return &NullableUserSvcIsAuthorizedResponse{value: val, isSet: true}
}

func (v NullableUserSvcIsAuthorizedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcIsAuthorizedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


