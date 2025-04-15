/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.37
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserSvcSaveEnrollsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcSaveEnrollsResponse{}

// UserSvcSaveEnrollsResponse struct for UserSvcSaveEnrollsResponse
type UserSvcSaveEnrollsResponse struct {
	Enrolls []UserSvcEnroll `json:"enrolls"`
}

type _UserSvcSaveEnrollsResponse UserSvcSaveEnrollsResponse

// NewUserSvcSaveEnrollsResponse instantiates a new UserSvcSaveEnrollsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcSaveEnrollsResponse(enrolls []UserSvcEnroll) *UserSvcSaveEnrollsResponse {
	this := UserSvcSaveEnrollsResponse{}
	this.Enrolls = enrolls
	return &this
}

// NewUserSvcSaveEnrollsResponseWithDefaults instantiates a new UserSvcSaveEnrollsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcSaveEnrollsResponseWithDefaults() *UserSvcSaveEnrollsResponse {
	this := UserSvcSaveEnrollsResponse{}
	return &this
}

// GetEnrolls returns the Enrolls field value
func (o *UserSvcSaveEnrollsResponse) GetEnrolls() []UserSvcEnroll {
	if o == nil {
		var ret []UserSvcEnroll
		return ret
	}

	return o.Enrolls
}

// GetEnrollsOk returns a tuple with the Enrolls field value
// and a boolean to check if the value has been set.
func (o *UserSvcSaveEnrollsResponse) GetEnrollsOk() ([]UserSvcEnroll, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enrolls, true
}

// SetEnrolls sets field value
func (o *UserSvcSaveEnrollsResponse) SetEnrolls(v []UserSvcEnroll) {
	o.Enrolls = v
}

func (o UserSvcSaveEnrollsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcSaveEnrollsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enrolls"] = o.Enrolls
	return toSerialize, nil
}

func (o *UserSvcSaveEnrollsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUserSvcSaveEnrollsResponse := _UserSvcSaveEnrollsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcSaveEnrollsResponse)

	if err != nil {
		return err
	}

	*o = UserSvcSaveEnrollsResponse(varUserSvcSaveEnrollsResponse)

	return err
}

type NullableUserSvcSaveEnrollsResponse struct {
	value *UserSvcSaveEnrollsResponse
	isSet bool
}

func (v NullableUserSvcSaveEnrollsResponse) Get() *UserSvcSaveEnrollsResponse {
	return v.value
}

func (v *NullableUserSvcSaveEnrollsResponse) Set(val *UserSvcSaveEnrollsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcSaveEnrollsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcSaveEnrollsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcSaveEnrollsResponse(val *UserSvcSaveEnrollsResponse) *NullableUserSvcSaveEnrollsResponse {
	return &NullableUserSvcSaveEnrollsResponse{value: val, isSet: true}
}

func (v NullableUserSvcSaveEnrollsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcSaveEnrollsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


