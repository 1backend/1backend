/*
1Backend

A unified backend development platform for microservices-based AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DataSvcCreateObjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSvcCreateObjectResponse{}

// DataSvcCreateObjectResponse struct for DataSvcCreateObjectResponse
type DataSvcCreateObjectResponse struct {
	Object *DataSvcObject `json:"object,omitempty"`
}

// NewDataSvcCreateObjectResponse instantiates a new DataSvcCreateObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSvcCreateObjectResponse() *DataSvcCreateObjectResponse {
	this := DataSvcCreateObjectResponse{}
	return &this
}

// NewDataSvcCreateObjectResponseWithDefaults instantiates a new DataSvcCreateObjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSvcCreateObjectResponseWithDefaults() *DataSvcCreateObjectResponse {
	this := DataSvcCreateObjectResponse{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *DataSvcCreateObjectResponse) GetObject() DataSvcObject {
	if o == nil || IsNil(o.Object) {
		var ret DataSvcObject
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcCreateObjectResponse) GetObjectOk() (*DataSvcObject, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *DataSvcCreateObjectResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given DataSvcObject and assigns it to the Object field.
func (o *DataSvcCreateObjectResponse) SetObject(v DataSvcObject) {
	o.Object = &v
}

func (o DataSvcCreateObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSvcCreateObjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableDataSvcCreateObjectResponse struct {
	value *DataSvcCreateObjectResponse
	isSet bool
}

func (v NullableDataSvcCreateObjectResponse) Get() *DataSvcCreateObjectResponse {
	return v.value
}

func (v *NullableDataSvcCreateObjectResponse) Set(val *DataSvcCreateObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSvcCreateObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSvcCreateObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSvcCreateObjectResponse(val *DataSvcCreateObjectResponse) *NullableDataSvcCreateObjectResponse {
	return &NullableDataSvcCreateObjectResponse{value: val, isSet: true}
}

func (v NullableDataSvcCreateObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSvcCreateObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


