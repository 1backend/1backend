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

// checks if the UserSvcListGrantsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcListGrantsRequest{}

// UserSvcListGrantsRequest struct for UserSvcListGrantsRequest
type UserSvcListGrantsRequest struct {
	Permission *string `json:"permission,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

// NewUserSvcListGrantsRequest instantiates a new UserSvcListGrantsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcListGrantsRequest() *UserSvcListGrantsRequest {
	this := UserSvcListGrantsRequest{}
	return &this
}

// NewUserSvcListGrantsRequestWithDefaults instantiates a new UserSvcListGrantsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcListGrantsRequestWithDefaults() *UserSvcListGrantsRequest {
	this := UserSvcListGrantsRequest{}
	return &this
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *UserSvcListGrantsRequest) GetPermission() string {
	if o == nil || IsNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcListGrantsRequest) GetPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *UserSvcListGrantsRequest) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *UserSvcListGrantsRequest) SetPermission(v string) {
	o.Permission = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UserSvcListGrantsRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcListGrantsRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UserSvcListGrantsRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UserSvcListGrantsRequest) SetSlug(v string) {
	o.Slug = &v
}

func (o UserSvcListGrantsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcListGrantsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	return toSerialize, nil
}

type NullableUserSvcListGrantsRequest struct {
	value *UserSvcListGrantsRequest
	isSet bool
}

func (v NullableUserSvcListGrantsRequest) Get() *UserSvcListGrantsRequest {
	return v.value
}

func (v *NullableUserSvcListGrantsRequest) Set(val *UserSvcListGrantsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcListGrantsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcListGrantsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcListGrantsRequest(val *UserSvcListGrantsRequest) *NullableUserSvcListGrantsRequest {
	return &NullableUserSvcListGrantsRequest{value: val, isSet: true}
}

func (v NullableUserSvcListGrantsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcListGrantsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


