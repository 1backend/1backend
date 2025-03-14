/*
1Backend

A common backend for your AI applications—microservices-based and built to scale.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserSvcGetUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcGetUsersRequest{}

// UserSvcGetUsersRequest struct for UserSvcGetUsersRequest
type UserSvcGetUsersRequest struct {
	Query *DatastoreQuery `json:"query,omitempty"`
}

// NewUserSvcGetUsersRequest instantiates a new UserSvcGetUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcGetUsersRequest() *UserSvcGetUsersRequest {
	this := UserSvcGetUsersRequest{}
	return &this
}

// NewUserSvcGetUsersRequestWithDefaults instantiates a new UserSvcGetUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcGetUsersRequestWithDefaults() *UserSvcGetUsersRequest {
	this := UserSvcGetUsersRequest{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *UserSvcGetUsersRequest) GetQuery() DatastoreQuery {
	if o == nil || IsNil(o.Query) {
		var ret DatastoreQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcGetUsersRequest) GetQueryOk() (*DatastoreQuery, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *UserSvcGetUsersRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given DatastoreQuery and assigns it to the Query field.
func (o *UserSvcGetUsersRequest) SetQuery(v DatastoreQuery) {
	o.Query = &v
}

func (o UserSvcGetUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcGetUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return toSerialize, nil
}

type NullableUserSvcGetUsersRequest struct {
	value *UserSvcGetUsersRequest
	isSet bool
}

func (v NullableUserSvcGetUsersRequest) Get() *UserSvcGetUsersRequest {
	return v.value
}

func (v *NullableUserSvcGetUsersRequest) Set(val *UserSvcGetUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcGetUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcGetUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcGetUsersRequest(val *UserSvcGetUsersRequest) *NullableUserSvcGetUsersRequest {
	return &NullableUserSvcGetUsersRequest{value: val, isSet: true}
}

func (v NullableUserSvcGetUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcGetUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


