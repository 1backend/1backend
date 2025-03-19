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

// checks if the UserSvcRegisterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcRegisterRequest{}

// UserSvcRegisterRequest struct for UserSvcRegisterRequest
type UserSvcRegisterRequest struct {
	Contact *UserSvcContact `json:"contact,omitempty"`
	Name *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

// NewUserSvcRegisterRequest instantiates a new UserSvcRegisterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcRegisterRequest() *UserSvcRegisterRequest {
	this := UserSvcRegisterRequest{}
	return &this
}

// NewUserSvcRegisterRequestWithDefaults instantiates a new UserSvcRegisterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcRegisterRequestWithDefaults() *UserSvcRegisterRequest {
	this := UserSvcRegisterRequest{}
	return &this
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *UserSvcRegisterRequest) GetContact() UserSvcContact {
	if o == nil || IsNil(o.Contact) {
		var ret UserSvcContact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcRegisterRequest) GetContactOk() (*UserSvcContact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *UserSvcRegisterRequest) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given UserSvcContact and assigns it to the Contact field.
func (o *UserSvcRegisterRequest) SetContact(v UserSvcContact) {
	o.Contact = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserSvcRegisterRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcRegisterRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserSvcRegisterRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserSvcRegisterRequest) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserSvcRegisterRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcRegisterRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserSvcRegisterRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserSvcRegisterRequest) SetPassword(v string) {
	o.Password = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UserSvcRegisterRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcRegisterRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UserSvcRegisterRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UserSvcRegisterRequest) SetSlug(v string) {
	o.Slug = &v
}

func (o UserSvcRegisterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcRegisterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	return toSerialize, nil
}

type NullableUserSvcRegisterRequest struct {
	value *UserSvcRegisterRequest
	isSet bool
}

func (v NullableUserSvcRegisterRequest) Get() *UserSvcRegisterRequest {
	return v.value
}

func (v *NullableUserSvcRegisterRequest) Set(val *UserSvcRegisterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcRegisterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcRegisterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcRegisterRequest(val *UserSvcRegisterRequest) *NullableUserSvcRegisterRequest {
	return &NullableUserSvcRegisterRequest{value: val, isSet: true}
}

func (v NullableUserSvcRegisterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcRegisterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


