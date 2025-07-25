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

// checks if the ConfigSvcListConfigsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigSvcListConfigsRequest{}

// ConfigSvcListConfigsRequest struct for ConfigSvcListConfigsRequest
type ConfigSvcListConfigsRequest struct {
	App *string `json:"app,omitempty"`
	// Keys are camelCased slugs of the config owners.
	Keys []string `json:"keys,omitempty"`
}

// NewConfigSvcListConfigsRequest instantiates a new ConfigSvcListConfigsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigSvcListConfigsRequest() *ConfigSvcListConfigsRequest {
	this := ConfigSvcListConfigsRequest{}
	return &this
}

// NewConfigSvcListConfigsRequestWithDefaults instantiates a new ConfigSvcListConfigsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigSvcListConfigsRequestWithDefaults() *ConfigSvcListConfigsRequest {
	this := ConfigSvcListConfigsRequest{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *ConfigSvcListConfigsRequest) GetApp() string {
	if o == nil || IsNil(o.App) {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSvcListConfigsRequest) GetAppOk() (*string, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *ConfigSvcListConfigsRequest) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *ConfigSvcListConfigsRequest) SetApp(v string) {
	o.App = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *ConfigSvcListConfigsRequest) GetKeys() []string {
	if o == nil || IsNil(o.Keys) {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSvcListConfigsRequest) GetKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.Keys) {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *ConfigSvcListConfigsRequest) HasKeys() bool {
	if o != nil && !IsNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *ConfigSvcListConfigsRequest) SetKeys(v []string) {
	o.Keys = v
}

func (o ConfigSvcListConfigsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigSvcListConfigsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Keys) {
		toSerialize["keys"] = o.Keys
	}
	return toSerialize, nil
}

type NullableConfigSvcListConfigsRequest struct {
	value *ConfigSvcListConfigsRequest
	isSet bool
}

func (v NullableConfigSvcListConfigsRequest) Get() *ConfigSvcListConfigsRequest {
	return v.value
}

func (v *NullableConfigSvcListConfigsRequest) Set(val *ConfigSvcListConfigsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigSvcListConfigsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigSvcListConfigsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigSvcListConfigsRequest(val *ConfigSvcListConfigsRequest) *NullableConfigSvcListConfigsRequest {
	return &NullableConfigSvcListConfigsRequest{value: val, isSet: true}
}

func (v NullableConfigSvcListConfigsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigSvcListConfigsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


