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
)

// checks if the PromptSvcLlamaCppParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptSvcLlamaCppParameters{}

// PromptSvcLlamaCppParameters struct for PromptSvcLlamaCppParameters
type PromptSvcLlamaCppParameters struct {
	// Template of the prompt. Optional. If not present it's derived from ModelId.
	Template *string `json:"template,omitempty"`
}

// NewPromptSvcLlamaCppParameters instantiates a new PromptSvcLlamaCppParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptSvcLlamaCppParameters() *PromptSvcLlamaCppParameters {
	this := PromptSvcLlamaCppParameters{}
	return &this
}

// NewPromptSvcLlamaCppParametersWithDefaults instantiates a new PromptSvcLlamaCppParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptSvcLlamaCppParametersWithDefaults() *PromptSvcLlamaCppParameters {
	this := PromptSvcLlamaCppParameters{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *PromptSvcLlamaCppParameters) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcLlamaCppParameters) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *PromptSvcLlamaCppParameters) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *PromptSvcLlamaCppParameters) SetTemplate(v string) {
	o.Template = &v
}

func (o PromptSvcLlamaCppParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptSvcLlamaCppParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

type NullablePromptSvcLlamaCppParameters struct {
	value *PromptSvcLlamaCppParameters
	isSet bool
}

func (v NullablePromptSvcLlamaCppParameters) Get() *PromptSvcLlamaCppParameters {
	return v.value
}

func (v *NullablePromptSvcLlamaCppParameters) Set(val *PromptSvcLlamaCppParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcLlamaCppParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcLlamaCppParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcLlamaCppParameters(val *PromptSvcLlamaCppParameters) *NullablePromptSvcLlamaCppParameters {
	return &NullablePromptSvcLlamaCppParameters{value: val, isSet: true}
}

func (v NullablePromptSvcLlamaCppParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcLlamaCppParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


