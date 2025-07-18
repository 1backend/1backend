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

// checks if the ModelSvcStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSvcStatusResponse{}

// ModelSvcStatusResponse struct for ModelSvcStatusResponse
type ModelSvcStatusResponse struct {
	Status *ModelSvcModelStatus `json:"status,omitempty"`
}

// NewModelSvcStatusResponse instantiates a new ModelSvcStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSvcStatusResponse() *ModelSvcStatusResponse {
	this := ModelSvcStatusResponse{}
	return &this
}

// NewModelSvcStatusResponseWithDefaults instantiates a new ModelSvcStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSvcStatusResponseWithDefaults() *ModelSvcStatusResponse {
	this := ModelSvcStatusResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelSvcStatusResponse) GetStatus() ModelSvcModelStatus {
	if o == nil || IsNil(o.Status) {
		var ret ModelSvcModelStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcStatusResponse) GetStatusOk() (*ModelSvcModelStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelSvcStatusResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ModelSvcModelStatus and assigns it to the Status field.
func (o *ModelSvcStatusResponse) SetStatus(v ModelSvcModelStatus) {
	o.Status = &v
}

func (o ModelSvcStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSvcStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableModelSvcStatusResponse struct {
	value *ModelSvcStatusResponse
	isSet bool
}

func (v NullableModelSvcStatusResponse) Get() *ModelSvcStatusResponse {
	return v.value
}

func (v *NullableModelSvcStatusResponse) Set(val *ModelSvcStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSvcStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSvcStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSvcStatusResponse(val *ModelSvcStatusResponse) *NullableModelSvcStatusResponse {
	return &NullableModelSvcStatusResponse{value: val, isSet: true}
}

func (v NullableModelSvcStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSvcStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


