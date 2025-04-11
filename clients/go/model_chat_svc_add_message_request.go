/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.35
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChatSvcAddMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSvcAddMessageRequest{}

// ChatSvcAddMessageRequest struct for ChatSvcAddMessageRequest
type ChatSvcAddMessageRequest struct {
	Message *ChatSvcMessage `json:"message,omitempty"`
}

// NewChatSvcAddMessageRequest instantiates a new ChatSvcAddMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSvcAddMessageRequest() *ChatSvcAddMessageRequest {
	this := ChatSvcAddMessageRequest{}
	return &this
}

// NewChatSvcAddMessageRequestWithDefaults instantiates a new ChatSvcAddMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSvcAddMessageRequestWithDefaults() *ChatSvcAddMessageRequest {
	this := ChatSvcAddMessageRequest{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ChatSvcAddMessageRequest) GetMessage() ChatSvcMessage {
	if o == nil || IsNil(o.Message) {
		var ret ChatSvcMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcAddMessageRequest) GetMessageOk() (*ChatSvcMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ChatSvcAddMessageRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ChatSvcMessage and assigns it to the Message field.
func (o *ChatSvcAddMessageRequest) SetMessage(v ChatSvcMessage) {
	o.Message = &v
}

func (o ChatSvcAddMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSvcAddMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableChatSvcAddMessageRequest struct {
	value *ChatSvcAddMessageRequest
	isSet bool
}

func (v NullableChatSvcAddMessageRequest) Get() *ChatSvcAddMessageRequest {
	return v.value
}

func (v *NullableChatSvcAddMessageRequest) Set(val *ChatSvcAddMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSvcAddMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSvcAddMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSvcAddMessageRequest(val *ChatSvcAddMessageRequest) *NullableChatSvcAddMessageRequest {
	return &NullableChatSvcAddMessageRequest{value: val, isSet: true}
}

func (v NullableChatSvcAddMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSvcAddMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


