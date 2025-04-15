/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.37
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RegistrySvcGPU type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcGPU{}

// RegistrySvcGPU struct for RegistrySvcGPU
type RegistrySvcGPU struct {
	BusId *string `json:"busId,omitempty"`
	ComputeMode *string `json:"computeMode,omitempty"`
	CudaVersion *string `json:"cudaVersion,omitempty"`
	GpuUtilization *float32 `json:"gpuUtilization,omitempty"`
	// Id Node.URL + IntraNodeId
	Id *string `json:"id,omitempty"`
	IntraNodeId *int32 `json:"intraNodeId,omitempty"`
	MemoryTotal *int32 `json:"memoryTotal,omitempty"`
	MemoryUsage *int32 `json:"memoryUsage,omitempty"`
	Name *string `json:"name,omitempty"`
	PerformanceState *string `json:"performanceState,omitempty"`
	PowerCap *float32 `json:"powerCap,omitempty"`
	PowerUsage *float32 `json:"powerUsage,omitempty"`
	ProcessDetails []RegistrySvcProcess `json:"processDetails,omitempty"`
	Temperature *float32 `json:"temperature,omitempty"`
}

// NewRegistrySvcGPU instantiates a new RegistrySvcGPU object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcGPU() *RegistrySvcGPU {
	this := RegistrySvcGPU{}
	return &this
}

// NewRegistrySvcGPUWithDefaults instantiates a new RegistrySvcGPU object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcGPUWithDefaults() *RegistrySvcGPU {
	this := RegistrySvcGPU{}
	return &this
}

// GetBusId returns the BusId field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetBusId() string {
	if o == nil || IsNil(o.BusId) {
		var ret string
		return ret
	}
	return *o.BusId
}

// GetBusIdOk returns a tuple with the BusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetBusIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusId) {
		return nil, false
	}
	return o.BusId, true
}

// HasBusId returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasBusId() bool {
	if o != nil && !IsNil(o.BusId) {
		return true
	}

	return false
}

// SetBusId gets a reference to the given string and assigns it to the BusId field.
func (o *RegistrySvcGPU) SetBusId(v string) {
	o.BusId = &v
}

// GetComputeMode returns the ComputeMode field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetComputeMode() string {
	if o == nil || IsNil(o.ComputeMode) {
		var ret string
		return ret
	}
	return *o.ComputeMode
}

// GetComputeModeOk returns a tuple with the ComputeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetComputeModeOk() (*string, bool) {
	if o == nil || IsNil(o.ComputeMode) {
		return nil, false
	}
	return o.ComputeMode, true
}

// HasComputeMode returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasComputeMode() bool {
	if o != nil && !IsNil(o.ComputeMode) {
		return true
	}

	return false
}

// SetComputeMode gets a reference to the given string and assigns it to the ComputeMode field.
func (o *RegistrySvcGPU) SetComputeMode(v string) {
	o.ComputeMode = &v
}

// GetCudaVersion returns the CudaVersion field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetCudaVersion() string {
	if o == nil || IsNil(o.CudaVersion) {
		var ret string
		return ret
	}
	return *o.CudaVersion
}

// GetCudaVersionOk returns a tuple with the CudaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetCudaVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CudaVersion) {
		return nil, false
	}
	return o.CudaVersion, true
}

// HasCudaVersion returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasCudaVersion() bool {
	if o != nil && !IsNil(o.CudaVersion) {
		return true
	}

	return false
}

// SetCudaVersion gets a reference to the given string and assigns it to the CudaVersion field.
func (o *RegistrySvcGPU) SetCudaVersion(v string) {
	o.CudaVersion = &v
}

// GetGpuUtilization returns the GpuUtilization field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetGpuUtilization() float32 {
	if o == nil || IsNil(o.GpuUtilization) {
		var ret float32
		return ret
	}
	return *o.GpuUtilization
}

// GetGpuUtilizationOk returns a tuple with the GpuUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetGpuUtilizationOk() (*float32, bool) {
	if o == nil || IsNil(o.GpuUtilization) {
		return nil, false
	}
	return o.GpuUtilization, true
}

// HasGpuUtilization returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasGpuUtilization() bool {
	if o != nil && !IsNil(o.GpuUtilization) {
		return true
	}

	return false
}

// SetGpuUtilization gets a reference to the given float32 and assigns it to the GpuUtilization field.
func (o *RegistrySvcGPU) SetGpuUtilization(v float32) {
	o.GpuUtilization = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegistrySvcGPU) SetId(v string) {
	o.Id = &v
}

// GetIntraNodeId returns the IntraNodeId field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetIntraNodeId() int32 {
	if o == nil || IsNil(o.IntraNodeId) {
		var ret int32
		return ret
	}
	return *o.IntraNodeId
}

// GetIntraNodeIdOk returns a tuple with the IntraNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetIntraNodeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.IntraNodeId) {
		return nil, false
	}
	return o.IntraNodeId, true
}

// HasIntraNodeId returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasIntraNodeId() bool {
	if o != nil && !IsNil(o.IntraNodeId) {
		return true
	}

	return false
}

// SetIntraNodeId gets a reference to the given int32 and assigns it to the IntraNodeId field.
func (o *RegistrySvcGPU) SetIntraNodeId(v int32) {
	o.IntraNodeId = &v
}

// GetMemoryTotal returns the MemoryTotal field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetMemoryTotal() int32 {
	if o == nil || IsNil(o.MemoryTotal) {
		var ret int32
		return ret
	}
	return *o.MemoryTotal
}

// GetMemoryTotalOk returns a tuple with the MemoryTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetMemoryTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.MemoryTotal) {
		return nil, false
	}
	return o.MemoryTotal, true
}

// HasMemoryTotal returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasMemoryTotal() bool {
	if o != nil && !IsNil(o.MemoryTotal) {
		return true
	}

	return false
}

// SetMemoryTotal gets a reference to the given int32 and assigns it to the MemoryTotal field.
func (o *RegistrySvcGPU) SetMemoryTotal(v int32) {
	o.MemoryTotal = &v
}

// GetMemoryUsage returns the MemoryUsage field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetMemoryUsage() int32 {
	if o == nil || IsNil(o.MemoryUsage) {
		var ret int32
		return ret
	}
	return *o.MemoryUsage
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetMemoryUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.MemoryUsage) {
		return nil, false
	}
	return o.MemoryUsage, true
}

// HasMemoryUsage returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasMemoryUsage() bool {
	if o != nil && !IsNil(o.MemoryUsage) {
		return true
	}

	return false
}

// SetMemoryUsage gets a reference to the given int32 and assigns it to the MemoryUsage field.
func (o *RegistrySvcGPU) SetMemoryUsage(v int32) {
	o.MemoryUsage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegistrySvcGPU) SetName(v string) {
	o.Name = &v
}

// GetPerformanceState returns the PerformanceState field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetPerformanceState() string {
	if o == nil || IsNil(o.PerformanceState) {
		var ret string
		return ret
	}
	return *o.PerformanceState
}

// GetPerformanceStateOk returns a tuple with the PerformanceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetPerformanceStateOk() (*string, bool) {
	if o == nil || IsNil(o.PerformanceState) {
		return nil, false
	}
	return o.PerformanceState, true
}

// HasPerformanceState returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasPerformanceState() bool {
	if o != nil && !IsNil(o.PerformanceState) {
		return true
	}

	return false
}

// SetPerformanceState gets a reference to the given string and assigns it to the PerformanceState field.
func (o *RegistrySvcGPU) SetPerformanceState(v string) {
	o.PerformanceState = &v
}

// GetPowerCap returns the PowerCap field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetPowerCap() float32 {
	if o == nil || IsNil(o.PowerCap) {
		var ret float32
		return ret
	}
	return *o.PowerCap
}

// GetPowerCapOk returns a tuple with the PowerCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetPowerCapOk() (*float32, bool) {
	if o == nil || IsNil(o.PowerCap) {
		return nil, false
	}
	return o.PowerCap, true
}

// HasPowerCap returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasPowerCap() bool {
	if o != nil && !IsNil(o.PowerCap) {
		return true
	}

	return false
}

// SetPowerCap gets a reference to the given float32 and assigns it to the PowerCap field.
func (o *RegistrySvcGPU) SetPowerCap(v float32) {
	o.PowerCap = &v
}

// GetPowerUsage returns the PowerUsage field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetPowerUsage() float32 {
	if o == nil || IsNil(o.PowerUsage) {
		var ret float32
		return ret
	}
	return *o.PowerUsage
}

// GetPowerUsageOk returns a tuple with the PowerUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetPowerUsageOk() (*float32, bool) {
	if o == nil || IsNil(o.PowerUsage) {
		return nil, false
	}
	return o.PowerUsage, true
}

// HasPowerUsage returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasPowerUsage() bool {
	if o != nil && !IsNil(o.PowerUsage) {
		return true
	}

	return false
}

// SetPowerUsage gets a reference to the given float32 and assigns it to the PowerUsage field.
func (o *RegistrySvcGPU) SetPowerUsage(v float32) {
	o.PowerUsage = &v
}

// GetProcessDetails returns the ProcessDetails field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetProcessDetails() []RegistrySvcProcess {
	if o == nil || IsNil(o.ProcessDetails) {
		var ret []RegistrySvcProcess
		return ret
	}
	return o.ProcessDetails
}

// GetProcessDetailsOk returns a tuple with the ProcessDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetProcessDetailsOk() ([]RegistrySvcProcess, bool) {
	if o == nil || IsNil(o.ProcessDetails) {
		return nil, false
	}
	return o.ProcessDetails, true
}

// HasProcessDetails returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasProcessDetails() bool {
	if o != nil && !IsNil(o.ProcessDetails) {
		return true
	}

	return false
}

// SetProcessDetails gets a reference to the given []RegistrySvcProcess and assigns it to the ProcessDetails field.
func (o *RegistrySvcGPU) SetProcessDetails(v []RegistrySvcProcess) {
	o.ProcessDetails = v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *RegistrySvcGPU) GetTemperature() float32 {
	if o == nil || IsNil(o.Temperature) {
		var ret float32
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcGPU) GetTemperatureOk() (*float32, bool) {
	if o == nil || IsNil(o.Temperature) {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *RegistrySvcGPU) HasTemperature() bool {
	if o != nil && !IsNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given float32 and assigns it to the Temperature field.
func (o *RegistrySvcGPU) SetTemperature(v float32) {
	o.Temperature = &v
}

func (o RegistrySvcGPU) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcGPU) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusId) {
		toSerialize["busId"] = o.BusId
	}
	if !IsNil(o.ComputeMode) {
		toSerialize["computeMode"] = o.ComputeMode
	}
	if !IsNil(o.CudaVersion) {
		toSerialize["cudaVersion"] = o.CudaVersion
	}
	if !IsNil(o.GpuUtilization) {
		toSerialize["gpuUtilization"] = o.GpuUtilization
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IntraNodeId) {
		toSerialize["intraNodeId"] = o.IntraNodeId
	}
	if !IsNil(o.MemoryTotal) {
		toSerialize["memoryTotal"] = o.MemoryTotal
	}
	if !IsNil(o.MemoryUsage) {
		toSerialize["memoryUsage"] = o.MemoryUsage
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PerformanceState) {
		toSerialize["performanceState"] = o.PerformanceState
	}
	if !IsNil(o.PowerCap) {
		toSerialize["powerCap"] = o.PowerCap
	}
	if !IsNil(o.PowerUsage) {
		toSerialize["powerUsage"] = o.PowerUsage
	}
	if !IsNil(o.ProcessDetails) {
		toSerialize["processDetails"] = o.ProcessDetails
	}
	if !IsNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	return toSerialize, nil
}

type NullableRegistrySvcGPU struct {
	value *RegistrySvcGPU
	isSet bool
}

func (v NullableRegistrySvcGPU) Get() *RegistrySvcGPU {
	return v.value
}

func (v *NullableRegistrySvcGPU) Set(val *RegistrySvcGPU) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcGPU) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcGPU) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcGPU(val *RegistrySvcGPU) *NullableRegistrySvcGPU {
	return &NullableRegistrySvcGPU{value: val, isSet: true}
}

func (v NullableRegistrySvcGPU) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcGPU) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


