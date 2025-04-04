/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.34
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcResetPasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcResetPasswordRequest{}

// UserSvcResetPasswordRequest struct for UserSvcResetPasswordRequest
type UserSvcResetPasswordRequest struct {
	NewPassword *string `json:"newPassword,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

// NewUserSvcResetPasswordRequest instantiates a new UserSvcResetPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcResetPasswordRequest() *UserSvcResetPasswordRequest {
	this := UserSvcResetPasswordRequest{}
	return &this
}

// NewUserSvcResetPasswordRequestWithDefaults instantiates a new UserSvcResetPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcResetPasswordRequestWithDefaults() *UserSvcResetPasswordRequest {
	this := UserSvcResetPasswordRequest{}
	return &this
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *UserSvcResetPasswordRequest) GetNewPassword() string {
	if o == nil || IsNil(o.NewPassword) {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcResetPasswordRequest) GetNewPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.NewPassword) {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *UserSvcResetPasswordRequest) HasNewPassword() bool {
	if o != nil && !IsNil(o.NewPassword) {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *UserSvcResetPasswordRequest) SetNewPassword(v string) {
	o.NewPassword = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UserSvcResetPasswordRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcResetPasswordRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UserSvcResetPasswordRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UserSvcResetPasswordRequest) SetSlug(v string) {
	o.Slug = &v
}

func (o UserSvcResetPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcResetPasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewPassword) {
		toSerialize["newPassword"] = o.NewPassword
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	return toSerialize, nil
}

type NullableUserSvcResetPasswordRequest struct {
	value *UserSvcResetPasswordRequest
	isSet bool
}

func (v NullableUserSvcResetPasswordRequest) Get() *UserSvcResetPasswordRequest {
	return v.value
}

func (v *NullableUserSvcResetPasswordRequest) Set(val *UserSvcResetPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcResetPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcResetPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcResetPasswordRequest(val *UserSvcResetPasswordRequest) *NullableUserSvcResetPasswordRequest {
	return &NullableUserSvcResetPasswordRequest{value: val, isSet: true}
}

func (v NullableUserSvcResetPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcResetPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


