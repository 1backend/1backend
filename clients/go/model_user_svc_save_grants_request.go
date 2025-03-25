/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.30
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcSaveGrantsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcSaveGrantsRequest{}

// UserSvcSaveGrantsRequest struct for UserSvcSaveGrantsRequest
type UserSvcSaveGrantsRequest struct {
	Grants []UserSvcGrant `json:"grants,omitempty"`
}

// NewUserSvcSaveGrantsRequest instantiates a new UserSvcSaveGrantsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcSaveGrantsRequest() *UserSvcSaveGrantsRequest {
	this := UserSvcSaveGrantsRequest{}
	return &this
}

// NewUserSvcSaveGrantsRequestWithDefaults instantiates a new UserSvcSaveGrantsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcSaveGrantsRequestWithDefaults() *UserSvcSaveGrantsRequest {
	this := UserSvcSaveGrantsRequest{}
	return &this
}

// GetGrants returns the Grants field value if set, zero value otherwise.
func (o *UserSvcSaveGrantsRequest) GetGrants() []UserSvcGrant {
	if o == nil || IsNil(o.Grants) {
		var ret []UserSvcGrant
		return ret
	}
	return o.Grants
}

// GetGrantsOk returns a tuple with the Grants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSaveGrantsRequest) GetGrantsOk() ([]UserSvcGrant, bool) {
	if o == nil || IsNil(o.Grants) {
		return nil, false
	}
	return o.Grants, true
}

// HasGrants returns a boolean if a field has been set.
func (o *UserSvcSaveGrantsRequest) HasGrants() bool {
	if o != nil && !IsNil(o.Grants) {
		return true
	}

	return false
}

// SetGrants gets a reference to the given []UserSvcGrant and assigns it to the Grants field.
func (o *UserSvcSaveGrantsRequest) SetGrants(v []UserSvcGrant) {
	o.Grants = v
}

func (o UserSvcSaveGrantsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcSaveGrantsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Grants) {
		toSerialize["grants"] = o.Grants
	}
	return toSerialize, nil
}

type NullableUserSvcSaveGrantsRequest struct {
	value *UserSvcSaveGrantsRequest
	isSet bool
}

func (v NullableUserSvcSaveGrantsRequest) Get() *UserSvcSaveGrantsRequest {
	return v.value
}

func (v *NullableUserSvcSaveGrantsRequest) Set(val *UserSvcSaveGrantsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcSaveGrantsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcSaveGrantsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcSaveGrantsRequest(val *UserSvcSaveGrantsRequest) *NullableUserSvcSaveGrantsRequest {
	return &NullableUserSvcSaveGrantsRequest{value: val, isSet: true}
}

func (v NullableUserSvcSaveGrantsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcSaveGrantsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


