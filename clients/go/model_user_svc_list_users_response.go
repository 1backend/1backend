/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.33
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcListUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcListUsersResponse{}

// UserSvcListUsersResponse struct for UserSvcListUsersResponse
type UserSvcListUsersResponse struct {
	After *string `json:"after,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Users []UserSvcUserRecord `json:"users,omitempty"`
}

// NewUserSvcListUsersResponse instantiates a new UserSvcListUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcListUsersResponse() *UserSvcListUsersResponse {
	this := UserSvcListUsersResponse{}
	return &this
}

// NewUserSvcListUsersResponseWithDefaults instantiates a new UserSvcListUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcListUsersResponseWithDefaults() *UserSvcListUsersResponse {
	this := UserSvcListUsersResponse{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *UserSvcListUsersResponse) GetAfter() string {
	if o == nil || IsNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcListUsersResponse) GetAfterOk() (*string, bool) {
	if o == nil || IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *UserSvcListUsersResponse) HasAfter() bool {
	if o != nil && !IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *UserSvcListUsersResponse) SetAfter(v string) {
	o.After = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UserSvcListUsersResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcListUsersResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UserSvcListUsersResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UserSvcListUsersResponse) SetCount(v int32) {
	o.Count = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *UserSvcListUsersResponse) GetUsers() []UserSvcUserRecord {
	if o == nil || IsNil(o.Users) {
		var ret []UserSvcUserRecord
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcListUsersResponse) GetUsersOk() ([]UserSvcUserRecord, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UserSvcListUsersResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserSvcUserRecord and assigns it to the Users field.
func (o *UserSvcListUsersResponse) SetUsers(v []UserSvcUserRecord) {
	o.Users = v
}

func (o UserSvcListUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcListUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.After) {
		toSerialize["after"] = o.After
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableUserSvcListUsersResponse struct {
	value *UserSvcListUsersResponse
	isSet bool
}

func (v NullableUserSvcListUsersResponse) Get() *UserSvcListUsersResponse {
	return v.value
}

func (v *NullableUserSvcListUsersResponse) Set(val *UserSvcListUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcListUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcListUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcListUsersResponse(val *UserSvcListUsersResponse) *NullableUserSvcListUsersResponse {
	return &NullableUserSvcListUsersResponse{value: val, isSet: true}
}

func (v NullableUserSvcListUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcListUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


