/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.31
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ContainerSvcLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcLog{}

// ContainerSvcLog struct for ContainerSvcLog
type ContainerSvcLog struct {
	// ContainerId is the raw underlying container ID. Eg. Docker container id. Node local.
	ContainerId *string `json:"containerId,omitempty"`
	Content *string `json:"content,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Id *string `json:"id,omitempty"`
	// Node Id Please see the documentation for the envar OB_NODE_ID
	NodeId *string `json:"nodeId,omitempty"`
}

// NewContainerSvcLog instantiates a new ContainerSvcLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcLog() *ContainerSvcLog {
	this := ContainerSvcLog{}
	return &this
}

// NewContainerSvcLogWithDefaults instantiates a new ContainerSvcLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcLogWithDefaults() *ContainerSvcLog {
	this := ContainerSvcLog{}
	return &this
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *ContainerSvcLog) GetContainerId() string {
	if o == nil || IsNil(o.ContainerId) {
		var ret string
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcLog) GetContainerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerId) {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *ContainerSvcLog) HasContainerId() bool {
	if o != nil && !IsNil(o.ContainerId) {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given string and assigns it to the ContainerId field.
func (o *ContainerSvcLog) SetContainerId(v string) {
	o.ContainerId = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ContainerSvcLog) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcLog) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ContainerSvcLog) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ContainerSvcLog) SetContent(v string) {
	o.Content = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ContainerSvcLog) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcLog) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ContainerSvcLog) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ContainerSvcLog) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerSvcLog) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcLog) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerSvcLog) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContainerSvcLog) SetId(v string) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ContainerSvcLog) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcLog) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ContainerSvcLog) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *ContainerSvcLog) SetNodeId(v string) {
	o.NodeId = &v
}

func (o ContainerSvcLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerId) {
		toSerialize["containerId"] = o.ContainerId
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NodeId) {
		toSerialize["nodeId"] = o.NodeId
	}
	return toSerialize, nil
}

type NullableContainerSvcLog struct {
	value *ContainerSvcLog
	isSet bool
}

func (v NullableContainerSvcLog) Get() *ContainerSvcLog {
	return v.value
}

func (v *NullableContainerSvcLog) Set(val *ContainerSvcLog) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcLog) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcLog(val *ContainerSvcLog) *NullableContainerSvcLog {
	return &NullableContainerSvcLog{value: val, isSet: true}
}

func (v NullableContainerSvcLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


