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

// checks if the UserSvcChangePasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcChangePasswordRequest{}

// UserSvcChangePasswordRequest struct for UserSvcChangePasswordRequest
type UserSvcChangePasswordRequest struct {
	CurrentPassword *string `json:"currentPassword,omitempty"`
	NewPassword *string `json:"newPassword,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

// NewUserSvcChangePasswordRequest instantiates a new UserSvcChangePasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcChangePasswordRequest() *UserSvcChangePasswordRequest {
	this := UserSvcChangePasswordRequest{}
	return &this
}

// NewUserSvcChangePasswordRequestWithDefaults instantiates a new UserSvcChangePasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcChangePasswordRequestWithDefaults() *UserSvcChangePasswordRequest {
	this := UserSvcChangePasswordRequest{}
	return &this
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise.
func (o *UserSvcChangePasswordRequest) GetCurrentPassword() string {
	if o == nil || IsNil(o.CurrentPassword) {
		var ret string
		return ret
	}
	return *o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcChangePasswordRequest) GetCurrentPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentPassword) {
		return nil, false
	}
	return o.CurrentPassword, true
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *UserSvcChangePasswordRequest) HasCurrentPassword() bool {
	if o != nil && !IsNil(o.CurrentPassword) {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given string and assigns it to the CurrentPassword field.
func (o *UserSvcChangePasswordRequest) SetCurrentPassword(v string) {
	o.CurrentPassword = &v
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *UserSvcChangePasswordRequest) GetNewPassword() string {
	if o == nil || IsNil(o.NewPassword) {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcChangePasswordRequest) GetNewPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.NewPassword) {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *UserSvcChangePasswordRequest) HasNewPassword() bool {
	if o != nil && !IsNil(o.NewPassword) {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *UserSvcChangePasswordRequest) SetNewPassword(v string) {
	o.NewPassword = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UserSvcChangePasswordRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcChangePasswordRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UserSvcChangePasswordRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UserSvcChangePasswordRequest) SetSlug(v string) {
	o.Slug = &v
}

func (o UserSvcChangePasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcChangePasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentPassword) {
		toSerialize["currentPassword"] = o.CurrentPassword
	}
	if !IsNil(o.NewPassword) {
		toSerialize["newPassword"] = o.NewPassword
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	return toSerialize, nil
}

type NullableUserSvcChangePasswordRequest struct {
	value *UserSvcChangePasswordRequest
	isSet bool
}

func (v NullableUserSvcChangePasswordRequest) Get() *UserSvcChangePasswordRequest {
	return v.value
}

func (v *NullableUserSvcChangePasswordRequest) Set(val *UserSvcChangePasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcChangePasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcChangePasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcChangePasswordRequest(val *UserSvcChangePasswordRequest) *NullableUserSvcChangePasswordRequest {
	return &NullableUserSvcChangePasswordRequest{value: val, isSet: true}
}

func (v NullableUserSvcChangePasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcChangePasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


