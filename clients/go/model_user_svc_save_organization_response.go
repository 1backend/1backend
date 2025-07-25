/*
1Backend

AI-native microservices platform.

API version: 0.7.6
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserSvcSaveOrganizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcSaveOrganizationResponse{}

// UserSvcSaveOrganizationResponse struct for UserSvcSaveOrganizationResponse
type UserSvcSaveOrganizationResponse struct {
	Organization UserSvcOrganization `json:"organization"`
	// Due to the nature of JWT tokens, the token must be refreshed after creating an organization, as dynamic organization roles are embedded in it.
	Token UserSvcAuthToken `json:"token"`
}

type _UserSvcSaveOrganizationResponse UserSvcSaveOrganizationResponse

// NewUserSvcSaveOrganizationResponse instantiates a new UserSvcSaveOrganizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcSaveOrganizationResponse(organization UserSvcOrganization, token UserSvcAuthToken) *UserSvcSaveOrganizationResponse {
	this := UserSvcSaveOrganizationResponse{}
	this.Organization = organization
	this.Token = token
	return &this
}

// NewUserSvcSaveOrganizationResponseWithDefaults instantiates a new UserSvcSaveOrganizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcSaveOrganizationResponseWithDefaults() *UserSvcSaveOrganizationResponse {
	this := UserSvcSaveOrganizationResponse{}
	return &this
}

// GetOrganization returns the Organization field value
func (o *UserSvcSaveOrganizationResponse) GetOrganization() UserSvcOrganization {
	if o == nil {
		var ret UserSvcOrganization
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *UserSvcSaveOrganizationResponse) GetOrganizationOk() (*UserSvcOrganization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *UserSvcSaveOrganizationResponse) SetOrganization(v UserSvcOrganization) {
	o.Organization = v
}

// GetToken returns the Token field value
func (o *UserSvcSaveOrganizationResponse) GetToken() UserSvcAuthToken {
	if o == nil {
		var ret UserSvcAuthToken
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UserSvcSaveOrganizationResponse) GetTokenOk() (*UserSvcAuthToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UserSvcSaveOrganizationResponse) SetToken(v UserSvcAuthToken) {
	o.Token = v
}

func (o UserSvcSaveOrganizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcSaveOrganizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["organization"] = o.Organization
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *UserSvcSaveOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"organization",
		"token",
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

	varUserSvcSaveOrganizationResponse := _UserSvcSaveOrganizationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcSaveOrganizationResponse)

	if err != nil {
		return err
	}

	*o = UserSvcSaveOrganizationResponse(varUserSvcSaveOrganizationResponse)

	return err
}

type NullableUserSvcSaveOrganizationResponse struct {
	value *UserSvcSaveOrganizationResponse
	isSet bool
}

func (v NullableUserSvcSaveOrganizationResponse) Get() *UserSvcSaveOrganizationResponse {
	return v.value
}

func (v *NullableUserSvcSaveOrganizationResponse) Set(val *UserSvcSaveOrganizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcSaveOrganizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcSaveOrganizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcSaveOrganizationResponse(val *UserSvcSaveOrganizationResponse) *NullableUserSvcSaveOrganizationResponse {
	return &NullableUserSvcSaveOrganizationResponse{value: val, isSet: true}
}

func (v NullableUserSvcSaveOrganizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcSaveOrganizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


