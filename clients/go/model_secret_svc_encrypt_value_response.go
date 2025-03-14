/*
1Backend

A common backend for your AI applications—microservices-based and built to scale.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SecretSvcEncryptValueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretSvcEncryptValueResponse{}

// SecretSvcEncryptValueResponse struct for SecretSvcEncryptValueResponse
type SecretSvcEncryptValueResponse struct {
	Value *string `json:"value,omitempty"`
	Values []string `json:"values,omitempty"`
}

// NewSecretSvcEncryptValueResponse instantiates a new SecretSvcEncryptValueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretSvcEncryptValueResponse() *SecretSvcEncryptValueResponse {
	this := SecretSvcEncryptValueResponse{}
	return &this
}

// NewSecretSvcEncryptValueResponseWithDefaults instantiates a new SecretSvcEncryptValueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretSvcEncryptValueResponseWithDefaults() *SecretSvcEncryptValueResponse {
	this := SecretSvcEncryptValueResponse{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SecretSvcEncryptValueResponse) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcEncryptValueResponse) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SecretSvcEncryptValueResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SecretSvcEncryptValueResponse) SetValue(v string) {
	o.Value = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SecretSvcEncryptValueResponse) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcEncryptValueResponse) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SecretSvcEncryptValueResponse) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *SecretSvcEncryptValueResponse) SetValues(v []string) {
	o.Values = v
}

func (o SecretSvcEncryptValueResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretSvcEncryptValueResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableSecretSvcEncryptValueResponse struct {
	value *SecretSvcEncryptValueResponse
	isSet bool
}

func (v NullableSecretSvcEncryptValueResponse) Get() *SecretSvcEncryptValueResponse {
	return v.value
}

func (v *NullableSecretSvcEncryptValueResponse) Set(val *SecretSvcEncryptValueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretSvcEncryptValueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretSvcEncryptValueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretSvcEncryptValueResponse(val *SecretSvcEncryptValueResponse) *NullableSecretSvcEncryptValueResponse {
	return &NullableSecretSvcEncryptValueResponse{value: val, isSet: true}
}

func (v NullableSecretSvcEncryptValueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretSvcEncryptValueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


