/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.39
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DeploySvcDeploymentStrategy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploySvcDeploymentStrategy{}

// DeploySvcDeploymentStrategy struct for DeploySvcDeploymentStrategy
type DeploySvcDeploymentStrategy struct {
	// Max extra replicas during update
	MaxSurge *int32 `json:"maxSurge,omitempty"`
	// Max unavailable replicas during update
	MaxUnavailable *int32 `json:"maxUnavailable,omitempty"`
	// Deployment strategy type (RollingUpdate, Recreate, etc.)
	Type *DeploySvcStrategyType `json:"type,omitempty"`
}

// NewDeploySvcDeploymentStrategy instantiates a new DeploySvcDeploymentStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploySvcDeploymentStrategy() *DeploySvcDeploymentStrategy {
	this := DeploySvcDeploymentStrategy{}
	return &this
}

// NewDeploySvcDeploymentStrategyWithDefaults instantiates a new DeploySvcDeploymentStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploySvcDeploymentStrategyWithDefaults() *DeploySvcDeploymentStrategy {
	this := DeploySvcDeploymentStrategy{}
	return &this
}

// GetMaxSurge returns the MaxSurge field value if set, zero value otherwise.
func (o *DeploySvcDeploymentStrategy) GetMaxSurge() int32 {
	if o == nil || IsNil(o.MaxSurge) {
		var ret int32
		return ret
	}
	return *o.MaxSurge
}

// GetMaxSurgeOk returns a tuple with the MaxSurge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeploymentStrategy) GetMaxSurgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSurge) {
		return nil, false
	}
	return o.MaxSurge, true
}

// HasMaxSurge returns a boolean if a field has been set.
func (o *DeploySvcDeploymentStrategy) HasMaxSurge() bool {
	if o != nil && !IsNil(o.MaxSurge) {
		return true
	}

	return false
}

// SetMaxSurge gets a reference to the given int32 and assigns it to the MaxSurge field.
func (o *DeploySvcDeploymentStrategy) SetMaxSurge(v int32) {
	o.MaxSurge = &v
}

// GetMaxUnavailable returns the MaxUnavailable field value if set, zero value otherwise.
func (o *DeploySvcDeploymentStrategy) GetMaxUnavailable() int32 {
	if o == nil || IsNil(o.MaxUnavailable) {
		var ret int32
		return ret
	}
	return *o.MaxUnavailable
}

// GetMaxUnavailableOk returns a tuple with the MaxUnavailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeploymentStrategy) GetMaxUnavailableOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUnavailable) {
		return nil, false
	}
	return o.MaxUnavailable, true
}

// HasMaxUnavailable returns a boolean if a field has been set.
func (o *DeploySvcDeploymentStrategy) HasMaxUnavailable() bool {
	if o != nil && !IsNil(o.MaxUnavailable) {
		return true
	}

	return false
}

// SetMaxUnavailable gets a reference to the given int32 and assigns it to the MaxUnavailable field.
func (o *DeploySvcDeploymentStrategy) SetMaxUnavailable(v int32) {
	o.MaxUnavailable = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeploySvcDeploymentStrategy) GetType() DeploySvcStrategyType {
	if o == nil || IsNil(o.Type) {
		var ret DeploySvcStrategyType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeploymentStrategy) GetTypeOk() (*DeploySvcStrategyType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeploySvcDeploymentStrategy) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DeploySvcStrategyType and assigns it to the Type field.
func (o *DeploySvcDeploymentStrategy) SetType(v DeploySvcStrategyType) {
	o.Type = &v
}

func (o DeploySvcDeploymentStrategy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploySvcDeploymentStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxSurge) {
		toSerialize["maxSurge"] = o.MaxSurge
	}
	if !IsNil(o.MaxUnavailable) {
		toSerialize["maxUnavailable"] = o.MaxUnavailable
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDeploySvcDeploymentStrategy struct {
	value *DeploySvcDeploymentStrategy
	isSet bool
}

func (v NullableDeploySvcDeploymentStrategy) Get() *DeploySvcDeploymentStrategy {
	return v.value
}

func (v *NullableDeploySvcDeploymentStrategy) Set(val *DeploySvcDeploymentStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploySvcDeploymentStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySvcDeploymentStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySvcDeploymentStrategy(val *DeploySvcDeploymentStrategy) *NullableDeploySvcDeploymentStrategy {
	return &NullableDeploySvcDeploymentStrategy{value: val, isSet: true}
}

func (v NullableDeploySvcDeploymentStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySvcDeploymentStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


