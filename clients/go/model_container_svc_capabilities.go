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

// checks if the ContainerSvcCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcCapabilities{}

// ContainerSvcCapabilities struct for ContainerSvcCapabilities
type ContainerSvcCapabilities struct {
	// GPUEnabled specifies whether GPU support is enabled for the container.
	GpuEnabled *bool `json:"gpuEnabled,omitempty"`
}

// NewContainerSvcCapabilities instantiates a new ContainerSvcCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcCapabilities() *ContainerSvcCapabilities {
	this := ContainerSvcCapabilities{}
	return &this
}

// NewContainerSvcCapabilitiesWithDefaults instantiates a new ContainerSvcCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcCapabilitiesWithDefaults() *ContainerSvcCapabilities {
	this := ContainerSvcCapabilities{}
	return &this
}

// GetGpuEnabled returns the GpuEnabled field value if set, zero value otherwise.
func (o *ContainerSvcCapabilities) GetGpuEnabled() bool {
	if o == nil || IsNil(o.GpuEnabled) {
		var ret bool
		return ret
	}
	return *o.GpuEnabled
}

// GetGpuEnabledOk returns a tuple with the GpuEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcCapabilities) GetGpuEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.GpuEnabled) {
		return nil, false
	}
	return o.GpuEnabled, true
}

// HasGpuEnabled returns a boolean if a field has been set.
func (o *ContainerSvcCapabilities) HasGpuEnabled() bool {
	if o != nil && !IsNil(o.GpuEnabled) {
		return true
	}

	return false
}

// SetGpuEnabled gets a reference to the given bool and assigns it to the GpuEnabled field.
func (o *ContainerSvcCapabilities) SetGpuEnabled(v bool) {
	o.GpuEnabled = &v
}

func (o ContainerSvcCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GpuEnabled) {
		toSerialize["gpuEnabled"] = o.GpuEnabled
	}
	return toSerialize, nil
}

type NullableContainerSvcCapabilities struct {
	value *ContainerSvcCapabilities
	isSet bool
}

func (v NullableContainerSvcCapabilities) Get() *ContainerSvcCapabilities {
	return v.value
}

func (v *NullableContainerSvcCapabilities) Set(val *ContainerSvcCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcCapabilities(val *ContainerSvcCapabilities) *NullableContainerSvcCapabilities {
	return &NullableContainerSvcCapabilities{value: val, isSet: true}
}

func (v NullableContainerSvcCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


