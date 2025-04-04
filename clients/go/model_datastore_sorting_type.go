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
	"fmt"
)

// DatastoreSortingType the model 'DatastoreSortingType'
type DatastoreSortingType string

// List of datastore.SortingType
const (
	SortingTypeDefault DatastoreSortingType = ""
	SortingTypeNumeric DatastoreSortingType = "numeric"
	SortingTypeText DatastoreSortingType = "text"
	SortingTypeDate DatastoreSortingType = "date"
)

// All allowed values of DatastoreSortingType enum
var AllowedDatastoreSortingTypeEnumValues = []DatastoreSortingType{
	"",
	"numeric",
	"text",
	"date",
}

func (v *DatastoreSortingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DatastoreSortingType(value)
	for _, existing := range AllowedDatastoreSortingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DatastoreSortingType", value)
}

// NewDatastoreSortingTypeFromValue returns a pointer to a valid DatastoreSortingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatastoreSortingTypeFromValue(v string) (*DatastoreSortingType, error) {
	ev := DatastoreSortingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DatastoreSortingType: valid values are %v", v, AllowedDatastoreSortingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DatastoreSortingType) IsValid() bool {
	for _, existing := range AllowedDatastoreSortingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to datastore.SortingType value
func (v DatastoreSortingType) Ptr() *DatastoreSortingType {
	return &v
}

type NullableDatastoreSortingType struct {
	value *DatastoreSortingType
	isSet bool
}

func (v NullableDatastoreSortingType) Get() *DatastoreSortingType {
	return v.value
}

func (v *NullableDatastoreSortingType) Set(val *DatastoreSortingType) {
	v.value = val
	v.isSet = true
}

func (v NullableDatastoreSortingType) IsSet() bool {
	return v.isSet
}

func (v *NullableDatastoreSortingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatastoreSortingType(val *DatastoreSortingType) *NullableDatastoreSortingType {
	return &NullableDatastoreSortingType{value: val, isSet: true}
}

func (v NullableDatastoreSortingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatastoreSortingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

