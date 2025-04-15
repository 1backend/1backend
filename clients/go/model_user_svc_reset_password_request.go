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

// checks if the UserSvcResetPasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcResetPasswordRequest{}

// UserSvcResetPasswordRequest struct for UserSvcResetPasswordRequest
type UserSvcResetPasswordRequest struct {
	NewPassword string `json:"newPassword"`
}

type _UserSvcResetPasswordRequest UserSvcResetPasswordRequest

// NewUserSvcResetPasswordRequest instantiates a new UserSvcResetPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcResetPasswordRequest(newPassword string) *UserSvcResetPasswordRequest {
	this := UserSvcResetPasswordRequest{}
	this.NewPassword = newPassword
	return &this
}

// NewUserSvcResetPasswordRequestWithDefaults instantiates a new UserSvcResetPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcResetPasswordRequestWithDefaults() *UserSvcResetPasswordRequest {
	this := UserSvcResetPasswordRequest{}
	return &this
}

// GetNewPassword returns the NewPassword field value
func (o *UserSvcResetPasswordRequest) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *UserSvcResetPasswordRequest) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *UserSvcResetPasswordRequest) SetNewPassword(v string) {
	o.NewPassword = v
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
	toSerialize["newPassword"] = o.NewPassword
	return toSerialize, nil
}

func (o *UserSvcResetPasswordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varUserSvcResetPasswordRequest := _UserSvcResetPasswordRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcResetPasswordRequest)

	if err != nil {
		return err
	}

	*o = UserSvcResetPasswordRequest(varUserSvcResetPasswordRequest)

	return err
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


