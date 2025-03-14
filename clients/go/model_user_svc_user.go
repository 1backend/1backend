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

// checks if the UserSvcUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcUser{}

// UserSvcUser struct for UserSvcUser
type UserSvcUser struct {
	// Contacts are used for login and identification purposes.
	Contacts []UserSvcContact `json:"contacts,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	Id *string `json:"id,omitempty"`
	// Full name of the organization.
	Name *string `json:"name,omitempty"`
	PasswordHash *string `json:"passwordHash,omitempty"`
	// URL-friendly unique (inside the Singularon platform) identifier for the `user`.
	Slug *string `json:"slug,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewUserSvcUser instantiates a new UserSvcUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcUser() *UserSvcUser {
	this := UserSvcUser{}
	return &this
}

// NewUserSvcUserWithDefaults instantiates a new UserSvcUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcUserWithDefaults() *UserSvcUser {
	this := UserSvcUser{}
	return &this
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *UserSvcUser) GetContacts() []UserSvcContact {
	if o == nil || IsNil(o.Contacts) {
		var ret []UserSvcContact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUser) GetContactsOk() ([]UserSvcContact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *UserSvcUser) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []UserSvcContact and assigns it to the Contacts field.
func (o *UserSvcUser) SetContacts(v []UserSvcContact) {
	o.Contacts = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserSvcUser) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUser) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserSvcUser) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *UserSvcUser) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *UserSvcUser) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUser) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *UserSvcUser) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *UserSvcUser) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserSvcUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserSvcUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserSvcUser) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserSvcUser) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUser) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserSvcUser) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserSvcUser) SetName(v string) {
	o.Name = &v
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise.
func (o *UserSvcUser) GetPasswordHash() string {
	if o == nil || IsNil(o.PasswordHash) {
		var ret string
		return ret
	}
	return *o.PasswordHash
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUser) GetPasswordHashOk() (*string, bool) {
	if o == nil || IsNil(o.PasswordHash) {
		return nil, false
	}
	return o.PasswordHash, true
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *UserSvcUser) HasPasswordHash() bool {
	if o != nil && !IsNil(o.PasswordHash) {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given string and assigns it to the PasswordHash field.
func (o *UserSvcUser) SetPasswordHash(v string) {
	o.PasswordHash = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UserSvcUser) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUser) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UserSvcUser) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UserSvcUser) SetSlug(v string) {
	o.Slug = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserSvcUser) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUser) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserSvcUser) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *UserSvcUser) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o UserSvcUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PasswordHash) {
		toSerialize["passwordHash"] = o.PasswordHash
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableUserSvcUser struct {
	value *UserSvcUser
	isSet bool
}

func (v NullableUserSvcUser) Get() *UserSvcUser {
	return v.value
}

func (v *NullableUserSvcUser) Set(val *UserSvcUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcUser(val *UserSvcUser) *NullableUserSvcUser {
	return &NullableUserSvcUser{value: val, isSet: true}
}

func (v NullableUserSvcUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


