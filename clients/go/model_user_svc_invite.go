/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.34
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserSvcInvite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcInvite{}

// UserSvcInvite struct for UserSvcInvite
type UserSvcInvite struct {
	AppliedAt *string `json:"appliedAt,omitempty"`
	// ContactId represents the recipient of the invite. If the user is already registered, the role is assigned immediately; otherwise, it is applied upon registration.
	ContactId string `json:"contactId"`
	CreatedAt string `json:"createdAt"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	ExpiresAt *string `json:"expiresAt,omitempty"`
	Id string `json:"id"`
	// OwnerIds specifies the users who created the invite. If you create an invite that already exists for a given role and contact ID, you get added to the list of owners.
	OwnerIds []string `json:"ownerIds"`
	// RoleId specifies the role to be assigned to the ContactId. Callers can only assign roles they own, identified by their service slug (e.g., if \"my-service\" creates an invite, the role must be \"my-service:admin\"). Dynamic organization roles can also be assigned (e.g., \"user-svc:org:{%orgId}:admin\" or \"user-svc:org:{%orgId}:user\"), but in this case, the caller must be an admin of the target organization.
	RoleId string `json:"roleId"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type _UserSvcInvite UserSvcInvite

// NewUserSvcInvite instantiates a new UserSvcInvite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcInvite(contactId string, createdAt string, id string, ownerIds []string, roleId string) *UserSvcInvite {
	this := UserSvcInvite{}
	this.ContactId = contactId
	this.CreatedAt = createdAt
	this.Id = id
	this.OwnerIds = ownerIds
	this.RoleId = roleId
	return &this
}

// NewUserSvcInviteWithDefaults instantiates a new UserSvcInvite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcInviteWithDefaults() *UserSvcInvite {
	this := UserSvcInvite{}
	return &this
}

// GetAppliedAt returns the AppliedAt field value if set, zero value otherwise.
func (o *UserSvcInvite) GetAppliedAt() string {
	if o == nil || IsNil(o.AppliedAt) {
		var ret string
		return ret
	}
	return *o.AppliedAt
}

// GetAppliedAtOk returns a tuple with the AppliedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcInvite) GetAppliedAtOk() (*string, bool) {
	if o == nil || IsNil(o.AppliedAt) {
		return nil, false
	}
	return o.AppliedAt, true
}

// HasAppliedAt returns a boolean if a field has been set.
func (o *UserSvcInvite) HasAppliedAt() bool {
	if o != nil && !IsNil(o.AppliedAt) {
		return true
	}

	return false
}

// SetAppliedAt gets a reference to the given string and assigns it to the AppliedAt field.
func (o *UserSvcInvite) SetAppliedAt(v string) {
	o.AppliedAt = &v
}

// GetContactId returns the ContactId field value
func (o *UserSvcInvite) GetContactId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactId
}

// GetContactIdOk returns a tuple with the ContactId field value
// and a boolean to check if the value has been set.
func (o *UserSvcInvite) GetContactIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactId, true
}

// SetContactId sets field value
func (o *UserSvcInvite) SetContactId(v string) {
	o.ContactId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *UserSvcInvite) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserSvcInvite) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UserSvcInvite) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *UserSvcInvite) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcInvite) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *UserSvcInvite) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *UserSvcInvite) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *UserSvcInvite) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcInvite) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *UserSvcInvite) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *UserSvcInvite) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value
func (o *UserSvcInvite) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserSvcInvite) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserSvcInvite) SetId(v string) {
	o.Id = v
}

// GetOwnerIds returns the OwnerIds field value
func (o *UserSvcInvite) GetOwnerIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OwnerIds
}

// GetOwnerIdsOk returns a tuple with the OwnerIds field value
// and a boolean to check if the value has been set.
func (o *UserSvcInvite) GetOwnerIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerIds, true
}

// SetOwnerIds sets field value
func (o *UserSvcInvite) SetOwnerIds(v []string) {
	o.OwnerIds = v
}

// GetRoleId returns the RoleId field value
func (o *UserSvcInvite) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *UserSvcInvite) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *UserSvcInvite) SetRoleId(v string) {
	o.RoleId = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserSvcInvite) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcInvite) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserSvcInvite) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *UserSvcInvite) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o UserSvcInvite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcInvite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppliedAt) {
		toSerialize["appliedAt"] = o.AppliedAt
	}
	toSerialize["contactId"] = o.ContactId
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.DeletedAt) {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	toSerialize["id"] = o.Id
	toSerialize["ownerIds"] = o.OwnerIds
	toSerialize["roleId"] = o.RoleId
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *UserSvcInvite) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contactId",
		"createdAt",
		"id",
		"ownerIds",
		"roleId",
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

	varUserSvcInvite := _UserSvcInvite{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcInvite)

	if err != nil {
		return err
	}

	*o = UserSvcInvite(varUserSvcInvite)

	return err
}

type NullableUserSvcInvite struct {
	value *UserSvcInvite
	isSet bool
}

func (v NullableUserSvcInvite) Get() *UserSvcInvite {
	return v.value
}

func (v *NullableUserSvcInvite) Set(val *UserSvcInvite) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcInvite) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcInvite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcInvite(val *UserSvcInvite) *NullableUserSvcInvite {
	return &NullableUserSvcInvite{value: val, isSet: true}
}

func (v NullableUserSvcInvite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcInvite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


