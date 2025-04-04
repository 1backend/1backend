/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.32
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DataSvcUpsertObjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSvcUpsertObjectRequest{}

// DataSvcUpsertObjectRequest struct for DataSvcUpsertObjectRequest
type DataSvcUpsertObjectRequest struct {
	Object *DataSvcCreateObjectFields `json:"object,omitempty"`
}

// NewDataSvcUpsertObjectRequest instantiates a new DataSvcUpsertObjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSvcUpsertObjectRequest() *DataSvcUpsertObjectRequest {
	this := DataSvcUpsertObjectRequest{}
	return &this
}

// NewDataSvcUpsertObjectRequestWithDefaults instantiates a new DataSvcUpsertObjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSvcUpsertObjectRequestWithDefaults() *DataSvcUpsertObjectRequest {
	this := DataSvcUpsertObjectRequest{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *DataSvcUpsertObjectRequest) GetObject() DataSvcCreateObjectFields {
	if o == nil || IsNil(o.Object) {
		var ret DataSvcCreateObjectFields
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcUpsertObjectRequest) GetObjectOk() (*DataSvcCreateObjectFields, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *DataSvcUpsertObjectRequest) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given DataSvcCreateObjectFields and assigns it to the Object field.
func (o *DataSvcUpsertObjectRequest) SetObject(v DataSvcCreateObjectFields) {
	o.Object = &v
}

func (o DataSvcUpsertObjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSvcUpsertObjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableDataSvcUpsertObjectRequest struct {
	value *DataSvcUpsertObjectRequest
	isSet bool
}

func (v NullableDataSvcUpsertObjectRequest) Get() *DataSvcUpsertObjectRequest {
	return v.value
}

func (v *NullableDataSvcUpsertObjectRequest) Set(val *DataSvcUpsertObjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSvcUpsertObjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSvcUpsertObjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSvcUpsertObjectRequest(val *DataSvcUpsertObjectRequest) *NullableDataSvcUpsertObjectRequest {
	return &NullableDataSvcUpsertObjectRequest{value: val, isSet: true}
}

func (v NullableDataSvcUpsertObjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSvcUpsertObjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


