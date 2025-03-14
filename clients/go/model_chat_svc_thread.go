/*
1Backend

A common backend for your AI applications—microservices-based and built to scale.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ChatSvcThread type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatSvcThread{}

// ChatSvcThread struct for ChatSvcThread
type ChatSvcThread struct {
	CreatedAt *string `json:"createdAt,omitempty"`
	Id string `json:"id"`
	// Title of the thread.
	Title *string `json:"title,omitempty"`
	// TopicIds defines which topics the thread belongs to. Topics can roughly be thought of as tags for threads.
	TopicIds []string `json:"topicIds,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// UserIds the ids of the users who can see this thread.
	UserIds []string `json:"userIds,omitempty"`
}

type _ChatSvcThread ChatSvcThread

// NewChatSvcThread instantiates a new ChatSvcThread object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatSvcThread(id string) *ChatSvcThread {
	this := ChatSvcThread{}
	this.Id = id
	return &this
}

// NewChatSvcThreadWithDefaults instantiates a new ChatSvcThread object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatSvcThreadWithDefaults() *ChatSvcThread {
	this := ChatSvcThread{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ChatSvcThread) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcThread) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ChatSvcThread) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ChatSvcThread) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value
func (o *ChatSvcThread) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChatSvcThread) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChatSvcThread) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ChatSvcThread) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcThread) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ChatSvcThread) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ChatSvcThread) SetTitle(v string) {
	o.Title = &v
}

// GetTopicIds returns the TopicIds field value if set, zero value otherwise.
func (o *ChatSvcThread) GetTopicIds() []string {
	if o == nil || IsNil(o.TopicIds) {
		var ret []string
		return ret
	}
	return o.TopicIds
}

// GetTopicIdsOk returns a tuple with the TopicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcThread) GetTopicIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TopicIds) {
		return nil, false
	}
	return o.TopicIds, true
}

// HasTopicIds returns a boolean if a field has been set.
func (o *ChatSvcThread) HasTopicIds() bool {
	if o != nil && !IsNil(o.TopicIds) {
		return true
	}

	return false
}

// SetTopicIds gets a reference to the given []string and assigns it to the TopicIds field.
func (o *ChatSvcThread) SetTopicIds(v []string) {
	o.TopicIds = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ChatSvcThread) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcThread) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ChatSvcThread) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ChatSvcThread) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUserIds returns the UserIds field value if set, zero value otherwise.
func (o *ChatSvcThread) GetUserIds() []string {
	if o == nil || IsNil(o.UserIds) {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatSvcThread) GetUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *ChatSvcThread) HasUserIds() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *ChatSvcThread) SetUserIds(v []string) {
	o.UserIds = v
}

func (o ChatSvcThread) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatSvcThread) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.TopicIds) {
		toSerialize["topicIds"] = o.TopicIds
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.UserIds) {
		toSerialize["userIds"] = o.UserIds
	}
	return toSerialize, nil
}

func (o *ChatSvcThread) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varChatSvcThread := _ChatSvcThread{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatSvcThread)

	if err != nil {
		return err
	}

	*o = ChatSvcThread(varChatSvcThread)

	return err
}

type NullableChatSvcThread struct {
	value *ChatSvcThread
	isSet bool
}

func (v NullableChatSvcThread) Get() *ChatSvcThread {
	return v.value
}

func (v *NullableChatSvcThread) Set(val *ChatSvcThread) {
	v.value = val
	v.isSet = true
}

func (v NullableChatSvcThread) IsSet() bool {
	return v.isSet
}

func (v *NullableChatSvcThread) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatSvcThread(val *ChatSvcThread) *NullableChatSvcThread {
	return &NullableChatSvcThread{value: val, isSet: true}
}

func (v NullableChatSvcThread) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatSvcThread) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


