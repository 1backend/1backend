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

// checks if the UserSvcGetPublicKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcGetPublicKeyResponse{}

// UserSvcGetPublicKeyResponse struct for UserSvcGetPublicKeyResponse
type UserSvcGetPublicKeyResponse struct {
	PublicKey *string `json:"publicKey,omitempty"`
}

// NewUserSvcGetPublicKeyResponse instantiates a new UserSvcGetPublicKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcGetPublicKeyResponse() *UserSvcGetPublicKeyResponse {
	this := UserSvcGetPublicKeyResponse{}
	return &this
}

// NewUserSvcGetPublicKeyResponseWithDefaults instantiates a new UserSvcGetPublicKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcGetPublicKeyResponseWithDefaults() *UserSvcGetPublicKeyResponse {
	this := UserSvcGetPublicKeyResponse{}
	return &this
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *UserSvcGetPublicKeyResponse) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcGetPublicKeyResponse) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *UserSvcGetPublicKeyResponse) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *UserSvcGetPublicKeyResponse) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o UserSvcGetPublicKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcGetPublicKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	return toSerialize, nil
}

type NullableUserSvcGetPublicKeyResponse struct {
	value *UserSvcGetPublicKeyResponse
	isSet bool
}

func (v NullableUserSvcGetPublicKeyResponse) Get() *UserSvcGetPublicKeyResponse {
	return v.value
}

func (v *NullableUserSvcGetPublicKeyResponse) Set(val *UserSvcGetPublicKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcGetPublicKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcGetPublicKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcGetPublicKeyResponse(val *UserSvcGetPublicKeyResponse) *NullableUserSvcGetPublicKeyResponse {
	return &NullableUserSvcGetPublicKeyResponse{value: val, isSet: true}
}

func (v NullableUserSvcGetPublicKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcGetPublicKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


