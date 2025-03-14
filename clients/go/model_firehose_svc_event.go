/*
1Backend

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FirehoseSvcEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirehoseSvcEvent{}

// FirehoseSvcEvent struct for FirehoseSvcEvent
type FirehoseSvcEvent struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewFirehoseSvcEvent instantiates a new FirehoseSvcEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirehoseSvcEvent() *FirehoseSvcEvent {
	this := FirehoseSvcEvent{}
	return &this
}

// NewFirehoseSvcEventWithDefaults instantiates a new FirehoseSvcEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirehoseSvcEventWithDefaults() *FirehoseSvcEvent {
	this := FirehoseSvcEvent{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FirehoseSvcEvent) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirehoseSvcEvent) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FirehoseSvcEvent) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *FirehoseSvcEvent) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FirehoseSvcEvent) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirehoseSvcEvent) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FirehoseSvcEvent) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FirehoseSvcEvent) SetName(v string) {
	o.Name = &v
}

func (o FirehoseSvcEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirehoseSvcEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableFirehoseSvcEvent struct {
	value *FirehoseSvcEvent
	isSet bool
}

func (v NullableFirehoseSvcEvent) Get() *FirehoseSvcEvent {
	return v.value
}

func (v *NullableFirehoseSvcEvent) Set(val *FirehoseSvcEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableFirehoseSvcEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableFirehoseSvcEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirehoseSvcEvent(val *FirehoseSvcEvent) *NullableFirehoseSvcEvent {
	return &NullableFirehoseSvcEvent{value: val, isSet: true}
}

func (v NullableFirehoseSvcEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirehoseSvcEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


