/*
OpenOrch

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// SecretSvcChecksumAlgorithm the model 'SecretSvcChecksumAlgorithm'
type SecretSvcChecksumAlgorithm string

// List of secret_svc.ChecksumAlgorithm
const (
	ChecksumAlgorithmUnspecified SecretSvcChecksumAlgorithm = ""
	ChecksumAlgorithmCRC32 SecretSvcChecksumAlgorithm = "CRC32"
	ChecksumAlgorithmBlake2s SecretSvcChecksumAlgorithm = "BLAKE2s"
	ChecksumAlgorithmSha256 SecretSvcChecksumAlgorithm = "SHA-256"
	ChecksumAlgorithmSha512 SecretSvcChecksumAlgorithm = "SHA-512"
)

// All allowed values of SecretSvcChecksumAlgorithm enum
var AllowedSecretSvcChecksumAlgorithmEnumValues = []SecretSvcChecksumAlgorithm{
	"",
	"CRC32",
	"BLAKE2s",
	"SHA-256",
	"SHA-512",
}

func (v *SecretSvcChecksumAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecretSvcChecksumAlgorithm(value)
	for _, existing := range AllowedSecretSvcChecksumAlgorithmEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecretSvcChecksumAlgorithm", value)
}

// NewSecretSvcChecksumAlgorithmFromValue returns a pointer to a valid SecretSvcChecksumAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecretSvcChecksumAlgorithmFromValue(v string) (*SecretSvcChecksumAlgorithm, error) {
	ev := SecretSvcChecksumAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecretSvcChecksumAlgorithm: valid values are %v", v, AllowedSecretSvcChecksumAlgorithmEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecretSvcChecksumAlgorithm) IsValid() bool {
	for _, existing := range AllowedSecretSvcChecksumAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to secret_svc.ChecksumAlgorithm value
func (v SecretSvcChecksumAlgorithm) Ptr() *SecretSvcChecksumAlgorithm {
	return &v
}

type NullableSecretSvcChecksumAlgorithm struct {
	value *SecretSvcChecksumAlgorithm
	isSet bool
}

func (v NullableSecretSvcChecksumAlgorithm) Get() *SecretSvcChecksumAlgorithm {
	return v.value
}

func (v *NullableSecretSvcChecksumAlgorithm) Set(val *SecretSvcChecksumAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretSvcChecksumAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretSvcChecksumAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretSvcChecksumAlgorithm(val *SecretSvcChecksumAlgorithm) *NullableSecretSvcChecksumAlgorithm {
	return &NullableSecretSvcChecksumAlgorithm{value: val, isSet: true}
}

func (v NullableSecretSvcChecksumAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretSvcChecksumAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

