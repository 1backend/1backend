/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DataSvcUpsertObjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSvcUpsertObjectResponse{}

// DataSvcUpsertObjectResponse struct for DataSvcUpsertObjectResponse
type DataSvcUpsertObjectResponse struct {
	Object *DataSvcObject `json:"object,omitempty"`
}

// NewDataSvcUpsertObjectResponse instantiates a new DataSvcUpsertObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSvcUpsertObjectResponse() *DataSvcUpsertObjectResponse {
	this := DataSvcUpsertObjectResponse{}
	return &this
}

// NewDataSvcUpsertObjectResponseWithDefaults instantiates a new DataSvcUpsertObjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSvcUpsertObjectResponseWithDefaults() *DataSvcUpsertObjectResponse {
	this := DataSvcUpsertObjectResponse{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *DataSvcUpsertObjectResponse) GetObject() DataSvcObject {
	if o == nil || IsNil(o.Object) {
		var ret DataSvcObject
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcUpsertObjectResponse) GetObjectOk() (*DataSvcObject, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *DataSvcUpsertObjectResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given DataSvcObject and assigns it to the Object field.
func (o *DataSvcUpsertObjectResponse) SetObject(v DataSvcObject) {
	o.Object = &v
}

func (o DataSvcUpsertObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSvcUpsertObjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableDataSvcUpsertObjectResponse struct {
	value *DataSvcUpsertObjectResponse
	isSet bool
}

func (v NullableDataSvcUpsertObjectResponse) Get() *DataSvcUpsertObjectResponse {
	return v.value
}

func (v *NullableDataSvcUpsertObjectResponse) Set(val *DataSvcUpsertObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSvcUpsertObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSvcUpsertObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSvcUpsertObjectResponse(val *DataSvcUpsertObjectResponse) *NullableDataSvcUpsertObjectResponse {
	return &NullableDataSvcUpsertObjectResponse{value: val, isSet: true}
}

func (v NullableDataSvcUpsertObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSvcUpsertObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


