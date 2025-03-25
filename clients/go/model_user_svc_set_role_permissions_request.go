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

// checks if the UserSvcSetRolePermissionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcSetRolePermissionsRequest{}

// UserSvcSetRolePermissionsRequest struct for UserSvcSetRolePermissionsRequest
type UserSvcSetRolePermissionsRequest struct {
	PermissionIds []string `json:"permissionIds,omitempty"`
}

// NewUserSvcSetRolePermissionsRequest instantiates a new UserSvcSetRolePermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcSetRolePermissionsRequest() *UserSvcSetRolePermissionsRequest {
	this := UserSvcSetRolePermissionsRequest{}
	return &this
}

// NewUserSvcSetRolePermissionsRequestWithDefaults instantiates a new UserSvcSetRolePermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcSetRolePermissionsRequestWithDefaults() *UserSvcSetRolePermissionsRequest {
	this := UserSvcSetRolePermissionsRequest{}
	return &this
}

// GetPermissionIds returns the PermissionIds field value if set, zero value otherwise.
func (o *UserSvcSetRolePermissionsRequest) GetPermissionIds() []string {
	if o == nil || IsNil(o.PermissionIds) {
		var ret []string
		return ret
	}
	return o.PermissionIds
}

// GetPermissionIdsOk returns a tuple with the PermissionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSetRolePermissionsRequest) GetPermissionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PermissionIds) {
		return nil, false
	}
	return o.PermissionIds, true
}

// HasPermissionIds returns a boolean if a field has been set.
func (o *UserSvcSetRolePermissionsRequest) HasPermissionIds() bool {
	if o != nil && !IsNil(o.PermissionIds) {
		return true
	}

	return false
}

// SetPermissionIds gets a reference to the given []string and assigns it to the PermissionIds field.
func (o *UserSvcSetRolePermissionsRequest) SetPermissionIds(v []string) {
	o.PermissionIds = v
}

func (o UserSvcSetRolePermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcSetRolePermissionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PermissionIds) {
		toSerialize["permissionIds"] = o.PermissionIds
	}
	return toSerialize, nil
}

type NullableUserSvcSetRolePermissionsRequest struct {
	value *UserSvcSetRolePermissionsRequest
	isSet bool
}

func (v NullableUserSvcSetRolePermissionsRequest) Get() *UserSvcSetRolePermissionsRequest {
	return v.value
}

func (v *NullableUserSvcSetRolePermissionsRequest) Set(val *UserSvcSetRolePermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcSetRolePermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcSetRolePermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcSetRolePermissionsRequest(val *UserSvcSetRolePermissionsRequest) *NullableUserSvcSetRolePermissionsRequest {
	return &NullableUserSvcSetRolePermissionsRequest{value: val, isSet: true}
}

func (v NullableUserSvcSetRolePermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcSetRolePermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


