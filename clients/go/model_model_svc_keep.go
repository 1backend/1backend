/*
OpenOrch

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ModelSvcKeep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSvcKeep{}

// ModelSvcKeep struct for ModelSvcKeep
type ModelSvcKeep struct {
	// Path is the absolute path inside the container for the folder that should persist across restarts.
	Path *string `json:"path,omitempty"`
	// ReadOnly indicates whether the keep is read-only.
	ReadOnly *bool `json:"readOnly,omitempty"`
}

// NewModelSvcKeep instantiates a new ModelSvcKeep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSvcKeep() *ModelSvcKeep {
	this := ModelSvcKeep{}
	return &this
}

// NewModelSvcKeepWithDefaults instantiates a new ModelSvcKeep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSvcKeepWithDefaults() *ModelSvcKeep {
	this := ModelSvcKeep{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ModelSvcKeep) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcKeep) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ModelSvcKeep) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ModelSvcKeep) SetPath(v string) {
	o.Path = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *ModelSvcKeep) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcKeep) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *ModelSvcKeep) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *ModelSvcKeep) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o ModelSvcKeep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSvcKeep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	return toSerialize, nil
}

type NullableModelSvcKeep struct {
	value *ModelSvcKeep
	isSet bool
}

func (v NullableModelSvcKeep) Get() *ModelSvcKeep {
	return v.value
}

func (v *NullableModelSvcKeep) Set(val *ModelSvcKeep) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSvcKeep) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSvcKeep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSvcKeep(val *ModelSvcKeep) *NullableModelSvcKeep {
	return &NullableModelSvcKeep{value: val, isSet: true}
}

func (v NullableModelSvcKeep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSvcKeep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


