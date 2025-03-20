/*
1Backend

AI-native microservices platform.

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

// checks if the PolicySvcInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicySvcInstance{}

// PolicySvcInstance struct for PolicySvcInstance
type PolicySvcInstance struct {
	Endpoint *string `json:"endpoint,omitempty"`
	Id *string `json:"id,omitempty"`
	Parameters PolicySvcParameters `json:"parameters"`
	TemplateId PolicySvcTemplateId `json:"templateId"`
}

type _PolicySvcInstance PolicySvcInstance

// NewPolicySvcInstance instantiates a new PolicySvcInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicySvcInstance(parameters PolicySvcParameters, templateId PolicySvcTemplateId) *PolicySvcInstance {
	this := PolicySvcInstance{}
	this.Parameters = parameters
	this.TemplateId = templateId
	return &this
}

// NewPolicySvcInstanceWithDefaults instantiates a new PolicySvcInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicySvcInstanceWithDefaults() *PolicySvcInstance {
	this := PolicySvcInstance{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *PolicySvcInstance) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySvcInstance) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *PolicySvcInstance) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *PolicySvcInstance) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PolicySvcInstance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySvcInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PolicySvcInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PolicySvcInstance) SetId(v string) {
	o.Id = &v
}

// GetParameters returns the Parameters field value
func (o *PolicySvcInstance) GetParameters() PolicySvcParameters {
	if o == nil {
		var ret PolicySvcParameters
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *PolicySvcInstance) GetParametersOk() (*PolicySvcParameters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *PolicySvcInstance) SetParameters(v PolicySvcParameters) {
	o.Parameters = v
}

// GetTemplateId returns the TemplateId field value
func (o *PolicySvcInstance) GetTemplateId() PolicySvcTemplateId {
	if o == nil {
		var ret PolicySvcTemplateId
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *PolicySvcInstance) GetTemplateIdOk() (*PolicySvcTemplateId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *PolicySvcInstance) SetTemplateId(v PolicySvcTemplateId) {
	o.TemplateId = v
}

func (o PolicySvcInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicySvcInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["parameters"] = o.Parameters
	toSerialize["templateId"] = o.TemplateId
	return toSerialize, nil
}

func (o *PolicySvcInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parameters",
		"templateId",
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

	varPolicySvcInstance := _PolicySvcInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPolicySvcInstance)

	if err != nil {
		return err
	}

	*o = PolicySvcInstance(varPolicySvcInstance)

	return err
}

type NullablePolicySvcInstance struct {
	value *PolicySvcInstance
	isSet bool
}

func (v NullablePolicySvcInstance) Get() *PolicySvcInstance {
	return v.value
}

func (v *NullablePolicySvcInstance) Set(val *PolicySvcInstance) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySvcInstance) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySvcInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySvcInstance(val *PolicySvcInstance) *NullablePolicySvcInstance {
	return &NullablePolicySvcInstance{value: val, isSet: true}
}

func (v NullablePolicySvcInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySvcInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


