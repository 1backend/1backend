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

// checks if the UserSvcGrant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcGrant{}

// UserSvcGrant struct for UserSvcGrant
type UserSvcGrant struct {
	Id *string `json:"id,omitempty"`
	PermissionId *string `json:"permissionId,omitempty"`
	// Slugs who are granted the PermissionId
	Slugs []string `json:"slugs,omitempty"`
}

// NewUserSvcGrant instantiates a new UserSvcGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcGrant() *UserSvcGrant {
	this := UserSvcGrant{}
	return &this
}

// NewUserSvcGrantWithDefaults instantiates a new UserSvcGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcGrantWithDefaults() *UserSvcGrant {
	this := UserSvcGrant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserSvcGrant) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcGrant) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserSvcGrant) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserSvcGrant) SetId(v string) {
	o.Id = &v
}

// GetPermissionId returns the PermissionId field value if set, zero value otherwise.
func (o *UserSvcGrant) GetPermissionId() string {
	if o == nil || IsNil(o.PermissionId) {
		var ret string
		return ret
	}
	return *o.PermissionId
}

// GetPermissionIdOk returns a tuple with the PermissionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcGrant) GetPermissionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PermissionId) {
		return nil, false
	}
	return o.PermissionId, true
}

// HasPermissionId returns a boolean if a field has been set.
func (o *UserSvcGrant) HasPermissionId() bool {
	if o != nil && !IsNil(o.PermissionId) {
		return true
	}

	return false
}

// SetPermissionId gets a reference to the given string and assigns it to the PermissionId field.
func (o *UserSvcGrant) SetPermissionId(v string) {
	o.PermissionId = &v
}

// GetSlugs returns the Slugs field value if set, zero value otherwise.
func (o *UserSvcGrant) GetSlugs() []string {
	if o == nil || IsNil(o.Slugs) {
		var ret []string
		return ret
	}
	return o.Slugs
}

// GetSlugsOk returns a tuple with the Slugs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcGrant) GetSlugsOk() ([]string, bool) {
	if o == nil || IsNil(o.Slugs) {
		return nil, false
	}
	return o.Slugs, true
}

// HasSlugs returns a boolean if a field has been set.
func (o *UserSvcGrant) HasSlugs() bool {
	if o != nil && !IsNil(o.Slugs) {
		return true
	}

	return false
}

// SetSlugs gets a reference to the given []string and assigns it to the Slugs field.
func (o *UserSvcGrant) SetSlugs(v []string) {
	o.Slugs = v
}

func (o UserSvcGrant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcGrant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PermissionId) {
		toSerialize["permissionId"] = o.PermissionId
	}
	if !IsNil(o.Slugs) {
		toSerialize["slugs"] = o.Slugs
	}
	return toSerialize, nil
}

type NullableUserSvcGrant struct {
	value *UserSvcGrant
	isSet bool
}

func (v NullableUserSvcGrant) Get() *UserSvcGrant {
	return v.value
}

func (v *NullableUserSvcGrant) Set(val *UserSvcGrant) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcGrant(val *UserSvcGrant) *NullableUserSvcGrant {
	return &NullableUserSvcGrant{value: val, isSet: true}
}

func (v NullableUserSvcGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


