/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.34
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DeploySvcAutoScalingConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploySvcAutoScalingConfig{}

// DeploySvcAutoScalingConfig struct for DeploySvcAutoScalingConfig
type DeploySvcAutoScalingConfig struct {
	// CPU usage threshold for scaling (as a percentage)
	CpuThreshold *int32 `json:"cpuThreshold,omitempty"`
	// Maximum number of replicas to run
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`
	// Minimum number of replicas to run
	MinReplicas *int32 `json:"minReplicas,omitempty"`
}

// NewDeploySvcAutoScalingConfig instantiates a new DeploySvcAutoScalingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploySvcAutoScalingConfig() *DeploySvcAutoScalingConfig {
	this := DeploySvcAutoScalingConfig{}
	return &this
}

// NewDeploySvcAutoScalingConfigWithDefaults instantiates a new DeploySvcAutoScalingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploySvcAutoScalingConfigWithDefaults() *DeploySvcAutoScalingConfig {
	this := DeploySvcAutoScalingConfig{}
	return &this
}

// GetCpuThreshold returns the CpuThreshold field value if set, zero value otherwise.
func (o *DeploySvcAutoScalingConfig) GetCpuThreshold() int32 {
	if o == nil || IsNil(o.CpuThreshold) {
		var ret int32
		return ret
	}
	return *o.CpuThreshold
}

// GetCpuThresholdOk returns a tuple with the CpuThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcAutoScalingConfig) GetCpuThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.CpuThreshold) {
		return nil, false
	}
	return o.CpuThreshold, true
}

// HasCpuThreshold returns a boolean if a field has been set.
func (o *DeploySvcAutoScalingConfig) HasCpuThreshold() bool {
	if o != nil && !IsNil(o.CpuThreshold) {
		return true
	}

	return false
}

// SetCpuThreshold gets a reference to the given int32 and assigns it to the CpuThreshold field.
func (o *DeploySvcAutoScalingConfig) SetCpuThreshold(v int32) {
	o.CpuThreshold = &v
}

// GetMaxReplicas returns the MaxReplicas field value if set, zero value otherwise.
func (o *DeploySvcAutoScalingConfig) GetMaxReplicas() int32 {
	if o == nil || IsNil(o.MaxReplicas) {
		var ret int32
		return ret
	}
	return *o.MaxReplicas
}

// GetMaxReplicasOk returns a tuple with the MaxReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcAutoScalingConfig) GetMaxReplicasOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxReplicas) {
		return nil, false
	}
	return o.MaxReplicas, true
}

// HasMaxReplicas returns a boolean if a field has been set.
func (o *DeploySvcAutoScalingConfig) HasMaxReplicas() bool {
	if o != nil && !IsNil(o.MaxReplicas) {
		return true
	}

	return false
}

// SetMaxReplicas gets a reference to the given int32 and assigns it to the MaxReplicas field.
func (o *DeploySvcAutoScalingConfig) SetMaxReplicas(v int32) {
	o.MaxReplicas = &v
}

// GetMinReplicas returns the MinReplicas field value if set, zero value otherwise.
func (o *DeploySvcAutoScalingConfig) GetMinReplicas() int32 {
	if o == nil || IsNil(o.MinReplicas) {
		var ret int32
		return ret
	}
	return *o.MinReplicas
}

// GetMinReplicasOk returns a tuple with the MinReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcAutoScalingConfig) GetMinReplicasOk() (*int32, bool) {
	if o == nil || IsNil(o.MinReplicas) {
		return nil, false
	}
	return o.MinReplicas, true
}

// HasMinReplicas returns a boolean if a field has been set.
func (o *DeploySvcAutoScalingConfig) HasMinReplicas() bool {
	if o != nil && !IsNil(o.MinReplicas) {
		return true
	}

	return false
}

// SetMinReplicas gets a reference to the given int32 and assigns it to the MinReplicas field.
func (o *DeploySvcAutoScalingConfig) SetMinReplicas(v int32) {
	o.MinReplicas = &v
}

func (o DeploySvcAutoScalingConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploySvcAutoScalingConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpuThreshold) {
		toSerialize["cpuThreshold"] = o.CpuThreshold
	}
	if !IsNil(o.MaxReplicas) {
		toSerialize["maxReplicas"] = o.MaxReplicas
	}
	if !IsNil(o.MinReplicas) {
		toSerialize["minReplicas"] = o.MinReplicas
	}
	return toSerialize, nil
}

type NullableDeploySvcAutoScalingConfig struct {
	value *DeploySvcAutoScalingConfig
	isSet bool
}

func (v NullableDeploySvcAutoScalingConfig) Get() *DeploySvcAutoScalingConfig {
	return v.value
}

func (v *NullableDeploySvcAutoScalingConfig) Set(val *DeploySvcAutoScalingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploySvcAutoScalingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySvcAutoScalingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySvcAutoScalingConfig(val *DeploySvcAutoScalingConfig) *NullableDeploySvcAutoScalingConfig {
	return &NullableDeploySvcAutoScalingConfig{value: val, isSet: true}
}

func (v NullableDeploySvcAutoScalingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySvcAutoScalingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


