/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.34
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ConfigSvcSaveConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigSvcSaveConfigRequest{}

// ConfigSvcSaveConfigRequest struct for ConfigSvcSaveConfigRequest
type ConfigSvcSaveConfigRequest struct {
	Config *ConfigSvcConfig `json:"config,omitempty"`
}

// NewConfigSvcSaveConfigRequest instantiates a new ConfigSvcSaveConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigSvcSaveConfigRequest() *ConfigSvcSaveConfigRequest {
	this := ConfigSvcSaveConfigRequest{}
	return &this
}

// NewConfigSvcSaveConfigRequestWithDefaults instantiates a new ConfigSvcSaveConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigSvcSaveConfigRequestWithDefaults() *ConfigSvcSaveConfigRequest {
	this := ConfigSvcSaveConfigRequest{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ConfigSvcSaveConfigRequest) GetConfig() ConfigSvcConfig {
	if o == nil || IsNil(o.Config) {
		var ret ConfigSvcConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSvcSaveConfigRequest) GetConfigOk() (*ConfigSvcConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ConfigSvcSaveConfigRequest) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ConfigSvcConfig and assigns it to the Config field.
func (o *ConfigSvcSaveConfigRequest) SetConfig(v ConfigSvcConfig) {
	o.Config = &v
}

func (o ConfigSvcSaveConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigSvcSaveConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableConfigSvcSaveConfigRequest struct {
	value *ConfigSvcSaveConfigRequest
	isSet bool
}

func (v NullableConfigSvcSaveConfigRequest) Get() *ConfigSvcSaveConfigRequest {
	return v.value
}

func (v *NullableConfigSvcSaveConfigRequest) Set(val *ConfigSvcSaveConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigSvcSaveConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigSvcSaveConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigSvcSaveConfigRequest(val *ConfigSvcSaveConfigRequest) *NullableConfigSvcSaveConfigRequest {
	return &NullableConfigSvcSaveConfigRequest{value: val, isSet: true}
}

func (v NullableConfigSvcSaveConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigSvcSaveConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


