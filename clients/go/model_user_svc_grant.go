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

// checks if the UserSvcGrant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSvcGrant{}

// UserSvcGrant struct for UserSvcGrant
type UserSvcGrant struct {
	Id *string `json:"id,omitempty"`
	Permission string `json:"permission"`
	// Role IDs that have been granted the specified permission.  Originally, grants were designed for slugs to facilitate service-to-service calls. Due to their convenience—especially with CLI and infrastructure-as-code support—they were later extended to roles.  Alternatively, permissions can be assigned to roles using UserSvc.AssignPermissions. Grants currently offer a more streamlined approach, though this may evolve over time.
	Roles []string `json:"roles,omitempty"`
	// Slugs that have been granted the specified permission.
	Slugs []string `json:"slugs,omitempty"`
}

type _UserSvcGrant UserSvcGrant

// NewUserSvcGrant instantiates a new UserSvcGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSvcGrant(permission string) *UserSvcGrant {
	this := UserSvcGrant{}
	this.Permission = permission
	return &this
}

// NewUserSvcGrantWithDefaults instantiates a new UserSvcGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSvcGrantWithDefaults() *UserSvcGrant {
	this := UserSvcGrant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserSvcGrant) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcGrant) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserSvcGrant) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserSvcGrant) SetId(v string) {
	o.Id = &v
}

// GetPermission returns the Permission field value
func (o *UserSvcGrant) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *UserSvcGrant) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *UserSvcGrant) SetPermission(v string) {
	o.Permission = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserSvcGrant) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcGrant) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserSvcGrant) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UserSvcGrant) SetRoles(v []string) {
	o.Roles = v
}

// GetSlugs returns the Slugs field value if set, zero value otherwise.
func (o *UserSvcGrant) GetSlugs() []string {
	if o == nil || IsNil(o.Slugs) {
		var ret []string
		return ret
	}
	return o.Slugs
}

// GetSlugsOk returns a tuple with the Slugs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSvcGrant) GetSlugsOk() ([]string, bool) {
	if o == nil || IsNil(o.Slugs) {
		return nil, false
	}
	return o.Slugs, true
}

// HasSlugs returns a boolean if a field has been set.
func (o *UserSvcGrant) HasSlugs() bool {
	if o != nil && !IsNil(o.Slugs) {
		return true
	}

	return false
}

// SetSlugs gets a reference to the given []string and assigns it to the Slugs field.
func (o *UserSvcGrant) SetSlugs(v []string) {
	o.Slugs = v
}

func (o UserSvcGrant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSvcGrant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["permission"] = o.Permission
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Slugs) {
		toSerialize["slugs"] = o.Slugs
	}
	return toSerialize, nil
}

func (o *UserSvcGrant) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"permission",
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

	varUserSvcGrant := _UserSvcGrant{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSvcGrant)

	if err != nil {
		return err
	}

	*o = UserSvcGrant(varUserSvcGrant)

	return err
}

type NullableUserSvcGrant struct {
	value *UserSvcGrant
	isSet bool
}

func (v NullableUserSvcGrant) Get() *UserSvcGrant {
	return v.value
}

func (v *NullableUserSvcGrant) Set(val *UserSvcGrant) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSvcGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSvcGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSvcGrant(val *UserSvcGrant) *NullableUserSvcGrant {
	return &NullableUserSvcGrant{value: val, isSet: true}
}

func (v NullableUserSvcGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSvcGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


