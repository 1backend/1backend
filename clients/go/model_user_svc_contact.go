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
	"bytes"
	"fmt"
)

// checks if the UserSvcContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcContact{}

// UserSvcContact struct for UserSvcContact
type UserSvcContact struct {
	CreatedAt *string `json:"createdAt,omitempty"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	// Handle is the platform local unique identifier. Ie. while the `id` of a Twitter contact is `twitter.com/thejoe`, the value will be only `thejoe`. For email and phones the `id` and the `value` will be the same. This field mostly exists for display purposes.  Example values: \"joe12\" (1backend username), \"thejoe\" (twitter username), \"joe@joesdomain.com\" (email)
	Handle string `json:"handle"`
	// The unique identifier, which can be a URL.  Example values: \"joe12\" (1backend username), \"twitter.com/thejoe\" (twitter url), \"joe@joesdomain.com\" (email)
	Id string `json:"id"`
	// If this is the primary contact method
	IsPrimary *bool `json:"isPrimary,omitempty"`
	// Platform of the contact (e.g., \"email\", \"phone\", \"twitter\")
	Platform string `json:"platform"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UserId string `json:"userId"`
	// Whether the contact is verified
	Verified *bool `json:"verified,omitempty"`
}

type _UserSvcContact UserSvcContact

// NewUserSvcContact instantiates a new UserSvcContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcContact(handle string, id string, platform string, userId string) *UserSvcContact {
	this := UserSvcContact{}
	this.Handle = handle
	this.Id = id
	this.Platform = platform
	this.UserId = userId
	return &this
}

// NewUserSvcContactWithDefaults instantiates a new UserSvcContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcContactWithDefaults() *UserSvcContact {
	this := UserSvcContact{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserSvcContact) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcContact) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserSvcContact) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *UserSvcContact) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *UserSvcContact) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcContact) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *UserSvcContact) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *UserSvcContact) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetHandle returns the Handle field value
func (o *UserSvcContact) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *UserSvcContact) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *UserSvcContact) SetHandle(v string) {
	o.Handle = v
}

// GetId returns the Id field value
func (o *UserSvcContact) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserSvcContact) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserSvcContact) SetId(v string) {
	o.Id = v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *UserSvcContact) GetIsPrimary() bool {
	if o == nil || IsNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcContact) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimary) {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *UserSvcContact) HasIsPrimary() bool {
	if o != nil && !IsNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *UserSvcContact) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetPlatform returns the Platform field value
func (o *UserSvcContact) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *UserSvcContact) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *UserSvcContact) SetPlatform(v string) {
	o.Platform = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserSvcContact) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcContact) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserSvcContact) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *UserSvcContact) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value
func (o *UserSvcContact) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserSvcContact) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserSvcContact) SetUserId(v string) {
	o.UserId = v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *UserSvcContact) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcContact) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *UserSvcContact) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *UserSvcContact) SetVerified(v bool) {
	o.Verified = &v
}

func (o UserSvcContact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	toSerialize["handle"] = o.Handle
	toSerialize["id"] = o.Id
	if !IsNil(o.IsPrimary) {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	toSerialize["platform"] = o.Platform
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	toSerialize["userId"] = o.UserId
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	return toSerialize, nil
}

func (o *UserSvcContact) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"handle",
		"id",
		"platform",
		"userId",
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

	varUserSvcContact := _UserSvcContact{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcContact)

	if err != nil {
		return err
	}

	*o = UserSvcContact(varUserSvcContact)

	return err
}

type NullableUserSvcContact struct {
	value *UserSvcContact
	isSet bool
}

func (v NullableUserSvcContact) Get() *UserSvcContact {
	return v.value
}

func (v *NullableUserSvcContact) Set(val *UserSvcContact) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcContact) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcContact(val *UserSvcContact) *NullableUserSvcContact {
	return &NullableUserSvcContact{value: val, isSet: true}
}

func (v NullableUserSvcContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


