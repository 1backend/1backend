/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.30
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RegistrySvcListNodesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcListNodesRequest{}

// RegistrySvcListNodesRequest struct for RegistrySvcListNodesRequest
type RegistrySvcListNodesRequest struct {
	// Node IDs to filter on
	Ids []string `json:"ids,omitempty"`
}

// NewRegistrySvcListNodesRequest instantiates a new RegistrySvcListNodesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcListNodesRequest() *RegistrySvcListNodesRequest {
	this := RegistrySvcListNodesRequest{}
	return &this
}

// NewRegistrySvcListNodesRequestWithDefaults instantiates a new RegistrySvcListNodesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcListNodesRequestWithDefaults() *RegistrySvcListNodesRequest {
	this := RegistrySvcListNodesRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *RegistrySvcListNodesRequest) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcListNodesRequest) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *RegistrySvcListNodesRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *RegistrySvcListNodesRequest) SetIds(v []string) {
	o.Ids = v
}

func (o RegistrySvcListNodesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcListNodesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableRegistrySvcListNodesRequest struct {
	value *RegistrySvcListNodesRequest
	isSet bool
}

func (v NullableRegistrySvcListNodesRequest) Get() *RegistrySvcListNodesRequest {
	return v.value
}

func (v *NullableRegistrySvcListNodesRequest) Set(val *RegistrySvcListNodesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcListNodesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcListNodesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcListNodesRequest(val *RegistrySvcListNodesRequest) *NullableRegistrySvcListNodesRequest {
	return &NullableRegistrySvcListNodesRequest{value: val, isSet: true}
}

func (v NullableRegistrySvcListNodesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcListNodesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


