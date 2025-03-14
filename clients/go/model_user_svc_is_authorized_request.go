/*
1Backend

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcIsAuthorizedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcIsAuthorizedRequest{}

// UserSvcIsAuthorizedRequest struct for UserSvcIsAuthorizedRequest
type UserSvcIsAuthorizedRequest struct {
	ContactsGranted []string `json:"contactsGranted,omitempty"`
	GrantedSlugs []string `json:"grantedSlugs,omitempty"`
}

// NewUserSvcIsAuthorizedRequest instantiates a new UserSvcIsAuthorizedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcIsAuthorizedRequest() *UserSvcIsAuthorizedRequest {
	this := UserSvcIsAuthorizedRequest{}
	return &this
}

// NewUserSvcIsAuthorizedRequestWithDefaults instantiates a new UserSvcIsAuthorizedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcIsAuthorizedRequestWithDefaults() *UserSvcIsAuthorizedRequest {
	this := UserSvcIsAuthorizedRequest{}
	return &this
}

// GetContactsGranted returns the ContactsGranted field value if set, zero value otherwise.
func (o *UserSvcIsAuthorizedRequest) GetContactsGranted() []string {
	if o == nil || IsNil(o.ContactsGranted) {
		var ret []string
		return ret
	}
	return o.ContactsGranted
}

// GetContactsGrantedOk returns a tuple with the ContactsGranted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcIsAuthorizedRequest) GetContactsGrantedOk() ([]string, bool) {
	if o == nil || IsNil(o.ContactsGranted) {
		return nil, false
	}
	return o.ContactsGranted, true
}

// HasContactsGranted returns a boolean if a field has been set.
func (o *UserSvcIsAuthorizedRequest) HasContactsGranted() bool {
	if o != nil && !IsNil(o.ContactsGranted) {
		return true
	}

	return false
}

// SetContactsGranted gets a reference to the given []string and assigns it to the ContactsGranted field.
func (o *UserSvcIsAuthorizedRequest) SetContactsGranted(v []string) {
	o.ContactsGranted = v
}

// GetGrantedSlugs returns the GrantedSlugs field value if set, zero value otherwise.
func (o *UserSvcIsAuthorizedRequest) GetGrantedSlugs() []string {
	if o == nil || IsNil(o.GrantedSlugs) {
		var ret []string
		return ret
	}
	return o.GrantedSlugs
}

// GetGrantedSlugsOk returns a tuple with the GrantedSlugs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcIsAuthorizedRequest) GetGrantedSlugsOk() ([]string, bool) {
	if o == nil || IsNil(o.GrantedSlugs) {
		return nil, false
	}
	return o.GrantedSlugs, true
}

// HasGrantedSlugs returns a boolean if a field has been set.
func (o *UserSvcIsAuthorizedRequest) HasGrantedSlugs() bool {
	if o != nil && !IsNil(o.GrantedSlugs) {
		return true
	}

	return false
}

// SetGrantedSlugs gets a reference to the given []string and assigns it to the GrantedSlugs field.
func (o *UserSvcIsAuthorizedRequest) SetGrantedSlugs(v []string) {
	o.GrantedSlugs = v
}

func (o UserSvcIsAuthorizedRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcIsAuthorizedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactsGranted) {
		toSerialize["contactsGranted"] = o.ContactsGranted
	}
	if !IsNil(o.GrantedSlugs) {
		toSerialize["grantedSlugs"] = o.GrantedSlugs
	}
	return toSerialize, nil
}

type NullableUserSvcIsAuthorizedRequest struct {
	value *UserSvcIsAuthorizedRequest
	isSet bool
}

func (v NullableUserSvcIsAuthorizedRequest) Get() *UserSvcIsAuthorizedRequest {
	return v.value
}

func (v *NullableUserSvcIsAuthorizedRequest) Set(val *UserSvcIsAuthorizedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcIsAuthorizedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcIsAuthorizedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcIsAuthorizedRequest(val *UserSvcIsAuthorizedRequest) *NullableUserSvcIsAuthorizedRequest {
	return &NullableUserSvcIsAuthorizedRequest{value: val, isSet: true}
}

func (v NullableUserSvcIsAuthorizedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcIsAuthorizedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


