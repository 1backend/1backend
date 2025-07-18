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
	"bytes"
	"fmt"
)

// checks if the UserSvcRefreshTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcRefreshTokenResponse{}

// UserSvcRefreshTokenResponse struct for UserSvcRefreshTokenResponse
type UserSvcRefreshTokenResponse struct {
	Token UserSvcAuthToken `json:"token"`
}

type _UserSvcRefreshTokenResponse UserSvcRefreshTokenResponse

// NewUserSvcRefreshTokenResponse instantiates a new UserSvcRefreshTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcRefreshTokenResponse(token UserSvcAuthToken) *UserSvcRefreshTokenResponse {
	this := UserSvcRefreshTokenResponse{}
	this.Token = token
	return &this
}

// NewUserSvcRefreshTokenResponseWithDefaults instantiates a new UserSvcRefreshTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcRefreshTokenResponseWithDefaults() *UserSvcRefreshTokenResponse {
	this := UserSvcRefreshTokenResponse{}
	return &this
}

// GetToken returns the Token field value
func (o *UserSvcRefreshTokenResponse) GetToken() UserSvcAuthToken {
	if o == nil {
		var ret UserSvcAuthToken
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UserSvcRefreshTokenResponse) GetTokenOk() (*UserSvcAuthToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UserSvcRefreshTokenResponse) SetToken(v UserSvcAuthToken) {
	o.Token = v
}

func (o UserSvcRefreshTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcRefreshTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *UserSvcRefreshTokenResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
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

	varUserSvcRefreshTokenResponse := _UserSvcRefreshTokenResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcRefreshTokenResponse)

	if err != nil {
		return err
	}

	*o = UserSvcRefreshTokenResponse(varUserSvcRefreshTokenResponse)

	return err
}

type NullableUserSvcRefreshTokenResponse struct {
	value *UserSvcRefreshTokenResponse
	isSet bool
}

func (v NullableUserSvcRefreshTokenResponse) Get() *UserSvcRefreshTokenResponse {
	return v.value
}

func (v *NullableUserSvcRefreshTokenResponse) Set(val *UserSvcRefreshTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcRefreshTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcRefreshTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcRefreshTokenResponse(val *UserSvcRefreshTokenResponse) *NullableUserSvcRefreshTokenResponse {
	return &NullableUserSvcRefreshTokenResponse{value: val, isSet: true}
}

func (v NullableUserSvcRefreshTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcRefreshTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


