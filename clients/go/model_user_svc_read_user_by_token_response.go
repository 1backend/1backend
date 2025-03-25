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

// checks if the UserSvcReadUserByTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcReadUserByTokenResponse{}

// UserSvcReadUserByTokenResponse struct for UserSvcReadUserByTokenResponse
type UserSvcReadUserByTokenResponse struct {
	ActiveOrganizationId *string `json:"activeOrganizationId,omitempty"`
	Organizations []UserSvcOrganization `json:"organizations,omitempty"`
	User *UserSvcUser `json:"user,omitempty"`
}

// NewUserSvcReadUserByTokenResponse instantiates a new UserSvcReadUserByTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcReadUserByTokenResponse() *UserSvcReadUserByTokenResponse {
	this := UserSvcReadUserByTokenResponse{}
	return &this
}

// NewUserSvcReadUserByTokenResponseWithDefaults instantiates a new UserSvcReadUserByTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcReadUserByTokenResponseWithDefaults() *UserSvcReadUserByTokenResponse {
	this := UserSvcReadUserByTokenResponse{}
	return &this
}

// GetActiveOrganizationId returns the ActiveOrganizationId field value if set, zero value otherwise.
func (o *UserSvcReadUserByTokenResponse) GetActiveOrganizationId() string {
	if o == nil || IsNil(o.ActiveOrganizationId) {
		var ret string
		return ret
	}
	return *o.ActiveOrganizationId
}

// GetActiveOrganizationIdOk returns a tuple with the ActiveOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcReadUserByTokenResponse) GetActiveOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActiveOrganizationId) {
		return nil, false
	}
	return o.ActiveOrganizationId, true
}

// HasActiveOrganizationId returns a boolean if a field has been set.
func (o *UserSvcReadUserByTokenResponse) HasActiveOrganizationId() bool {
	if o != nil && !IsNil(o.ActiveOrganizationId) {
		return true
	}

	return false
}

// SetActiveOrganizationId gets a reference to the given string and assigns it to the ActiveOrganizationId field.
func (o *UserSvcReadUserByTokenResponse) SetActiveOrganizationId(v string) {
	o.ActiveOrganizationId = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *UserSvcReadUserByTokenResponse) GetOrganizations() []UserSvcOrganization {
	if o == nil || IsNil(o.Organizations) {
		var ret []UserSvcOrganization
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcReadUserByTokenResponse) GetOrganizationsOk() ([]UserSvcOrganization, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *UserSvcReadUserByTokenResponse) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []UserSvcOrganization and assigns it to the Organizations field.
func (o *UserSvcReadUserByTokenResponse) SetOrganizations(v []UserSvcOrganization) {
	o.Organizations = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserSvcReadUserByTokenResponse) GetUser() UserSvcUser {
	if o == nil || IsNil(o.User) {
		var ret UserSvcUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcReadUserByTokenResponse) GetUserOk() (*UserSvcUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserSvcReadUserByTokenResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserSvcUser and assigns it to the User field.
func (o *UserSvcReadUserByTokenResponse) SetUser(v UserSvcUser) {
	o.User = &v
}

func (o UserSvcReadUserByTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcReadUserByTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveOrganizationId) {
		toSerialize["activeOrganizationId"] = o.ActiveOrganizationId
	}
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUserSvcReadUserByTokenResponse struct {
	value *UserSvcReadUserByTokenResponse
	isSet bool
}

func (v NullableUserSvcReadUserByTokenResponse) Get() *UserSvcReadUserByTokenResponse {
	return v.value
}

func (v *NullableUserSvcReadUserByTokenResponse) Set(val *UserSvcReadUserByTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcReadUserByTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcReadUserByTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcReadUserByTokenResponse(val *UserSvcReadUserByTokenResponse) *NullableUserSvcReadUserByTokenResponse {
	return &NullableUserSvcReadUserByTokenResponse{value: val, isSet: true}
}

func (v NullableUserSvcReadUserByTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcReadUserByTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


