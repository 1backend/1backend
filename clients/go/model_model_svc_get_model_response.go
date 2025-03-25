/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.30
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelSvcGetModelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSvcGetModelResponse{}

// ModelSvcGetModelResponse struct for ModelSvcGetModelResponse
type ModelSvcGetModelResponse struct {
	Exists bool `json:"exists"`
	Model ModelSvcModel `json:"model"`
	Platform ModelSvcPlatform `json:"platform"`
}

type _ModelSvcGetModelResponse ModelSvcGetModelResponse

// NewModelSvcGetModelResponse instantiates a new ModelSvcGetModelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSvcGetModelResponse(exists bool, model ModelSvcModel, platform ModelSvcPlatform) *ModelSvcGetModelResponse {
	this := ModelSvcGetModelResponse{}
	this.Exists = exists
	this.Model = model
	this.Platform = platform
	return &this
}

// NewModelSvcGetModelResponseWithDefaults instantiates a new ModelSvcGetModelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSvcGetModelResponseWithDefaults() *ModelSvcGetModelResponse {
	this := ModelSvcGetModelResponse{}
	return &this
}

// GetExists returns the Exists field value
func (o *ModelSvcGetModelResponse) GetExists() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exists
}

// GetExistsOk returns a tuple with the Exists field value
// and a boolean to check if the value has been set.
func (o *ModelSvcGetModelResponse) GetExistsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exists, true
}

// SetExists sets field value
func (o *ModelSvcGetModelResponse) SetExists(v bool) {
	o.Exists = v
}

// GetModel returns the Model field value
func (o *ModelSvcGetModelResponse) GetModel() ModelSvcModel {
	if o == nil {
		var ret ModelSvcModel
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ModelSvcGetModelResponse) GetModelOk() (*ModelSvcModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ModelSvcGetModelResponse) SetModel(v ModelSvcModel) {
	o.Model = v
}

// GetPlatform returns the Platform field value
func (o *ModelSvcGetModelResponse) GetPlatform() ModelSvcPlatform {
	if o == nil {
		var ret ModelSvcPlatform
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *ModelSvcGetModelResponse) GetPlatformOk() (*ModelSvcPlatform, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *ModelSvcGetModelResponse) SetPlatform(v ModelSvcPlatform) {
	o.Platform = v
}

func (o ModelSvcGetModelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSvcGetModelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exists"] = o.Exists
	toSerialize["model"] = o.Model
	toSerialize["platform"] = o.Platform
	return toSerialize, nil
}

func (o *ModelSvcGetModelResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"exists",
		"model",
		"platform",
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

	varModelSvcGetModelResponse := _ModelSvcGetModelResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelSvcGetModelResponse)

	if err != nil {
		return err
	}

	*o = ModelSvcGetModelResponse(varModelSvcGetModelResponse)

	return err
}

type NullableModelSvcGetModelResponse struct {
	value *ModelSvcGetModelResponse
	isSet bool
}

func (v NullableModelSvcGetModelResponse) Get() *ModelSvcGetModelResponse {
	return v.value
}

func (v *NullableModelSvcGetModelResponse) Set(val *ModelSvcGetModelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSvcGetModelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSvcGetModelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSvcGetModelResponse(val *ModelSvcGetModelResponse) *NullableModelSvcGetModelResponse {
	return &NullableModelSvcGetModelResponse{value: val, isSet: true}
}

func (v NullableModelSvcGetModelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSvcGetModelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


