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

// checks if the RegistrySvcNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrySvcNode{}

// RegistrySvcNode struct for RegistrySvcNode
type RegistrySvcNode struct {
	// The availability zone of the node
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	// List of GPUs available on the node
	Gpus []RegistrySvcGPU `json:"gpus,omitempty"`
	// Required: ID of the instance
	Id string `json:"id"`
	// Last time the instance gave a sign of life
	LastHeartbeat *string `json:"lastHeartbeat,omitempty"`
	// The region of the node
	Region *string `json:"region,omitempty"`
	// URL of the daemon running on the node. If not configured defaults to hostname + default 1Backend server port.
	Url string `json:"url"`
	// Resource usage metrics of the node.
	Usage *RegistrySvcResourceUsage `json:"usage,omitempty"`
}

type _RegistrySvcNode RegistrySvcNode

// NewRegistrySvcNode instantiates a new RegistrySvcNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySvcNode(id string, url string) *RegistrySvcNode {
	this := RegistrySvcNode{}
	this.Id = id
	this.Url = url
	return &this
}

// NewRegistrySvcNodeWithDefaults instantiates a new RegistrySvcNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySvcNodeWithDefaults() *RegistrySvcNode {
	this := RegistrySvcNode{}
	return &this
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *RegistrySvcNode) GetAvailabilityZone() string {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcNode) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *RegistrySvcNode) HasAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *RegistrySvcNode) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

// GetGpus returns the Gpus field value if set, zero value otherwise.
func (o *RegistrySvcNode) GetGpus() []RegistrySvcGPU {
	if o == nil || IsNil(o.Gpus) {
		var ret []RegistrySvcGPU
		return ret
	}
	return o.Gpus
}

// GetGpusOk returns a tuple with the Gpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcNode) GetGpusOk() ([]RegistrySvcGPU, bool) {
	if o == nil || IsNil(o.Gpus) {
		return nil, false
	}
	return o.Gpus, true
}

// HasGpus returns a boolean if a field has been set.
func (o *RegistrySvcNode) HasGpus() bool {
	if o != nil && !IsNil(o.Gpus) {
		return true
	}

	return false
}

// SetGpus gets a reference to the given []RegistrySvcGPU and assigns it to the Gpus field.
func (o *RegistrySvcNode) SetGpus(v []RegistrySvcGPU) {
	o.Gpus = v
}

// GetId returns the Id field value
func (o *RegistrySvcNode) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcNode) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RegistrySvcNode) SetId(v string) {
	o.Id = v
}

// GetLastHeartbeat returns the LastHeartbeat field value if set, zero value otherwise.
func (o *RegistrySvcNode) GetLastHeartbeat() string {
	if o == nil || IsNil(o.LastHeartbeat) {
		var ret string
		return ret
	}
	return *o.LastHeartbeat
}

// GetLastHeartbeatOk returns a tuple with the LastHeartbeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcNode) GetLastHeartbeatOk() (*string, bool) {
	if o == nil || IsNil(o.LastHeartbeat) {
		return nil, false
	}
	return o.LastHeartbeat, true
}

// HasLastHeartbeat returns a boolean if a field has been set.
func (o *RegistrySvcNode) HasLastHeartbeat() bool {
	if o != nil && !IsNil(o.LastHeartbeat) {
		return true
	}

	return false
}

// SetLastHeartbeat gets a reference to the given string and assigns it to the LastHeartbeat field.
func (o *RegistrySvcNode) SetLastHeartbeat(v string) {
	o.LastHeartbeat = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *RegistrySvcNode) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcNode) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *RegistrySvcNode) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *RegistrySvcNode) SetRegion(v string) {
	o.Region = &v
}

// GetUrl returns the Url field value
func (o *RegistrySvcNode) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RegistrySvcNode) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RegistrySvcNode) SetUrl(v string) {
	o.Url = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *RegistrySvcNode) GetUsage() RegistrySvcResourceUsage {
	if o == nil || IsNil(o.Usage) {
		var ret RegistrySvcResourceUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySvcNode) GetUsageOk() (*RegistrySvcResourceUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *RegistrySvcNode) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given RegistrySvcResourceUsage and assigns it to the Usage field.
func (o *RegistrySvcNode) SetUsage(v RegistrySvcResourceUsage) {
	o.Usage = &v
}

func (o RegistrySvcNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrySvcNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailabilityZone) {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if !IsNil(o.Gpus) {
		toSerialize["gpus"] = o.Gpus
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.LastHeartbeat) {
		toSerialize["lastHeartbeat"] = o.LastHeartbeat
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	toSerialize["url"] = o.Url
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

func (o *RegistrySvcNode) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varRegistrySvcNode := _RegistrySvcNode{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrySvcNode)

	if err != nil {
		return err
	}

	*o = RegistrySvcNode(varRegistrySvcNode)

	return err
}

type NullableRegistrySvcNode struct {
	value *RegistrySvcNode
	isSet bool
}

func (v NullableRegistrySvcNode) Get() *RegistrySvcNode {
	return v.value
}

func (v *NullableRegistrySvcNode) Set(val *RegistrySvcNode) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySvcNode) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySvcNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySvcNode(val *RegistrySvcNode) *NullableRegistrySvcNode {
	return &NullableRegistrySvcNode{value: val, isSet: true}
}

func (v NullableRegistrySvcNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySvcNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


