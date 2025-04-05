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

// checks if the ConfigSvcGetConfigResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigSvcGetConfigResponse{}

// ConfigSvcGetConfigResponse struct for ConfigSvcGetConfigResponse
type ConfigSvcGetConfigResponse struct {
	Config *ConfigSvcConfig `json:"config,omitempty"`
}

// NewConfigSvcGetConfigResponse instantiates a new ConfigSvcGetConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigSvcGetConfigResponse() *ConfigSvcGetConfigResponse {
	this := ConfigSvcGetConfigResponse{}
	return &this
}

// NewConfigSvcGetConfigResponseWithDefaults instantiates a new ConfigSvcGetConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigSvcGetConfigResponseWithDefaults() *ConfigSvcGetConfigResponse {
	this := ConfigSvcGetConfigResponse{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ConfigSvcGetConfigResponse) GetConfig() ConfigSvcConfig {
	if o == nil || IsNil(o.Config) {
		var ret ConfigSvcConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSvcGetConfigResponse) GetConfigOk() (*ConfigSvcConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ConfigSvcGetConfigResponse) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ConfigSvcConfig and assigns it to the Config field.
func (o *ConfigSvcGetConfigResponse) SetConfig(v ConfigSvcConfig) {
	o.Config = &v
}

func (o ConfigSvcGetConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigSvcGetConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableConfigSvcGetConfigResponse struct {
	value *ConfigSvcGetConfigResponse
	isSet bool
}

func (v NullableConfigSvcGetConfigResponse) Get() *ConfigSvcGetConfigResponse {
	return v.value
}

func (v *NullableConfigSvcGetConfigResponse) Set(val *ConfigSvcGetConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigSvcGetConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigSvcGetConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigSvcGetConfigResponse(val *ConfigSvcGetConfigResponse) *NullableConfigSvcGetConfigResponse {
	return &NullableConfigSvcGetConfigResponse{value: val, isSet: true}
}

func (v NullableConfigSvcGetConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigSvcGetConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


