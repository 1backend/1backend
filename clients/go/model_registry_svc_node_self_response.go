/*
1Backend

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RegistrySvcNodeSelfResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcNodeSelfResponse{}

// RegistrySvcNodeSelfResponse struct for RegistrySvcNodeSelfResponse
type RegistrySvcNodeSelfResponse struct {
	Node RegistrySvcNode `json:"node"`
}

type _RegistrySvcNodeSelfResponse RegistrySvcNodeSelfResponse

// NewRegistrySvcNodeSelfResponse instantiates a new RegistrySvcNodeSelfResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcNodeSelfResponse(node RegistrySvcNode) *RegistrySvcNodeSelfResponse {
	this := RegistrySvcNodeSelfResponse{}
	this.Node = node
	return &this
}

// NewRegistrySvcNodeSelfResponseWithDefaults instantiates a new RegistrySvcNodeSelfResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcNodeSelfResponseWithDefaults() *RegistrySvcNodeSelfResponse {
	this := RegistrySvcNodeSelfResponse{}
	return &this
}

// GetNode returns the Node field value
func (o *RegistrySvcNodeSelfResponse) GetNode() RegistrySvcNode {
	if o == nil {
		var ret RegistrySvcNode
		return ret
	}

	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcNodeSelfResponse) GetNodeOk() (*RegistrySvcNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value
func (o *RegistrySvcNodeSelfResponse) SetNode(v RegistrySvcNode) {
	o.Node = v
}

func (o RegistrySvcNodeSelfResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcNodeSelfResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node"] = o.Node
	return toSerialize, nil
}

func (o *RegistrySvcNodeSelfResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRegistrySvcNodeSelfResponse := _RegistrySvcNodeSelfResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrySvcNodeSelfResponse)

	if err != nil {
		return err
	}

	*o = RegistrySvcNodeSelfResponse(varRegistrySvcNodeSelfResponse)

	return err
}

type NullableRegistrySvcNodeSelfResponse struct {
	value *RegistrySvcNodeSelfResponse
	isSet bool
}

func (v NullableRegistrySvcNodeSelfResponse) Get() *RegistrySvcNodeSelfResponse {
	return v.value
}

func (v *NullableRegistrySvcNodeSelfResponse) Set(val *RegistrySvcNodeSelfResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcNodeSelfResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcNodeSelfResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcNodeSelfResponse(val *RegistrySvcNodeSelfResponse) *NullableRegistrySvcNodeSelfResponse {
	return &NullableRegistrySvcNodeSelfResponse{value: val, isSet: true}
}

func (v NullableRegistrySvcNodeSelfResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcNodeSelfResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


