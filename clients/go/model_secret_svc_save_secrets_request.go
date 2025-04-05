/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.33
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SecretSvcSaveSecretsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretSvcSaveSecretsRequest{}

// SecretSvcSaveSecretsRequest struct for SecretSvcSaveSecretsRequest
type SecretSvcSaveSecretsRequest struct {
	Secrets []SecretSvcSecret `json:"secrets,omitempty"`
}

// NewSecretSvcSaveSecretsRequest instantiates a new SecretSvcSaveSecretsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretSvcSaveSecretsRequest() *SecretSvcSaveSecretsRequest {
	this := SecretSvcSaveSecretsRequest{}
	return &this
}

// NewSecretSvcSaveSecretsRequestWithDefaults instantiates a new SecretSvcSaveSecretsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretSvcSaveSecretsRequestWithDefaults() *SecretSvcSaveSecretsRequest {
	this := SecretSvcSaveSecretsRequest{}
	return &this
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *SecretSvcSaveSecretsRequest) GetSecrets() []SecretSvcSecret {
	if o == nil || IsNil(o.Secrets) {
		var ret []SecretSvcSecret
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSvcSaveSecretsRequest) GetSecretsOk() ([]SecretSvcSecret, bool) {
	if o == nil || IsNil(o.Secrets) {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *SecretSvcSaveSecretsRequest) HasSecrets() bool {
	if o != nil && !IsNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []SecretSvcSecret and assigns it to the Secrets field.
func (o *SecretSvcSaveSecretsRequest) SetSecrets(v []SecretSvcSecret) {
	o.Secrets = v
}

func (o SecretSvcSaveSecretsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretSvcSaveSecretsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	return toSerialize, nil
}

type NullableSecretSvcSaveSecretsRequest struct {
	value *SecretSvcSaveSecretsRequest
	isSet bool
}

func (v NullableSecretSvcSaveSecretsRequest) Get() *SecretSvcSaveSecretsRequest {
	return v.value
}

func (v *NullableSecretSvcSaveSecretsRequest) Set(val *SecretSvcSaveSecretsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretSvcSaveSecretsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretSvcSaveSecretsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretSvcSaveSecretsRequest(val *SecretSvcSaveSecretsRequest) *NullableSecretSvcSaveSecretsRequest {
	return &NullableSecretSvcSaveSecretsRequest{value: val, isSet: true}
}

func (v NullableSecretSvcSaveSecretsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretSvcSaveSecretsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


