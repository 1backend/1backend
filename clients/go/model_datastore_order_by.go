/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.38
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DatastoreOrderBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatastoreOrderBy{}

// DatastoreOrderBy struct for DatastoreOrderBy
type DatastoreOrderBy struct {
	// Desc indicates whether the sorting should be in descending order.
	Desc *bool `json:"desc,omitempty"`
	// The field by which to order the results
	Field *string `json:"field,omitempty"`
	// Randomize indicates that the results should be randomized instead of ordered by the `field` and `desc` criteria
	Randomize *bool `json:"randomize,omitempty"`
	// Defines the type of sorting to apply (numeric, text, date, etc.)
	SortingType *DatastoreSortingType `json:"sortingType,omitempty"`
}

// NewDatastoreOrderBy instantiates a new DatastoreOrderBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatastoreOrderBy() *DatastoreOrderBy {
	this := DatastoreOrderBy{}
	return &this
}

// NewDatastoreOrderByWithDefaults instantiates a new DatastoreOrderBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatastoreOrderByWithDefaults() *DatastoreOrderBy {
	this := DatastoreOrderBy{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *DatastoreOrderBy) GetDesc() bool {
	if o == nil || IsNil(o.Desc) {
		var ret bool
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreOrderBy) GetDescOk() (*bool, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *DatastoreOrderBy) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given bool and assigns it to the Desc field.
func (o *DatastoreOrderBy) SetDesc(v bool) {
	o.Desc = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *DatastoreOrderBy) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreOrderBy) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *DatastoreOrderBy) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *DatastoreOrderBy) SetField(v string) {
	o.Field = &v
}

// GetRandomize returns the Randomize field value if set, zero value otherwise.
func (o *DatastoreOrderBy) GetRandomize() bool {
	if o == nil || IsNil(o.Randomize) {
		var ret bool
		return ret
	}
	return *o.Randomize
}

// GetRandomizeOk returns a tuple with the Randomize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreOrderBy) GetRandomizeOk() (*bool, bool) {
	if o == nil || IsNil(o.Randomize) {
		return nil, false
	}
	return o.Randomize, true
}

// HasRandomize returns a boolean if a field has been set.
func (o *DatastoreOrderBy) HasRandomize() bool {
	if o != nil && !IsNil(o.Randomize) {
		return true
	}

	return false
}

// SetRandomize gets a reference to the given bool and assigns it to the Randomize field.
func (o *DatastoreOrderBy) SetRandomize(v bool) {
	o.Randomize = &v
}

// GetSortingType returns the SortingType field value if set, zero value otherwise.
func (o *DatastoreOrderBy) GetSortingType() DatastoreSortingType {
	if o == nil || IsNil(o.SortingType) {
		var ret DatastoreSortingType
		return ret
	}
	return *o.SortingType
}

// GetSortingTypeOk returns a tuple with the SortingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreOrderBy) GetSortingTypeOk() (*DatastoreSortingType, bool) {
	if o == nil || IsNil(o.SortingType) {
		return nil, false
	}
	return o.SortingType, true
}

// HasSortingType returns a boolean if a field has been set.
func (o *DatastoreOrderBy) HasSortingType() bool {
	if o != nil && !IsNil(o.SortingType) {
		return true
	}

	return false
}

// SetSortingType gets a reference to the given DatastoreSortingType and assigns it to the SortingType field.
func (o *DatastoreOrderBy) SetSortingType(v DatastoreSortingType) {
	o.SortingType = &v
}

func (o DatastoreOrderBy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatastoreOrderBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Randomize) {
		toSerialize["randomize"] = o.Randomize
	}
	if !IsNil(o.SortingType) {
		toSerialize["sortingType"] = o.SortingType
	}
	return toSerialize, nil
}

type NullableDatastoreOrderBy struct {
	value *DatastoreOrderBy
	isSet bool
}

func (v NullableDatastoreOrderBy) Get() *DatastoreOrderBy {
	return v.value
}

func (v *NullableDatastoreOrderBy) Set(val *DatastoreOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableDatastoreOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableDatastoreOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatastoreOrderBy(val *DatastoreOrderBy) *NullableDatastoreOrderBy {
	return &NullableDatastoreOrderBy{value: val, isSet: true}
}

func (v NullableDatastoreOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatastoreOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


