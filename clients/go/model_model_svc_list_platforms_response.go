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
	"bytes"
	"fmt"
)

// checks if the ModelSvcListPlatformsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSvcListPlatformsResponse{}

// ModelSvcListPlatformsResponse struct for ModelSvcListPlatformsResponse
type ModelSvcListPlatformsResponse struct {
	Platforms []ModelSvcPlatform `json:"platforms"`
}

type _ModelSvcListPlatformsResponse ModelSvcListPlatformsResponse

// NewModelSvcListPlatformsResponse instantiates a new ModelSvcListPlatformsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSvcListPlatformsResponse(platforms []ModelSvcPlatform) *ModelSvcListPlatformsResponse {
	this := ModelSvcListPlatformsResponse{}
	this.Platforms = platforms
	return &this
}

// NewModelSvcListPlatformsResponseWithDefaults instantiates a new ModelSvcListPlatformsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSvcListPlatformsResponseWithDefaults() *ModelSvcListPlatformsResponse {
	this := ModelSvcListPlatformsResponse{}
	return &this
}

// GetPlatforms returns the Platforms field value
func (o *ModelSvcListPlatformsResponse) GetPlatforms() []ModelSvcPlatform {
	if o == nil {
		var ret []ModelSvcPlatform
		return ret
	}

	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value
// and a boolean to check if the value has been set.
func (o *ModelSvcListPlatformsResponse) GetPlatformsOk() ([]ModelSvcPlatform, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platforms, true
}

// SetPlatforms sets field value
func (o *ModelSvcListPlatformsResponse) SetPlatforms(v []ModelSvcPlatform) {
	o.Platforms = v
}

func (o ModelSvcListPlatformsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSvcListPlatformsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["platforms"] = o.Platforms
	return toSerialize, nil
}

func (o *ModelSvcListPlatformsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"platforms",
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

	varModelSvcListPlatformsResponse := _ModelSvcListPlatformsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelSvcListPlatformsResponse)

	if err != nil {
		return err
	}

	*o = ModelSvcListPlatformsResponse(varModelSvcListPlatformsResponse)

	return err
}

type NullableModelSvcListPlatformsResponse struct {
	value *ModelSvcListPlatformsResponse
	isSet bool
}

func (v NullableModelSvcListPlatformsResponse) Get() *ModelSvcListPlatformsResponse {
	return v.value
}

func (v *NullableModelSvcListPlatformsResponse) Set(val *ModelSvcListPlatformsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSvcListPlatformsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSvcListPlatformsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSvcListPlatformsResponse(val *ModelSvcListPlatformsResponse) *NullableModelSvcListPlatformsResponse {
	return &NullableModelSvcListPlatformsResponse{value: val, isSet: true}
}

func (v NullableModelSvcListPlatformsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSvcListPlatformsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


