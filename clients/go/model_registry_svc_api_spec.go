/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.35
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RegistrySvcAPISpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcAPISpec{}

// RegistrySvcAPISpec struct for RegistrySvcAPISpec
type RegistrySvcAPISpec struct {
	// Additional metadata about the API (e.g., author, license, etc.)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Protocol type (e.g., OpenAPI, Swagger, etc.)
	ProtocolType *string `json:"protocolType,omitempty"`
	// URL to the OpenAPI file or other API definition
	Url *string `json:"url,omitempty"`
	// Version of the API specification (e.g., 3.0.0)
	Version *string `json:"version,omitempty"`
}

// NewRegistrySvcAPISpec instantiates a new RegistrySvcAPISpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcAPISpec() *RegistrySvcAPISpec {
	this := RegistrySvcAPISpec{}
	return &this
}

// NewRegistrySvcAPISpecWithDefaults instantiates a new RegistrySvcAPISpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcAPISpecWithDefaults() *RegistrySvcAPISpec {
	this := RegistrySvcAPISpec{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RegistrySvcAPISpec) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcAPISpec) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RegistrySvcAPISpec) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *RegistrySvcAPISpec) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetProtocolType returns the ProtocolType field value if set, zero value otherwise.
func (o *RegistrySvcAPISpec) GetProtocolType() string {
	if o == nil || IsNil(o.ProtocolType) {
		var ret string
		return ret
	}
	return *o.ProtocolType
}

// GetProtocolTypeOk returns a tuple with the ProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcAPISpec) GetProtocolTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolType) {
		return nil, false
	}
	return o.ProtocolType, true
}

// HasProtocolType returns a boolean if a field has been set.
func (o *RegistrySvcAPISpec) HasProtocolType() bool {
	if o != nil && !IsNil(o.ProtocolType) {
		return true
	}

	return false
}

// SetProtocolType gets a reference to the given string and assigns it to the ProtocolType field.
func (o *RegistrySvcAPISpec) SetProtocolType(v string) {
	o.ProtocolType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RegistrySvcAPISpec) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcAPISpec) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RegistrySvcAPISpec) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RegistrySvcAPISpec) SetUrl(v string) {
	o.Url = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RegistrySvcAPISpec) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcAPISpec) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RegistrySvcAPISpec) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RegistrySvcAPISpec) SetVersion(v string) {
	o.Version = &v
}

func (o RegistrySvcAPISpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcAPISpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.ProtocolType) {
		toSerialize["protocolType"] = o.ProtocolType
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableRegistrySvcAPISpec struct {
	value *RegistrySvcAPISpec
	isSet bool
}

func (v NullableRegistrySvcAPISpec) Get() *RegistrySvcAPISpec {
	return v.value
}

func (v *NullableRegistrySvcAPISpec) Set(val *RegistrySvcAPISpec) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcAPISpec) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcAPISpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcAPISpec(val *RegistrySvcAPISpec) *NullableRegistrySvcAPISpec {
	return &NullableRegistrySvcAPISpec{value: val, isSet: true}
}

func (v NullableRegistrySvcAPISpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcAPISpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


