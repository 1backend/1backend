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

// checks if the DeploySvcDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploySvcDeployment{}

// DeploySvcDeployment struct for DeploySvcDeployment
type DeploySvcDeployment struct {
	// Optional: Auto-scaling rules
	AutoScaling *DeploySvcAutoScalingConfig `json:"autoScaling,omitempty"`
	// DefinitionId is the id of the definition
	DefinitionId string `json:"definitionId"`
	// Description of what this deployment does
	Description *string `json:"description,omitempty"`
	// Details provides additional information about the deployment's current state, including both success and failure conditions (e.g., \"Deployment in progress\", \"Error pulling image\").
	Details *string `json:"details,omitempty"`
	// Envars is a map of environment variables that will be passed down to service instances (see Registry Svc Instance) Also see the Registry Svc Definition for required envars.
	Envars *map[string]string `json:"envars,omitempty"`
	// ID of the deployment (e.g., \"depl_dbOdi5eLQK\")
	Id string `json:"id"`
	// Short name for easy reference (e.g., \"user-service-v2\")
	Name *string `json:"name,omitempty"`
	// Number of container instances to run
	Replicas *int32 `json:"replicas,omitempty"`
	// Resource requirements for each replica
	Resources *DeploySvcResourceLimits `json:"resources,omitempty"`
	// Current status of the deployment (e.g., \"OK\", \"Error\", \"Pending\")
	Status *DeploySvcDeploymentStatus `json:"status,omitempty"`
	// Deployment strategy (e.g., rolling update)
	Strategy *DeploySvcDeploymentStrategy `json:"strategy,omitempty"`
	// Target deployment regions or clusters
	TargetRegions []DeploySvcTargetRegion `json:"targetRegions,omitempty"`
}

type _DeploySvcDeployment DeploySvcDeployment

// NewDeploySvcDeployment instantiates a new DeploySvcDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploySvcDeployment(definitionId string, id string) *DeploySvcDeployment {
	this := DeploySvcDeployment{}
	this.DefinitionId = definitionId
	this.Id = id
	return &this
}

// NewDeploySvcDeploymentWithDefaults instantiates a new DeploySvcDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploySvcDeploymentWithDefaults() *DeploySvcDeployment {
	this := DeploySvcDeployment{}
	return &this
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise.
func (o *DeploySvcDeployment) GetAutoScaling() DeploySvcAutoScalingConfig {
	if o == nil || IsNil(o.AutoScaling) {
		var ret DeploySvcAutoScalingConfig
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeployment) GetAutoScalingOk() (*DeploySvcAutoScalingConfig, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}
	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *DeploySvcDeployment) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given DeploySvcAutoScalingConfig and assigns it to the AutoScaling field.
func (o *DeploySvcDeployment) SetAutoScaling(v DeploySvcAutoScalingConfig) {
	o.AutoScaling = &v
}

// GetDefinitionId returns the DefinitionId field value
func (o *DeploySvcDeployment) GetDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefinitionId
}

// GetDefinitionIdOk returns a tuple with the DefinitionId field value
// and a boolean to check if the value has been set.
func (o *DeploySvcDeployment) GetDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefinitionId, true
}

// SetDefinitionId sets field value
func (o *DeploySvcDeployment) SetDefinitionId(v string) {
	o.DefinitionId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeploySvcDeployment) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeployment) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeploySvcDeployment) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeploySvcDeployment) SetDescription(v string) {
	o.Description = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *DeploySvcDeployment) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeployment) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *DeploySvcDeployment) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *DeploySvcDeployment) SetDetails(v string) {
	o.Details = &v
}

// GetEnvars returns the Envars field value if set, zero value otherwise.
func (o *DeploySvcDeployment) GetEnvars() map[string]string {
	if o == nil || IsNil(o.Envars) {
		var ret map[string]string
		return ret
	}
	return *o.Envars
}

// GetEnvarsOk returns a tuple with the Envars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeployment) GetEnvarsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Envars) {
		return nil, false
	}
	return o.Envars, true
}

// HasEnvars returns a boolean if a field has been set.
func (o *DeploySvcDeployment) HasEnvars() bool {
	if o != nil && !IsNil(o.Envars) {
		return true
	}

	return false
}

// SetEnvars gets a reference to the given map[string]string and assigns it to the Envars field.
func (o *DeploySvcDeployment) SetEnvars(v map[string]string) {
	o.Envars = &v
}

// GetId returns the Id field value
func (o *DeploySvcDeployment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploySvcDeployment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeploySvcDeployment) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploySvcDeployment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeployment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploySvcDeployment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploySvcDeployment) SetName(v string) {
	o.Name = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *DeploySvcDeployment) GetReplicas() int32 {
	if o == nil || IsNil(o.Replicas) {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeployment) GetReplicasOk() (*int32, bool) {
	if o == nil || IsNil(o.Replicas) {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *DeploySvcDeployment) HasReplicas() bool {
	if o != nil && !IsNil(o.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *DeploySvcDeployment) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *DeploySvcDeployment) GetResources() DeploySvcResourceLimits {
	if o == nil || IsNil(o.Resources) {
		var ret DeploySvcResourceLimits
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeployment) GetResourcesOk() (*DeploySvcResourceLimits, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *DeploySvcDeployment) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given DeploySvcResourceLimits and assigns it to the Resources field.
func (o *DeploySvcDeployment) SetResources(v DeploySvcResourceLimits) {
	o.Resources = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeploySvcDeployment) GetStatus() DeploySvcDeploymentStatus {
	if o == nil || IsNil(o.Status) {
		var ret DeploySvcDeploymentStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeployment) GetStatusOk() (*DeploySvcDeploymentStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeploySvcDeployment) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeploySvcDeploymentStatus and assigns it to the Status field.
func (o *DeploySvcDeployment) SetStatus(v DeploySvcDeploymentStatus) {
	o.Status = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *DeploySvcDeployment) GetStrategy() DeploySvcDeploymentStrategy {
	if o == nil || IsNil(o.Strategy) {
		var ret DeploySvcDeploymentStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeployment) GetStrategyOk() (*DeploySvcDeploymentStrategy, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *DeploySvcDeployment) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given DeploySvcDeploymentStrategy and assigns it to the Strategy field.
func (o *DeploySvcDeployment) SetStrategy(v DeploySvcDeploymentStrategy) {
	o.Strategy = &v
}

// GetTargetRegions returns the TargetRegions field value if set, zero value otherwise.
func (o *DeploySvcDeployment) GetTargetRegions() []DeploySvcTargetRegion {
	if o == nil || IsNil(o.TargetRegions) {
		var ret []DeploySvcTargetRegion
		return ret
	}
	return o.TargetRegions
}

// GetTargetRegionsOk returns a tuple with the TargetRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcDeployment) GetTargetRegionsOk() ([]DeploySvcTargetRegion, bool) {
	if o == nil || IsNil(o.TargetRegions) {
		return nil, false
	}
	return o.TargetRegions, true
}

// HasTargetRegions returns a boolean if a field has been set.
func (o *DeploySvcDeployment) HasTargetRegions() bool {
	if o != nil && !IsNil(o.TargetRegions) {
		return true
	}

	return false
}

// SetTargetRegions gets a reference to the given []DeploySvcTargetRegion and assigns it to the TargetRegions field.
func (o *DeploySvcDeployment) SetTargetRegions(v []DeploySvcTargetRegion) {
	o.TargetRegions = v
}

func (o DeploySvcDeployment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploySvcDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoScaling) {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	toSerialize["definitionId"] = o.DefinitionId
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Envars) {
		toSerialize["envars"] = o.Envars
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Replicas) {
		toSerialize["replicas"] = o.Replicas
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !IsNil(o.TargetRegions) {
		toSerialize["targetRegions"] = o.TargetRegions
	}
	return toSerialize, nil
}

func (o *DeploySvcDeployment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"definitionId",
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

	varDeploySvcDeployment := _DeploySvcDeployment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeploySvcDeployment)

	if err != nil {
		return err
	}

	*o = DeploySvcDeployment(varDeploySvcDeployment)

	return err
}

type NullableDeploySvcDeployment struct {
	value *DeploySvcDeployment
	isSet bool
}

func (v NullableDeploySvcDeployment) Get() *DeploySvcDeployment {
	return v.value
}

func (v *NullableDeploySvcDeployment) Set(val *DeploySvcDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploySvcDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySvcDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySvcDeployment(val *DeploySvcDeployment) *NullableDeploySvcDeployment {
	return &NullableDeploySvcDeployment{value: val, isSet: true}
}

func (v NullableDeploySvcDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySvcDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


