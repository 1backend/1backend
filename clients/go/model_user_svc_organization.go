/*
OpenOrch

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcOrganization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcOrganization{}

// UserSvcOrganization struct for UserSvcOrganization
type UserSvcOrganization struct {
	CreatedAt *string `json:"createdAt,omitempty"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	Id *string `json:"id,omitempty"`
	// Full name of the organization
	Name *string `json:"name,omitempty"`
	// URL-friendly unique (inside the Singularon platform) identifier for the `organization`.
	Slug *string `json:"slug,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewUserSvcOrganization instantiates a new UserSvcOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcOrganization() *UserSvcOrganization {
	this := UserSvcOrganization{}
	return &this
}

// NewUserSvcOrganizationWithDefaults instantiates a new UserSvcOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcOrganizationWithDefaults() *UserSvcOrganization {
	this := UserSvcOrganization{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserSvcOrganization) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcOrganization) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserSvcOrganization) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *UserSvcOrganization) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *UserSvcOrganization) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcOrganization) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *UserSvcOrganization) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *UserSvcOrganization) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserSvcOrganization) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcOrganization) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserSvcOrganization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserSvcOrganization) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserSvcOrganization) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcOrganization) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserSvcOrganization) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserSvcOrganization) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UserSvcOrganization) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcOrganization) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UserSvcOrganization) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UserSvcOrganization) SetSlug(v string) {
	o.Slug = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserSvcOrganization) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcOrganization) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserSvcOrganization) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *UserSvcOrganization) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o UserSvcOrganization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcOrganization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableUserSvcOrganization struct {
	value *UserSvcOrganization
	isSet bool
}

func (v NullableUserSvcOrganization) Get() *UserSvcOrganization {
	return v.value
}

func (v *NullableUserSvcOrganization) Set(val *UserSvcOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcOrganization(val *UserSvcOrganization) *NullableUserSvcOrganization {
	return &NullableUserSvcOrganization{value: val, isSet: true}
}

func (v NullableUserSvcOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


