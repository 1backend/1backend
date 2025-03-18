/*
1Backend

A unified backend development platform for your AI applications—microservices-based and built to scale.

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

// checks if the ConfigSvcConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigSvcConfig{}

// ConfigSvcConfig struct for ConfigSvcConfig
type ConfigSvcConfig struct {
	Data map[string]interface{} `json:"data"`
	DataJson *string `json:"dataJson,omitempty"`
	Id *string `json:"id,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

type _ConfigSvcConfig ConfigSvcConfig

// NewConfigSvcConfig instantiates a new ConfigSvcConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigSvcConfig(data map[string]interface{}) *ConfigSvcConfig {
	this := ConfigSvcConfig{}
	this.Data = data
	return &this
}

// NewConfigSvcConfigWithDefaults instantiates a new ConfigSvcConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigSvcConfigWithDefaults() *ConfigSvcConfig {
	this := ConfigSvcConfig{}
	return &this
}

// GetData returns the Data field value
func (o *ConfigSvcConfig) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ConfigSvcConfig) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ConfigSvcConfig) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetDataJson returns the DataJson field value if set, zero value otherwise.
func (o *ConfigSvcConfig) GetDataJson() string {
	if o == nil || IsNil(o.DataJson) {
		var ret string
		return ret
	}
	return *o.DataJson
}

// GetDataJsonOk returns a tuple with the DataJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSvcConfig) GetDataJsonOk() (*string, bool) {
	if o == nil || IsNil(o.DataJson) {
		return nil, false
	}
	return o.DataJson, true
}

// HasDataJson returns a boolean if a field has been set.
func (o *ConfigSvcConfig) HasDataJson() bool {
	if o != nil && !IsNil(o.DataJson) {
		return true
	}

	return false
}

// SetDataJson gets a reference to the given string and assigns it to the DataJson field.
func (o *ConfigSvcConfig) SetDataJson(v string) {
	o.DataJson = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigSvcConfig) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSvcConfig) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigSvcConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigSvcConfig) SetId(v string) {
	o.Id = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ConfigSvcConfig) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSvcConfig) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ConfigSvcConfig) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ConfigSvcConfig) SetNamespace(v string) {
	o.Namespace = &v
}

func (o ConfigSvcConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigSvcConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.DataJson) {
		toSerialize["dataJson"] = o.DataJson
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

func (o *ConfigSvcConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varConfigSvcConfig := _ConfigSvcConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigSvcConfig)

	if err != nil {
		return err
	}

	*o = ConfigSvcConfig(varConfigSvcConfig)

	return err
}

type NullableConfigSvcConfig struct {
	value *ConfigSvcConfig
	isSet bool
}

func (v NullableConfigSvcConfig) Get() *ConfigSvcConfig {
	return v.value
}

func (v *NullableConfigSvcConfig) Set(val *ConfigSvcConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigSvcConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigSvcConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigSvcConfig(val *ConfigSvcConfig) *NullableConfigSvcConfig {
	return &NullableConfigSvcConfig{value: val, isSet: true}
}

func (v NullableConfigSvcConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigSvcConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


