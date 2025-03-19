/*
1Backend

A unified backend development platform for microservices-based AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcAssignPermissionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcAssignPermissionsRequest{}

// UserSvcAssignPermissionsRequest struct for UserSvcAssignPermissionsRequest
type UserSvcAssignPermissionsRequest struct {
	PermissionLinks []UserSvcPermissionLink `json:"permissionLinks,omitempty"`
}

// NewUserSvcAssignPermissionsRequest instantiates a new UserSvcAssignPermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcAssignPermissionsRequest() *UserSvcAssignPermissionsRequest {
	this := UserSvcAssignPermissionsRequest{}
	return &this
}

// NewUserSvcAssignPermissionsRequestWithDefaults instantiates a new UserSvcAssignPermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcAssignPermissionsRequestWithDefaults() *UserSvcAssignPermissionsRequest {
	this := UserSvcAssignPermissionsRequest{}
	return &this
}

// GetPermissionLinks returns the PermissionLinks field value if set, zero value otherwise.
func (o *UserSvcAssignPermissionsRequest) GetPermissionLinks() []UserSvcPermissionLink {
	if o == nil || IsNil(o.PermissionLinks) {
		var ret []UserSvcPermissionLink
		return ret
	}
	return o.PermissionLinks
}

// GetPermissionLinksOk returns a tuple with the PermissionLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcAssignPermissionsRequest) GetPermissionLinksOk() ([]UserSvcPermissionLink, bool) {
	if o == nil || IsNil(o.PermissionLinks) {
		return nil, false
	}
	return o.PermissionLinks, true
}

// HasPermissionLinks returns a boolean if a field has been set.
func (o *UserSvcAssignPermissionsRequest) HasPermissionLinks() bool {
	if o != nil && !IsNil(o.PermissionLinks) {
		return true
	}

	return false
}

// SetPermissionLinks gets a reference to the given []UserSvcPermissionLink and assigns it to the PermissionLinks field.
func (o *UserSvcAssignPermissionsRequest) SetPermissionLinks(v []UserSvcPermissionLink) {
	o.PermissionLinks = v
}

func (o UserSvcAssignPermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcAssignPermissionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PermissionLinks) {
		toSerialize["permissionLinks"] = o.PermissionLinks
	}
	return toSerialize, nil
}

type NullableUserSvcAssignPermissionsRequest struct {
	value *UserSvcAssignPermissionsRequest
	isSet bool
}

func (v NullableUserSvcAssignPermissionsRequest) Get() *UserSvcAssignPermissionsRequest {
	return v.value
}

func (v *NullableUserSvcAssignPermissionsRequest) Set(val *UserSvcAssignPermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcAssignPermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcAssignPermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcAssignPermissionsRequest(val *UserSvcAssignPermissionsRequest) *NullableUserSvcAssignPermissionsRequest {
	return &NullableUserSvcAssignPermissionsRequest{value: val, isSet: true}
}

func (v NullableUserSvcAssignPermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcAssignPermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


