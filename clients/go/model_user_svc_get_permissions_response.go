/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.31
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcGetPermissionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcGetPermissionsResponse{}

// UserSvcGetPermissionsResponse struct for UserSvcGetPermissionsResponse
type UserSvcGetPermissionsResponse struct {
	Permissions []UserSvcPermission `json:"permissions,omitempty"`
}

// NewUserSvcGetPermissionsResponse instantiates a new UserSvcGetPermissionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcGetPermissionsResponse() *UserSvcGetPermissionsResponse {
	this := UserSvcGetPermissionsResponse{}
	return &this
}

// NewUserSvcGetPermissionsResponseWithDefaults instantiates a new UserSvcGetPermissionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcGetPermissionsResponseWithDefaults() *UserSvcGetPermissionsResponse {
	this := UserSvcGetPermissionsResponse{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UserSvcGetPermissionsResponse) GetPermissions() []UserSvcPermission {
	if o == nil || IsNil(o.Permissions) {
		var ret []UserSvcPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcGetPermissionsResponse) GetPermissionsOk() ([]UserSvcPermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UserSvcGetPermissionsResponse) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []UserSvcPermission and assigns it to the Permissions field.
func (o *UserSvcGetPermissionsResponse) SetPermissions(v []UserSvcPermission) {
	o.Permissions = v
}

func (o UserSvcGetPermissionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcGetPermissionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableUserSvcGetPermissionsResponse struct {
	value *UserSvcGetPermissionsResponse
	isSet bool
}

func (v NullableUserSvcGetPermissionsResponse) Get() *UserSvcGetPermissionsResponse {
	return v.value
}

func (v *NullableUserSvcGetPermissionsResponse) Set(val *UserSvcGetPermissionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcGetPermissionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcGetPermissionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcGetPermissionsResponse(val *UserSvcGetPermissionsResponse) *NullableUserSvcGetPermissionsResponse {
	return &NullableUserSvcGetPermissionsResponse{value: val, isSet: true}
}

func (v NullableUserSvcGetPermissionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcGetPermissionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


