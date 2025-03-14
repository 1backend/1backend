/*
1Backend

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RegistrySvcRepositorySpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcRepositorySpec{}

// RegistrySvcRepositorySpec struct for RegistrySvcRepositorySpec
type RegistrySvcRepositorySpec struct {
	// Context is the path to the image build context
	BuildContext *string `json:"buildContext,omitempty"`
	// ContainerFile is the path to the file that contains the container build instructions Relative from the build context. By default, it is assumed to be a Dockerfile.
	ContainerFile *string `json:"containerFile,omitempty"`
	// Ports the container will listen on internally
	Ports []int32 `json:"ports"`
	// URL is the URL to the repository
	Url string `json:"url"`
	// Version of the code to use
	Version *string `json:"version,omitempty"`
}

type _RegistrySvcRepositorySpec RegistrySvcRepositorySpec

// NewRegistrySvcRepositorySpec instantiates a new RegistrySvcRepositorySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcRepositorySpec(ports []int32, url string) *RegistrySvcRepositorySpec {
	this := RegistrySvcRepositorySpec{}
	this.Ports = ports
	this.Url = url
	return &this
}

// NewRegistrySvcRepositorySpecWithDefaults instantiates a new RegistrySvcRepositorySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcRepositorySpecWithDefaults() *RegistrySvcRepositorySpec {
	this := RegistrySvcRepositorySpec{}
	return &this
}

// GetBuildContext returns the BuildContext field value if set, zero value otherwise.
func (o *RegistrySvcRepositorySpec) GetBuildContext() string {
	if o == nil || IsNil(o.BuildContext) {
		var ret string
		return ret
	}
	return *o.BuildContext
}

// GetBuildContextOk returns a tuple with the BuildContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcRepositorySpec) GetBuildContextOk() (*string, bool) {
	if o == nil || IsNil(o.BuildContext) {
		return nil, false
	}
	return o.BuildContext, true
}

// HasBuildContext returns a boolean if a field has been set.
func (o *RegistrySvcRepositorySpec) HasBuildContext() bool {
	if o != nil && !IsNil(o.BuildContext) {
		return true
	}

	return false
}

// SetBuildContext gets a reference to the given string and assigns it to the BuildContext field.
func (o *RegistrySvcRepositorySpec) SetBuildContext(v string) {
	o.BuildContext = &v
}

// GetContainerFile returns the ContainerFile field value if set, zero value otherwise.
func (o *RegistrySvcRepositorySpec) GetContainerFile() string {
	if o == nil || IsNil(o.ContainerFile) {
		var ret string
		return ret
	}
	return *o.ContainerFile
}

// GetContainerFileOk returns a tuple with the ContainerFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcRepositorySpec) GetContainerFileOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerFile) {
		return nil, false
	}
	return o.ContainerFile, true
}

// HasContainerFile returns a boolean if a field has been set.
func (o *RegistrySvcRepositorySpec) HasContainerFile() bool {
	if o != nil && !IsNil(o.ContainerFile) {
		return true
	}

	return false
}

// SetContainerFile gets a reference to the given string and assigns it to the ContainerFile field.
func (o *RegistrySvcRepositorySpec) SetContainerFile(v string) {
	o.ContainerFile = &v
}

// GetPorts returns the Ports field value
func (o *RegistrySvcRepositorySpec) GetPorts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcRepositorySpec) GetPortsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *RegistrySvcRepositorySpec) SetPorts(v []int32) {
	o.Ports = v
}

// GetUrl returns the Url field value
func (o *RegistrySvcRepositorySpec) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcRepositorySpec) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RegistrySvcRepositorySpec) SetUrl(v string) {
	o.Url = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RegistrySvcRepositorySpec) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcRepositorySpec) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RegistrySvcRepositorySpec) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RegistrySvcRepositorySpec) SetVersion(v string) {
	o.Version = &v
}

func (o RegistrySvcRepositorySpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcRepositorySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildContext) {
		toSerialize["buildContext"] = o.BuildContext
	}
	if !IsNil(o.ContainerFile) {
		toSerialize["containerFile"] = o.ContainerFile
	}
	toSerialize["ports"] = o.Ports
	toSerialize["url"] = o.Url
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

func (o *RegistrySvcRepositorySpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ports",
		"url",
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

	varRegistrySvcRepositorySpec := _RegistrySvcRepositorySpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrySvcRepositorySpec)

	if err != nil {
		return err
	}

	*o = RegistrySvcRepositorySpec(varRegistrySvcRepositorySpec)

	return err
}

type NullableRegistrySvcRepositorySpec struct {
	value *RegistrySvcRepositorySpec
	isSet bool
}

func (v NullableRegistrySvcRepositorySpec) Get() *RegistrySvcRepositorySpec {
	return v.value
}

func (v *NullableRegistrySvcRepositorySpec) Set(val *RegistrySvcRepositorySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcRepositorySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcRepositorySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcRepositorySpec(val *RegistrySvcRepositorySpec) *NullableRegistrySvcRepositorySpec {
	return &NullableRegistrySvcRepositorySpec{value: val, isSet: true}
}

func (v NullableRegistrySvcRepositorySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcRepositorySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


