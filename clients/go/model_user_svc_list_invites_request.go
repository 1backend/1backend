/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.34
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcListInvitesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcListInvitesRequest{}

// UserSvcListInvitesRequest struct for UserSvcListInvitesRequest
type UserSvcListInvitesRequest struct {
	ContactId *string `json:"contactId,omitempty"`
	RoleId *string `json:"roleId,omitempty"`
}

// NewUserSvcListInvitesRequest instantiates a new UserSvcListInvitesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcListInvitesRequest() *UserSvcListInvitesRequest {
	this := UserSvcListInvitesRequest{}
	return &this
}

// NewUserSvcListInvitesRequestWithDefaults instantiates a new UserSvcListInvitesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcListInvitesRequestWithDefaults() *UserSvcListInvitesRequest {
	this := UserSvcListInvitesRequest{}
	return &this
}

// GetContactId returns the ContactId field value if set, zero value otherwise.
func (o *UserSvcListInvitesRequest) GetContactId() string {
	if o == nil || IsNil(o.ContactId) {
		var ret string
		return ret
	}
	return *o.ContactId
}

// GetContactIdOk returns a tuple with the ContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcListInvitesRequest) GetContactIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContactId) {
		return nil, false
	}
	return o.ContactId, true
}

// HasContactId returns a boolean if a field has been set.
func (o *UserSvcListInvitesRequest) HasContactId() bool {
	if o != nil && !IsNil(o.ContactId) {
		return true
	}

	return false
}

// SetContactId gets a reference to the given string and assigns it to the ContactId field.
func (o *UserSvcListInvitesRequest) SetContactId(v string) {
	o.ContactId = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *UserSvcListInvitesRequest) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcListInvitesRequest) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *UserSvcListInvitesRequest) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *UserSvcListInvitesRequest) SetRoleId(v string) {
	o.RoleId = &v
}

func (o UserSvcListInvitesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcListInvitesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactId) {
		toSerialize["contactId"] = o.ContactId
	}
	if !IsNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	return toSerialize, nil
}

type NullableUserSvcListInvitesRequest struct {
	value *UserSvcListInvitesRequest
	isSet bool
}

func (v NullableUserSvcListInvitesRequest) Get() *UserSvcListInvitesRequest {
	return v.value
}

func (v *NullableUserSvcListInvitesRequest) Set(val *UserSvcListInvitesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcListInvitesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcListInvitesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcListInvitesRequest(val *UserSvcListInvitesRequest) *NullableUserSvcListInvitesRequest {
	return &NullableUserSvcListInvitesRequest{value: val, isSet: true}
}

func (v NullableUserSvcListInvitesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcListInvitesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


