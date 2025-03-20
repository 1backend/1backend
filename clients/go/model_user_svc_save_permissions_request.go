/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcSavePermissionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcSavePermissionsRequest{}

// UserSvcSavePermissionsRequest struct for UserSvcSavePermissionsRequest
type UserSvcSavePermissionsRequest struct {
	Permissions []UserSvcPermission `json:"permissions,omitempty"`
}

// NewUserSvcSavePermissionsRequest instantiates a new UserSvcSavePermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcSavePermissionsRequest() *UserSvcSavePermissionsRequest {
	this := UserSvcSavePermissionsRequest{}
	return &this
}

// NewUserSvcSavePermissionsRequestWithDefaults instantiates a new UserSvcSavePermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcSavePermissionsRequestWithDefaults() *UserSvcSavePermissionsRequest {
	this := UserSvcSavePermissionsRequest{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UserSvcSavePermissionsRequest) GetPermissions() []UserSvcPermission {
	if o == nil || IsNil(o.Permissions) {
		var ret []UserSvcPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSavePermissionsRequest) GetPermissionsOk() ([]UserSvcPermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UserSvcSavePermissionsRequest) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []UserSvcPermission and assigns it to the Permissions field.
func (o *UserSvcSavePermissionsRequest) SetPermissions(v []UserSvcPermission) {
	o.Permissions = v
}

func (o UserSvcSavePermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcSavePermissionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableUserSvcSavePermissionsRequest struct {
	value *UserSvcSavePermissionsRequest
	isSet bool
}

func (v NullableUserSvcSavePermissionsRequest) Get() *UserSvcSavePermissionsRequest {
	return v.value
}

func (v *NullableUserSvcSavePermissionsRequest) Set(val *UserSvcSavePermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcSavePermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcSavePermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcSavePermissionsRequest(val *UserSvcSavePermissionsRequest) *NullableUserSvcSavePermissionsRequest {
	return &NullableUserSvcSavePermissionsRequest{value: val, isSet: true}
}

func (v NullableUserSvcSavePermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcSavePermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


