/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.30
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RegistrySvcInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcInstance{}

// RegistrySvcInstance struct for RegistrySvcInstance
type RegistrySvcInstance struct {
	// The ID of the deployment that this instance is an instance of. Only instances deployed by 1Backend have a DeploymentId. Services can be deployed through other means (Docker Compose, K8s, anything), in that case they self-register and will not have a DeploymentId.
	DeploymentId *string `json:"deploymentId,omitempty"`
	// Details
	Details *string `json:"details,omitempty"`
	// Host of the instance address. Required if URL is not provided
	Host *string `json:"host,omitempty"`
	// Required: ID of the instance
	Id string `json:"id"`
	// IP of the instance address. Optional: to register by IP instead of host
	Ip *string `json:"ip,omitempty"`
	// Last time the instance gave a sign of life
	LastHeartbeat *string `json:"lastHeartbeat,omitempty"`
	// NodeURL is the URL of the 1Backend server the instance is running on. To have a NodeURL the instance must either: - Be deployed by 1Backend - Declare the 1Backend server URL when registering its instance
	NodeUrl *string `json:"nodeUrl,omitempty"`
	// Path of the instance address. Optional (e.g., \"/api\")
	Path *string `json:"path,omitempty"`
	// Port of the instance address. Required if URL is not provided
	Port *int32 `json:"port,omitempty"`
	// Scheme of the instance address. Required if URL is not provided.
	Scheme *string `json:"scheme,omitempty"`
	// Slug of the account that owns this instance Services that want to be proxied by their slug are advised to self register their instance at startup. Keep in mind, instances might be deployed by 1Backend yet they still won't be 1Backend services and they won't have slugs. Think NGINX, MySQL, etc.
	Slug *string `json:"slug,omitempty"`
	// Status
	Status RegistrySvcInstanceStatus `json:"status"`
	// Tags are used to filter instances
	Tags []string `json:"tags,omitempty"`
	// Full address URL of the instance.
	Url string `json:"url"`
}

type _RegistrySvcInstance RegistrySvcInstance

// NewRegistrySvcInstance instantiates a new RegistrySvcInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcInstance(id string, status RegistrySvcInstanceStatus, url string) *RegistrySvcInstance {
	this := RegistrySvcInstance{}
	this.Id = id
	this.Status = status
	this.Url = url
	return &this
}

// NewRegistrySvcInstanceWithDefaults instantiates a new RegistrySvcInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcInstanceWithDefaults() *RegistrySvcInstance {
	this := RegistrySvcInstance{}
	return &this
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *RegistrySvcInstance) GetDeploymentId() string {
	if o == nil || IsNil(o.DeploymentId) {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetDeploymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentId) {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *RegistrySvcInstance) HasDeploymentId() bool {
	if o != nil && !IsNil(o.DeploymentId) {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *RegistrySvcInstance) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RegistrySvcInstance) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RegistrySvcInstance) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *RegistrySvcInstance) SetDetails(v string) {
	o.Details = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *RegistrySvcInstance) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *RegistrySvcInstance) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *RegistrySvcInstance) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value
func (o *RegistrySvcInstance) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RegistrySvcInstance) SetId(v string) {
	o.Id = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *RegistrySvcInstance) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *RegistrySvcInstance) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *RegistrySvcInstance) SetIp(v string) {
	o.Ip = &v
}

// GetLastHeartbeat returns the LastHeartbeat field value if set, zero value otherwise.
func (o *RegistrySvcInstance) GetLastHeartbeat() string {
	if o == nil || IsNil(o.LastHeartbeat) {
		var ret string
		return ret
	}
	return *o.LastHeartbeat
}

// GetLastHeartbeatOk returns a tuple with the LastHeartbeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetLastHeartbeatOk() (*string, bool) {
	if o == nil || IsNil(o.LastHeartbeat) {
		return nil, false
	}
	return o.LastHeartbeat, true
}

// HasLastHeartbeat returns a boolean if a field has been set.
func (o *RegistrySvcInstance) HasLastHeartbeat() bool {
	if o != nil && !IsNil(o.LastHeartbeat) {
		return true
	}

	return false
}

// SetLastHeartbeat gets a reference to the given string and assigns it to the LastHeartbeat field.
func (o *RegistrySvcInstance) SetLastHeartbeat(v string) {
	o.LastHeartbeat = &v
}

// GetNodeUrl returns the NodeUrl field value if set, zero value otherwise.
func (o *RegistrySvcInstance) GetNodeUrl() string {
	if o == nil || IsNil(o.NodeUrl) {
		var ret string
		return ret
	}
	return *o.NodeUrl
}

// GetNodeUrlOk returns a tuple with the NodeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetNodeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.NodeUrl) {
		return nil, false
	}
	return o.NodeUrl, true
}

// HasNodeUrl returns a boolean if a field has been set.
func (o *RegistrySvcInstance) HasNodeUrl() bool {
	if o != nil && !IsNil(o.NodeUrl) {
		return true
	}

	return false
}

// SetNodeUrl gets a reference to the given string and assigns it to the NodeUrl field.
func (o *RegistrySvcInstance) SetNodeUrl(v string) {
	o.NodeUrl = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *RegistrySvcInstance) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *RegistrySvcInstance) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *RegistrySvcInstance) SetPath(v string) {
	o.Path = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *RegistrySvcInstance) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *RegistrySvcInstance) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *RegistrySvcInstance) SetPort(v int32) {
	o.Port = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *RegistrySvcInstance) GetScheme() string {
	if o == nil || IsNil(o.Scheme) {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *RegistrySvcInstance) HasScheme() bool {
	if o != nil && !IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *RegistrySvcInstance) SetScheme(v string) {
	o.Scheme = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *RegistrySvcInstance) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *RegistrySvcInstance) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *RegistrySvcInstance) SetSlug(v string) {
	o.Slug = &v
}

// GetStatus returns the Status field value
func (o *RegistrySvcInstance) GetStatus() RegistrySvcInstanceStatus {
	if o == nil {
		var ret RegistrySvcInstanceStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetStatusOk() (*RegistrySvcInstanceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RegistrySvcInstance) SetStatus(v RegistrySvcInstanceStatus) {
	o.Status = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RegistrySvcInstance) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RegistrySvcInstance) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RegistrySvcInstance) SetTags(v []string) {
	o.Tags = v
}

// GetUrl returns the Url field value
func (o *RegistrySvcInstance) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcInstance) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RegistrySvcInstance) SetUrl(v string) {
	o.Url = v
}

func (o RegistrySvcInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeploymentId) {
		toSerialize["deploymentId"] = o.DeploymentId
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.LastHeartbeat) {
		toSerialize["lastHeartbeat"] = o.LastHeartbeat
	}
	if !IsNil(o.NodeUrl) {
		toSerialize["nodeUrl"] = o.NodeUrl
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *RegistrySvcInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
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

	varRegistrySvcInstance := _RegistrySvcInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrySvcInstance)

	if err != nil {
		return err
	}

	*o = RegistrySvcInstance(varRegistrySvcInstance)

	return err
}

type NullableRegistrySvcInstance struct {
	value *RegistrySvcInstance
	isSet bool
}

func (v NullableRegistrySvcInstance) Get() *RegistrySvcInstance {
	return v.value
}

func (v *NullableRegistrySvcInstance) Set(val *RegistrySvcInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcInstance(val *RegistrySvcInstance) *NullableRegistrySvcInstance {
	return &NullableRegistrySvcInstance{value: val, isSet: true}
}

func (v NullableRegistrySvcInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


