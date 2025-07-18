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
)

// checks if the ContainerSvcNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcNetwork{}

// ContainerSvcNetwork struct for ContainerSvcNetwork
type ContainerSvcNetwork struct {
	// IPAddress is the assigned IP address of the container.
	IpAddress *string `json:"ipAddress,omitempty"`
	// MacAddress is the container's MAC address if applicable.
	MacAddress *string `json:"macAddress,omitempty"`
	// Mode specifies the container's network mode (e.g., bridge, host, none, custom).
	Mode *string `json:"mode,omitempty"`
}

// NewContainerSvcNetwork instantiates a new ContainerSvcNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcNetwork() *ContainerSvcNetwork {
	this := ContainerSvcNetwork{}
	return &this
}

// NewContainerSvcNetworkWithDefaults instantiates a new ContainerSvcNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcNetworkWithDefaults() *ContainerSvcNetwork {
	this := ContainerSvcNetwork{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ContainerSvcNetwork) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcNetwork) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ContainerSvcNetwork) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ContainerSvcNetwork) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *ContainerSvcNetwork) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcNetwork) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *ContainerSvcNetwork) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *ContainerSvcNetwork) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ContainerSvcNetwork) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcNetwork) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ContainerSvcNetwork) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ContainerSvcNetwork) SetMode(v string) {
	o.Mode = &v
}

func (o ContainerSvcNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableContainerSvcNetwork struct {
	value *ContainerSvcNetwork
	isSet bool
}

func (v NullableContainerSvcNetwork) Get() *ContainerSvcNetwork {
	return v.value
}

func (v *NullableContainerSvcNetwork) Set(val *ContainerSvcNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcNetwork(val *ContainerSvcNetwork) *NullableContainerSvcNetwork {
	return &NullableContainerSvcNetwork{value: val, isSet: true}
}

func (v NullableContainerSvcNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


