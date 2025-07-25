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
	"bytes"
	"fmt"
)

// checks if the UserSvcSaveEnrollsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcSaveEnrollsRequest{}

// UserSvcSaveEnrollsRequest struct for UserSvcSaveEnrollsRequest
type UserSvcSaveEnrollsRequest struct {
	Enrolls []UserSvcEnrollInput `json:"enrolls"`
}

type _UserSvcSaveEnrollsRequest UserSvcSaveEnrollsRequest

// NewUserSvcSaveEnrollsRequest instantiates a new UserSvcSaveEnrollsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcSaveEnrollsRequest(enrolls []UserSvcEnrollInput) *UserSvcSaveEnrollsRequest {
	this := UserSvcSaveEnrollsRequest{}
	this.Enrolls = enrolls
	return &this
}

// NewUserSvcSaveEnrollsRequestWithDefaults instantiates a new UserSvcSaveEnrollsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcSaveEnrollsRequestWithDefaults() *UserSvcSaveEnrollsRequest {
	this := UserSvcSaveEnrollsRequest{}
	return &this
}

// GetEnrolls returns the Enrolls field value
func (o *UserSvcSaveEnrollsRequest) GetEnrolls() []UserSvcEnrollInput {
	if o == nil {
		var ret []UserSvcEnrollInput
		return ret
	}

	return o.Enrolls
}

// GetEnrollsOk returns a tuple with the Enrolls field value
// and a boolean to check if the value has been set.
func (o *UserSvcSaveEnrollsRequest) GetEnrollsOk() ([]UserSvcEnrollInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enrolls, true
}

// SetEnrolls sets field value
func (o *UserSvcSaveEnrollsRequest) SetEnrolls(v []UserSvcEnrollInput) {
	o.Enrolls = v
}

func (o UserSvcSaveEnrollsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcSaveEnrollsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enrolls"] = o.Enrolls
	return toSerialize, nil
}

func (o *UserSvcSaveEnrollsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enrolls",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserSvcSaveEnrollsRequest := _UserSvcSaveEnrollsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcSaveEnrollsRequest)

	if err != nil {
		return err
	}

	*o = UserSvcSaveEnrollsRequest(varUserSvcSaveEnrollsRequest)

	return err
}

type NullableUserSvcSaveEnrollsRequest struct {
	value *UserSvcSaveEnrollsRequest
	isSet bool
}

func (v NullableUserSvcSaveEnrollsRequest) Get() *UserSvcSaveEnrollsRequest {
	return v.value
}

func (v *NullableUserSvcSaveEnrollsRequest) Set(val *UserSvcSaveEnrollsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcSaveEnrollsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcSaveEnrollsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcSaveEnrollsRequest(val *UserSvcSaveEnrollsRequest) *NullableUserSvcSaveEnrollsRequest {
	return &NullableUserSvcSaveEnrollsRequest{value: val, isSet: true}
}

func (v NullableUserSvcSaveEnrollsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcSaveEnrollsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


