/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.39
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcCreateUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcCreateUserRequest{}

// UserSvcCreateUserRequest struct for UserSvcCreateUserRequest
type UserSvcCreateUserRequest struct {
	Contacts []UserSvcContact `json:"contacts,omitempty"`
	Password *string `json:"password,omitempty"`
	RoleIds []string `json:"roleIds,omitempty"`
	User *UserSvcUser `json:"user,omitempty"`
}

// NewUserSvcCreateUserRequest instantiates a new UserSvcCreateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcCreateUserRequest() *UserSvcCreateUserRequest {
	this := UserSvcCreateUserRequest{}
	return &this
}

// NewUserSvcCreateUserRequestWithDefaults instantiates a new UserSvcCreateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcCreateUserRequestWithDefaults() *UserSvcCreateUserRequest {
	this := UserSvcCreateUserRequest{}
	return &this
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *UserSvcCreateUserRequest) GetContacts() []UserSvcContact {
	if o == nil || IsNil(o.Contacts) {
		var ret []UserSvcContact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcCreateUserRequest) GetContactsOk() ([]UserSvcContact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *UserSvcCreateUserRequest) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []UserSvcContact and assigns it to the Contacts field.
func (o *UserSvcCreateUserRequest) SetContacts(v []UserSvcContact) {
	o.Contacts = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserSvcCreateUserRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcCreateUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserSvcCreateUserRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserSvcCreateUserRequest) SetPassword(v string) {
	o.Password = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *UserSvcCreateUserRequest) GetRoleIds() []string {
	if o == nil || IsNil(o.RoleIds) {
		var ret []string
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcCreateUserRequest) GetRoleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *UserSvcCreateUserRequest) HasRoleIds() bool {
	if o != nil && !IsNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *UserSvcCreateUserRequest) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserSvcCreateUserRequest) GetUser() UserSvcUser {
	if o == nil || IsNil(o.User) {
		var ret UserSvcUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcCreateUserRequest) GetUserOk() (*UserSvcUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserSvcCreateUserRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserSvcUser and assigns it to the User field.
func (o *UserSvcCreateUserRequest) SetUser(v UserSvcUser) {
	o.User = &v
}

func (o UserSvcCreateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcCreateUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.RoleIds) {
		toSerialize["roleIds"] = o.RoleIds
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUserSvcCreateUserRequest struct {
	value *UserSvcCreateUserRequest
	isSet bool
}

func (v NullableUserSvcCreateUserRequest) Get() *UserSvcCreateUserRequest {
	return v.value
}

func (v *NullableUserSvcCreateUserRequest) Set(val *UserSvcCreateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcCreateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcCreateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcCreateUserRequest(val *UserSvcCreateUserRequest) *NullableUserSvcCreateUserRequest {
	return &NullableUserSvcCreateUserRequest{value: val, isSet: true}
}

func (v NullableUserSvcCreateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcCreateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


