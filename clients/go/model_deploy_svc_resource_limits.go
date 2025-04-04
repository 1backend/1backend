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

// checks if the DeploySvcResourceLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploySvcResourceLimits{}

// DeploySvcResourceLimits struct for DeploySvcResourceLimits
type DeploySvcResourceLimits struct {
	// CPU limit, e.g., \"500m\" for 0.5 cores
	Cpu *string `json:"cpu,omitempty"`
	// Memory limit, e.g., \"128Mi\"
	Memory *string `json:"memory,omitempty"`
	// Optional: GPU VRAM requirement, e.g., \"48GB\"
	Vram *string `json:"vram,omitempty"`
}

// NewDeploySvcResourceLimits instantiates a new DeploySvcResourceLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploySvcResourceLimits() *DeploySvcResourceLimits {
	this := DeploySvcResourceLimits{}
	return &this
}

// NewDeploySvcResourceLimitsWithDefaults instantiates a new DeploySvcResourceLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploySvcResourceLimitsWithDefaults() *DeploySvcResourceLimits {
	this := DeploySvcResourceLimits{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *DeploySvcResourceLimits) GetCpu() string {
	if o == nil || IsNil(o.Cpu) {
		var ret string
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcResourceLimits) GetCpuOk() (*string, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *DeploySvcResourceLimits) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given string and assigns it to the Cpu field.
func (o *DeploySvcResourceLimits) SetCpu(v string) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *DeploySvcResourceLimits) GetMemory() string {
	if o == nil || IsNil(o.Memory) {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcResourceLimits) GetMemoryOk() (*string, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *DeploySvcResourceLimits) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the Memory field.
func (o *DeploySvcResourceLimits) SetMemory(v string) {
	o.Memory = &v
}

// GetVram returns the Vram field value if set, zero value otherwise.
func (o *DeploySvcResourceLimits) GetVram() string {
	if o == nil || IsNil(o.Vram) {
		var ret string
		return ret
	}
	return *o.Vram
}

// GetVramOk returns a tuple with the Vram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcResourceLimits) GetVramOk() (*string, bool) {
	if o == nil || IsNil(o.Vram) {
		return nil, false
	}
	return o.Vram, true
}

// HasVram returns a boolean if a field has been set.
func (o *DeploySvcResourceLimits) HasVram() bool {
	if o != nil && !IsNil(o.Vram) {
		return true
	}

	return false
}

// SetVram gets a reference to the given string and assigns it to the Vram field.
func (o *DeploySvcResourceLimits) SetVram(v string) {
	o.Vram = &v
}

func (o DeploySvcResourceLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploySvcResourceLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.Vram) {
		toSerialize["vram"] = o.Vram
	}
	return toSerialize, nil
}

type NullableDeploySvcResourceLimits struct {
	value *DeploySvcResourceLimits
	isSet bool
}

func (v NullableDeploySvcResourceLimits) Get() *DeploySvcResourceLimits {
	return v.value
}

func (v *NullableDeploySvcResourceLimits) Set(val *DeploySvcResourceLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploySvcResourceLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySvcResourceLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySvcResourceLimits(val *DeploySvcResourceLimits) *NullableDeploySvcResourceLimits {
	return &NullableDeploySvcResourceLimits{value: val, isSet: true}
}

func (v NullableDeploySvcResourceLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySvcResourceLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


