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

// checks if the UserSvcSavePermissionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcSavePermissionsResponse{}

// UserSvcSavePermissionsResponse struct for UserSvcSavePermissionsResponse
type UserSvcSavePermissionsResponse struct {
	Permissions []UserSvcPermission `json:"permissions,omitempty"`
}

// NewUserSvcSavePermissionsResponse instantiates a new UserSvcSavePermissionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcSavePermissionsResponse() *UserSvcSavePermissionsResponse {
	this := UserSvcSavePermissionsResponse{}
	return &this
}

// NewUserSvcSavePermissionsResponseWithDefaults instantiates a new UserSvcSavePermissionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcSavePermissionsResponseWithDefaults() *UserSvcSavePermissionsResponse {
	this := UserSvcSavePermissionsResponse{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UserSvcSavePermissionsResponse) GetPermissions() []UserSvcPermission {
	if o == nil || IsNil(o.Permissions) {
		var ret []UserSvcPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSavePermissionsResponse) GetPermissionsOk() ([]UserSvcPermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UserSvcSavePermissionsResponse) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []UserSvcPermission and assigns it to the Permissions field.
func (o *UserSvcSavePermissionsResponse) SetPermissions(v []UserSvcPermission) {
	o.Permissions = v
}

func (o UserSvcSavePermissionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcSavePermissionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableUserSvcSavePermissionsResponse struct {
	value *UserSvcSavePermissionsResponse
	isSet bool
}

func (v NullableUserSvcSavePermissionsResponse) Get() *UserSvcSavePermissionsResponse {
	return v.value
}

func (v *NullableUserSvcSavePermissionsResponse) Set(val *UserSvcSavePermissionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcSavePermissionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcSavePermissionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcSavePermissionsResponse(val *UserSvcSavePermissionsResponse) *NullableUserSvcSavePermissionsResponse {
	return &NullableUserSvcSavePermissionsResponse{value: val, isSet: true}
}

func (v NullableUserSvcSavePermissionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcSavePermissionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


