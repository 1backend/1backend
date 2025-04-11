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

// checks if the ChatSvcGetThreadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSvcGetThreadResponse{}

// ChatSvcGetThreadResponse struct for ChatSvcGetThreadResponse
type ChatSvcGetThreadResponse struct {
	Exists *bool `json:"exists,omitempty"`
	Thread *ChatSvcThread `json:"thread,omitempty"`
}

// NewChatSvcGetThreadResponse instantiates a new ChatSvcGetThreadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSvcGetThreadResponse() *ChatSvcGetThreadResponse {
	this := ChatSvcGetThreadResponse{}
	return &this
}

// NewChatSvcGetThreadResponseWithDefaults instantiates a new ChatSvcGetThreadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSvcGetThreadResponseWithDefaults() *ChatSvcGetThreadResponse {
	this := ChatSvcGetThreadResponse{}
	return &this
}

// GetExists returns the Exists field value if set, zero value otherwise.
func (o *ChatSvcGetThreadResponse) GetExists() bool {
	if o == nil || IsNil(o.Exists) {
		var ret bool
		return ret
	}
	return *o.Exists
}

// GetExistsOk returns a tuple with the Exists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcGetThreadResponse) GetExistsOk() (*bool, bool) {
	if o == nil || IsNil(o.Exists) {
		return nil, false
	}
	return o.Exists, true
}

// HasExists returns a boolean if a field has been set.
func (o *ChatSvcGetThreadResponse) HasExists() bool {
	if o != nil && !IsNil(o.Exists) {
		return true
	}

	return false
}

// SetExists gets a reference to the given bool and assigns it to the Exists field.
func (o *ChatSvcGetThreadResponse) SetExists(v bool) {
	o.Exists = &v
}

// GetThread returns the Thread field value if set, zero value otherwise.
func (o *ChatSvcGetThreadResponse) GetThread() ChatSvcThread {
	if o == nil || IsNil(o.Thread) {
		var ret ChatSvcThread
		return ret
	}
	return *o.Thread
}

// GetThreadOk returns a tuple with the Thread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcGetThreadResponse) GetThreadOk() (*ChatSvcThread, bool) {
	if o == nil || IsNil(o.Thread) {
		return nil, false
	}
	return o.Thread, true
}

// HasThread returns a boolean if a field has been set.
func (o *ChatSvcGetThreadResponse) HasThread() bool {
	if o != nil && !IsNil(o.Thread) {
		return true
	}

	return false
}

// SetThread gets a reference to the given ChatSvcThread and assigns it to the Thread field.
func (o *ChatSvcGetThreadResponse) SetThread(v ChatSvcThread) {
	o.Thread = &v
}

func (o ChatSvcGetThreadResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSvcGetThreadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Exists) {
		toSerialize["exists"] = o.Exists
	}
	if !IsNil(o.Thread) {
		toSerialize["thread"] = o.Thread
	}
	return toSerialize, nil
}

type NullableChatSvcGetThreadResponse struct {
	value *ChatSvcGetThreadResponse
	isSet bool
}

func (v NullableChatSvcGetThreadResponse) Get() *ChatSvcGetThreadResponse {
	return v.value
}

func (v *NullableChatSvcGetThreadResponse) Set(val *ChatSvcGetThreadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSvcGetThreadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSvcGetThreadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSvcGetThreadResponse(val *ChatSvcGetThreadResponse) *NullableChatSvcGetThreadResponse {
	return &NullableChatSvcGetThreadResponse{value: val, isSet: true}
}

func (v NullableChatSvcGetThreadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSvcGetThreadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


