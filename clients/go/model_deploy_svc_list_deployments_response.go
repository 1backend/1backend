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
)

// checks if the DeploySvcListDeploymentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploySvcListDeploymentsResponse{}

// DeploySvcListDeploymentsResponse struct for DeploySvcListDeploymentsResponse
type DeploySvcListDeploymentsResponse struct {
	Deployments []DeploySvcDeployment `json:"deployments,omitempty"`
}

// NewDeploySvcListDeploymentsResponse instantiates a new DeploySvcListDeploymentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploySvcListDeploymentsResponse() *DeploySvcListDeploymentsResponse {
	this := DeploySvcListDeploymentsResponse{}
	return &this
}

// NewDeploySvcListDeploymentsResponseWithDefaults instantiates a new DeploySvcListDeploymentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploySvcListDeploymentsResponseWithDefaults() *DeploySvcListDeploymentsResponse {
	this := DeploySvcListDeploymentsResponse{}
	return &this
}

// GetDeployments returns the Deployments field value if set, zero value otherwise.
func (o *DeploySvcListDeploymentsResponse) GetDeployments() []DeploySvcDeployment {
	if o == nil || IsNil(o.Deployments) {
		var ret []DeploySvcDeployment
		return ret
	}
	return o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploySvcListDeploymentsResponse) GetDeploymentsOk() ([]DeploySvcDeployment, bool) {
	if o == nil || IsNil(o.Deployments) {
		return nil, false
	}
	return o.Deployments, true
}

// HasDeployments returns a boolean if a field has been set.
func (o *DeploySvcListDeploymentsResponse) HasDeployments() bool {
	if o != nil && !IsNil(o.Deployments) {
		return true
	}

	return false
}

// SetDeployments gets a reference to the given []DeploySvcDeployment and assigns it to the Deployments field.
func (o *DeploySvcListDeploymentsResponse) SetDeployments(v []DeploySvcDeployment) {
	o.Deployments = v
}

func (o DeploySvcListDeploymentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploySvcListDeploymentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deployments) {
		toSerialize["deployments"] = o.Deployments
	}
	return toSerialize, nil
}

type NullableDeploySvcListDeploymentsResponse struct {
	value *DeploySvcListDeploymentsResponse
	isSet bool
}

func (v NullableDeploySvcListDeploymentsResponse) Get() *DeploySvcListDeploymentsResponse {
	return v.value
}

func (v *NullableDeploySvcListDeploymentsResponse) Set(val *DeploySvcListDeploymentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploySvcListDeploymentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySvcListDeploymentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySvcListDeploymentsResponse(val *DeploySvcListDeploymentsResponse) *NullableDeploySvcListDeploymentsResponse {
	return &NullableDeploySvcListDeploymentsResponse{value: val, isSet: true}
}

func (v NullableDeploySvcListDeploymentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySvcListDeploymentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


