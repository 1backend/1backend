/*
1Backend

A unified backend development platform for microservices-based AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcLoginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcLoginRequest{}

// UserSvcLoginRequest struct for UserSvcLoginRequest
type UserSvcLoginRequest struct {
	Contact *string `json:"contact,omitempty"`
	Password *string `json:"password,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

// NewUserSvcLoginRequest instantiates a new UserSvcLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcLoginRequest() *UserSvcLoginRequest {
	this := UserSvcLoginRequest{}
	return &this
}

// NewUserSvcLoginRequestWithDefaults instantiates a new UserSvcLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcLoginRequestWithDefaults() *UserSvcLoginRequest {
	this := UserSvcLoginRequest{}
	return &this
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *UserSvcLoginRequest) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcLoginRequest) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *UserSvcLoginRequest) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *UserSvcLoginRequest) SetContact(v string) {
	o.Contact = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserSvcLoginRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcLoginRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserSvcLoginRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserSvcLoginRequest) SetPassword(v string) {
	o.Password = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UserSvcLoginRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcLoginRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UserSvcLoginRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UserSvcLoginRequest) SetSlug(v string) {
	o.Slug = &v
}

func (o UserSvcLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcLoginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	return toSerialize, nil
}

type NullableUserSvcLoginRequest struct {
	value *UserSvcLoginRequest
	isSet bool
}

func (v NullableUserSvcLoginRequest) Get() *UserSvcLoginRequest {
	return v.value
}

func (v *NullableUserSvcLoginRequest) Set(val *UserSvcLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcLoginRequest(val *UserSvcLoginRequest) *NullableUserSvcLoginRequest {
	return &NullableUserSvcLoginRequest{value: val, isSet: true}
}

func (v NullableUserSvcLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


