/*
1Backend

A unified backend for your AI applications—microservices-based and built to scale.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RegistrySvcResourceUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcResourceUsage{}

// RegistrySvcResourceUsage struct for RegistrySvcResourceUsage
type RegistrySvcResourceUsage struct {
	// CPU usage metrics.
	Cpu *RegistrySvcUsage `json:"cpu,omitempty"`
	// Disk usage metrics.
	Disk *RegistrySvcUsage `json:"disk,omitempty"`
	// Memory usage metrics.
	Memory *RegistrySvcUsage `json:"memory,omitempty"`
}

// NewRegistrySvcResourceUsage instantiates a new RegistrySvcResourceUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcResourceUsage() *RegistrySvcResourceUsage {
	this := RegistrySvcResourceUsage{}
	return &this
}

// NewRegistrySvcResourceUsageWithDefaults instantiates a new RegistrySvcResourceUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcResourceUsageWithDefaults() *RegistrySvcResourceUsage {
	this := RegistrySvcResourceUsage{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *RegistrySvcResourceUsage) GetCpu() RegistrySvcUsage {
	if o == nil || IsNil(o.Cpu) {
		var ret RegistrySvcUsage
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcResourceUsage) GetCpuOk() (*RegistrySvcUsage, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *RegistrySvcResourceUsage) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given RegistrySvcUsage and assigns it to the Cpu field.
func (o *RegistrySvcResourceUsage) SetCpu(v RegistrySvcUsage) {
	o.Cpu = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *RegistrySvcResourceUsage) GetDisk() RegistrySvcUsage {
	if o == nil || IsNil(o.Disk) {
		var ret RegistrySvcUsage
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcResourceUsage) GetDiskOk() (*RegistrySvcUsage, bool) {
	if o == nil || IsNil(o.Disk) {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *RegistrySvcResourceUsage) HasDisk() bool {
	if o != nil && !IsNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given RegistrySvcUsage and assigns it to the Disk field.
func (o *RegistrySvcResourceUsage) SetDisk(v RegistrySvcUsage) {
	o.Disk = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *RegistrySvcResourceUsage) GetMemory() RegistrySvcUsage {
	if o == nil || IsNil(o.Memory) {
		var ret RegistrySvcUsage
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcResourceUsage) GetMemoryOk() (*RegistrySvcUsage, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *RegistrySvcResourceUsage) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given RegistrySvcUsage and assigns it to the Memory field.
func (o *RegistrySvcResourceUsage) SetMemory(v RegistrySvcUsage) {
	o.Memory = &v
}

func (o RegistrySvcResourceUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcResourceUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Disk) {
		toSerialize["disk"] = o.Disk
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return toSerialize, nil
}

type NullableRegistrySvcResourceUsage struct {
	value *RegistrySvcResourceUsage
	isSet bool
}

func (v NullableRegistrySvcResourceUsage) Get() *RegistrySvcResourceUsage {
	return v.value
}

func (v *NullableRegistrySvcResourceUsage) Set(val *RegistrySvcResourceUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcResourceUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcResourceUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcResourceUsage(val *RegistrySvcResourceUsage) *NullableRegistrySvcResourceUsage {
	return &NullableRegistrySvcResourceUsage{value: val, isSet: true}
}

func (v NullableRegistrySvcResourceUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcResourceUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


