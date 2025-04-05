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

// checks if the PromptSvcStableDiffusionParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptSvcStableDiffusionParameters{}

// PromptSvcStableDiffusionParameters struct for PromptSvcStableDiffusionParameters
type PromptSvcStableDiffusionParameters struct {
	// Text to image parameters
	Txt2Img *StableDiffusionTxt2ImgRequest `json:"txt2Img,omitempty"`
}

// NewPromptSvcStableDiffusionParameters instantiates a new PromptSvcStableDiffusionParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptSvcStableDiffusionParameters() *PromptSvcStableDiffusionParameters {
	this := PromptSvcStableDiffusionParameters{}
	return &this
}

// NewPromptSvcStableDiffusionParametersWithDefaults instantiates a new PromptSvcStableDiffusionParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptSvcStableDiffusionParametersWithDefaults() *PromptSvcStableDiffusionParameters {
	this := PromptSvcStableDiffusionParameters{}
	return &this
}

// GetTxt2Img returns the Txt2Img field value if set, zero value otherwise.
func (o *PromptSvcStableDiffusionParameters) GetTxt2Img() StableDiffusionTxt2ImgRequest {
	if o == nil || IsNil(o.Txt2Img) {
		var ret StableDiffusionTxt2ImgRequest
		return ret
	}
	return *o.Txt2Img
}

// GetTxt2ImgOk returns a tuple with the Txt2Img field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcStableDiffusionParameters) GetTxt2ImgOk() (*StableDiffusionTxt2ImgRequest, bool) {
	if o == nil || IsNil(o.Txt2Img) {
		return nil, false
	}
	return o.Txt2Img, true
}

// HasTxt2Img returns a boolean if a field has been set.
func (o *PromptSvcStableDiffusionParameters) HasTxt2Img() bool {
	if o != nil && !IsNil(o.Txt2Img) {
		return true
	}

	return false
}

// SetTxt2Img gets a reference to the given StableDiffusionTxt2ImgRequest and assigns it to the Txt2Img field.
func (o *PromptSvcStableDiffusionParameters) SetTxt2Img(v StableDiffusionTxt2ImgRequest) {
	o.Txt2Img = &v
}

func (o PromptSvcStableDiffusionParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptSvcStableDiffusionParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Txt2Img) {
		toSerialize["txt2Img"] = o.Txt2Img
	}
	return toSerialize, nil
}

type NullablePromptSvcStableDiffusionParameters struct {
	value *PromptSvcStableDiffusionParameters
	isSet bool
}

func (v NullablePromptSvcStableDiffusionParameters) Get() *PromptSvcStableDiffusionParameters {
	return v.value
}

func (v *NullablePromptSvcStableDiffusionParameters) Set(val *PromptSvcStableDiffusionParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcStableDiffusionParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcStableDiffusionParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcStableDiffusionParameters(val *PromptSvcStableDiffusionParameters) *NullablePromptSvcStableDiffusionParameters {
	return &NullablePromptSvcStableDiffusionParameters{value: val, isSet: true}
}

func (v NullablePromptSvcStableDiffusionParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcStableDiffusionParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


