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

// checks if the ModelSvcErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSvcErrorResponse{}

// ModelSvcErrorResponse struct for ModelSvcErrorResponse
type ModelSvcErrorResponse struct {
	Error *string `json:"error,omitempty"`
}

// NewModelSvcErrorResponse instantiates a new ModelSvcErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSvcErrorResponse() *ModelSvcErrorResponse {
	this := ModelSvcErrorResponse{}
	return &this
}

// NewModelSvcErrorResponseWithDefaults instantiates a new ModelSvcErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSvcErrorResponseWithDefaults() *ModelSvcErrorResponse {
	this := ModelSvcErrorResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ModelSvcErrorResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcErrorResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ModelSvcErrorResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ModelSvcErrorResponse) SetError(v string) {
	o.Error = &v
}

func (o ModelSvcErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSvcErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableModelSvcErrorResponse struct {
	value *ModelSvcErrorResponse
	isSet bool
}

func (v NullableModelSvcErrorResponse) Get() *ModelSvcErrorResponse {
	return v.value
}

func (v *NullableModelSvcErrorResponse) Set(val *ModelSvcErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSvcErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSvcErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSvcErrorResponse(val *ModelSvcErrorResponse) *NullableModelSvcErrorResponse {
	return &NullableModelSvcErrorResponse{value: val, isSet: true}
}

func (v NullableModelSvcErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSvcErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


