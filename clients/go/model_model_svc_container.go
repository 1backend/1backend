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

// checks if the ModelSvcContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSvcContainer{}

// ModelSvcContainer struct for ModelSvcContainer
type ModelSvcContainer struct {
	// Environment variables to be passed to the container (e.g., \"DEVICES=all\").
	Envars []ModelSvcEnvVar `json:"envars,omitempty"`
	// Template for constructing the container image name.
	ImageTemplate *string `json:"imageTemplate,omitempty"`
	// List of container paths that should persist across restarts.
	Keeps []ModelSvcKeep `json:"keeps,omitempty"`
	// Internal port exposed by the container.
	Port *int32 `json:"port,omitempty"`
}

// NewModelSvcContainer instantiates a new ModelSvcContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSvcContainer() *ModelSvcContainer {
	this := ModelSvcContainer{}
	return &this
}

// NewModelSvcContainerWithDefaults instantiates a new ModelSvcContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSvcContainerWithDefaults() *ModelSvcContainer {
	this := ModelSvcContainer{}
	return &this
}

// GetEnvars returns the Envars field value if set, zero value otherwise.
func (o *ModelSvcContainer) GetEnvars() []ModelSvcEnvVar {
	if o == nil || IsNil(o.Envars) {
		var ret []ModelSvcEnvVar
		return ret
	}
	return o.Envars
}

// GetEnvarsOk returns a tuple with the Envars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcContainer) GetEnvarsOk() ([]ModelSvcEnvVar, bool) {
	if o == nil || IsNil(o.Envars) {
		return nil, false
	}
	return o.Envars, true
}

// HasEnvars returns a boolean if a field has been set.
func (o *ModelSvcContainer) HasEnvars() bool {
	if o != nil && !IsNil(o.Envars) {
		return true
	}

	return false
}

// SetEnvars gets a reference to the given []ModelSvcEnvVar and assigns it to the Envars field.
func (o *ModelSvcContainer) SetEnvars(v []ModelSvcEnvVar) {
	o.Envars = v
}

// GetImageTemplate returns the ImageTemplate field value if set, zero value otherwise.
func (o *ModelSvcContainer) GetImageTemplate() string {
	if o == nil || IsNil(o.ImageTemplate) {
		var ret string
		return ret
	}
	return *o.ImageTemplate
}

// GetImageTemplateOk returns a tuple with the ImageTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcContainer) GetImageTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.ImageTemplate) {
		return nil, false
	}
	return o.ImageTemplate, true
}

// HasImageTemplate returns a boolean if a field has been set.
func (o *ModelSvcContainer) HasImageTemplate() bool {
	if o != nil && !IsNil(o.ImageTemplate) {
		return true
	}

	return false
}

// SetImageTemplate gets a reference to the given string and assigns it to the ImageTemplate field.
func (o *ModelSvcContainer) SetImageTemplate(v string) {
	o.ImageTemplate = &v
}

// GetKeeps returns the Keeps field value if set, zero value otherwise.
func (o *ModelSvcContainer) GetKeeps() []ModelSvcKeep {
	if o == nil || IsNil(o.Keeps) {
		var ret []ModelSvcKeep
		return ret
	}
	return o.Keeps
}

// GetKeepsOk returns a tuple with the Keeps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcContainer) GetKeepsOk() ([]ModelSvcKeep, bool) {
	if o == nil || IsNil(o.Keeps) {
		return nil, false
	}
	return o.Keeps, true
}

// HasKeeps returns a boolean if a field has been set.
func (o *ModelSvcContainer) HasKeeps() bool {
	if o != nil && !IsNil(o.Keeps) {
		return true
	}

	return false
}

// SetKeeps gets a reference to the given []ModelSvcKeep and assigns it to the Keeps field.
func (o *ModelSvcContainer) SetKeeps(v []ModelSvcKeep) {
	o.Keeps = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ModelSvcContainer) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcContainer) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ModelSvcContainer) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ModelSvcContainer) SetPort(v int32) {
	o.Port = &v
}

func (o ModelSvcContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSvcContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Envars) {
		toSerialize["envars"] = o.Envars
	}
	if !IsNil(o.ImageTemplate) {
		toSerialize["imageTemplate"] = o.ImageTemplate
	}
	if !IsNil(o.Keeps) {
		toSerialize["keeps"] = o.Keeps
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableModelSvcContainer struct {
	value *ModelSvcContainer
	isSet bool
}

func (v NullableModelSvcContainer) Get() *ModelSvcContainer {
	return v.value
}

func (v *NullableModelSvcContainer) Set(val *ModelSvcContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSvcContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSvcContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSvcContainer(val *ModelSvcContainer) *NullableModelSvcContainer {
	return &NullableModelSvcContainer{value: val, isSet: true}
}

func (v NullableModelSvcContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSvcContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


