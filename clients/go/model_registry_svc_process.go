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

// checks if the RegistrySvcProcess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcProcess{}

// RegistrySvcProcess struct for RegistrySvcProcess
type RegistrySvcProcess struct {
	MemoryUsage *int32 `json:"memoryUsage,omitempty"`
	Pid *int32 `json:"pid,omitempty"`
	ProcessName *string `json:"processName,omitempty"`
}

// NewRegistrySvcProcess instantiates a new RegistrySvcProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcProcess() *RegistrySvcProcess {
	this := RegistrySvcProcess{}
	return &this
}

// NewRegistrySvcProcessWithDefaults instantiates a new RegistrySvcProcess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcProcessWithDefaults() *RegistrySvcProcess {
	this := RegistrySvcProcess{}
	return &this
}

// GetMemoryUsage returns the MemoryUsage field value if set, zero value otherwise.
func (o *RegistrySvcProcess) GetMemoryUsage() int32 {
	if o == nil || IsNil(o.MemoryUsage) {
		var ret int32
		return ret
	}
	return *o.MemoryUsage
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcProcess) GetMemoryUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.MemoryUsage) {
		return nil, false
	}
	return o.MemoryUsage, true
}

// HasMemoryUsage returns a boolean if a field has been set.
func (o *RegistrySvcProcess) HasMemoryUsage() bool {
	if o != nil && !IsNil(o.MemoryUsage) {
		return true
	}

	return false
}

// SetMemoryUsage gets a reference to the given int32 and assigns it to the MemoryUsage field.
func (o *RegistrySvcProcess) SetMemoryUsage(v int32) {
	o.MemoryUsage = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *RegistrySvcProcess) GetPid() int32 {
	if o == nil || IsNil(o.Pid) {
		var ret int32
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcProcess) GetPidOk() (*int32, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *RegistrySvcProcess) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given int32 and assigns it to the Pid field.
func (o *RegistrySvcProcess) SetPid(v int32) {
	o.Pid = &v
}

// GetProcessName returns the ProcessName field value if set, zero value otherwise.
func (o *RegistrySvcProcess) GetProcessName() string {
	if o == nil || IsNil(o.ProcessName) {
		var ret string
		return ret
	}
	return *o.ProcessName
}

// GetProcessNameOk returns a tuple with the ProcessName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcProcess) GetProcessNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessName) {
		return nil, false
	}
	return o.ProcessName, true
}

// HasProcessName returns a boolean if a field has been set.
func (o *RegistrySvcProcess) HasProcessName() bool {
	if o != nil && !IsNil(o.ProcessName) {
		return true
	}

	return false
}

// SetProcessName gets a reference to the given string and assigns it to the ProcessName field.
func (o *RegistrySvcProcess) SetProcessName(v string) {
	o.ProcessName = &v
}

func (o RegistrySvcProcess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcProcess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemoryUsage) {
		toSerialize["memoryUsage"] = o.MemoryUsage
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.ProcessName) {
		toSerialize["processName"] = o.ProcessName
	}
	return toSerialize, nil
}

type NullableRegistrySvcProcess struct {
	value *RegistrySvcProcess
	isSet bool
}

func (v NullableRegistrySvcProcess) Get() *RegistrySvcProcess {
	return v.value
}

func (v *NullableRegistrySvcProcess) Set(val *RegistrySvcProcess) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcProcess) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcProcess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcProcess(val *RegistrySvcProcess) *NullableRegistrySvcProcess {
	return &NullableRegistrySvcProcess{value: val, isSet: true}
}

func (v NullableRegistrySvcProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcProcess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


