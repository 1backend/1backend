/*
1Backend

A unified backend for your AI applications—microservices-based and built to scale.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcCreateOrganizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcCreateOrganizationRequest{}

// UserSvcCreateOrganizationRequest struct for UserSvcCreateOrganizationRequest
type UserSvcCreateOrganizationRequest struct {
	Id *string `json:"id,omitempty"`
	// Full name of the organization.
	Name *string `json:"name,omitempty"`
	// URL-friendly unique (inside the Singularon platform) identifier for the `organization`.
	Slug *string `json:"slug,omitempty"`
}

// NewUserSvcCreateOrganizationRequest instantiates a new UserSvcCreateOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcCreateOrganizationRequest() *UserSvcCreateOrganizationRequest {
	this := UserSvcCreateOrganizationRequest{}
	return &this
}

// NewUserSvcCreateOrganizationRequestWithDefaults instantiates a new UserSvcCreateOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcCreateOrganizationRequestWithDefaults() *UserSvcCreateOrganizationRequest {
	this := UserSvcCreateOrganizationRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserSvcCreateOrganizationRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcCreateOrganizationRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserSvcCreateOrganizationRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserSvcCreateOrganizationRequest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserSvcCreateOrganizationRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcCreateOrganizationRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserSvcCreateOrganizationRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserSvcCreateOrganizationRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UserSvcCreateOrganizationRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcCreateOrganizationRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UserSvcCreateOrganizationRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UserSvcCreateOrganizationRequest) SetSlug(v string) {
	o.Slug = &v
}

func (o UserSvcCreateOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcCreateOrganizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	return toSerialize, nil
}

type NullableUserSvcCreateOrganizationRequest struct {
	value *UserSvcCreateOrganizationRequest
	isSet bool
}

func (v NullableUserSvcCreateOrganizationRequest) Get() *UserSvcCreateOrganizationRequest {
	return v.value
}

func (v *NullableUserSvcCreateOrganizationRequest) Set(val *UserSvcCreateOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcCreateOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcCreateOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcCreateOrganizationRequest(val *UserSvcCreateOrganizationRequest) *NullableUserSvcCreateOrganizationRequest {
	return &NullableUserSvcCreateOrganizationRequest{value: val, isSet: true}
}

func (v NullableUserSvcCreateOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcCreateOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


