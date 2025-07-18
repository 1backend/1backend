/*
1Backend

AI-native microservices platform.

API version: 0.7.6
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcLoginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcLoginResponse{}

// UserSvcLoginResponse struct for UserSvcLoginResponse
type UserSvcLoginResponse struct {
	Token *UserSvcAuthToken `json:"token,omitempty"`
}

// NewUserSvcLoginResponse instantiates a new UserSvcLoginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcLoginResponse() *UserSvcLoginResponse {
	this := UserSvcLoginResponse{}
	return &this
}

// NewUserSvcLoginResponseWithDefaults instantiates a new UserSvcLoginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcLoginResponseWithDefaults() *UserSvcLoginResponse {
	this := UserSvcLoginResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UserSvcLoginResponse) GetToken() UserSvcAuthToken {
	if o == nil || IsNil(o.Token) {
		var ret UserSvcAuthToken
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcLoginResponse) GetTokenOk() (*UserSvcAuthToken, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UserSvcLoginResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given UserSvcAuthToken and assigns it to the Token field.
func (o *UserSvcLoginResponse) SetToken(v UserSvcAuthToken) {
	o.Token = &v
}

func (o UserSvcLoginResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcLoginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableUserSvcLoginResponse struct {
	value *UserSvcLoginResponse
	isSet bool
}

func (v NullableUserSvcLoginResponse) Get() *UserSvcLoginResponse {
	return v.value
}

func (v *NullableUserSvcLoginResponse) Set(val *UserSvcLoginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcLoginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcLoginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcLoginResponse(val *UserSvcLoginResponse) *NullableUserSvcLoginResponse {
	return &NullableUserSvcLoginResponse{value: val, isSet: true}
}

func (v NullableUserSvcLoginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcLoginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


