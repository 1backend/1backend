/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.38
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ContainerSvcContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcContainer{}

// ContainerSvcContainer struct for ContainerSvcContainer
type ContainerSvcContainer struct {
	// Assets maps environment variable names to file URLs. Example: {\"MODEL\": \"https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf\"} These files are downloaded by the File Svc and mounted in the container. The environment variable `MODEL` will point to the local file path in the container.
	Assets []ContainerSvcAsset `json:"assets,omitempty"`
	// Capabilities define additional runtime features, such as GPU support.
	Capabilities *ContainerSvcCapabilities `json:"capabilities,omitempty"`
	// Envs are environment variables set within the container.
	Envs []ContainerSvcEnvVar `json:"envs,omitempty"`
	// Hash is a unique identifier associated with the container.
	Hash *string `json:"hash,omitempty"`
	// Id is the unique identifier for the container instance.
	Id *string `json:"id,omitempty"`
	// Image is the Docker image used to create the container.
	Image *string `json:"image,omitempty"`
	// Keeps are paths that persist across container restarts. They function like mounts or volumes, but their external storage location is irrelevant.
	Keeps []ContainerSvcKeep `json:"keeps,omitempty"`
	// Labels are metadata tags assigned to the container.
	Labels []ContainerSvcLabel `json:"labels,omitempty"`
	// Names are the human-readable aliases assigned to the container.
	Names []string `json:"names,omitempty"`
	// Network contains networking-related information for the container.
	Network *ContainerSvcNetwork `json:"network,omitempty"`
	// Node Id Please see the documentation for the envar OB_NODE_ID
	NodeId *string `json:"nodeId,omitempty"`
	// Ports maps host ports (keys) to container ports (values).
	Ports []ContainerSvcPortMapping `json:"ports,omitempty"`
	// Resources defines CPU, memory, and disk constraints for the container.
	Resources *ContainerSvcResources `json:"resources,omitempty"`
	// Runtime specifies the container runtime (e.g., Docker, containerd, etc.).
	Runtime *string `json:"runtime,omitempty"`
	// Status indicates the current state of the container (e.g., running, stopped).
	Status *string `json:"status,omitempty"`
	// Volumes mounted by the container.
	Volumes []ContainerSvcVolume `json:"volumes,omitempty"`
}

// NewContainerSvcContainer instantiates a new ContainerSvcContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcContainer() *ContainerSvcContainer {
	this := ContainerSvcContainer{}
	return &this
}

// NewContainerSvcContainerWithDefaults instantiates a new ContainerSvcContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcContainerWithDefaults() *ContainerSvcContainer {
	this := ContainerSvcContainer{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetAssets() []ContainerSvcAsset {
	if o == nil || IsNil(o.Assets) {
		var ret []ContainerSvcAsset
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetAssetsOk() ([]ContainerSvcAsset, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []ContainerSvcAsset and assigns it to the Assets field.
func (o *ContainerSvcContainer) SetAssets(v []ContainerSvcAsset) {
	o.Assets = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetCapabilities() ContainerSvcCapabilities {
	if o == nil || IsNil(o.Capabilities) {
		var ret ContainerSvcCapabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetCapabilitiesOk() (*ContainerSvcCapabilities, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given ContainerSvcCapabilities and assigns it to the Capabilities field.
func (o *ContainerSvcContainer) SetCapabilities(v ContainerSvcCapabilities) {
	o.Capabilities = &v
}

// GetEnvs returns the Envs field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetEnvs() []ContainerSvcEnvVar {
	if o == nil || IsNil(o.Envs) {
		var ret []ContainerSvcEnvVar
		return ret
	}
	return o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetEnvsOk() ([]ContainerSvcEnvVar, bool) {
	if o == nil || IsNil(o.Envs) {
		return nil, false
	}
	return o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasEnvs() bool {
	if o != nil && !IsNil(o.Envs) {
		return true
	}

	return false
}

// SetEnvs gets a reference to the given []ContainerSvcEnvVar and assigns it to the Envs field.
func (o *ContainerSvcContainer) SetEnvs(v []ContainerSvcEnvVar) {
	o.Envs = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ContainerSvcContainer) SetHash(v string) {
	o.Hash = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContainerSvcContainer) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ContainerSvcContainer) SetImage(v string) {
	o.Image = &v
}

// GetKeeps returns the Keeps field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetKeeps() []ContainerSvcKeep {
	if o == nil || IsNil(o.Keeps) {
		var ret []ContainerSvcKeep
		return ret
	}
	return o.Keeps
}

// GetKeepsOk returns a tuple with the Keeps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetKeepsOk() ([]ContainerSvcKeep, bool) {
	if o == nil || IsNil(o.Keeps) {
		return nil, false
	}
	return o.Keeps, true
}

// HasKeeps returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasKeeps() bool {
	if o != nil && !IsNil(o.Keeps) {
		return true
	}

	return false
}

// SetKeeps gets a reference to the given []ContainerSvcKeep and assigns it to the Keeps field.
func (o *ContainerSvcContainer) SetKeeps(v []ContainerSvcKeep) {
	o.Keeps = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetLabels() []ContainerSvcLabel {
	if o == nil || IsNil(o.Labels) {
		var ret []ContainerSvcLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetLabelsOk() ([]ContainerSvcLabel, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []ContainerSvcLabel and assigns it to the Labels field.
func (o *ContainerSvcContainer) SetLabels(v []ContainerSvcLabel) {
	o.Labels = v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetNames() []string {
	if o == nil || IsNil(o.Names) {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *ContainerSvcContainer) SetNames(v []string) {
	o.Names = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetNetwork() ContainerSvcNetwork {
	if o == nil || IsNil(o.Network) {
		var ret ContainerSvcNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetNetworkOk() (*ContainerSvcNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ContainerSvcNetwork and assigns it to the Network field.
func (o *ContainerSvcContainer) SetNetwork(v ContainerSvcNetwork) {
	o.Network = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *ContainerSvcContainer) SetNodeId(v string) {
	o.NodeId = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetPorts() []ContainerSvcPortMapping {
	if o == nil || IsNil(o.Ports) {
		var ret []ContainerSvcPortMapping
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetPortsOk() ([]ContainerSvcPortMapping, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ContainerSvcPortMapping and assigns it to the Ports field.
func (o *ContainerSvcContainer) SetPorts(v []ContainerSvcPortMapping) {
	o.Ports = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetResources() ContainerSvcResources {
	if o == nil || IsNil(o.Resources) {
		var ret ContainerSvcResources
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetResourcesOk() (*ContainerSvcResources, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given ContainerSvcResources and assigns it to the Resources field.
func (o *ContainerSvcContainer) SetResources(v ContainerSvcResources) {
	o.Resources = &v
}

// GetRuntime returns the Runtime field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetRuntime() string {
	if o == nil || IsNil(o.Runtime) {
		var ret string
		return ret
	}
	return *o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetRuntimeOk() (*string, bool) {
	if o == nil || IsNil(o.Runtime) {
		return nil, false
	}
	return o.Runtime, true
}

// HasRuntime returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasRuntime() bool {
	if o != nil && !IsNil(o.Runtime) {
		return true
	}

	return false
}

// SetRuntime gets a reference to the given string and assigns it to the Runtime field.
func (o *ContainerSvcContainer) SetRuntime(v string) {
	o.Runtime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ContainerSvcContainer) SetStatus(v string) {
	o.Status = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *ContainerSvcContainer) GetVolumes() []ContainerSvcVolume {
	if o == nil || IsNil(o.Volumes) {
		var ret []ContainerSvcVolume
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcContainer) GetVolumesOk() ([]ContainerSvcVolume, bool) {
	if o == nil || IsNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *ContainerSvcContainer) HasVolumes() bool {
	if o != nil && !IsNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []ContainerSvcVolume and assigns it to the Volumes field.
func (o *ContainerSvcContainer) SetVolumes(v []ContainerSvcVolume) {
	o.Volumes = v
}

func (o ContainerSvcContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcContainer) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Keeps) {
		toSerialize["keeps"] = o.Keeps
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.NodeId) {
		toSerialize["nodeId"] = o.NodeId
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Runtime) {
		toSerialize["runtime"] = o.Runtime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	return toSerialize, nil
}

type NullableContainerSvcContainer struct {
	value *ContainerSvcContainer
	isSet bool
}

func (v NullableContainerSvcContainer) Get() *ContainerSvcContainer {
	return v.value
}

func (v *NullableContainerSvcContainer) Set(val *ContainerSvcContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcContainer(val *ContainerSvcContainer) *NullableContainerSvcContainer {
	return &NullableContainerSvcContainer{value: val, isSet: true}
}

func (v NullableContainerSvcContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


