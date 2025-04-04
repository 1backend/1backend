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

// checks if the PolicySvcUpsertInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicySvcUpsertInstanceRequest{}

// PolicySvcUpsertInstanceRequest struct for PolicySvcUpsertInstanceRequest
type PolicySvcUpsertInstanceRequest struct {
	Instance *PolicySvcInstance `json:"instance,omitempty"`
}

// NewPolicySvcUpsertInstanceRequest instantiates a new PolicySvcUpsertInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicySvcUpsertInstanceRequest() *PolicySvcUpsertInstanceRequest {
	this := PolicySvcUpsertInstanceRequest{}
	return &this
}

// NewPolicySvcUpsertInstanceRequestWithDefaults instantiates a new PolicySvcUpsertInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicySvcUpsertInstanceRequestWithDefaults() *PolicySvcUpsertInstanceRequest {
	this := PolicySvcUpsertInstanceRequest{}
	return &this
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *PolicySvcUpsertInstanceRequest) GetInstance() PolicySvcInstance {
	if o == nil || IsNil(o.Instance) {
		var ret PolicySvcInstance
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySvcUpsertInstanceRequest) GetInstanceOk() (*PolicySvcInstance, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *PolicySvcUpsertInstanceRequest) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given PolicySvcInstance and assigns it to the Instance field.
func (o *PolicySvcUpsertInstanceRequest) SetInstance(v PolicySvcInstance) {
	o.Instance = &v
}

func (o PolicySvcUpsertInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicySvcUpsertInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	return toSerialize, nil
}

type NullablePolicySvcUpsertInstanceRequest struct {
	value *PolicySvcUpsertInstanceRequest
	isSet bool
}

func (v NullablePolicySvcUpsertInstanceRequest) Get() *PolicySvcUpsertInstanceRequest {
	return v.value
}

func (v *NullablePolicySvcUpsertInstanceRequest) Set(val *PolicySvcUpsertInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySvcUpsertInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySvcUpsertInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySvcUpsertInstanceRequest(val *PolicySvcUpsertInstanceRequest) *NullablePolicySvcUpsertInstanceRequest {
	return &NullablePolicySvcUpsertInstanceRequest{value: val, isSet: true}
}

func (v NullablePolicySvcUpsertInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySvcUpsertInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


