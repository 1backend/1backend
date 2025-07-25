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

// checks if the ContainerSvcBuildImageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcBuildImageRequest{}

// ContainerSvcBuildImageRequest struct for ContainerSvcBuildImageRequest
type ContainerSvcBuildImageRequest struct {
	// ContextPath is the local path to the build context
	ContextPath string `json:"contextPath"`
	// DockerfilePath is the local path to the Dockerfile
	DockerfilePath *string `json:"dockerfilePath,omitempty"`
	// Name is the name of the image to build
	Name string `json:"name"`
}

type _ContainerSvcBuildImageRequest ContainerSvcBuildImageRequest

// NewContainerSvcBuildImageRequest instantiates a new ContainerSvcBuildImageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcBuildImageRequest(contextPath string, name string) *ContainerSvcBuildImageRequest {
	this := ContainerSvcBuildImageRequest{}
	this.ContextPath = contextPath
	this.Name = name
	return &this
}

// NewContainerSvcBuildImageRequestWithDefaults instantiates a new ContainerSvcBuildImageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcBuildImageRequestWithDefaults() *ContainerSvcBuildImageRequest {
	this := ContainerSvcBuildImageRequest{}
	return &this
}

// GetContextPath returns the ContextPath field value
func (o *ContainerSvcBuildImageRequest) GetContextPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContextPath
}

// GetContextPathOk returns a tuple with the ContextPath field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcBuildImageRequest) GetContextPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContextPath, true
}

// SetContextPath sets field value
func (o *ContainerSvcBuildImageRequest) SetContextPath(v string) {
	o.ContextPath = v
}

// GetDockerfilePath returns the DockerfilePath field value if set, zero value otherwise.
func (o *ContainerSvcBuildImageRequest) GetDockerfilePath() string {
	if o == nil || IsNil(o.DockerfilePath) {
		var ret string
		return ret
	}
	return *o.DockerfilePath
}

// GetDockerfilePathOk returns a tuple with the DockerfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcBuildImageRequest) GetDockerfilePathOk() (*string, bool) {
	if o == nil || IsNil(o.DockerfilePath) {
		return nil, false
	}
	return o.DockerfilePath, true
}

// HasDockerfilePath returns a boolean if a field has been set.
func (o *ContainerSvcBuildImageRequest) HasDockerfilePath() bool {
	if o != nil && !IsNil(o.DockerfilePath) {
		return true
	}

	return false
}

// SetDockerfilePath gets a reference to the given string and assigns it to the DockerfilePath field.
func (o *ContainerSvcBuildImageRequest) SetDockerfilePath(v string) {
	o.DockerfilePath = &v
}

// GetName returns the Name field value
func (o *ContainerSvcBuildImageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcBuildImageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerSvcBuildImageRequest) SetName(v string) {
	o.Name = v
}

func (o ContainerSvcBuildImageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcBuildImageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contextPath"] = o.ContextPath
	if !IsNil(o.DockerfilePath) {
		toSerialize["dockerfilePath"] = o.DockerfilePath
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ContainerSvcBuildImageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contextPath",
		"name",
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

	varContainerSvcBuildImageRequest := _ContainerSvcBuildImageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerSvcBuildImageRequest)

	if err != nil {
		return err
	}

	*o = ContainerSvcBuildImageRequest(varContainerSvcBuildImageRequest)

	return err
}

type NullableContainerSvcBuildImageRequest struct {
	value *ContainerSvcBuildImageRequest
	isSet bool
}

func (v NullableContainerSvcBuildImageRequest) Get() *ContainerSvcBuildImageRequest {
	return v.value
}

func (v *NullableContainerSvcBuildImageRequest) Set(val *ContainerSvcBuildImageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcBuildImageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcBuildImageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcBuildImageRequest(val *ContainerSvcBuildImageRequest) *NullableContainerSvcBuildImageRequest {
	return &NullableContainerSvcBuildImageRequest{value: val, isSet: true}
}

func (v NullableContainerSvcBuildImageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcBuildImageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


