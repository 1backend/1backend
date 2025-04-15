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
)

// checks if the DataSvcQueryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSvcQueryResponse{}

// DataSvcQueryResponse struct for DataSvcQueryResponse
type DataSvcQueryResponse struct {
	Objects []DataSvcObject `json:"objects,omitempty"`
}

// NewDataSvcQueryResponse instantiates a new DataSvcQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSvcQueryResponse() *DataSvcQueryResponse {
	this := DataSvcQueryResponse{}
	return &this
}

// NewDataSvcQueryResponseWithDefaults instantiates a new DataSvcQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSvcQueryResponseWithDefaults() *DataSvcQueryResponse {
	this := DataSvcQueryResponse{}
	return &this
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *DataSvcQueryResponse) GetObjects() []DataSvcObject {
	if o == nil || IsNil(o.Objects) {
		var ret []DataSvcObject
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcQueryResponse) GetObjectsOk() ([]DataSvcObject, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *DataSvcQueryResponse) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []DataSvcObject and assigns it to the Objects field.
func (o *DataSvcQueryResponse) SetObjects(v []DataSvcObject) {
	o.Objects = v
}

func (o DataSvcQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSvcQueryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	return toSerialize, nil
}

type NullableDataSvcQueryResponse struct {
	value *DataSvcQueryResponse
	isSet bool
}

func (v NullableDataSvcQueryResponse) Get() *DataSvcQueryResponse {
	return v.value
}

func (v *NullableDataSvcQueryResponse) Set(val *DataSvcQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSvcQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSvcQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSvcQueryResponse(val *DataSvcQueryResponse) *NullableDataSvcQueryResponse {
	return &NullableDataSvcQueryResponse{value: val, isSet: true}
}

func (v NullableDataSvcQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSvcQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


