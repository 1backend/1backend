/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.38
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FileSvcGetDownloadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSvcGetDownloadResponse{}

// FileSvcGetDownloadResponse struct for FileSvcGetDownloadResponse
type FileSvcGetDownloadResponse struct {
	Download *FileSvcDownload `json:"download,omitempty"`
	Exists bool `json:"exists"`
}

type _FileSvcGetDownloadResponse FileSvcGetDownloadResponse

// NewFileSvcGetDownloadResponse instantiates a new FileSvcGetDownloadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSvcGetDownloadResponse(exists bool) *FileSvcGetDownloadResponse {
	this := FileSvcGetDownloadResponse{}
	this.Exists = exists
	return &this
}

// NewFileSvcGetDownloadResponseWithDefaults instantiates a new FileSvcGetDownloadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSvcGetDownloadResponseWithDefaults() *FileSvcGetDownloadResponse {
	this := FileSvcGetDownloadResponse{}
	return &this
}

// GetDownload returns the Download field value if set, zero value otherwise.
func (o *FileSvcGetDownloadResponse) GetDownload() FileSvcDownload {
	if o == nil || IsNil(o.Download) {
		var ret FileSvcDownload
		return ret
	}
	return *o.Download
}

// GetDownloadOk returns a tuple with the Download field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcGetDownloadResponse) GetDownloadOk() (*FileSvcDownload, bool) {
	if o == nil || IsNil(o.Download) {
		return nil, false
	}
	return o.Download, true
}

// HasDownload returns a boolean if a field has been set.
func (o *FileSvcGetDownloadResponse) HasDownload() bool {
	if o != nil && !IsNil(o.Download) {
		return true
	}

	return false
}

// SetDownload gets a reference to the given FileSvcDownload and assigns it to the Download field.
func (o *FileSvcGetDownloadResponse) SetDownload(v FileSvcDownload) {
	o.Download = &v
}

// GetExists returns the Exists field value
func (o *FileSvcGetDownloadResponse) GetExists() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exists
}

// GetExistsOk returns a tuple with the Exists field value
// and a boolean to check if the value has been set.
func (o *FileSvcGetDownloadResponse) GetExistsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exists, true
}

// SetExists sets field value
func (o *FileSvcGetDownloadResponse) SetExists(v bool) {
	o.Exists = v
}

func (o FileSvcGetDownloadResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSvcGetDownloadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Download) {
		toSerialize["download"] = o.Download
	}
	toSerialize["exists"] = o.Exists
	return toSerialize, nil
}

func (o *FileSvcGetDownloadResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"exists",
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

	varFileSvcGetDownloadResponse := _FileSvcGetDownloadResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileSvcGetDownloadResponse)

	if err != nil {
		return err
	}

	*o = FileSvcGetDownloadResponse(varFileSvcGetDownloadResponse)

	return err
}

type NullableFileSvcGetDownloadResponse struct {
	value *FileSvcGetDownloadResponse
	isSet bool
}

func (v NullableFileSvcGetDownloadResponse) Get() *FileSvcGetDownloadResponse {
	return v.value
}

func (v *NullableFileSvcGetDownloadResponse) Set(val *FileSvcGetDownloadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSvcGetDownloadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSvcGetDownloadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSvcGetDownloadResponse(val *FileSvcGetDownloadResponse) *NullableFileSvcGetDownloadResponse {
	return &NullableFileSvcGetDownloadResponse{value: val, isSet: true}
}

func (v NullableFileSvcGetDownloadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSvcGetDownloadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


