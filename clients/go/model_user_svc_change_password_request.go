/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.35
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserSvcChangePasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcChangePasswordRequest{}

// UserSvcChangePasswordRequest struct for UserSvcChangePasswordRequest
type UserSvcChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword string `json:"newPassword"`
}

type _UserSvcChangePasswordRequest UserSvcChangePasswordRequest

// NewUserSvcChangePasswordRequest instantiates a new UserSvcChangePasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcChangePasswordRequest(currentPassword string, newPassword string) *UserSvcChangePasswordRequest {
	this := UserSvcChangePasswordRequest{}
	this.CurrentPassword = currentPassword
	this.NewPassword = newPassword
	return &this
}

// NewUserSvcChangePasswordRequestWithDefaults instantiates a new UserSvcChangePasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcChangePasswordRequestWithDefaults() *UserSvcChangePasswordRequest {
	this := UserSvcChangePasswordRequest{}
	return &this
}

// GetCurrentPassword returns the CurrentPassword field value
func (o *UserSvcChangePasswordRequest) GetCurrentPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value
// and a boolean to check if the value has been set.
func (o *UserSvcChangePasswordRequest) GetCurrentPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPassword, true
}

// SetCurrentPassword sets field value
func (o *UserSvcChangePasswordRequest) SetCurrentPassword(v string) {
	o.CurrentPassword = v
}

// GetNewPassword returns the NewPassword field value
func (o *UserSvcChangePasswordRequest) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *UserSvcChangePasswordRequest) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *UserSvcChangePasswordRequest) SetNewPassword(v string) {
	o.NewPassword = v
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
	toSerialize["currentPassword"] = o.CurrentPassword
	toSerialize["newPassword"] = o.NewPassword
	return toSerialize, nil
}

func (o *UserSvcChangePasswordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currentPassword",
		"newPassword",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserSvcChangePasswordRequest := _UserSvcChangePasswordRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcChangePasswordRequest)

	if err != nil {
		return err
	}

	*o = UserSvcChangePasswordRequest(varUserSvcChangePasswordRequest)

	return err
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


