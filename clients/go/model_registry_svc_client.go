/*
1Backend

A unified backend development platform for microservices-based AI applications.

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

// checks if the RegistrySvcClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcClient{}

// RegistrySvcClient struct for RegistrySvcClient
type RegistrySvcClient struct {
	// Programming language.
	Language RegistrySvcLanguage `json:"language"`
	// The URL of the client.
	Url string `json:"url"`
}

type _RegistrySvcClient RegistrySvcClient

// NewRegistrySvcClient instantiates a new RegistrySvcClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcClient(language RegistrySvcLanguage, url string) *RegistrySvcClient {
	this := RegistrySvcClient{}
	this.Language = language
	this.Url = url
	return &this
}

// NewRegistrySvcClientWithDefaults instantiates a new RegistrySvcClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcClientWithDefaults() *RegistrySvcClient {
	this := RegistrySvcClient{}
	return &this
}

// GetLanguage returns the Language field value
func (o *RegistrySvcClient) GetLanguage() RegistrySvcLanguage {
	if o == nil {
		var ret RegistrySvcLanguage
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcClient) GetLanguageOk() (*RegistrySvcLanguage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *RegistrySvcClient) SetLanguage(v RegistrySvcLanguage) {
	o.Language = v
}

// GetUrl returns the Url field value
func (o *RegistrySvcClient) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcClient) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RegistrySvcClient) SetUrl(v string) {
	o.Url = v
}

func (o RegistrySvcClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["language"] = o.Language
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *RegistrySvcClient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"language",
		"url",
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

	varRegistrySvcClient := _RegistrySvcClient{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrySvcClient)

	if err != nil {
		return err
	}

	*o = RegistrySvcClient(varRegistrySvcClient)

	return err
}

type NullableRegistrySvcClient struct {
	value *RegistrySvcClient
	isSet bool
}

func (v NullableRegistrySvcClient) Get() *RegistrySvcClient {
	return v.value
}

func (v *NullableRegistrySvcClient) Set(val *RegistrySvcClient) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcClient) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcClient(val *RegistrySvcClient) *NullableRegistrySvcClient {
	return &NullableRegistrySvcClient{value: val, isSet: true}
}

func (v NullableRegistrySvcClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


