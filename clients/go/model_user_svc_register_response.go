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

// checks if the UserSvcRegisterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcRegisterResponse{}

// UserSvcRegisterResponse struct for UserSvcRegisterResponse
type UserSvcRegisterResponse struct {
	Token *UserSvcAuthToken `json:"token,omitempty"`
}

// NewUserSvcRegisterResponse instantiates a new UserSvcRegisterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcRegisterResponse() *UserSvcRegisterResponse {
	this := UserSvcRegisterResponse{}
	return &this
}

// NewUserSvcRegisterResponseWithDefaults instantiates a new UserSvcRegisterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcRegisterResponseWithDefaults() *UserSvcRegisterResponse {
	this := UserSvcRegisterResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UserSvcRegisterResponse) GetToken() UserSvcAuthToken {
	if o == nil || IsNil(o.Token) {
		var ret UserSvcAuthToken
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcRegisterResponse) GetTokenOk() (*UserSvcAuthToken, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UserSvcRegisterResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given UserSvcAuthToken and assigns it to the Token field.
func (o *UserSvcRegisterResponse) SetToken(v UserSvcAuthToken) {
	o.Token = &v
}

func (o UserSvcRegisterResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcRegisterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableUserSvcRegisterResponse struct {
	value *UserSvcRegisterResponse
	isSet bool
}

func (v NullableUserSvcRegisterResponse) Get() *UserSvcRegisterResponse {
	return v.value
}

func (v *NullableUserSvcRegisterResponse) Set(val *UserSvcRegisterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcRegisterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcRegisterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcRegisterResponse(val *UserSvcRegisterResponse) *NullableUserSvcRegisterResponse {
	return &NullableUserSvcRegisterResponse{value: val, isSet: true}
}

func (v NullableUserSvcRegisterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcRegisterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


