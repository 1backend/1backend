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

// checks if the ChatSvcAddThreadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSvcAddThreadResponse{}

// ChatSvcAddThreadResponse struct for ChatSvcAddThreadResponse
type ChatSvcAddThreadResponse struct {
	Thread *ChatSvcThread `json:"thread,omitempty"`
}

// NewChatSvcAddThreadResponse instantiates a new ChatSvcAddThreadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSvcAddThreadResponse() *ChatSvcAddThreadResponse {
	this := ChatSvcAddThreadResponse{}
	return &this
}

// NewChatSvcAddThreadResponseWithDefaults instantiates a new ChatSvcAddThreadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSvcAddThreadResponseWithDefaults() *ChatSvcAddThreadResponse {
	this := ChatSvcAddThreadResponse{}
	return &this
}

// GetThread returns the Thread field value if set, zero value otherwise.
func (o *ChatSvcAddThreadResponse) GetThread() ChatSvcThread {
	if o == nil || IsNil(o.Thread) {
		var ret ChatSvcThread
		return ret
	}
	return *o.Thread
}

// GetThreadOk returns a tuple with the Thread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcAddThreadResponse) GetThreadOk() (*ChatSvcThread, bool) {
	if o == nil || IsNil(o.Thread) {
		return nil, false
	}
	return o.Thread, true
}

// HasThread returns a boolean if a field has been set.
func (o *ChatSvcAddThreadResponse) HasThread() bool {
	if o != nil && !IsNil(o.Thread) {
		return true
	}

	return false
}

// SetThread gets a reference to the given ChatSvcThread and assigns it to the Thread field.
func (o *ChatSvcAddThreadResponse) SetThread(v ChatSvcThread) {
	o.Thread = &v
}

func (o ChatSvcAddThreadResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSvcAddThreadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Thread) {
		toSerialize["thread"] = o.Thread
	}
	return toSerialize, nil
}

type NullableChatSvcAddThreadResponse struct {
	value *ChatSvcAddThreadResponse
	isSet bool
}

func (v NullableChatSvcAddThreadResponse) Get() *ChatSvcAddThreadResponse {
	return v.value
}

func (v *NullableChatSvcAddThreadResponse) Set(val *ChatSvcAddThreadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSvcAddThreadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSvcAddThreadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSvcAddThreadResponse(val *ChatSvcAddThreadResponse) *NullableChatSvcAddThreadResponse {
	return &NullableChatSvcAddThreadResponse{value: val, isSet: true}
}

func (v NullableChatSvcAddThreadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSvcAddThreadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


