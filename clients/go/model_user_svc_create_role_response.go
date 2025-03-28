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

// checks if the UserSvcCreateRoleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcCreateRoleResponse{}

// UserSvcCreateRoleResponse struct for UserSvcCreateRoleResponse
type UserSvcCreateRoleResponse struct {
	Role *UserSvcRole `json:"role,omitempty"`
}

// NewUserSvcCreateRoleResponse instantiates a new UserSvcCreateRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcCreateRoleResponse() *UserSvcCreateRoleResponse {
	this := UserSvcCreateRoleResponse{}
	return &this
}

// NewUserSvcCreateRoleResponseWithDefaults instantiates a new UserSvcCreateRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcCreateRoleResponseWithDefaults() *UserSvcCreateRoleResponse {
	this := UserSvcCreateRoleResponse{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UserSvcCreateRoleResponse) GetRole() UserSvcRole {
	if o == nil || IsNil(o.Role) {
		var ret UserSvcRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcCreateRoleResponse) GetRoleOk() (*UserSvcRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UserSvcCreateRoleResponse) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given UserSvcRole and assigns it to the Role field.
func (o *UserSvcCreateRoleResponse) SetRole(v UserSvcRole) {
	o.Role = &v
}

func (o UserSvcCreateRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcCreateRoleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableUserSvcCreateRoleResponse struct {
	value *UserSvcCreateRoleResponse
	isSet bool
}

func (v NullableUserSvcCreateRoleResponse) Get() *UserSvcCreateRoleResponse {
	return v.value
}

func (v *NullableUserSvcCreateRoleResponse) Set(val *UserSvcCreateRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcCreateRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcCreateRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcCreateRoleResponse(val *UserSvcCreateRoleResponse) *NullableUserSvcCreateRoleResponse {
	return &NullableUserSvcCreateRoleResponse{value: val, isSet: true}
}

func (v NullableUserSvcCreateRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcCreateRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


