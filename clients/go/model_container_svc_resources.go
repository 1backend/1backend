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

// checks if the ContainerSvcResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcResources{}

// ContainerSvcResources struct for ContainerSvcResources
type ContainerSvcResources struct {
	// CPU cores allocated to the container (e.g., 0.5 = 500m, 2 = 2 cores).
	Cpu *float32 `json:"cpu,omitempty"`
	// Disk space allocated to the container in megabytes.
	DiskMB *int32 `json:"diskMB,omitempty"`
	// Memory allocated to the container in megabytes.
	MemoryMB *int32 `json:"memoryMB,omitempty"`
}

// NewContainerSvcResources instantiates a new ContainerSvcResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcResources() *ContainerSvcResources {
	this := ContainerSvcResources{}
	return &this
}

// NewContainerSvcResourcesWithDefaults instantiates a new ContainerSvcResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcResourcesWithDefaults() *ContainerSvcResources {
	this := ContainerSvcResources{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ContainerSvcResources) GetCpu() float32 {
	if o == nil || IsNil(o.Cpu) {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcResources) GetCpuOk() (*float32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ContainerSvcResources) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *ContainerSvcResources) SetCpu(v float32) {
	o.Cpu = &v
}

// GetDiskMB returns the DiskMB field value if set, zero value otherwise.
func (o *ContainerSvcResources) GetDiskMB() int32 {
	if o == nil || IsNil(o.DiskMB) {
		var ret int32
		return ret
	}
	return *o.DiskMB
}

// GetDiskMBOk returns a tuple with the DiskMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcResources) GetDiskMBOk() (*int32, bool) {
	if o == nil || IsNil(o.DiskMB) {
		return nil, false
	}
	return o.DiskMB, true
}

// HasDiskMB returns a boolean if a field has been set.
func (o *ContainerSvcResources) HasDiskMB() bool {
	if o != nil && !IsNil(o.DiskMB) {
		return true
	}

	return false
}

// SetDiskMB gets a reference to the given int32 and assigns it to the DiskMB field.
func (o *ContainerSvcResources) SetDiskMB(v int32) {
	o.DiskMB = &v
}

// GetMemoryMB returns the MemoryMB field value if set, zero value otherwise.
func (o *ContainerSvcResources) GetMemoryMB() int32 {
	if o == nil || IsNil(o.MemoryMB) {
		var ret int32
		return ret
	}
	return *o.MemoryMB
}

// GetMemoryMBOk returns a tuple with the MemoryMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcResources) GetMemoryMBOk() (*int32, bool) {
	if o == nil || IsNil(o.MemoryMB) {
		return nil, false
	}
	return o.MemoryMB, true
}

// HasMemoryMB returns a boolean if a field has been set.
func (o *ContainerSvcResources) HasMemoryMB() bool {
	if o != nil && !IsNil(o.MemoryMB) {
		return true
	}

	return false
}

// SetMemoryMB gets a reference to the given int32 and assigns it to the MemoryMB field.
func (o *ContainerSvcResources) SetMemoryMB(v int32) {
	o.MemoryMB = &v
}

func (o ContainerSvcResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.DiskMB) {
		toSerialize["diskMB"] = o.DiskMB
	}
	if !IsNil(o.MemoryMB) {
		toSerialize["memoryMB"] = o.MemoryMB
	}
	return toSerialize, nil
}

type NullableContainerSvcResources struct {
	value *ContainerSvcResources
	isSet bool
}

func (v NullableContainerSvcResources) Get() *ContainerSvcResources {
	return v.value
}

func (v *NullableContainerSvcResources) Set(val *ContainerSvcResources) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcResources) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcResources(val *ContainerSvcResources) *NullableContainerSvcResources {
	return &NullableContainerSvcResources{value: val, isSet: true}
}

func (v NullableContainerSvcResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


