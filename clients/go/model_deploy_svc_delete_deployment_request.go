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

// checks if the DeploySvcDeleteDeploymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploySvcDeleteDeploymentRequest{}

// DeploySvcDeleteDeploymentRequest struct for DeploySvcDeleteDeploymentRequest
type DeploySvcDeleteDeploymentRequest struct {
	DeploymentId string `json:"deploymentId"`
}

type _DeploySvcDeleteDeploymentRequest DeploySvcDeleteDeploymentRequest

// NewDeploySvcDeleteDeploymentRequest instantiates a new DeploySvcDeleteDeploymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploySvcDeleteDeploymentRequest(deploymentId string) *DeploySvcDeleteDeploymentRequest {
	this := DeploySvcDeleteDeploymentRequest{}
	this.DeploymentId = deploymentId
	return &this
}

// NewDeploySvcDeleteDeploymentRequestWithDefaults instantiates a new DeploySvcDeleteDeploymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploySvcDeleteDeploymentRequestWithDefaults() *DeploySvcDeleteDeploymentRequest {
	this := DeploySvcDeleteDeploymentRequest{}
	return &this
}

// GetDeploymentId returns the DeploymentId field value
func (o *DeploySvcDeleteDeploymentRequest) GetDeploymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value
// and a boolean to check if the value has been set.
func (o *DeploySvcDeleteDeploymentRequest) GetDeploymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentId, true
}

// SetDeploymentId sets field value
func (o *DeploySvcDeleteDeploymentRequest) SetDeploymentId(v string) {
	o.DeploymentId = v
}

func (o DeploySvcDeleteDeploymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploySvcDeleteDeploymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deploymentId"] = o.DeploymentId
	return toSerialize, nil
}

func (o *DeploySvcDeleteDeploymentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deploymentId",
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

	varDeploySvcDeleteDeploymentRequest := _DeploySvcDeleteDeploymentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeploySvcDeleteDeploymentRequest)

	if err != nil {
		return err
	}

	*o = DeploySvcDeleteDeploymentRequest(varDeploySvcDeleteDeploymentRequest)

	return err
}

type NullableDeploySvcDeleteDeploymentRequest struct {
	value *DeploySvcDeleteDeploymentRequest
	isSet bool
}

func (v NullableDeploySvcDeleteDeploymentRequest) Get() *DeploySvcDeleteDeploymentRequest {
	return v.value
}

func (v *NullableDeploySvcDeleteDeploymentRequest) Set(val *DeploySvcDeleteDeploymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploySvcDeleteDeploymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySvcDeleteDeploymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySvcDeleteDeploymentRequest(val *DeploySvcDeleteDeploymentRequest) *NullableDeploySvcDeleteDeploymentRequest {
	return &NullableDeploySvcDeleteDeploymentRequest{value: val, isSet: true}
}

func (v NullableDeploySvcDeleteDeploymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySvcDeleteDeploymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


