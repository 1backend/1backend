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
)

// checks if the PolicySvcErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicySvcErrorResponse{}

// PolicySvcErrorResponse struct for PolicySvcErrorResponse
type PolicySvcErrorResponse struct {
	Error *string `json:"error,omitempty"`
}

// NewPolicySvcErrorResponse instantiates a new PolicySvcErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicySvcErrorResponse() *PolicySvcErrorResponse {
	this := PolicySvcErrorResponse{}
	return &this
}

// NewPolicySvcErrorResponseWithDefaults instantiates a new PolicySvcErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicySvcErrorResponseWithDefaults() *PolicySvcErrorResponse {
	this := PolicySvcErrorResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *PolicySvcErrorResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySvcErrorResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *PolicySvcErrorResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *PolicySvcErrorResponse) SetError(v string) {
	o.Error = &v
}

func (o PolicySvcErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicySvcErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullablePolicySvcErrorResponse struct {
	value *PolicySvcErrorResponse
	isSet bool
}

func (v NullablePolicySvcErrorResponse) Get() *PolicySvcErrorResponse {
	return v.value
}

func (v *NullablePolicySvcErrorResponse) Set(val *PolicySvcErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySvcErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySvcErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySvcErrorResponse(val *PolicySvcErrorResponse) *NullablePolicySvcErrorResponse {
	return &NullablePolicySvcErrorResponse{value: val, isSet: true}
}

func (v NullablePolicySvcErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySvcErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


