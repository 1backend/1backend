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

// checks if the ContainerSvcDaemonInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcDaemonInfoResponse{}

// ContainerSvcDaemonInfoResponse struct for ContainerSvcDaemonInfoResponse
type ContainerSvcDaemonInfoResponse struct {
	Address *string `json:"address,omitempty"`
	Available bool `json:"available"`
	Error *string `json:"error,omitempty"`
}

type _ContainerSvcDaemonInfoResponse ContainerSvcDaemonInfoResponse

// NewContainerSvcDaemonInfoResponse instantiates a new ContainerSvcDaemonInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcDaemonInfoResponse(available bool) *ContainerSvcDaemonInfoResponse {
	this := ContainerSvcDaemonInfoResponse{}
	this.Available = available
	return &this
}

// NewContainerSvcDaemonInfoResponseWithDefaults instantiates a new ContainerSvcDaemonInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcDaemonInfoResponseWithDefaults() *ContainerSvcDaemonInfoResponse {
	this := ContainerSvcDaemonInfoResponse{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ContainerSvcDaemonInfoResponse) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcDaemonInfoResponse) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ContainerSvcDaemonInfoResponse) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ContainerSvcDaemonInfoResponse) SetAddress(v string) {
	o.Address = &v
}

// GetAvailable returns the Available field value
func (o *ContainerSvcDaemonInfoResponse) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcDaemonInfoResponse) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *ContainerSvcDaemonInfoResponse) SetAvailable(v bool) {
	o.Available = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ContainerSvcDaemonInfoResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcDaemonInfoResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ContainerSvcDaemonInfoResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ContainerSvcDaemonInfoResponse) SetError(v string) {
	o.Error = &v
}

func (o ContainerSvcDaemonInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcDaemonInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	toSerialize["available"] = o.Available
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

func (o *ContainerSvcDaemonInfoResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"available",
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

	varContainerSvcDaemonInfoResponse := _ContainerSvcDaemonInfoResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerSvcDaemonInfoResponse)

	if err != nil {
		return err
	}

	*o = ContainerSvcDaemonInfoResponse(varContainerSvcDaemonInfoResponse)

	return err
}

type NullableContainerSvcDaemonInfoResponse struct {
	value *ContainerSvcDaemonInfoResponse
	isSet bool
}

func (v NullableContainerSvcDaemonInfoResponse) Get() *ContainerSvcDaemonInfoResponse {
	return v.value
}

func (v *NullableContainerSvcDaemonInfoResponse) Set(val *ContainerSvcDaemonInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcDaemonInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcDaemonInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcDaemonInfoResponse(val *ContainerSvcDaemonInfoResponse) *NullableContainerSvcDaemonInfoResponse {
	return &NullableContainerSvcDaemonInfoResponse{value: val, isSet: true}
}

func (v NullableContainerSvcDaemonInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcDaemonInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


