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

// checks if the DataSvcDeleteObjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSvcDeleteObjectRequest{}

// DataSvcDeleteObjectRequest struct for DataSvcDeleteObjectRequest
type DataSvcDeleteObjectRequest struct {
	Filters []DatastoreFilter `json:"filters,omitempty"`
	Table *string `json:"table,omitempty"`
}

// NewDataSvcDeleteObjectRequest instantiates a new DataSvcDeleteObjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSvcDeleteObjectRequest() *DataSvcDeleteObjectRequest {
	this := DataSvcDeleteObjectRequest{}
	return &this
}

// NewDataSvcDeleteObjectRequestWithDefaults instantiates a new DataSvcDeleteObjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSvcDeleteObjectRequestWithDefaults() *DataSvcDeleteObjectRequest {
	this := DataSvcDeleteObjectRequest{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *DataSvcDeleteObjectRequest) GetFilters() []DatastoreFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []DatastoreFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcDeleteObjectRequest) GetFiltersOk() ([]DatastoreFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *DataSvcDeleteObjectRequest) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []DatastoreFilter and assigns it to the Filters field.
func (o *DataSvcDeleteObjectRequest) SetFilters(v []DatastoreFilter) {
	o.Filters = v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *DataSvcDeleteObjectRequest) GetTable() string {
	if o == nil || IsNil(o.Table) {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcDeleteObjectRequest) GetTableOk() (*string, bool) {
	if o == nil || IsNil(o.Table) {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *DataSvcDeleteObjectRequest) HasTable() bool {
	if o != nil && !IsNil(o.Table) {
		return true
	}

	return false
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *DataSvcDeleteObjectRequest) SetTable(v string) {
	o.Table = &v
}

func (o DataSvcDeleteObjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSvcDeleteObjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Table) {
		toSerialize["table"] = o.Table
	}
	return toSerialize, nil
}

type NullableDataSvcDeleteObjectRequest struct {
	value *DataSvcDeleteObjectRequest
	isSet bool
}

func (v NullableDataSvcDeleteObjectRequest) Get() *DataSvcDeleteObjectRequest {
	return v.value
}

func (v *NullableDataSvcDeleteObjectRequest) Set(val *DataSvcDeleteObjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSvcDeleteObjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSvcDeleteObjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSvcDeleteObjectRequest(val *DataSvcDeleteObjectRequest) *NullableDataSvcDeleteObjectRequest {
	return &NullableDataSvcDeleteObjectRequest{value: val, isSet: true}
}

func (v NullableDataSvcDeleteObjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSvcDeleteObjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


