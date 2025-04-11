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
	"fmt"
)

// DatastoreOp the model 'DatastoreOp'
type DatastoreOp string

// List of datastore.Op
const (
	OpOr DatastoreOp = "or"
	OpEquals DatastoreOp = "equals"
	OpContainsSubstring DatastoreOp = "containsSubstring"
	OpStartsWith DatastoreOp = "startsWith"
	OpIntersects DatastoreOp = "intersects"
	OpIsInList DatastoreOp = "isInList"
)

// All allowed values of DatastoreOp enum
var AllowedDatastoreOpEnumValues = []DatastoreOp{
	"or",
	"equals",
	"containsSubstring",
	"startsWith",
	"intersects",
	"isInList",
}

func (v *DatastoreOp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DatastoreOp(value)
	for _, existing := range AllowedDatastoreOpEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DatastoreOp", value)
}

// NewDatastoreOpFromValue returns a pointer to a valid DatastoreOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatastoreOpFromValue(v string) (*DatastoreOp, error) {
	ev := DatastoreOp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DatastoreOp: valid values are %v", v, AllowedDatastoreOpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DatastoreOp) IsValid() bool {
	for _, existing := range AllowedDatastoreOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to datastore.Op value
func (v DatastoreOp) Ptr() *DatastoreOp {
	return &v
}

type NullableDatastoreOp struct {
	value *DatastoreOp
	isSet bool
}

func (v NullableDatastoreOp) Get() *DatastoreOp {
	return v.value
}

func (v *NullableDatastoreOp) Set(val *DatastoreOp) {
	v.value = val
	v.isSet = true
}

func (v NullableDatastoreOp) IsSet() bool {
	return v.isSet
}

func (v *NullableDatastoreOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatastoreOp(val *DatastoreOp) *NullableDatastoreOp {
	return &NullableDatastoreOp{value: val, isSet: true}
}

func (v NullableDatastoreOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatastoreOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

