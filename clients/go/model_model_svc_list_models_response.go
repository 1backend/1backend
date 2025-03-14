/*
1Backend

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ModelSvcListModelsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSvcListModelsResponse{}

// ModelSvcListModelsResponse struct for ModelSvcListModelsResponse
type ModelSvcListModelsResponse struct {
	Models []ModelSvcModel `json:"models,omitempty"`
}

// NewModelSvcListModelsResponse instantiates a new ModelSvcListModelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSvcListModelsResponse() *ModelSvcListModelsResponse {
	this := ModelSvcListModelsResponse{}
	return &this
}

// NewModelSvcListModelsResponseWithDefaults instantiates a new ModelSvcListModelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSvcListModelsResponseWithDefaults() *ModelSvcListModelsResponse {
	this := ModelSvcListModelsResponse{}
	return &this
}

// GetModels returns the Models field value if set, zero value otherwise.
func (o *ModelSvcListModelsResponse) GetModels() []ModelSvcModel {
	if o == nil || IsNil(o.Models) {
		var ret []ModelSvcModel
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcListModelsResponse) GetModelsOk() ([]ModelSvcModel, bool) {
	if o == nil || IsNil(o.Models) {
		return nil, false
	}
	return o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *ModelSvcListModelsResponse) HasModels() bool {
	if o != nil && !IsNil(o.Models) {
		return true
	}

	return false
}

// SetModels gets a reference to the given []ModelSvcModel and assigns it to the Models field.
func (o *ModelSvcListModelsResponse) SetModels(v []ModelSvcModel) {
	o.Models = v
}

func (o ModelSvcListModelsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSvcListModelsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Models) {
		toSerialize["models"] = o.Models
	}
	return toSerialize, nil
}

type NullableModelSvcListModelsResponse struct {
	value *ModelSvcListModelsResponse
	isSet bool
}

func (v NullableModelSvcListModelsResponse) Get() *ModelSvcListModelsResponse {
	return v.value
}

func (v *NullableModelSvcListModelsResponse) Set(val *ModelSvcListModelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSvcListModelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSvcListModelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSvcListModelsResponse(val *ModelSvcListModelsResponse) *NullableModelSvcListModelsResponse {
	return &NullableModelSvcListModelsResponse{value: val, isSet: true}
}

func (v NullableModelSvcListModelsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSvcListModelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


