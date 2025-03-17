/*
1Backend

A unified backend for your AI applications—microservices-based and built to scale.

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

// checks if the ContainerSvcRunContainerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcRunContainerRequest{}

// ContainerSvcRunContainerRequest struct for ContainerSvcRunContainerRequest
type ContainerSvcRunContainerRequest struct {
	// Assets maps environment variable names to file URLs. Example: {\"MODEL\": \"https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf\"} These files are downloaded by the File Svc and mounted in the container. The environment variable `MODEL` will point to the local file path in the container.
	Assets []ContainerSvcAsset `json:"assets,omitempty"`
	// Capabilities define additional runtime features, such as GPU support.
	Capabilities *ContainerSvcCapabilities `json:"capabilities,omitempty"`
	// Envs are environment variables set within the container.
	Envs []ContainerSvcEnvVar `json:"envs,omitempty"`
	// Hash is a unique identifier for the container
	Hash *string `json:"hash,omitempty"`
	// Image is the Docker image to use for the container
	Image string `json:"image"`
	// Keeps are paths that persist across container restarts. They function like mounts or volumes, but their external storage location is irrelevant.
	Keeps []ContainerSvcKeep `json:"keeps,omitempty"`
	// Labels are metadata tags assigned to the container.
	Labels []ContainerSvcLabel `json:"labels,omitempty"`
	// Names are the human-readable aliases assigned to the container.
	Names []string `json:"names,omitempty"`
	// Ports maps host ports (keys) to container ports (values).
	Ports []ContainerSvcPortMapping `json:"ports,omitempty"`
}

type _ContainerSvcRunContainerRequest ContainerSvcRunContainerRequest

// NewContainerSvcRunContainerRequest instantiates a new ContainerSvcRunContainerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcRunContainerRequest(image string) *ContainerSvcRunContainerRequest {
	this := ContainerSvcRunContainerRequest{}
	this.Image = image
	return &this
}

// NewContainerSvcRunContainerRequestWithDefaults instantiates a new ContainerSvcRunContainerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcRunContainerRequestWithDefaults() *ContainerSvcRunContainerRequest {
	this := ContainerSvcRunContainerRequest{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *ContainerSvcRunContainerRequest) GetAssets() []ContainerSvcAsset {
	if o == nil || IsNil(o.Assets) {
		var ret []ContainerSvcAsset
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcRunContainerRequest) GetAssetsOk() ([]ContainerSvcAsset, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *ContainerSvcRunContainerRequest) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []ContainerSvcAsset and assigns it to the Assets field.
func (o *ContainerSvcRunContainerRequest) SetAssets(v []ContainerSvcAsset) {
	o.Assets = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *ContainerSvcRunContainerRequest) GetCapabilities() ContainerSvcCapabilities {
	if o == nil || IsNil(o.Capabilities) {
		var ret ContainerSvcCapabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcRunContainerRequest) GetCapabilitiesOk() (*ContainerSvcCapabilities, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ContainerSvcRunContainerRequest) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given ContainerSvcCapabilities and assigns it to the Capabilities field.
func (o *ContainerSvcRunContainerRequest) SetCapabilities(v ContainerSvcCapabilities) {
	o.Capabilities = &v
}

// GetEnvs returns the Envs field value if set, zero value otherwise.
func (o *ContainerSvcRunContainerRequest) GetEnvs() []ContainerSvcEnvVar {
	if o == nil || IsNil(o.Envs) {
		var ret []ContainerSvcEnvVar
		return ret
	}
	return o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcRunContainerRequest) GetEnvsOk() ([]ContainerSvcEnvVar, bool) {
	if o == nil || IsNil(o.Envs) {
		return nil, false
	}
	return o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *ContainerSvcRunContainerRequest) HasEnvs() bool {
	if o != nil && !IsNil(o.Envs) {
		return true
	}

	return false
}

// SetEnvs gets a reference to the given []ContainerSvcEnvVar and assigns it to the Envs field.
func (o *ContainerSvcRunContainerRequest) SetEnvs(v []ContainerSvcEnvVar) {
	o.Envs = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ContainerSvcRunContainerRequest) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcRunContainerRequest) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ContainerSvcRunContainerRequest) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ContainerSvcRunContainerRequest) SetHash(v string) {
	o.Hash = &v
}

// GetImage returns the Image field value
func (o *ContainerSvcRunContainerRequest) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcRunContainerRequest) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ContainerSvcRunContainerRequest) SetImage(v string) {
	o.Image = v
}

// GetKeeps returns the Keeps field value if set, zero value otherwise.
func (o *ContainerSvcRunContainerRequest) GetKeeps() []ContainerSvcKeep {
	if o == nil || IsNil(o.Keeps) {
		var ret []ContainerSvcKeep
		return ret
	}
	return o.Keeps
}

// GetKeepsOk returns a tuple with the Keeps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcRunContainerRequest) GetKeepsOk() ([]ContainerSvcKeep, bool) {
	if o == nil || IsNil(o.Keeps) {
		return nil, false
	}
	return o.Keeps, true
}

// HasKeeps returns a boolean if a field has been set.
func (o *ContainerSvcRunContainerRequest) HasKeeps() bool {
	if o != nil && !IsNil(o.Keeps) {
		return true
	}

	return false
}

// SetKeeps gets a reference to the given []ContainerSvcKeep and assigns it to the Keeps field.
func (o *ContainerSvcRunContainerRequest) SetKeeps(v []ContainerSvcKeep) {
	o.Keeps = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ContainerSvcRunContainerRequest) GetLabels() []ContainerSvcLabel {
	if o == nil || IsNil(o.Labels) {
		var ret []ContainerSvcLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcRunContainerRequest) GetLabelsOk() ([]ContainerSvcLabel, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ContainerSvcRunContainerRequest) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []ContainerSvcLabel and assigns it to the Labels field.
func (o *ContainerSvcRunContainerRequest) SetLabels(v []ContainerSvcLabel) {
	o.Labels = v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *ContainerSvcRunContainerRequest) GetNames() []string {
	if o == nil || IsNil(o.Names) {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcRunContainerRequest) GetNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *ContainerSvcRunContainerRequest) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *ContainerSvcRunContainerRequest) SetNames(v []string) {
	o.Names = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ContainerSvcRunContainerRequest) GetPorts() []ContainerSvcPortMapping {
	if o == nil || IsNil(o.Ports) {
		var ret []ContainerSvcPortMapping
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcRunContainerRequest) GetPortsOk() ([]ContainerSvcPortMapping, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ContainerSvcRunContainerRequest) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ContainerSvcPortMapping and assigns it to the Ports field.
func (o *ContainerSvcRunContainerRequest) SetPorts(v []ContainerSvcPortMapping) {
	o.Ports = v
}

func (o ContainerSvcRunContainerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcRunContainerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.Envs) {
		toSerialize["envs"] = o.Envs
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	toSerialize["image"] = o.Image
	if !IsNil(o.Keeps) {
		toSerialize["keeps"] = o.Keeps
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return toSerialize, nil
}

func (o *ContainerSvcRunContainerRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image",
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

	varContainerSvcRunContainerRequest := _ContainerSvcRunContainerRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerSvcRunContainerRequest)

	if err != nil {
		return err
	}

	*o = ContainerSvcRunContainerRequest(varContainerSvcRunContainerRequest)

	return err
}

type NullableContainerSvcRunContainerRequest struct {
	value *ContainerSvcRunContainerRequest
	isSet bool
}

func (v NullableContainerSvcRunContainerRequest) Get() *ContainerSvcRunContainerRequest {
	return v.value
}

func (v *NullableContainerSvcRunContainerRequest) Set(val *ContainerSvcRunContainerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcRunContainerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcRunContainerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcRunContainerRequest(val *ContainerSvcRunContainerRequest) *NullableContainerSvcRunContainerRequest {
	return &NullableContainerSvcRunContainerRequest{value: val, isSet: true}
}

func (v NullableContainerSvcRunContainerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcRunContainerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


