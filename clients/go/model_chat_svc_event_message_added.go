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

// checks if the ChatSvcEventMessageAdded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSvcEventMessageAdded{}

// ChatSvcEventMessageAdded struct for ChatSvcEventMessageAdded
type ChatSvcEventMessageAdded struct {
	ThreadId *string `json:"threadId,omitempty"`
}

// NewChatSvcEventMessageAdded instantiates a new ChatSvcEventMessageAdded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSvcEventMessageAdded() *ChatSvcEventMessageAdded {
	this := ChatSvcEventMessageAdded{}
	return &this
}

// NewChatSvcEventMessageAddedWithDefaults instantiates a new ChatSvcEventMessageAdded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSvcEventMessageAddedWithDefaults() *ChatSvcEventMessageAdded {
	this := ChatSvcEventMessageAdded{}
	return &this
}

// GetThreadId returns the ThreadId field value if set, zero value otherwise.
func (o *ChatSvcEventMessageAdded) GetThreadId() string {
	if o == nil || IsNil(o.ThreadId) {
		var ret string
		return ret
	}
	return *o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcEventMessageAdded) GetThreadIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThreadId) {
		return nil, false
	}
	return o.ThreadId, true
}

// HasThreadId returns a boolean if a field has been set.
func (o *ChatSvcEventMessageAdded) HasThreadId() bool {
	if o != nil && !IsNil(o.ThreadId) {
		return true
	}

	return false
}

// SetThreadId gets a reference to the given string and assigns it to the ThreadId field.
func (o *ChatSvcEventMessageAdded) SetThreadId(v string) {
	o.ThreadId = &v
}

func (o ChatSvcEventMessageAdded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSvcEventMessageAdded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThreadId) {
		toSerialize["threadId"] = o.ThreadId
	}
	return toSerialize, nil
}

type NullableChatSvcEventMessageAdded struct {
	value *ChatSvcEventMessageAdded
	isSet bool
}

func (v NullableChatSvcEventMessageAdded) Get() *ChatSvcEventMessageAdded {
	return v.value
}

func (v *NullableChatSvcEventMessageAdded) Set(val *ChatSvcEventMessageAdded) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSvcEventMessageAdded) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSvcEventMessageAdded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSvcEventMessageAdded(val *ChatSvcEventMessageAdded) *NullableChatSvcEventMessageAdded {
	return &NullableChatSvcEventMessageAdded{value: val, isSet: true}
}

func (v NullableChatSvcEventMessageAdded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSvcEventMessageAdded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


