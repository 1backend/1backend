/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.30
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChatSvcEventThreadUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSvcEventThreadUpdate{}

// ChatSvcEventThreadUpdate struct for ChatSvcEventThreadUpdate
type ChatSvcEventThreadUpdate struct {
	ThreadId *string `json:"threadId,omitempty"`
}

// NewChatSvcEventThreadUpdate instantiates a new ChatSvcEventThreadUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSvcEventThreadUpdate() *ChatSvcEventThreadUpdate {
	this := ChatSvcEventThreadUpdate{}
	return &this
}

// NewChatSvcEventThreadUpdateWithDefaults instantiates a new ChatSvcEventThreadUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSvcEventThreadUpdateWithDefaults() *ChatSvcEventThreadUpdate {
	this := ChatSvcEventThreadUpdate{}
	return &this
}

// GetThreadId returns the ThreadId field value if set, zero value otherwise.
func (o *ChatSvcEventThreadUpdate) GetThreadId() string {
	if o == nil || IsNil(o.ThreadId) {
		var ret string
		return ret
	}
	return *o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcEventThreadUpdate) GetThreadIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThreadId) {
		return nil, false
	}
	return o.ThreadId, true
}

// HasThreadId returns a boolean if a field has been set.
func (o *ChatSvcEventThreadUpdate) HasThreadId() bool {
	if o != nil && !IsNil(o.ThreadId) {
		return true
	}

	return false
}

// SetThreadId gets a reference to the given string and assigns it to the ThreadId field.
func (o *ChatSvcEventThreadUpdate) SetThreadId(v string) {
	o.ThreadId = &v
}

func (o ChatSvcEventThreadUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSvcEventThreadUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThreadId) {
		toSerialize["threadId"] = o.ThreadId
	}
	return toSerialize, nil
}

type NullableChatSvcEventThreadUpdate struct {
	value *ChatSvcEventThreadUpdate
	isSet bool
}

func (v NullableChatSvcEventThreadUpdate) Get() *ChatSvcEventThreadUpdate {
	return v.value
}

func (v *NullableChatSvcEventThreadUpdate) Set(val *ChatSvcEventThreadUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSvcEventThreadUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSvcEventThreadUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSvcEventThreadUpdate(val *ChatSvcEventThreadUpdate) *NullableChatSvcEventThreadUpdate {
	return &NullableChatSvcEventThreadUpdate{value: val, isSet: true}
}

func (v NullableChatSvcEventThreadUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSvcEventThreadUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


