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

// checks if the SecretSvcDecryptValueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretSvcDecryptValueResponse{}

// SecretSvcDecryptValueResponse struct for SecretSvcDecryptValueResponse
type SecretSvcDecryptValueResponse struct {
	Value *string `json:"value,omitempty"`
	Values []string `json:"values,omitempty"`
}

// NewSecretSvcDecryptValueResponse instantiates a new SecretSvcDecryptValueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretSvcDecryptValueResponse() *SecretSvcDecryptValueResponse {
	this := SecretSvcDecryptValueResponse{}
	return &this
}

// NewSecretSvcDecryptValueResponseWithDefaults instantiates a new SecretSvcDecryptValueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretSvcDecryptValueResponseWithDefaults() *SecretSvcDecryptValueResponse {
	this := SecretSvcDecryptValueResponse{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SecretSvcDecryptValueResponse) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcDecryptValueResponse) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SecretSvcDecryptValueResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SecretSvcDecryptValueResponse) SetValue(v string) {
	o.Value = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SecretSvcDecryptValueResponse) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcDecryptValueResponse) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SecretSvcDecryptValueResponse) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *SecretSvcDecryptValueResponse) SetValues(v []string) {
	o.Values = v
}

func (o SecretSvcDecryptValueResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretSvcDecryptValueResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableSecretSvcDecryptValueResponse struct {
	value *SecretSvcDecryptValueResponse
	isSet bool
}

func (v NullableSecretSvcDecryptValueResponse) Get() *SecretSvcDecryptValueResponse {
	return v.value
}

func (v *NullableSecretSvcDecryptValueResponse) Set(val *SecretSvcDecryptValueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretSvcDecryptValueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretSvcDecryptValueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretSvcDecryptValueResponse(val *SecretSvcDecryptValueResponse) *NullableSecretSvcDecryptValueResponse {
	return &NullableSecretSvcDecryptValueResponse{value: val, isSet: true}
}

func (v NullableSecretSvcDecryptValueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretSvcDecryptValueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


