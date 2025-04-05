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

// checks if the PromptSvcParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptSvcParameters{}

// PromptSvcParameters struct for PromptSvcParameters
type PromptSvcParameters struct {
	TextToImage *PromptSvcTextToImageParameters `json:"textToImage,omitempty"`
	TextToText *PromptSvcTextToTextParameters `json:"textToText,omitempty"`
}

// NewPromptSvcParameters instantiates a new PromptSvcParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptSvcParameters() *PromptSvcParameters {
	this := PromptSvcParameters{}
	return &this
}

// NewPromptSvcParametersWithDefaults instantiates a new PromptSvcParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptSvcParametersWithDefaults() *PromptSvcParameters {
	this := PromptSvcParameters{}
	return &this
}

// GetTextToImage returns the TextToImage field value if set, zero value otherwise.
func (o *PromptSvcParameters) GetTextToImage() PromptSvcTextToImageParameters {
	if o == nil || IsNil(o.TextToImage) {
		var ret PromptSvcTextToImageParameters
		return ret
	}
	return *o.TextToImage
}

// GetTextToImageOk returns a tuple with the TextToImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcParameters) GetTextToImageOk() (*PromptSvcTextToImageParameters, bool) {
	if o == nil || IsNil(o.TextToImage) {
		return nil, false
	}
	return o.TextToImage, true
}

// HasTextToImage returns a boolean if a field has been set.
func (o *PromptSvcParameters) HasTextToImage() bool {
	if o != nil && !IsNil(o.TextToImage) {
		return true
	}

	return false
}

// SetTextToImage gets a reference to the given PromptSvcTextToImageParameters and assigns it to the TextToImage field.
func (o *PromptSvcParameters) SetTextToImage(v PromptSvcTextToImageParameters) {
	o.TextToImage = &v
}

// GetTextToText returns the TextToText field value if set, zero value otherwise.
func (o *PromptSvcParameters) GetTextToText() PromptSvcTextToTextParameters {
	if o == nil || IsNil(o.TextToText) {
		var ret PromptSvcTextToTextParameters
		return ret
	}
	return *o.TextToText
}

// GetTextToTextOk returns a tuple with the TextToText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcParameters) GetTextToTextOk() (*PromptSvcTextToTextParameters, bool) {
	if o == nil || IsNil(o.TextToText) {
		return nil, false
	}
	return o.TextToText, true
}

// HasTextToText returns a boolean if a field has been set.
func (o *PromptSvcParameters) HasTextToText() bool {
	if o != nil && !IsNil(o.TextToText) {
		return true
	}

	return false
}

// SetTextToText gets a reference to the given PromptSvcTextToTextParameters and assigns it to the TextToText field.
func (o *PromptSvcParameters) SetTextToText(v PromptSvcTextToTextParameters) {
	o.TextToText = &v
}

func (o PromptSvcParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptSvcParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TextToImage) {
		toSerialize["textToImage"] = o.TextToImage
	}
	if !IsNil(o.TextToText) {
		toSerialize["textToText"] = o.TextToText
	}
	return toSerialize, nil
}

type NullablePromptSvcParameters struct {
	value *PromptSvcParameters
	isSet bool
}

func (v NullablePromptSvcParameters) Get() *PromptSvcParameters {
	return v.value
}

func (v *NullablePromptSvcParameters) Set(val *PromptSvcParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcParameters(val *PromptSvcParameters) *NullablePromptSvcParameters {
	return &NullablePromptSvcParameters{value: val, isSet: true}
}

func (v NullablePromptSvcParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


