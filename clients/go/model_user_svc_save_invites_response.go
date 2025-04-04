/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.31
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserSvcSaveInvitesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcSaveInvitesResponse{}

// UserSvcSaveInvitesResponse struct for UserSvcSaveInvitesResponse
type UserSvcSaveInvitesResponse struct {
	Invites []UserSvcInvite `json:"invites"`
}

type _UserSvcSaveInvitesResponse UserSvcSaveInvitesResponse

// NewUserSvcSaveInvitesResponse instantiates a new UserSvcSaveInvitesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcSaveInvitesResponse(invites []UserSvcInvite) *UserSvcSaveInvitesResponse {
	this := UserSvcSaveInvitesResponse{}
	this.Invites = invites
	return &this
}

// NewUserSvcSaveInvitesResponseWithDefaults instantiates a new UserSvcSaveInvitesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcSaveInvitesResponseWithDefaults() *UserSvcSaveInvitesResponse {
	this := UserSvcSaveInvitesResponse{}
	return &this
}

// GetInvites returns the Invites field value
func (o *UserSvcSaveInvitesResponse) GetInvites() []UserSvcInvite {
	if o == nil {
		var ret []UserSvcInvite
		return ret
	}

	return o.Invites
}

// GetInvitesOk returns a tuple with the Invites field value
// and a boolean to check if the value has been set.
func (o *UserSvcSaveInvitesResponse) GetInvitesOk() ([]UserSvcInvite, bool) {
	if o == nil {
		return nil, false
	}
	return o.Invites, true
}

// SetInvites sets field value
func (o *UserSvcSaveInvitesResponse) SetInvites(v []UserSvcInvite) {
	o.Invites = v
}

func (o UserSvcSaveInvitesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcSaveInvitesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invites"] = o.Invites
	return toSerialize, nil
}

func (o *UserSvcSaveInvitesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invites",
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

	varUserSvcSaveInvitesResponse := _UserSvcSaveInvitesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcSaveInvitesResponse)

	if err != nil {
		return err
	}

	*o = UserSvcSaveInvitesResponse(varUserSvcSaveInvitesResponse)

	return err
}

type NullableUserSvcSaveInvitesResponse struct {
	value *UserSvcSaveInvitesResponse
	isSet bool
}

func (v NullableUserSvcSaveInvitesResponse) Get() *UserSvcSaveInvitesResponse {
	return v.value
}

func (v *NullableUserSvcSaveInvitesResponse) Set(val *UserSvcSaveInvitesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcSaveInvitesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcSaveInvitesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcSaveInvitesResponse(val *UserSvcSaveInvitesResponse) *NullableUserSvcSaveInvitesResponse {
	return &NullableUserSvcSaveInvitesResponse{value: val, isSet: true}
}

func (v NullableUserSvcSaveInvitesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcSaveInvitesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


