/*
OpenOrch

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

// checks if the RegistrySvcDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcDefinition{}

// RegistrySvcDefinition struct for RegistrySvcDefinition
type RegistrySvcDefinition struct {
	// API Specs such as OpenAPI definitions etc.
	ApiSpecs []RegistrySvcAPISpec `json:"apiSpecs,omitempty"`
	// Programming language clients such as on npm or GitHub.
	Clients []RegistrySvcClient `json:"clients,omitempty"`
	// Envars is a map of Renvironment variables that a deployment (see Deploy Svc Deployment) of this definition will REQUIRE to run. E.g., {\"DB_URL\": \"mysql://user:password@host:port/db\"} These will be injected into the service instances (see Registry Svc Instance) at runtime. The value of a key here is the default value. The actual value can be overridden at deployment time.
	Envars []RegistrySvcEnvVar `json:"envars,omitempty"`
	Id string `json:"id"`
	// Container specifications for Docker, K8s, etc. Use this to deploy already built images.
	Image *RegistrySvcImageSpec `json:"image,omitempty"`
	// Ports have host ports and internal ports currently but they really only should have internal ports as host ports should be assigned by the system. Host ports might go away in the future.
	Ports []RegistrySvcPortMapping `json:"ports,omitempty"`
	// Repository based definitions is an alternative to Image definitions. Instead of deploying an already built and pushed image, a source code repository url can be provided. The container will be built from the source.
	Repository *RegistrySvcRepositorySpec `json:"repository,omitempty"`
}

type _RegistrySvcDefinition RegistrySvcDefinition

// NewRegistrySvcDefinition instantiates a new RegistrySvcDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcDefinition(id string) *RegistrySvcDefinition {
	this := RegistrySvcDefinition{}
	this.Id = id
	return &this
}

// NewRegistrySvcDefinitionWithDefaults instantiates a new RegistrySvcDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcDefinitionWithDefaults() *RegistrySvcDefinition {
	this := RegistrySvcDefinition{}
	return &this
}

// GetApiSpecs returns the ApiSpecs field value if set, zero value otherwise.
func (o *RegistrySvcDefinition) GetApiSpecs() []RegistrySvcAPISpec {
	if o == nil || IsNil(o.ApiSpecs) {
		var ret []RegistrySvcAPISpec
		return ret
	}
	return o.ApiSpecs
}

// GetApiSpecsOk returns a tuple with the ApiSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcDefinition) GetApiSpecsOk() ([]RegistrySvcAPISpec, bool) {
	if o == nil || IsNil(o.ApiSpecs) {
		return nil, false
	}
	return o.ApiSpecs, true
}

// HasApiSpecs returns a boolean if a field has been set.
func (o *RegistrySvcDefinition) HasApiSpecs() bool {
	if o != nil && !IsNil(o.ApiSpecs) {
		return true
	}

	return false
}

// SetApiSpecs gets a reference to the given []RegistrySvcAPISpec and assigns it to the ApiSpecs field.
func (o *RegistrySvcDefinition) SetApiSpecs(v []RegistrySvcAPISpec) {
	o.ApiSpecs = v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *RegistrySvcDefinition) GetClients() []RegistrySvcClient {
	if o == nil || IsNil(o.Clients) {
		var ret []RegistrySvcClient
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcDefinition) GetClientsOk() ([]RegistrySvcClient, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *RegistrySvcDefinition) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []RegistrySvcClient and assigns it to the Clients field.
func (o *RegistrySvcDefinition) SetClients(v []RegistrySvcClient) {
	o.Clients = v
}

// GetEnvars returns the Envars field value if set, zero value otherwise.
func (o *RegistrySvcDefinition) GetEnvars() []RegistrySvcEnvVar {
	if o == nil || IsNil(o.Envars) {
		var ret []RegistrySvcEnvVar
		return ret
	}
	return o.Envars
}

// GetEnvarsOk returns a tuple with the Envars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcDefinition) GetEnvarsOk() ([]RegistrySvcEnvVar, bool) {
	if o == nil || IsNil(o.Envars) {
		return nil, false
	}
	return o.Envars, true
}

// HasEnvars returns a boolean if a field has been set.
func (o *RegistrySvcDefinition) HasEnvars() bool {
	if o != nil && !IsNil(o.Envars) {
		return true
	}

	return false
}

// SetEnvars gets a reference to the given []RegistrySvcEnvVar and assigns it to the Envars field.
func (o *RegistrySvcDefinition) SetEnvars(v []RegistrySvcEnvVar) {
	o.Envars = v
}

// GetId returns the Id field value
func (o *RegistrySvcDefinition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcDefinition) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RegistrySvcDefinition) SetId(v string) {
	o.Id = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *RegistrySvcDefinition) GetImage() RegistrySvcImageSpec {
	if o == nil || IsNil(o.Image) {
		var ret RegistrySvcImageSpec
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcDefinition) GetImageOk() (*RegistrySvcImageSpec, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *RegistrySvcDefinition) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given RegistrySvcImageSpec and assigns it to the Image field.
func (o *RegistrySvcDefinition) SetImage(v RegistrySvcImageSpec) {
	o.Image = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *RegistrySvcDefinition) GetPorts() []RegistrySvcPortMapping {
	if o == nil || IsNil(o.Ports) {
		var ret []RegistrySvcPortMapping
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcDefinition) GetPortsOk() ([]RegistrySvcPortMapping, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *RegistrySvcDefinition) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []RegistrySvcPortMapping and assigns it to the Ports field.
func (o *RegistrySvcDefinition) SetPorts(v []RegistrySvcPortMapping) {
	o.Ports = v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *RegistrySvcDefinition) GetRepository() RegistrySvcRepositorySpec {
	if o == nil || IsNil(o.Repository) {
		var ret RegistrySvcRepositorySpec
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcDefinition) GetRepositoryOk() (*RegistrySvcRepositorySpec, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *RegistrySvcDefinition) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given RegistrySvcRepositorySpec and assigns it to the Repository field.
func (o *RegistrySvcDefinition) SetRepository(v RegistrySvcRepositorySpec) {
	o.Repository = &v
}

func (o RegistrySvcDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiSpecs) {
		toSerialize["apiSpecs"] = o.ApiSpecs
	}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !IsNil(o.Envars) {
		toSerialize["envars"] = o.Envars
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	return toSerialize, nil
}

func (o *RegistrySvcDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varRegistrySvcDefinition := _RegistrySvcDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrySvcDefinition)

	if err != nil {
		return err
	}

	*o = RegistrySvcDefinition(varRegistrySvcDefinition)

	return err
}

type NullableRegistrySvcDefinition struct {
	value *RegistrySvcDefinition
	isSet bool
}

func (v NullableRegistrySvcDefinition) Get() *RegistrySvcDefinition {
	return v.value
}

func (v *NullableRegistrySvcDefinition) Set(val *RegistrySvcDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcDefinition(val *RegistrySvcDefinition) *NullableRegistrySvcDefinition {
	return &NullableRegistrySvcDefinition{value: val, isSet: true}
}

func (v NullableRegistrySvcDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


