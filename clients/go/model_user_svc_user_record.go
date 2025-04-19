/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.39
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserSvcUserRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcUserRecord{}

// UserSvcUserRecord struct for UserSvcUserRecord
type UserSvcUserRecord struct {
	ContactIds []string `json:"contactIds,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Id string `json:"id"`
	// Full name of the user.
	Name *string `json:"name,omitempty"`
	Roles []string `json:"roles,omitempty"`
	// URL-friendly unique (inside the 1Backend platform) identifier for the `user`.
	Slug string `json:"slug"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type _UserSvcUserRecord UserSvcUserRecord

// NewUserSvcUserRecord instantiates a new UserSvcUserRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcUserRecord(id string, slug string) *UserSvcUserRecord {
	this := UserSvcUserRecord{}
	this.Id = id
	this.Slug = slug
	return &this
}

// NewUserSvcUserRecordWithDefaults instantiates a new UserSvcUserRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcUserRecordWithDefaults() *UserSvcUserRecord {
	this := UserSvcUserRecord{}
	return &this
}

// GetContactIds returns the ContactIds field value if set, zero value otherwise.
func (o *UserSvcUserRecord) GetContactIds() []string {
	if o == nil || IsNil(o.ContactIds) {
		var ret []string
		return ret
	}
	return o.ContactIds
}

// GetContactIdsOk returns a tuple with the ContactIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUserRecord) GetContactIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ContactIds) {
		return nil, false
	}
	return o.ContactIds, true
}

// HasContactIds returns a boolean if a field has been set.
func (o *UserSvcUserRecord) HasContactIds() bool {
	if o != nil && !IsNil(o.ContactIds) {
		return true
	}

	return false
}

// SetContactIds gets a reference to the given []string and assigns it to the ContactIds field.
func (o *UserSvcUserRecord) SetContactIds(v []string) {
	o.ContactIds = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserSvcUserRecord) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUserRecord) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserSvcUserRecord) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *UserSvcUserRecord) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value
func (o *UserSvcUserRecord) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserSvcUserRecord) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserSvcUserRecord) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserSvcUserRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUserRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserSvcUserRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserSvcUserRecord) SetName(v string) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserSvcUserRecord) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUserRecord) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserSvcUserRecord) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UserSvcUserRecord) SetRoles(v []string) {
	o.Roles = v
}

// GetSlug returns the Slug field value
func (o *UserSvcUserRecord) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *UserSvcUserRecord) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *UserSvcUserRecord) SetSlug(v string) {
	o.Slug = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserSvcUserRecord) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcUserRecord) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserSvcUserRecord) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *UserSvcUserRecord) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o UserSvcUserRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcUserRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactIds) {
		toSerialize["contactIds"] = o.ContactIds
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	toSerialize["slug"] = o.Slug
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *UserSvcUserRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varUserSvcUserRecord := _UserSvcUserRecord{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcUserRecord)

	if err != nil {
		return err
	}

	*o = UserSvcUserRecord(varUserSvcUserRecord)

	return err
}

type NullableUserSvcUserRecord struct {
	value *UserSvcUserRecord
	isSet bool
}

func (v NullableUserSvcUserRecord) Get() *UserSvcUserRecord {
	return v.value
}

func (v *NullableUserSvcUserRecord) Set(val *UserSvcUserRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcUserRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcUserRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcUserRecord(val *UserSvcUserRecord) *NullableUserSvcUserRecord {
	return &NullableUserSvcUserRecord{value: val, isSet: true}
}

func (v NullableUserSvcUserRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcUserRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


