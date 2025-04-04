/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.31
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ContainerSvcAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcAsset{}

// ContainerSvcAsset struct for ContainerSvcAsset
type ContainerSvcAsset struct {
	EnvVarKey string `json:"envVarKey"`
	Url string `json:"url"`
}

type _ContainerSvcAsset ContainerSvcAsset

// NewContainerSvcAsset instantiates a new ContainerSvcAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcAsset(envVarKey string, url string) *ContainerSvcAsset {
	this := ContainerSvcAsset{}
	this.EnvVarKey = envVarKey
	this.Url = url
	return &this
}

// NewContainerSvcAssetWithDefaults instantiates a new ContainerSvcAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcAssetWithDefaults() *ContainerSvcAsset {
	this := ContainerSvcAsset{}
	return &this
}

// GetEnvVarKey returns the EnvVarKey field value
func (o *ContainerSvcAsset) GetEnvVarKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvVarKey
}

// GetEnvVarKeyOk returns a tuple with the EnvVarKey field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcAsset) GetEnvVarKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvVarKey, true
}

// SetEnvVarKey sets field value
func (o *ContainerSvcAsset) SetEnvVarKey(v string) {
	o.EnvVarKey = v
}

// GetUrl returns the Url field value
func (o *ContainerSvcAsset) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcAsset) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ContainerSvcAsset) SetUrl(v string) {
	o.Url = v
}

func (o ContainerSvcAsset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["envVarKey"] = o.EnvVarKey
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *ContainerSvcAsset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"envVarKey",
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

	varContainerSvcAsset := _ContainerSvcAsset{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerSvcAsset)

	if err != nil {
		return err
	}

	*o = ContainerSvcAsset(varContainerSvcAsset)

	return err
}

type NullableContainerSvcAsset struct {
	value *ContainerSvcAsset
	isSet bool
}

func (v NullableContainerSvcAsset) Get() *ContainerSvcAsset {
	return v.value
}

func (v *NullableContainerSvcAsset) Set(val *ContainerSvcAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcAsset(val *ContainerSvcAsset) *NullableContainerSvcAsset {
	return &NullableContainerSvcAsset{value: val, isSet: true}
}

func (v NullableContainerSvcAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


