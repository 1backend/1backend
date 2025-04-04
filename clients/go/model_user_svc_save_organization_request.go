/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.32
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserSvcSaveOrganizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcSaveOrganizationRequest{}

// UserSvcSaveOrganizationRequest struct for UserSvcSaveOrganizationRequest
type UserSvcSaveOrganizationRequest struct {
	Id *string `json:"id,omitempty"`
	// Full name of the organization.
	Name *string `json:"name,omitempty"`
	// URL-friendly unique (inside the Singularon platform) identifier for the `organization`.
	Slug string `json:"slug"`
	ThumbnailFileId *string `json:"thumbnailFileId,omitempty"`
}

type _UserSvcSaveOrganizationRequest UserSvcSaveOrganizationRequest

// NewUserSvcSaveOrganizationRequest instantiates a new UserSvcSaveOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcSaveOrganizationRequest(slug string) *UserSvcSaveOrganizationRequest {
	this := UserSvcSaveOrganizationRequest{}
	this.Slug = slug
	return &this
}

// NewUserSvcSaveOrganizationRequestWithDefaults instantiates a new UserSvcSaveOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcSaveOrganizationRequestWithDefaults() *UserSvcSaveOrganizationRequest {
	this := UserSvcSaveOrganizationRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserSvcSaveOrganizationRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSaveOrganizationRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserSvcSaveOrganizationRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserSvcSaveOrganizationRequest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserSvcSaveOrganizationRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSaveOrganizationRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserSvcSaveOrganizationRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserSvcSaveOrganizationRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value
func (o *UserSvcSaveOrganizationRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *UserSvcSaveOrganizationRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *UserSvcSaveOrganizationRequest) SetSlug(v string) {
	o.Slug = v
}

// GetThumbnailFileId returns the ThumbnailFileId field value if set, zero value otherwise.
func (o *UserSvcSaveOrganizationRequest) GetThumbnailFileId() string {
	if o == nil || IsNil(o.ThumbnailFileId) {
		var ret string
		return ret
	}
	return *o.ThumbnailFileId
}

// GetThumbnailFileIdOk returns a tuple with the ThumbnailFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSaveOrganizationRequest) GetThumbnailFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailFileId) {
		return nil, false
	}
	return o.ThumbnailFileId, true
}

// HasThumbnailFileId returns a boolean if a field has been set.
func (o *UserSvcSaveOrganizationRequest) HasThumbnailFileId() bool {
	if o != nil && !IsNil(o.ThumbnailFileId) {
		return true
	}

	return false
}

// SetThumbnailFileId gets a reference to the given string and assigns it to the ThumbnailFileId field.
func (o *UserSvcSaveOrganizationRequest) SetThumbnailFileId(v string) {
	o.ThumbnailFileId = &v
}

func (o UserSvcSaveOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcSaveOrganizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["slug"] = o.Slug
	if !IsNil(o.ThumbnailFileId) {
		toSerialize["thumbnailFileId"] = o.ThumbnailFileId
	}
	return toSerialize, nil
}

func (o *UserSvcSaveOrganizationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"slug",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserSvcSaveOrganizationRequest := _UserSvcSaveOrganizationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcSaveOrganizationRequest)

	if err != nil {
		return err
	}

	*o = UserSvcSaveOrganizationRequest(varUserSvcSaveOrganizationRequest)

	return err
}

type NullableUserSvcSaveOrganizationRequest struct {
	value *UserSvcSaveOrganizationRequest
	isSet bool
}

func (v NullableUserSvcSaveOrganizationRequest) Get() *UserSvcSaveOrganizationRequest {
	return v.value
}

func (v *NullableUserSvcSaveOrganizationRequest) Set(val *UserSvcSaveOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcSaveOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcSaveOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcSaveOrganizationRequest(val *UserSvcSaveOrganizationRequest) *NullableUserSvcSaveOrganizationRequest {
	return &NullableUserSvcSaveOrganizationRequest{value: val, isSet: true}
}

func (v NullableUserSvcSaveOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcSaveOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


