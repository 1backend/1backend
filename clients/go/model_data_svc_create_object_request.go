/*
OpenOrch

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DataSvcCreateObjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSvcCreateObjectRequest{}

// DataSvcCreateObjectRequest struct for DataSvcCreateObjectRequest
type DataSvcCreateObjectRequest struct {
	Object *DataSvcCreateObjectFields `json:"object,omitempty"`
}

// NewDataSvcCreateObjectRequest instantiates a new DataSvcCreateObjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSvcCreateObjectRequest() *DataSvcCreateObjectRequest {
	this := DataSvcCreateObjectRequest{}
	return &this
}

// NewDataSvcCreateObjectRequestWithDefaults instantiates a new DataSvcCreateObjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSvcCreateObjectRequestWithDefaults() *DataSvcCreateObjectRequest {
	this := DataSvcCreateObjectRequest{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *DataSvcCreateObjectRequest) GetObject() DataSvcCreateObjectFields {
	if o == nil || IsNil(o.Object) {
		var ret DataSvcCreateObjectFields
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcCreateObjectRequest) GetObjectOk() (*DataSvcCreateObjectFields, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *DataSvcCreateObjectRequest) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given DataSvcCreateObjectFields and assigns it to the Object field.
func (o *DataSvcCreateObjectRequest) SetObject(v DataSvcCreateObjectFields) {
	o.Object = &v
}

func (o DataSvcCreateObjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSvcCreateObjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableDataSvcCreateObjectRequest struct {
	value *DataSvcCreateObjectRequest
	isSet bool
}

func (v NullableDataSvcCreateObjectRequest) Get() *DataSvcCreateObjectRequest {
	return v.value
}

func (v *NullableDataSvcCreateObjectRequest) Set(val *DataSvcCreateObjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSvcCreateObjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSvcCreateObjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSvcCreateObjectRequest(val *DataSvcCreateObjectRequest) *NullableDataSvcCreateObjectRequest {
	return &NullableDataSvcCreateObjectRequest{value: val, isSet: true}
}

func (v NullableDataSvcCreateObjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSvcCreateObjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


