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
	"fmt"
)

// UserSvcOrderDirection the model 'UserSvcOrderDirection'
type UserSvcOrderDirection string

// List of user_svc.OrderDirection
const (
	OrderDirectionAsc UserSvcOrderDirection = "asc"
	OrderDirectionDesc UserSvcOrderDirection = "desc"
)

// All allowed values of UserSvcOrderDirection enum
var AllowedUserSvcOrderDirectionEnumValues = []UserSvcOrderDirection{
	"asc",
	"desc",
}

func (v *UserSvcOrderDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserSvcOrderDirection(value)
	for _, existing := range AllowedUserSvcOrderDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserSvcOrderDirection", value)
}

// NewUserSvcOrderDirectionFromValue returns a pointer to a valid UserSvcOrderDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserSvcOrderDirectionFromValue(v string) (*UserSvcOrderDirection, error) {
	ev := UserSvcOrderDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserSvcOrderDirection: valid values are %v", v, AllowedUserSvcOrderDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserSvcOrderDirection) IsValid() bool {
	for _, existing := range AllowedUserSvcOrderDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to user_svc.OrderDirection value
func (v UserSvcOrderDirection) Ptr() *UserSvcOrderDirection {
	return &v
}

type NullableUserSvcOrderDirection struct {
	value *UserSvcOrderDirection
	isSet bool
}

func (v NullableUserSvcOrderDirection) Get() *UserSvcOrderDirection {
	return v.value
}

func (v *NullableUserSvcOrderDirection) Set(val *UserSvcOrderDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcOrderDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcOrderDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcOrderDirection(val *UserSvcOrderDirection) *NullableUserSvcOrderDirection {
	return &NullableUserSvcOrderDirection{value: val, isSet: true}
}

func (v NullableUserSvcOrderDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcOrderDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

