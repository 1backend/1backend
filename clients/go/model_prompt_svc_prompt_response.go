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

// checks if the PromptSvcPromptResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptSvcPromptResponse{}

// PromptSvcPromptResponse struct for PromptSvcPromptResponse
type PromptSvcPromptResponse struct {
	// Prompt contains the details of the prompt that was just created by this request. This includes the ID, prompt text, status, and other associated metadata.
	Prompt *PromptSvcPrompt `json:"prompt,omitempty"`
	// Response message contains the response text and files. This field is populated only for synchronous prompts (`sync = true`). For asynchronous prompts, the response will provided in the associated message identified by the `responseMessageId` of the `promptSvc.prompt` object once the prompt completes.
	ResponseMessage *ChatSvcMessage `json:"responseMessage,omitempty"`
}

// NewPromptSvcPromptResponse instantiates a new PromptSvcPromptResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptSvcPromptResponse() *PromptSvcPromptResponse {
	this := PromptSvcPromptResponse{}
	return &this
}

// NewPromptSvcPromptResponseWithDefaults instantiates a new PromptSvcPromptResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptSvcPromptResponseWithDefaults() *PromptSvcPromptResponse {
	this := PromptSvcPromptResponse{}
	return &this
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *PromptSvcPromptResponse) GetPrompt() PromptSvcPrompt {
	if o == nil || IsNil(o.Prompt) {
		var ret PromptSvcPrompt
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPromptResponse) GetPromptOk() (*PromptSvcPrompt, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *PromptSvcPromptResponse) HasPrompt() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given PromptSvcPrompt and assigns it to the Prompt field.
func (o *PromptSvcPromptResponse) SetPrompt(v PromptSvcPrompt) {
	o.Prompt = &v
}

// GetResponseMessage returns the ResponseMessage field value if set, zero value otherwise.
func (o *PromptSvcPromptResponse) GetResponseMessage() ChatSvcMessage {
	if o == nil || IsNil(o.ResponseMessage) {
		var ret ChatSvcMessage
		return ret
	}
	return *o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcPromptResponse) GetResponseMessageOk() (*ChatSvcMessage, bool) {
	if o == nil || IsNil(o.ResponseMessage) {
		return nil, false
	}
	return o.ResponseMessage, true
}

// HasResponseMessage returns a boolean if a field has been set.
func (o *PromptSvcPromptResponse) HasResponseMessage() bool {
	if o != nil && !IsNil(o.ResponseMessage) {
		return true
	}

	return false
}

// SetResponseMessage gets a reference to the given ChatSvcMessage and assigns it to the ResponseMessage field.
func (o *PromptSvcPromptResponse) SetResponseMessage(v ChatSvcMessage) {
	o.ResponseMessage = &v
}

func (o PromptSvcPromptResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptSvcPromptResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	if !IsNil(o.ResponseMessage) {
		toSerialize["responseMessage"] = o.ResponseMessage
	}
	return toSerialize, nil
}

type NullablePromptSvcPromptResponse struct {
	value *PromptSvcPromptResponse
	isSet bool
}

func (v NullablePromptSvcPromptResponse) Get() *PromptSvcPromptResponse {
	return v.value
}

func (v *NullablePromptSvcPromptResponse) Set(val *PromptSvcPromptResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcPromptResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcPromptResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcPromptResponse(val *PromptSvcPromptResponse) *NullablePromptSvcPromptResponse {
	return &NullablePromptSvcPromptResponse{value: val, isSet: true}
}

func (v NullablePromptSvcPromptResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcPromptResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


