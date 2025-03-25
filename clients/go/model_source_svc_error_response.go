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

// checks if the SourceSvcErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceSvcErrorResponse{}

// SourceSvcErrorResponse struct for SourceSvcErrorResponse
type SourceSvcErrorResponse struct {
	Error *string `json:"error,omitempty"`
}

// NewSourceSvcErrorResponse instantiates a new SourceSvcErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceSvcErrorResponse() *SourceSvcErrorResponse {
	this := SourceSvcErrorResponse{}
	return &this
}

// NewSourceSvcErrorResponseWithDefaults instantiates a new SourceSvcErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceSvcErrorResponseWithDefaults() *SourceSvcErrorResponse {
	this := SourceSvcErrorResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SourceSvcErrorResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSvcErrorResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SourceSvcErrorResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *SourceSvcErrorResponse) SetError(v string) {
	o.Error = &v
}

func (o SourceSvcErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceSvcErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableSourceSvcErrorResponse struct {
	value *SourceSvcErrorResponse
	isSet bool
}

func (v NullableSourceSvcErrorResponse) Get() *SourceSvcErrorResponse {
	return v.value
}

func (v *NullableSourceSvcErrorResponse) Set(val *SourceSvcErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceSvcErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceSvcErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceSvcErrorResponse(val *SourceSvcErrorResponse) *NullableSourceSvcErrorResponse {
	return &NullableSourceSvcErrorResponse{value: val, isSet: true}
}

func (v NullableSourceSvcErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceSvcErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


