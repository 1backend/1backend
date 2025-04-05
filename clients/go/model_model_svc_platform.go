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
)

// checks if the ModelSvcPlatform type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSvcPlatform{}

// ModelSvcPlatform struct for ModelSvcPlatform
type ModelSvcPlatform struct {
	Architectures *ModelSvcArchitectures `json:"architectures,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// List of prompt types that the AI engine supports.
	Types []PromptSvcPromptType `json:"types,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewModelSvcPlatform instantiates a new ModelSvcPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSvcPlatform() *ModelSvcPlatform {
	this := ModelSvcPlatform{}
	return &this
}

// NewModelSvcPlatformWithDefaults instantiates a new ModelSvcPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSvcPlatformWithDefaults() *ModelSvcPlatform {
	this := ModelSvcPlatform{}
	return &this
}

// GetArchitectures returns the Architectures field value if set, zero value otherwise.
func (o *ModelSvcPlatform) GetArchitectures() ModelSvcArchitectures {
	if o == nil || IsNil(o.Architectures) {
		var ret ModelSvcArchitectures
		return ret
	}
	return *o.Architectures
}

// GetArchitecturesOk returns a tuple with the Architectures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcPlatform) GetArchitecturesOk() (*ModelSvcArchitectures, bool) {
	if o == nil || IsNil(o.Architectures) {
		return nil, false
	}
	return o.Architectures, true
}

// HasArchitectures returns a boolean if a field has been set.
func (o *ModelSvcPlatform) HasArchitectures() bool {
	if o != nil && !IsNil(o.Architectures) {
		return true
	}

	return false
}

// SetArchitectures gets a reference to the given ModelSvcArchitectures and assigns it to the Architectures field.
func (o *ModelSvcPlatform) SetArchitectures(v ModelSvcArchitectures) {
	o.Architectures = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelSvcPlatform) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcPlatform) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelSvcPlatform) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelSvcPlatform) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelSvcPlatform) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcPlatform) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelSvcPlatform) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelSvcPlatform) SetName(v string) {
	o.Name = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *ModelSvcPlatform) GetTypes() []PromptSvcPromptType {
	if o == nil || IsNil(o.Types) {
		var ret []PromptSvcPromptType
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcPlatform) GetTypesOk() ([]PromptSvcPromptType, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *ModelSvcPlatform) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []PromptSvcPromptType and assigns it to the Types field.
func (o *ModelSvcPlatform) SetTypes(v []PromptSvcPromptType) {
	o.Types = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ModelSvcPlatform) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSvcPlatform) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ModelSvcPlatform) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ModelSvcPlatform) SetVersion(v int32) {
	o.Version = &v
}

func (o ModelSvcPlatform) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSvcPlatform) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Architectures) {
		toSerialize["architectures"] = o.Architectures
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableModelSvcPlatform struct {
	value *ModelSvcPlatform
	isSet bool
}

func (v NullableModelSvcPlatform) Get() *ModelSvcPlatform {
	return v.value
}

func (v *NullableModelSvcPlatform) Set(val *ModelSvcPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSvcPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSvcPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSvcPlatform(val *ModelSvcPlatform) *NullableModelSvcPlatform {
	return &NullableModelSvcPlatform{value: val, isSet: true}
}

func (v NullableModelSvcPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSvcPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


