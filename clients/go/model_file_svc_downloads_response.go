/*
OpenOrch

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FileSvcDownloadsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSvcDownloadsResponse{}

// FileSvcDownloadsResponse struct for FileSvcDownloadsResponse
type FileSvcDownloadsResponse struct {
	Downloads []FileSvcDownload `json:"downloads,omitempty"`
}

// NewFileSvcDownloadsResponse instantiates a new FileSvcDownloadsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSvcDownloadsResponse() *FileSvcDownloadsResponse {
	this := FileSvcDownloadsResponse{}
	return &this
}

// NewFileSvcDownloadsResponseWithDefaults instantiates a new FileSvcDownloadsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSvcDownloadsResponseWithDefaults() *FileSvcDownloadsResponse {
	this := FileSvcDownloadsResponse{}
	return &this
}

// GetDownloads returns the Downloads field value if set, zero value otherwise.
func (o *FileSvcDownloadsResponse) GetDownloads() []FileSvcDownload {
	if o == nil || IsNil(o.Downloads) {
		var ret []FileSvcDownload
		return ret
	}
	return o.Downloads
}

// GetDownloadsOk returns a tuple with the Downloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownloadsResponse) GetDownloadsOk() ([]FileSvcDownload, bool) {
	if o == nil || IsNil(o.Downloads) {
		return nil, false
	}
	return o.Downloads, true
}

// HasDownloads returns a boolean if a field has been set.
func (o *FileSvcDownloadsResponse) HasDownloads() bool {
	if o != nil && !IsNil(o.Downloads) {
		return true
	}

	return false
}

// SetDownloads gets a reference to the given []FileSvcDownload and assigns it to the Downloads field.
func (o *FileSvcDownloadsResponse) SetDownloads(v []FileSvcDownload) {
	o.Downloads = v
}

func (o FileSvcDownloadsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSvcDownloadsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Downloads) {
		toSerialize["downloads"] = o.Downloads
	}
	return toSerialize, nil
}

type NullableFileSvcDownloadsResponse struct {
	value *FileSvcDownloadsResponse
	isSet bool
}

func (v NullableFileSvcDownloadsResponse) Get() *FileSvcDownloadsResponse {
	return v.value
}

func (v *NullableFileSvcDownloadsResponse) Set(val *FileSvcDownloadsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSvcDownloadsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSvcDownloadsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSvcDownloadsResponse(val *FileSvcDownloadsResponse) *NullableFileSvcDownloadsResponse {
	return &NullableFileSvcDownloadsResponse{value: val, isSet: true}
}

func (v NullableFileSvcDownloadsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSvcDownloadsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


