/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.39
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChatSvcGetMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSvcGetMessageResponse{}

// ChatSvcGetMessageResponse struct for ChatSvcGetMessageResponse
type ChatSvcGetMessageResponse struct {
	Exists *bool `json:"exists,omitempty"`
	Message *ChatSvcMessage `json:"message,omitempty"`
}

// NewChatSvcGetMessageResponse instantiates a new ChatSvcGetMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSvcGetMessageResponse() *ChatSvcGetMessageResponse {
	this := ChatSvcGetMessageResponse{}
	return &this
}

// NewChatSvcGetMessageResponseWithDefaults instantiates a new ChatSvcGetMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSvcGetMessageResponseWithDefaults() *ChatSvcGetMessageResponse {
	this := ChatSvcGetMessageResponse{}
	return &this
}

// GetExists returns the Exists field value if set, zero value otherwise.
func (o *ChatSvcGetMessageResponse) GetExists() bool {
	if o == nil || IsNil(o.Exists) {
		var ret bool
		return ret
	}
	return *o.Exists
}

// GetExistsOk returns a tuple with the Exists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcGetMessageResponse) GetExistsOk() (*bool, bool) {
	if o == nil || IsNil(o.Exists) {
		return nil, false
	}
	return o.Exists, true
}

// HasExists returns a boolean if a field has been set.
func (o *ChatSvcGetMessageResponse) HasExists() bool {
	if o != nil && !IsNil(o.Exists) {
		return true
	}

	return false
}

// SetExists gets a reference to the given bool and assigns it to the Exists field.
func (o *ChatSvcGetMessageResponse) SetExists(v bool) {
	o.Exists = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ChatSvcGetMessageResponse) GetMessage() ChatSvcMessage {
	if o == nil || IsNil(o.Message) {
		var ret ChatSvcMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcGetMessageResponse) GetMessageOk() (*ChatSvcMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ChatSvcGetMessageResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ChatSvcMessage and assigns it to the Message field.
func (o *ChatSvcGetMessageResponse) SetMessage(v ChatSvcMessage) {
	o.Message = &v
}

func (o ChatSvcGetMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSvcGetMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Exists) {
		toSerialize["exists"] = o.Exists
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableChatSvcGetMessageResponse struct {
	value *ChatSvcGetMessageResponse
	isSet bool
}

func (v NullableChatSvcGetMessageResponse) Get() *ChatSvcGetMessageResponse {
	return v.value
}

func (v *NullableChatSvcGetMessageResponse) Set(val *ChatSvcGetMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSvcGetMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSvcGetMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSvcGetMessageResponse(val *ChatSvcGetMessageResponse) *NullableChatSvcGetMessageResponse {
	return &NullableChatSvcGetMessageResponse{value: val, isSet: true}
}

func (v NullableChatSvcGetMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSvcGetMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


