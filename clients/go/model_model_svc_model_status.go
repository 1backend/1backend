/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.31
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelSvcModelStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSvcModelStatus{}

// ModelSvcModelStatus struct for ModelSvcModelStatus
type ModelSvcModelStatus struct {
	Address string `json:"address"`
	AssetsReady bool `json:"assetsReady"`
	// Running triggers onModelLaunch on the frontend.  Running is true when the model is both running and answering  - fully loaded.
	Running bool `json:"running"`
}

type _ModelSvcModelStatus ModelSvcModelStatus

// NewModelSvcModelStatus instantiates a new ModelSvcModelStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSvcModelStatus(address string, assetsReady bool, running bool) *ModelSvcModelStatus {
	this := ModelSvcModelStatus{}
	this.Address = address
	this.AssetsReady = assetsReady
	this.Running = running
	return &this
}

// NewModelSvcModelStatusWithDefaults instantiates a new ModelSvcModelStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSvcModelStatusWithDefaults() *ModelSvcModelStatus {
	this := ModelSvcModelStatus{}
	return &this
}

// GetAddress returns the Address field value
func (o *ModelSvcModelStatus) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ModelSvcModelStatus) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ModelSvcModelStatus) SetAddress(v string) {
	o.Address = v
}

// GetAssetsReady returns the AssetsReady field value
func (o *ModelSvcModelStatus) GetAssetsReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AssetsReady
}

// GetAssetsReadyOk returns a tuple with the AssetsReady field value
// and a boolean to check if the value has been set.
func (o *ModelSvcModelStatus) GetAssetsReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetsReady, true
}

// SetAssetsReady sets field value
func (o *ModelSvcModelStatus) SetAssetsReady(v bool) {
	o.AssetsReady = v
}

// GetRunning returns the Running field value
func (o *ModelSvcModelStatus) GetRunning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Running
}

// GetRunningOk returns a tuple with the Running field value
// and a boolean to check if the value has been set.
func (o *ModelSvcModelStatus) GetRunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Running, true
}

// SetRunning sets field value
func (o *ModelSvcModelStatus) SetRunning(v bool) {
	o.Running = v
}

func (o ModelSvcModelStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSvcModelStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["assetsReady"] = o.AssetsReady
	toSerialize["running"] = o.Running
	return toSerialize, nil
}

func (o *ModelSvcModelStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"assetsReady",
		"running",
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

	varModelSvcModelStatus := _ModelSvcModelStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelSvcModelStatus)

	if err != nil {
		return err
	}

	*o = ModelSvcModelStatus(varModelSvcModelStatus)

	return err
}

type NullableModelSvcModelStatus struct {
	value *ModelSvcModelStatus
	isSet bool
}

func (v NullableModelSvcModelStatus) Get() *ModelSvcModelStatus {
	return v.value
}

func (v *NullableModelSvcModelStatus) Set(val *ModelSvcModelStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSvcModelStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSvcModelStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSvcModelStatus(val *ModelSvcModelStatus) *NullableModelSvcModelStatus {
	return &NullableModelSvcModelStatus{value: val, isSet: true}
}

func (v NullableModelSvcModelStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSvcModelStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


