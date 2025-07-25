/*
1Backend

AI-native microservices platform.

API version: 0.7.6
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PolicySvcParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicySvcParameters{}

// PolicySvcParameters struct for PolicySvcParameters
type PolicySvcParameters struct {
	Blocklist *PolicySvcBlocklistParameters `json:"blocklist,omitempty"`
	RateLimit *PolicySvcRateLimitParameters `json:"rateLimit,omitempty"`
}

// NewPolicySvcParameters instantiates a new PolicySvcParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicySvcParameters() *PolicySvcParameters {
	this := PolicySvcParameters{}
	return &this
}

// NewPolicySvcParametersWithDefaults instantiates a new PolicySvcParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicySvcParametersWithDefaults() *PolicySvcParameters {
	this := PolicySvcParameters{}
	return &this
}

// GetBlocklist returns the Blocklist field value if set, zero value otherwise.
func (o *PolicySvcParameters) GetBlocklist() PolicySvcBlocklistParameters {
	if o == nil || IsNil(o.Blocklist) {
		var ret PolicySvcBlocklistParameters
		return ret
	}
	return *o.Blocklist
}

// GetBlocklistOk returns a tuple with the Blocklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySvcParameters) GetBlocklistOk() (*PolicySvcBlocklistParameters, bool) {
	if o == nil || IsNil(o.Blocklist) {
		return nil, false
	}
	return o.Blocklist, true
}

// HasBlocklist returns a boolean if a field has been set.
func (o *PolicySvcParameters) HasBlocklist() bool {
	if o != nil && !IsNil(o.Blocklist) {
		return true
	}

	return false
}

// SetBlocklist gets a reference to the given PolicySvcBlocklistParameters and assigns it to the Blocklist field.
func (o *PolicySvcParameters) SetBlocklist(v PolicySvcBlocklistParameters) {
	o.Blocklist = &v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *PolicySvcParameters) GetRateLimit() PolicySvcRateLimitParameters {
	if o == nil || IsNil(o.RateLimit) {
		var ret PolicySvcRateLimitParameters
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySvcParameters) GetRateLimitOk() (*PolicySvcRateLimitParameters, bool) {
	if o == nil || IsNil(o.RateLimit) {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *PolicySvcParameters) HasRateLimit() bool {
	if o != nil && !IsNil(o.RateLimit) {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given PolicySvcRateLimitParameters and assigns it to the RateLimit field.
func (o *PolicySvcParameters) SetRateLimit(v PolicySvcRateLimitParameters) {
	o.RateLimit = &v
}

func (o PolicySvcParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicySvcParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Blocklist) {
		toSerialize["blocklist"] = o.Blocklist
	}
	if !IsNil(o.RateLimit) {
		toSerialize["rateLimit"] = o.RateLimit
	}
	return toSerialize, nil
}

type NullablePolicySvcParameters struct {
	value *PolicySvcParameters
	isSet bool
}

func (v NullablePolicySvcParameters) Get() *PolicySvcParameters {
	return v.value
}

func (v *NullablePolicySvcParameters) Set(val *PolicySvcParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySvcParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySvcParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySvcParameters(val *PolicySvcParameters) *NullablePolicySvcParameters {
	return &NullablePolicySvcParameters{value: val, isSet: true}
}

func (v NullablePolicySvcParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySvcParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


