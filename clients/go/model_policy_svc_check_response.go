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
	"bytes"
	"fmt"
)

// checks if the PolicySvcCheckResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicySvcCheckResponse{}

// PolicySvcCheckResponse struct for PolicySvcCheckResponse
type PolicySvcCheckResponse struct {
	Allowed bool `json:"allowed"`
}

type _PolicySvcCheckResponse PolicySvcCheckResponse

// NewPolicySvcCheckResponse instantiates a new PolicySvcCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicySvcCheckResponse(allowed bool) *PolicySvcCheckResponse {
	this := PolicySvcCheckResponse{}
	this.Allowed = allowed
	return &this
}

// NewPolicySvcCheckResponseWithDefaults instantiates a new PolicySvcCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicySvcCheckResponseWithDefaults() *PolicySvcCheckResponse {
	this := PolicySvcCheckResponse{}
	return &this
}

// GetAllowed returns the Allowed field value
func (o *PolicySvcCheckResponse) GetAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value
// and a boolean to check if the value has been set.
func (o *PolicySvcCheckResponse) GetAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allowed, true
}

// SetAllowed sets field value
func (o *PolicySvcCheckResponse) SetAllowed(v bool) {
	o.Allowed = v
}

func (o PolicySvcCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicySvcCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowed"] = o.Allowed
	return toSerialize, nil
}

func (o *PolicySvcCheckResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allowed",
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

	varPolicySvcCheckResponse := _PolicySvcCheckResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPolicySvcCheckResponse)

	if err != nil {
		return err
	}

	*o = PolicySvcCheckResponse(varPolicySvcCheckResponse)

	return err
}

type NullablePolicySvcCheckResponse struct {
	value *PolicySvcCheckResponse
	isSet bool
}

func (v NullablePolicySvcCheckResponse) Get() *PolicySvcCheckResponse {
	return v.value
}

func (v *NullablePolicySvcCheckResponse) Set(val *PolicySvcCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySvcCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySvcCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySvcCheckResponse(val *PolicySvcCheckResponse) *NullablePolicySvcCheckResponse {
	return &NullablePolicySvcCheckResponse{value: val, isSet: true}
}

func (v NullablePolicySvcCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySvcCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


