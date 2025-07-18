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

// checks if the ChatSvcListThreadsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSvcListThreadsRequest{}

// ChatSvcListThreadsRequest struct for ChatSvcListThreadsRequest
type ChatSvcListThreadsRequest struct {
	Ids []string `json:"ids,omitempty"`
}

// NewChatSvcListThreadsRequest instantiates a new ChatSvcListThreadsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSvcListThreadsRequest() *ChatSvcListThreadsRequest {
	this := ChatSvcListThreadsRequest{}
	return &this
}

// NewChatSvcListThreadsRequestWithDefaults instantiates a new ChatSvcListThreadsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSvcListThreadsRequestWithDefaults() *ChatSvcListThreadsRequest {
	this := ChatSvcListThreadsRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ChatSvcListThreadsRequest) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcListThreadsRequest) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ChatSvcListThreadsRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *ChatSvcListThreadsRequest) SetIds(v []string) {
	o.Ids = v
}

func (o ChatSvcListThreadsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSvcListThreadsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableChatSvcListThreadsRequest struct {
	value *ChatSvcListThreadsRequest
	isSet bool
}

func (v NullableChatSvcListThreadsRequest) Get() *ChatSvcListThreadsRequest {
	return v.value
}

func (v *NullableChatSvcListThreadsRequest) Set(val *ChatSvcListThreadsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSvcListThreadsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSvcListThreadsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSvcListThreadsRequest(val *ChatSvcListThreadsRequest) *NullableChatSvcListThreadsRequest {
	return &NullableChatSvcListThreadsRequest{value: val, isSet: true}
}

func (v NullableChatSvcListThreadsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSvcListThreadsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


