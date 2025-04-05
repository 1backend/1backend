/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.33
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// RegistrySvcLanguage the model 'RegistrySvcLanguage'
type RegistrySvcLanguage string

// List of registry_svc.Language
const (
	JavaScript RegistrySvcLanguage = "JavaScript"
	Python RegistrySvcLanguage = "Python"
	Java RegistrySvcLanguage = "Java"
	CSharp RegistrySvcLanguage = "C#"
	CPlusPlus RegistrySvcLanguage = "C++"
	Ruby RegistrySvcLanguage = "Ruby"
	Go RegistrySvcLanguage = "Go"
	Swift RegistrySvcLanguage = "Swift"
	PHP RegistrySvcLanguage = "PHP"
	TypeScript RegistrySvcLanguage = "TypeScript"
	Kotlin RegistrySvcLanguage = "Kotlin"
	Scala RegistrySvcLanguage = "Scala"
	Perl RegistrySvcLanguage = "Perl"
	Rust RegistrySvcLanguage = "Rust"
	Haskell RegistrySvcLanguage = "Haskell"
	Clojure RegistrySvcLanguage = "Clojure"
	Elixir RegistrySvcLanguage = "Elixir"
	ObjectiveC RegistrySvcLanguage = "Objective-C"
	FSharp RegistrySvcLanguage = "F#"
)

// All allowed values of RegistrySvcLanguage enum
var AllowedRegistrySvcLanguageEnumValues = []RegistrySvcLanguage{
	"JavaScript",
	"Python",
	"Java",
	"C#",
	"C++",
	"Ruby",
	"Go",
	"Swift",
	"PHP",
	"TypeScript",
	"Kotlin",
	"Scala",
	"Perl",
	"Rust",
	"Haskell",
	"Clojure",
	"Elixir",
	"Objective-C",
	"F#",
}

func (v *RegistrySvcLanguage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegistrySvcLanguage(value)
	for _, existing := range AllowedRegistrySvcLanguageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegistrySvcLanguage", value)
}

// NewRegistrySvcLanguageFromValue returns a pointer to a valid RegistrySvcLanguage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegistrySvcLanguageFromValue(v string) (*RegistrySvcLanguage, error) {
	ev := RegistrySvcLanguage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegistrySvcLanguage: valid values are %v", v, AllowedRegistrySvcLanguageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegistrySvcLanguage) IsValid() bool {
	for _, existing := range AllowedRegistrySvcLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to registry_svc.Language value
func (v RegistrySvcLanguage) Ptr() *RegistrySvcLanguage {
	return &v
}

type NullableRegistrySvcLanguage struct {
	value *RegistrySvcLanguage
	isSet bool
}

func (v NullableRegistrySvcLanguage) Get() *RegistrySvcLanguage {
	return v.value
}

func (v *NullableRegistrySvcLanguage) Set(val *RegistrySvcLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcLanguage(val *RegistrySvcLanguage) *NullableRegistrySvcLanguage {
	return &NullableRegistrySvcLanguage{value: val, isSet: true}
}

func (v NullableRegistrySvcLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

