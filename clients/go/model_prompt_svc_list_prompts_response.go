/*
1Backend

A unified backend for your AI applications—microservices-based and built to scale.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PromptSvcListPromptsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptSvcListPromptsResponse{}

// PromptSvcListPromptsResponse struct for PromptSvcListPromptsResponse
type PromptSvcListPromptsResponse struct {
	After map[string]interface{} `json:"after,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Prompts []PromptSvcPrompt `json:"prompts,omitempty"`
}

// NewPromptSvcListPromptsResponse instantiates a new PromptSvcListPromptsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptSvcListPromptsResponse() *PromptSvcListPromptsResponse {
	this := PromptSvcListPromptsResponse{}
	return &this
}

// NewPromptSvcListPromptsResponseWithDefaults instantiates a new PromptSvcListPromptsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptSvcListPromptsResponseWithDefaults() *PromptSvcListPromptsResponse {
	this := PromptSvcListPromptsResponse{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *PromptSvcListPromptsResponse) GetAfter() map[string]interface{} {
	if o == nil || IsNil(o.After) {
		var ret map[string]interface{}
		return ret
	}
	return o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcListPromptsResponse) GetAfterOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.After) {
		return map[string]interface{}{}, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *PromptSvcListPromptsResponse) HasAfter() bool {
	if o != nil && !IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given map[string]interface{} and assigns it to the After field.
func (o *PromptSvcListPromptsResponse) SetAfter(v map[string]interface{}) {
	o.After = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PromptSvcListPromptsResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcListPromptsResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PromptSvcListPromptsResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PromptSvcListPromptsResponse) SetCount(v int32) {
	o.Count = &v
}

// GetPrompts returns the Prompts field value if set, zero value otherwise.
func (o *PromptSvcListPromptsResponse) GetPrompts() []PromptSvcPrompt {
	if o == nil || IsNil(o.Prompts) {
		var ret []PromptSvcPrompt
		return ret
	}
	return o.Prompts
}

// GetPromptsOk returns a tuple with the Prompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcListPromptsResponse) GetPromptsOk() ([]PromptSvcPrompt, bool) {
	if o == nil || IsNil(o.Prompts) {
		return nil, false
	}
	return o.Prompts, true
}

// HasPrompts returns a boolean if a field has been set.
func (o *PromptSvcListPromptsResponse) HasPrompts() bool {
	if o != nil && !IsNil(o.Prompts) {
		return true
	}

	return false
}

// SetPrompts gets a reference to the given []PromptSvcPrompt and assigns it to the Prompts field.
func (o *PromptSvcListPromptsResponse) SetPrompts(v []PromptSvcPrompt) {
	o.Prompts = v
}

func (o PromptSvcListPromptsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptSvcListPromptsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.After) {
		toSerialize["after"] = o.After
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Prompts) {
		toSerialize["prompts"] = o.Prompts
	}
	return toSerialize, nil
}

type NullablePromptSvcListPromptsResponse struct {
	value *PromptSvcListPromptsResponse
	isSet bool
}

func (v NullablePromptSvcListPromptsResponse) Get() *PromptSvcListPromptsResponse {
	return v.value
}

func (v *NullablePromptSvcListPromptsResponse) Set(val *PromptSvcListPromptsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcListPromptsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcListPromptsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcListPromptsResponse(val *PromptSvcListPromptsResponse) *NullablePromptSvcListPromptsResponse {
	return &NullablePromptSvcListPromptsResponse{value: val, isSet: true}
}

func (v NullablePromptSvcListPromptsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcListPromptsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


