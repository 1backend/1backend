/*
1Backend

A unified backend development platform for your AI applications—microservices-based and built to scale.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChatSvcAddThreadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSvcAddThreadRequest{}

// ChatSvcAddThreadRequest struct for ChatSvcAddThreadRequest
type ChatSvcAddThreadRequest struct {
	Thread *ChatSvcThread `json:"thread,omitempty"`
}

// NewChatSvcAddThreadRequest instantiates a new ChatSvcAddThreadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSvcAddThreadRequest() *ChatSvcAddThreadRequest {
	this := ChatSvcAddThreadRequest{}
	return &this
}

// NewChatSvcAddThreadRequestWithDefaults instantiates a new ChatSvcAddThreadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSvcAddThreadRequestWithDefaults() *ChatSvcAddThreadRequest {
	this := ChatSvcAddThreadRequest{}
	return &this
}

// GetThread returns the Thread field value if set, zero value otherwise.
func (o *ChatSvcAddThreadRequest) GetThread() ChatSvcThread {
	if o == nil || IsNil(o.Thread) {
		var ret ChatSvcThread
		return ret
	}
	return *o.Thread
}

// GetThreadOk returns a tuple with the Thread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcAddThreadRequest) GetThreadOk() (*ChatSvcThread, bool) {
	if o == nil || IsNil(o.Thread) {
		return nil, false
	}
	return o.Thread, true
}

// HasThread returns a boolean if a field has been set.
func (o *ChatSvcAddThreadRequest) HasThread() bool {
	if o != nil && !IsNil(o.Thread) {
		return true
	}

	return false
}

// SetThread gets a reference to the given ChatSvcThread and assigns it to the Thread field.
func (o *ChatSvcAddThreadRequest) SetThread(v ChatSvcThread) {
	o.Thread = &v
}

func (o ChatSvcAddThreadRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSvcAddThreadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Thread) {
		toSerialize["thread"] = o.Thread
	}
	return toSerialize, nil
}

type NullableChatSvcAddThreadRequest struct {
	value *ChatSvcAddThreadRequest
	isSet bool
}

func (v NullableChatSvcAddThreadRequest) Get() *ChatSvcAddThreadRequest {
	return v.value
}

func (v *NullableChatSvcAddThreadRequest) Set(val *ChatSvcAddThreadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSvcAddThreadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSvcAddThreadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSvcAddThreadRequest(val *ChatSvcAddThreadRequest) *NullableChatSvcAddThreadRequest {
	return &NullableChatSvcAddThreadRequest{value: val, isSet: true}
}

func (v NullableChatSvcAddThreadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSvcAddThreadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


