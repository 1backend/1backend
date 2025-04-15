/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.37
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcSaveSelfRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcSaveSelfRequest{}

// UserSvcSaveSelfRequest struct for UserSvcSaveSelfRequest
type UserSvcSaveSelfRequest struct {
	Labels *map[string]string `json:"labels,omitempty"`
	Name *string `json:"name,omitempty"`
	ThumbnailFileId *string `json:"thumbnailFileId,omitempty"`
}

// NewUserSvcSaveSelfRequest instantiates a new UserSvcSaveSelfRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcSaveSelfRequest() *UserSvcSaveSelfRequest {
	this := UserSvcSaveSelfRequest{}
	return &this
}

// NewUserSvcSaveSelfRequestWithDefaults instantiates a new UserSvcSaveSelfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcSaveSelfRequestWithDefaults() *UserSvcSaveSelfRequest {
	this := UserSvcSaveSelfRequest{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UserSvcSaveSelfRequest) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSaveSelfRequest) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *UserSvcSaveSelfRequest) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *UserSvcSaveSelfRequest) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserSvcSaveSelfRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSaveSelfRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserSvcSaveSelfRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserSvcSaveSelfRequest) SetName(v string) {
	o.Name = &v
}

// GetThumbnailFileId returns the ThumbnailFileId field value if set, zero value otherwise.
func (o *UserSvcSaveSelfRequest) GetThumbnailFileId() string {
	if o == nil || IsNil(o.ThumbnailFileId) {
		var ret string
		return ret
	}
	return *o.ThumbnailFileId
}

// GetThumbnailFileIdOk returns a tuple with the ThumbnailFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcSaveSelfRequest) GetThumbnailFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailFileId) {
		return nil, false
	}
	return o.ThumbnailFileId, true
}

// HasThumbnailFileId returns a boolean if a field has been set.
func (o *UserSvcSaveSelfRequest) HasThumbnailFileId() bool {
	if o != nil && !IsNil(o.ThumbnailFileId) {
		return true
	}

	return false
}

// SetThumbnailFileId gets a reference to the given string and assigns it to the ThumbnailFileId field.
func (o *UserSvcSaveSelfRequest) SetThumbnailFileId(v string) {
	o.ThumbnailFileId = &v
}

func (o UserSvcSaveSelfRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcSaveSelfRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ThumbnailFileId) {
		toSerialize["thumbnailFileId"] = o.ThumbnailFileId
	}
	return toSerialize, nil
}

type NullableUserSvcSaveSelfRequest struct {
	value *UserSvcSaveSelfRequest
	isSet bool
}

func (v NullableUserSvcSaveSelfRequest) Get() *UserSvcSaveSelfRequest {
	return v.value
}

func (v *NullableUserSvcSaveSelfRequest) Set(val *UserSvcSaveSelfRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcSaveSelfRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcSaveSelfRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcSaveSelfRequest(val *UserSvcSaveSelfRequest) *NullableUserSvcSaveSelfRequest {
	return &NullableUserSvcSaveSelfRequest{value: val, isSet: true}
}

func (v NullableUserSvcSaveSelfRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcSaveSelfRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


