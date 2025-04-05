/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.34
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DatastoreFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatastoreFilter{}

// DatastoreFilter struct for DatastoreFilter
type DatastoreFilter struct {
	Fields []string `json:"fields,omitempty"`
	// JSONValues is a JSON marshalled array of values. It's JSON marhalled due to the limitations of the Swaggo -> OpenAPI 2.0 -> OpenAPI Go generator toolchain.
	JsonValues *string `json:"jsonValues,omitempty"`
	Op *DatastoreOp `json:"op,omitempty"`
	// SubFilters is used for operations like OR where multiple filters are combined.
	SubFilters []DatastoreFilter `json:"subFilters,omitempty"`
}

// NewDatastoreFilter instantiates a new DatastoreFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatastoreFilter() *DatastoreFilter {
	this := DatastoreFilter{}
	return &this
}

// NewDatastoreFilterWithDefaults instantiates a new DatastoreFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatastoreFilterWithDefaults() *DatastoreFilter {
	this := DatastoreFilter{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *DatastoreFilter) GetFields() []string {
	if o == nil || IsNil(o.Fields) {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreFilter) GetFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *DatastoreFilter) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *DatastoreFilter) SetFields(v []string) {
	o.Fields = v
}

// GetJsonValues returns the JsonValues field value if set, zero value otherwise.
func (o *DatastoreFilter) GetJsonValues() string {
	if o == nil || IsNil(o.JsonValues) {
		var ret string
		return ret
	}
	return *o.JsonValues
}

// GetJsonValuesOk returns a tuple with the JsonValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreFilter) GetJsonValuesOk() (*string, bool) {
	if o == nil || IsNil(o.JsonValues) {
		return nil, false
	}
	return o.JsonValues, true
}

// HasJsonValues returns a boolean if a field has been set.
func (o *DatastoreFilter) HasJsonValues() bool {
	if o != nil && !IsNil(o.JsonValues) {
		return true
	}

	return false
}

// SetJsonValues gets a reference to the given string and assigns it to the JsonValues field.
func (o *DatastoreFilter) SetJsonValues(v string) {
	o.JsonValues = &v
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *DatastoreFilter) GetOp() DatastoreOp {
	if o == nil || IsNil(o.Op) {
		var ret DatastoreOp
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreFilter) GetOpOk() (*DatastoreOp, bool) {
	if o == nil || IsNil(o.Op) {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *DatastoreFilter) HasOp() bool {
	if o != nil && !IsNil(o.Op) {
		return true
	}

	return false
}

// SetOp gets a reference to the given DatastoreOp and assigns it to the Op field.
func (o *DatastoreFilter) SetOp(v DatastoreOp) {
	o.Op = &v
}

// GetSubFilters returns the SubFilters field value if set, zero value otherwise.
func (o *DatastoreFilter) GetSubFilters() []DatastoreFilter {
	if o == nil || IsNil(o.SubFilters) {
		var ret []DatastoreFilter
		return ret
	}
	return o.SubFilters
}

// GetSubFiltersOk returns a tuple with the SubFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreFilter) GetSubFiltersOk() ([]DatastoreFilter, bool) {
	if o == nil || IsNil(o.SubFilters) {
		return nil, false
	}
	return o.SubFilters, true
}

// HasSubFilters returns a boolean if a field has been set.
func (o *DatastoreFilter) HasSubFilters() bool {
	if o != nil && !IsNil(o.SubFilters) {
		return true
	}

	return false
}

// SetSubFilters gets a reference to the given []DatastoreFilter and assigns it to the SubFilters field.
func (o *DatastoreFilter) SetSubFilters(v []DatastoreFilter) {
	o.SubFilters = v
}

func (o DatastoreFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatastoreFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.JsonValues) {
		toSerialize["jsonValues"] = o.JsonValues
	}
	if !IsNil(o.Op) {
		toSerialize["op"] = o.Op
	}
	if !IsNil(o.SubFilters) {
		toSerialize["subFilters"] = o.SubFilters
	}
	return toSerialize, nil
}

type NullableDatastoreFilter struct {
	value *DatastoreFilter
	isSet bool
}

func (v NullableDatastoreFilter) Get() *DatastoreFilter {
	return v.value
}

func (v *NullableDatastoreFilter) Set(val *DatastoreFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableDatastoreFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableDatastoreFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatastoreFilter(val *DatastoreFilter) *NullableDatastoreFilter {
	return &NullableDatastoreFilter{value: val, isSet: true}
}

func (v NullableDatastoreFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatastoreFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


