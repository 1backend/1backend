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

// checks if the ContainerSvcGetContainerSummaryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerSvcGetContainerSummaryResponse{}

// ContainerSvcGetContainerSummaryResponse struct for ContainerSvcGetContainerSummaryResponse
type ContainerSvcGetContainerSummaryResponse struct {
	Logs string `json:"logs"`
	Status string `json:"status"`
	// DEPRECATED. Summary contains both Status and Logs.
	Summary string `json:"summary"`
}

type _ContainerSvcGetContainerSummaryResponse ContainerSvcGetContainerSummaryResponse

// NewContainerSvcGetContainerSummaryResponse instantiates a new ContainerSvcGetContainerSummaryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerSvcGetContainerSummaryResponse(logs string, status string, summary string) *ContainerSvcGetContainerSummaryResponse {
	this := ContainerSvcGetContainerSummaryResponse{}
	this.Logs = logs
	this.Status = status
	this.Summary = summary
	return &this
}

// NewContainerSvcGetContainerSummaryResponseWithDefaults instantiates a new ContainerSvcGetContainerSummaryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerSvcGetContainerSummaryResponseWithDefaults() *ContainerSvcGetContainerSummaryResponse {
	this := ContainerSvcGetContainerSummaryResponse{}
	return &this
}

// GetLogs returns the Logs field value
func (o *ContainerSvcGetContainerSummaryResponse) GetLogs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcGetContainerSummaryResponse) GetLogsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logs, true
}

// SetLogs sets field value
func (o *ContainerSvcGetContainerSummaryResponse) SetLogs(v string) {
	o.Logs = v
}

// GetStatus returns the Status field value
func (o *ContainerSvcGetContainerSummaryResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcGetContainerSummaryResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ContainerSvcGetContainerSummaryResponse) SetStatus(v string) {
	o.Status = v
}

// GetSummary returns the Summary field value
func (o *ContainerSvcGetContainerSummaryResponse) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ContainerSvcGetContainerSummaryResponse) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *ContainerSvcGetContainerSummaryResponse) SetSummary(v string) {
	o.Summary = v
}

func (o ContainerSvcGetContainerSummaryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerSvcGetContainerSummaryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logs"] = o.Logs
	toSerialize["status"] = o.Status
	toSerialize["summary"] = o.Summary
	return toSerialize, nil
}

func (o *ContainerSvcGetContainerSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logs",
		"status",
		"summary",
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

	varContainerSvcGetContainerSummaryResponse := _ContainerSvcGetContainerSummaryResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerSvcGetContainerSummaryResponse)

	if err != nil {
		return err
	}

	*o = ContainerSvcGetContainerSummaryResponse(varContainerSvcGetContainerSummaryResponse)

	return err
}

type NullableContainerSvcGetContainerSummaryResponse struct {
	value *ContainerSvcGetContainerSummaryResponse
	isSet bool
}

func (v NullableContainerSvcGetContainerSummaryResponse) Get() *ContainerSvcGetContainerSummaryResponse {
	return v.value
}

func (v *NullableContainerSvcGetContainerSummaryResponse) Set(val *ContainerSvcGetContainerSummaryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerSvcGetContainerSummaryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerSvcGetContainerSummaryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerSvcGetContainerSummaryResponse(val *ContainerSvcGetContainerSummaryResponse) *NullableContainerSvcGetContainerSummaryResponse {
	return &NullableContainerSvcGetContainerSummaryResponse{value: val, isSet: true}
}

func (v NullableContainerSvcGetContainerSummaryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerSvcGetContainerSummaryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


