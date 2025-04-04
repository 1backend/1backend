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

// checks if the UserSvcSaveProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcSaveProfileRequest{}

// UserSvcSaveProfileRequest struct for UserSvcSaveProfileRequest
type UserSvcSaveProfileRequest struct {
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
	ThumbnailFileId *string `json:"thumbnailFileId,omitempty"`
}

// NewUserSvcSaveProfileRequest instantiates a new UserSvcSaveProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcSaveProfileRequest() *UserSvcSaveProfileRequest {
	this := UserSvcSaveProfileRequest{}
	return &this
}

// NewUserSvcSaveProfileRequestWithDefaults instantiates a new UserSvcSaveProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcSaveProfileRequestWithDefaults() *UserSvcSaveProfileRequest {
	this := UserSvcSaveProfileRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserSvcSaveProfileRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSaveProfileRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserSvcSaveProfileRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserSvcSaveProfileRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UserSvcSaveProfileRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSaveProfileRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UserSvcSaveProfileRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UserSvcSaveProfileRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetThumbnailFileId returns the ThumbnailFileId field value if set, zero value otherwise.
func (o *UserSvcSaveProfileRequest) GetThumbnailFileId() string {
	if o == nil || IsNil(o.ThumbnailFileId) {
		var ret string
		return ret
	}
	return *o.ThumbnailFileId
}

// GetThumbnailFileIdOk returns a tuple with the ThumbnailFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSaveProfileRequest) GetThumbnailFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailFileId) {
		return nil, false
	}
	return o.ThumbnailFileId, true
}

// HasThumbnailFileId returns a boolean if a field has been set.
func (o *UserSvcSaveProfileRequest) HasThumbnailFileId() bool {
	if o != nil && !IsNil(o.ThumbnailFileId) {
		return true
	}

	return false
}

// SetThumbnailFileId gets a reference to the given string and assigns it to the ThumbnailFileId field.
func (o *UserSvcSaveProfileRequest) SetThumbnailFileId(v string) {
	o.ThumbnailFileId = &v
}

func (o UserSvcSaveProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcSaveProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.ThumbnailFileId) {
		toSerialize["thumbnailFileId"] = o.ThumbnailFileId
	}
	return toSerialize, nil
}

type NullableUserSvcSaveProfileRequest struct {
	value *UserSvcSaveProfileRequest
	isSet bool
}

func (v NullableUserSvcSaveProfileRequest) Get() *UserSvcSaveProfileRequest {
	return v.value
}

func (v *NullableUserSvcSaveProfileRequest) Set(val *UserSvcSaveProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcSaveProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcSaveProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcSaveProfileRequest(val *UserSvcSaveProfileRequest) *NullableUserSvcSaveProfileRequest {
	return &NullableUserSvcSaveProfileRequest{value: val, isSet: true}
}

func (v NullableUserSvcSaveProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcSaveProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


