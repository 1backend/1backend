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
	"bytes"
	"fmt"
)

// checks if the ChatSvcMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSvcMessage{}

// ChatSvcMessage struct for ChatSvcMessage
type ChatSvcMessage struct {
	CreatedAt string `json:"createdAt"`
	// FileIds defines the file attachments the message has.
	FileIds []string `json:"fileIds,omitempty"`
	Id string `json:"id"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	// Text content of the message eg. \"Hi, what's up?\"
	Text *string `json:"text,omitempty"`
	// ThreadId of the message.
	ThreadId string `json:"threadId"`
	UpdatedAt string `json:"updatedAt"`
	// UserId is the id of the user who wrote the message. For AI messages this field is empty.
	UserId *string `json:"userId,omitempty"`
}

type _ChatSvcMessage ChatSvcMessage

// NewChatSvcMessage instantiates a new ChatSvcMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSvcMessage(createdAt string, id string, threadId string, updatedAt string) *ChatSvcMessage {
	this := ChatSvcMessage{}
	this.CreatedAt = createdAt
	this.Id = id
	this.ThreadId = threadId
	this.UpdatedAt = updatedAt
	return &this
}

// NewChatSvcMessageWithDefaults instantiates a new ChatSvcMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSvcMessageWithDefaults() *ChatSvcMessage {
	this := ChatSvcMessage{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ChatSvcMessage) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ChatSvcMessage) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ChatSvcMessage) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetFileIds returns the FileIds field value if set, zero value otherwise.
func (o *ChatSvcMessage) GetFileIds() []string {
	if o == nil || IsNil(o.FileIds) {
		var ret []string
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcMessage) GetFileIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *ChatSvcMessage) HasFileIds() bool {
	if o != nil && !IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []string and assigns it to the FileIds field.
func (o *ChatSvcMessage) SetFileIds(v []string) {
	o.FileIds = v
}

// GetId returns the Id field value
func (o *ChatSvcMessage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChatSvcMessage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChatSvcMessage) SetId(v string) {
	o.Id = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ChatSvcMessage) GetMeta() map[string]interface{} {
	if o == nil || IsNil(o.Meta) {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcMessage) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ChatSvcMessage) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *ChatSvcMessage) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ChatSvcMessage) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcMessage) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ChatSvcMessage) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *ChatSvcMessage) SetText(v string) {
	o.Text = &v
}

// GetThreadId returns the ThreadId field value
func (o *ChatSvcMessage) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *ChatSvcMessage) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *ChatSvcMessage) SetThreadId(v string) {
	o.ThreadId = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ChatSvcMessage) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ChatSvcMessage) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ChatSvcMessage) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ChatSvcMessage) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcMessage) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ChatSvcMessage) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ChatSvcMessage) SetUserId(v string) {
	o.UserId = &v
}

func (o ChatSvcMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSvcMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.FileIds) {
		toSerialize["fileIds"] = o.FileIds
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	toSerialize["threadId"] = o.ThreadId
	toSerialize["updatedAt"] = o.UpdatedAt
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

func (o *ChatSvcMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"id",
		"threadId",
		"updatedAt",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varChatSvcMessage := _ChatSvcMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatSvcMessage)

	if err != nil {
		return err
	}

	*o = ChatSvcMessage(varChatSvcMessage)

	return err
}

type NullableChatSvcMessage struct {
	value *ChatSvcMessage
	isSet bool
}

func (v NullableChatSvcMessage) Get() *ChatSvcMessage {
	return v.value
}

func (v *NullableChatSvcMessage) Set(val *ChatSvcMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSvcMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSvcMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSvcMessage(val *ChatSvcMessage) *NullableChatSvcMessage {
	return &NullableChatSvcMessage{value: val, isSet: true}
}

func (v NullableChatSvcMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSvcMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


