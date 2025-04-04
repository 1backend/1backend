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

// checks if the PromptSvcStreamChunk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptSvcStreamChunk{}

// PromptSvcStreamChunk struct for PromptSvcStreamChunk
type PromptSvcStreamChunk struct {
	// MessageId is the ChatSvc Message id that the chunk is part of. Might only be available for \"done\" chunks.
	MessageId *string `json:"messageId,omitempty"`
	// TextChunk contains a part of the text output from the stream.
	Text *string `json:"text,omitempty"`
	// Type indicates the type of the stream event (e.g., text, done).
	Type *PromptSvcStreamChunkType `json:"type,omitempty"`
}

// NewPromptSvcStreamChunk instantiates a new PromptSvcStreamChunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptSvcStreamChunk() *PromptSvcStreamChunk {
	this := PromptSvcStreamChunk{}
	return &this
}

// NewPromptSvcStreamChunkWithDefaults instantiates a new PromptSvcStreamChunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptSvcStreamChunkWithDefaults() *PromptSvcStreamChunk {
	this := PromptSvcStreamChunk{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *PromptSvcStreamChunk) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcStreamChunk) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *PromptSvcStreamChunk) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *PromptSvcStreamChunk) SetMessageId(v string) {
	o.MessageId = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *PromptSvcStreamChunk) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcStreamChunk) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *PromptSvcStreamChunk) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *PromptSvcStreamChunk) SetText(v string) {
	o.Text = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PromptSvcStreamChunk) GetType() PromptSvcStreamChunkType {
	if o == nil || IsNil(o.Type) {
		var ret PromptSvcStreamChunkType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcStreamChunk) GetTypeOk() (*PromptSvcStreamChunkType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PromptSvcStreamChunk) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PromptSvcStreamChunkType and assigns it to the Type field.
func (o *PromptSvcStreamChunk) SetType(v PromptSvcStreamChunkType) {
	o.Type = &v
}

func (o PromptSvcStreamChunk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptSvcStreamChunk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePromptSvcStreamChunk struct {
	value *PromptSvcStreamChunk
	isSet bool
}

func (v NullablePromptSvcStreamChunk) Get() *PromptSvcStreamChunk {
	return v.value
}

func (v *NullablePromptSvcStreamChunk) Set(val *PromptSvcStreamChunk) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcStreamChunk) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcStreamChunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcStreamChunk(val *PromptSvcStreamChunk) *NullablePromptSvcStreamChunk {
	return &NullablePromptSvcStreamChunk{value: val, isSet: true}
}

func (v NullablePromptSvcStreamChunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcStreamChunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


