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

// checks if the ChatSvcUpdateThreadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSvcUpdateThreadRequest{}

// ChatSvcUpdateThreadRequest struct for ChatSvcUpdateThreadRequest
type ChatSvcUpdateThreadRequest struct {
	Thread *ChatSvcThread `json:"thread,omitempty"`
}

// NewChatSvcUpdateThreadRequest instantiates a new ChatSvcUpdateThreadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSvcUpdateThreadRequest() *ChatSvcUpdateThreadRequest {
	this := ChatSvcUpdateThreadRequest{}
	return &this
}

// NewChatSvcUpdateThreadRequestWithDefaults instantiates a new ChatSvcUpdateThreadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSvcUpdateThreadRequestWithDefaults() *ChatSvcUpdateThreadRequest {
	this := ChatSvcUpdateThreadRequest{}
	return &this
}

// GetThread returns the Thread field value if set, zero value otherwise.
func (o *ChatSvcUpdateThreadRequest) GetThread() ChatSvcThread {
	if o == nil || IsNil(o.Thread) {
		var ret ChatSvcThread
		return ret
	}
	return *o.Thread
}

// GetThreadOk returns a tuple with the Thread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcUpdateThreadRequest) GetThreadOk() (*ChatSvcThread, bool) {
	if o == nil || IsNil(o.Thread) {
		return nil, false
	}
	return o.Thread, true
}

// HasThread returns a boolean if a field has been set.
func (o *ChatSvcUpdateThreadRequest) HasThread() bool {
	if o != nil && !IsNil(o.Thread) {
		return true
	}

	return false
}

// SetThread gets a reference to the given ChatSvcThread and assigns it to the Thread field.
func (o *ChatSvcUpdateThreadRequest) SetThread(v ChatSvcThread) {
	o.Thread = &v
}

func (o ChatSvcUpdateThreadRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSvcUpdateThreadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Thread) {
		toSerialize["thread"] = o.Thread
	}
	return toSerialize, nil
}

type NullableChatSvcUpdateThreadRequest struct {
	value *ChatSvcUpdateThreadRequest
	isSet bool
}

func (v NullableChatSvcUpdateThreadRequest) Get() *ChatSvcUpdateThreadRequest {
	return v.value
}

func (v *NullableChatSvcUpdateThreadRequest) Set(val *ChatSvcUpdateThreadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSvcUpdateThreadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSvcUpdateThreadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSvcUpdateThreadRequest(val *ChatSvcUpdateThreadRequest) *NullableChatSvcUpdateThreadRequest {
	return &NullableChatSvcUpdateThreadRequest{value: val, isSet: true}
}

func (v NullableChatSvcUpdateThreadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSvcUpdateThreadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


