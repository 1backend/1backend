/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.32
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChatSvcGetThreadsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSvcGetThreadsResponse{}

// ChatSvcGetThreadsResponse struct for ChatSvcGetThreadsResponse
type ChatSvcGetThreadsResponse struct {
	Threads []ChatSvcThread `json:"threads,omitempty"`
}

// NewChatSvcGetThreadsResponse instantiates a new ChatSvcGetThreadsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSvcGetThreadsResponse() *ChatSvcGetThreadsResponse {
	this := ChatSvcGetThreadsResponse{}
	return &this
}

// NewChatSvcGetThreadsResponseWithDefaults instantiates a new ChatSvcGetThreadsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSvcGetThreadsResponseWithDefaults() *ChatSvcGetThreadsResponse {
	this := ChatSvcGetThreadsResponse{}
	return &this
}

// GetThreads returns the Threads field value if set, zero value otherwise.
func (o *ChatSvcGetThreadsResponse) GetThreads() []ChatSvcThread {
	if o == nil || IsNil(o.Threads) {
		var ret []ChatSvcThread
		return ret
	}
	return o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcGetThreadsResponse) GetThreadsOk() ([]ChatSvcThread, bool) {
	if o == nil || IsNil(o.Threads) {
		return nil, false
	}
	return o.Threads, true
}

// HasThreads returns a boolean if a field has been set.
func (o *ChatSvcGetThreadsResponse) HasThreads() bool {
	if o != nil && !IsNil(o.Threads) {
		return true
	}

	return false
}

// SetThreads gets a reference to the given []ChatSvcThread and assigns it to the Threads field.
func (o *ChatSvcGetThreadsResponse) SetThreads(v []ChatSvcThread) {
	o.Threads = v
}

func (o ChatSvcGetThreadsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSvcGetThreadsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Threads) {
		toSerialize["threads"] = o.Threads
	}
	return toSerialize, nil
}

type NullableChatSvcGetThreadsResponse struct {
	value *ChatSvcGetThreadsResponse
	isSet bool
}

func (v NullableChatSvcGetThreadsResponse) Get() *ChatSvcGetThreadsResponse {
	return v.value
}

func (v *NullableChatSvcGetThreadsResponse) Set(val *ChatSvcGetThreadsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSvcGetThreadsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSvcGetThreadsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSvcGetThreadsResponse(val *ChatSvcGetThreadsResponse) *NullableChatSvcGetThreadsResponse {
	return &NullableChatSvcGetThreadsResponse{value: val, isSet: true}
}

func (v NullableChatSvcGetThreadsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSvcGetThreadsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


