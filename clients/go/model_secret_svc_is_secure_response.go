/*
OpenOrch

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SecretSvcIsSecureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretSvcIsSecureResponse{}

// SecretSvcIsSecureResponse struct for SecretSvcIsSecureResponse
type SecretSvcIsSecureResponse struct {
	IsSecure bool `json:"isSecure"`
}

type _SecretSvcIsSecureResponse SecretSvcIsSecureResponse

// NewSecretSvcIsSecureResponse instantiates a new SecretSvcIsSecureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretSvcIsSecureResponse(isSecure bool) *SecretSvcIsSecureResponse {
	this := SecretSvcIsSecureResponse{}
	this.IsSecure = isSecure
	return &this
}

// NewSecretSvcIsSecureResponseWithDefaults instantiates a new SecretSvcIsSecureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretSvcIsSecureResponseWithDefaults() *SecretSvcIsSecureResponse {
	this := SecretSvcIsSecureResponse{}
	return &this
}

// GetIsSecure returns the IsSecure field value
func (o *SecretSvcIsSecureResponse) GetIsSecure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSecure
}

// GetIsSecureOk returns a tuple with the IsSecure field value
// and a boolean to check if the value has been set.
func (o *SecretSvcIsSecureResponse) GetIsSecureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSecure, true
}

// SetIsSecure sets field value
func (o *SecretSvcIsSecureResponse) SetIsSecure(v bool) {
	o.IsSecure = v
}

func (o SecretSvcIsSecureResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretSvcIsSecureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isSecure"] = o.IsSecure
	return toSerialize, nil
}

func (o *SecretSvcIsSecureResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isSecure",
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

	varSecretSvcIsSecureResponse := _SecretSvcIsSecureResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSecretSvcIsSecureResponse)

	if err != nil {
		return err
	}

	*o = SecretSvcIsSecureResponse(varSecretSvcIsSecureResponse)

	return err
}

type NullableSecretSvcIsSecureResponse struct {
	value *SecretSvcIsSecureResponse
	isSet bool
}

func (v NullableSecretSvcIsSecureResponse) Get() *SecretSvcIsSecureResponse {
	return v.value
}

func (v *NullableSecretSvcIsSecureResponse) Set(val *SecretSvcIsSecureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretSvcIsSecureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretSvcIsSecureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretSvcIsSecureResponse(val *SecretSvcIsSecureResponse) *NullableSecretSvcIsSecureResponse {
	return &NullableSecretSvcIsSecureResponse{value: val, isSet: true}
}

func (v NullableSecretSvcIsSecureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretSvcIsSecureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


