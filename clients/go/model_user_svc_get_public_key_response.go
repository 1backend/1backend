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

// checks if the UserSvcGetPublicKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcGetPublicKeyResponse{}

// UserSvcGetPublicKeyResponse struct for UserSvcGetPublicKeyResponse
type UserSvcGetPublicKeyResponse struct {
	PublicKey string `json:"publicKey"`
}

type _UserSvcGetPublicKeyResponse UserSvcGetPublicKeyResponse

// NewUserSvcGetPublicKeyResponse instantiates a new UserSvcGetPublicKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcGetPublicKeyResponse(publicKey string) *UserSvcGetPublicKeyResponse {
	this := UserSvcGetPublicKeyResponse{}
	this.PublicKey = publicKey
	return &this
}

// NewUserSvcGetPublicKeyResponseWithDefaults instantiates a new UserSvcGetPublicKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcGetPublicKeyResponseWithDefaults() *UserSvcGetPublicKeyResponse {
	this := UserSvcGetPublicKeyResponse{}
	return &this
}

// GetPublicKey returns the PublicKey field value
func (o *UserSvcGetPublicKeyResponse) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *UserSvcGetPublicKeyResponse) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *UserSvcGetPublicKeyResponse) SetPublicKey(v string) {
	o.PublicKey = v
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
	toSerialize["publicKey"] = o.PublicKey
	return toSerialize, nil
}

func (o *UserSvcGetPublicKeyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"publicKey",
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

	varUserSvcGetPublicKeyResponse := _UserSvcGetPublicKeyResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcGetPublicKeyResponse)

	if err != nil {
		return err
	}

	*o = UserSvcGetPublicKeyResponse(varUserSvcGetPublicKeyResponse)

	return err
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


