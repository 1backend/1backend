/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.37
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ContainerSvcListLogsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcListLogsResponse{}

// ContainerSvcListLogsResponse struct for ContainerSvcListLogsResponse
type ContainerSvcListLogsResponse struct {
	Logs []ContainerSvcLog `json:"logs,omitempty"`
}

// NewContainerSvcListLogsResponse instantiates a new ContainerSvcListLogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcListLogsResponse() *ContainerSvcListLogsResponse {
	this := ContainerSvcListLogsResponse{}
	return &this
}

// NewContainerSvcListLogsResponseWithDefaults instantiates a new ContainerSvcListLogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcListLogsResponseWithDefaults() *ContainerSvcListLogsResponse {
	this := ContainerSvcListLogsResponse{}
	return &this
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *ContainerSvcListLogsResponse) GetLogs() []ContainerSvcLog {
	if o == nil || IsNil(o.Logs) {
		var ret []ContainerSvcLog
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerSvcListLogsResponse) GetLogsOk() ([]ContainerSvcLog, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *ContainerSvcListLogsResponse) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []ContainerSvcLog and assigns it to the Logs field.
func (o *ContainerSvcListLogsResponse) SetLogs(v []ContainerSvcLog) {
	o.Logs = v
}

func (o ContainerSvcListLogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcListLogsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	return toSerialize, nil
}

type NullableContainerSvcListLogsResponse struct {
	value *ContainerSvcListLogsResponse
	isSet bool
}

func (v NullableContainerSvcListLogsResponse) Get() *ContainerSvcListLogsResponse {
	return v.value
}

func (v *NullableContainerSvcListLogsResponse) Set(val *ContainerSvcListLogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcListLogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcListLogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcListLogsResponse(val *ContainerSvcListLogsResponse) *NullableContainerSvcListLogsResponse {
	return &NullableContainerSvcListLogsResponse{value: val, isSet: true}
}

func (v NullableContainerSvcListLogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcListLogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


