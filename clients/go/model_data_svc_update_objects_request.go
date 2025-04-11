/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.35
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DataSvcUpdateObjectsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSvcUpdateObjectsRequest{}

// DataSvcUpdateObjectsRequest struct for DataSvcUpdateObjectsRequest
type DataSvcUpdateObjectsRequest struct {
	// Filters to determine which objects will be updated. Only objects matching all filters will be modified.
	Filters []DatastoreFilter `json:"filters,omitempty"`
	// The object containing the fields to update in matching objects.
	Object *DataSvcObject `json:"object,omitempty"`
	Table *string `json:"table,omitempty"`
}

// NewDataSvcUpdateObjectsRequest instantiates a new DataSvcUpdateObjectsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSvcUpdateObjectsRequest() *DataSvcUpdateObjectsRequest {
	this := DataSvcUpdateObjectsRequest{}
	return &this
}

// NewDataSvcUpdateObjectsRequestWithDefaults instantiates a new DataSvcUpdateObjectsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSvcUpdateObjectsRequestWithDefaults() *DataSvcUpdateObjectsRequest {
	this := DataSvcUpdateObjectsRequest{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *DataSvcUpdateObjectsRequest) GetFilters() []DatastoreFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []DatastoreFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcUpdateObjectsRequest) GetFiltersOk() ([]DatastoreFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *DataSvcUpdateObjectsRequest) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []DatastoreFilter and assigns it to the Filters field.
func (o *DataSvcUpdateObjectsRequest) SetFilters(v []DatastoreFilter) {
	o.Filters = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *DataSvcUpdateObjectsRequest) GetObject() DataSvcObject {
	if o == nil || IsNil(o.Object) {
		var ret DataSvcObject
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcUpdateObjectsRequest) GetObjectOk() (*DataSvcObject, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *DataSvcUpdateObjectsRequest) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given DataSvcObject and assigns it to the Object field.
func (o *DataSvcUpdateObjectsRequest) SetObject(v DataSvcObject) {
	o.Object = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *DataSvcUpdateObjectsRequest) GetTable() string {
	if o == nil || IsNil(o.Table) {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcUpdateObjectsRequest) GetTableOk() (*string, bool) {
	if o == nil || IsNil(o.Table) {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *DataSvcUpdateObjectsRequest) HasTable() bool {
	if o != nil && !IsNil(o.Table) {
		return true
	}

	return false
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *DataSvcUpdateObjectsRequest) SetTable(v string) {
	o.Table = &v
}

func (o DataSvcUpdateObjectsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSvcUpdateObjectsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Table) {
		toSerialize["table"] = o.Table
	}
	return toSerialize, nil
}

type NullableDataSvcUpdateObjectsRequest struct {
	value *DataSvcUpdateObjectsRequest
	isSet bool
}

func (v NullableDataSvcUpdateObjectsRequest) Get() *DataSvcUpdateObjectsRequest {
	return v.value
}

func (v *NullableDataSvcUpdateObjectsRequest) Set(val *DataSvcUpdateObjectsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSvcUpdateObjectsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSvcUpdateObjectsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSvcUpdateObjectsRequest(val *DataSvcUpdateObjectsRequest) *NullableDataSvcUpdateObjectsRequest {
	return &NullableDataSvcUpdateObjectsRequest{value: val, isSet: true}
}

func (v NullableDataSvcUpdateObjectsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSvcUpdateObjectsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


