/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PolicySvcBlocklistParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicySvcBlocklistParameters{}

// PolicySvcBlocklistParameters struct for PolicySvcBlocklistParameters
type PolicySvcBlocklistParameters struct {
	BlockedIPs []string `json:"blockedIPs,omitempty"`
}

// NewPolicySvcBlocklistParameters instantiates a new PolicySvcBlocklistParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicySvcBlocklistParameters() *PolicySvcBlocklistParameters {
	this := PolicySvcBlocklistParameters{}
	return &this
}

// NewPolicySvcBlocklistParametersWithDefaults instantiates a new PolicySvcBlocklistParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicySvcBlocklistParametersWithDefaults() *PolicySvcBlocklistParameters {
	this := PolicySvcBlocklistParameters{}
	return &this
}

// GetBlockedIPs returns the BlockedIPs field value if set, zero value otherwise.
func (o *PolicySvcBlocklistParameters) GetBlockedIPs() []string {
	if o == nil || IsNil(o.BlockedIPs) {
		var ret []string
		return ret
	}
	return o.BlockedIPs
}

// GetBlockedIPsOk returns a tuple with the BlockedIPs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySvcBlocklistParameters) GetBlockedIPsOk() ([]string, bool) {
	if o == nil || IsNil(o.BlockedIPs) {
		return nil, false
	}
	return o.BlockedIPs, true
}

// HasBlockedIPs returns a boolean if a field has been set.
func (o *PolicySvcBlocklistParameters) HasBlockedIPs() bool {
	if o != nil && !IsNil(o.BlockedIPs) {
		return true
	}

	return false
}

// SetBlockedIPs gets a reference to the given []string and assigns it to the BlockedIPs field.
func (o *PolicySvcBlocklistParameters) SetBlockedIPs(v []string) {
	o.BlockedIPs = v
}

func (o PolicySvcBlocklistParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicySvcBlocklistParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockedIPs) {
		toSerialize["blockedIPs"] = o.BlockedIPs
	}
	return toSerialize, nil
}

type NullablePolicySvcBlocklistParameters struct {
	value *PolicySvcBlocklistParameters
	isSet bool
}

func (v NullablePolicySvcBlocklistParameters) Get() *PolicySvcBlocklistParameters {
	return v.value
}

func (v *NullablePolicySvcBlocklistParameters) Set(val *PolicySvcBlocklistParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySvcBlocklistParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySvcBlocklistParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySvcBlocklistParameters(val *PolicySvcBlocklistParameters) *NullablePolicySvcBlocklistParameters {
	return &NullablePolicySvcBlocklistParameters{value: val, isSet: true}
}

func (v NullablePolicySvcBlocklistParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySvcBlocklistParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


