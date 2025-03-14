/*
1Backend

A common backend for your AI applications—microservices-based and built to scale.

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

// checks if the FileSvcDownloadFileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSvcDownloadFileRequest{}

// FileSvcDownloadFileRequest struct for FileSvcDownloadFileRequest
type FileSvcDownloadFileRequest struct {
	FolderPath *string `json:"folderPath,omitempty"`
	Url string `json:"url"`
}

type _FileSvcDownloadFileRequest FileSvcDownloadFileRequest

// NewFileSvcDownloadFileRequest instantiates a new FileSvcDownloadFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSvcDownloadFileRequest(url string) *FileSvcDownloadFileRequest {
	this := FileSvcDownloadFileRequest{}
	this.Url = url
	return &this
}

// NewFileSvcDownloadFileRequestWithDefaults instantiates a new FileSvcDownloadFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSvcDownloadFileRequestWithDefaults() *FileSvcDownloadFileRequest {
	this := FileSvcDownloadFileRequest{}
	return &this
}

// GetFolderPath returns the FolderPath field value if set, zero value otherwise.
func (o *FileSvcDownloadFileRequest) GetFolderPath() string {
	if o == nil || IsNil(o.FolderPath) {
		var ret string
		return ret
	}
	return *o.FolderPath
}

// GetFolderPathOk returns a tuple with the FolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownloadFileRequest) GetFolderPathOk() (*string, bool) {
	if o == nil || IsNil(o.FolderPath) {
		return nil, false
	}
	return o.FolderPath, true
}

// HasFolderPath returns a boolean if a field has been set.
func (o *FileSvcDownloadFileRequest) HasFolderPath() bool {
	if o != nil && !IsNil(o.FolderPath) {
		return true
	}

	return false
}

// SetFolderPath gets a reference to the given string and assigns it to the FolderPath field.
func (o *FileSvcDownloadFileRequest) SetFolderPath(v string) {
	o.FolderPath = &v
}

// GetUrl returns the Url field value
func (o *FileSvcDownloadFileRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FileSvcDownloadFileRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FileSvcDownloadFileRequest) SetUrl(v string) {
	o.Url = v
}

func (o FileSvcDownloadFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSvcDownloadFileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FolderPath) {
		toSerialize["folderPath"] = o.FolderPath
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *FileSvcDownloadFileRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varFileSvcDownloadFileRequest := _FileSvcDownloadFileRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileSvcDownloadFileRequest)

	if err != nil {
		return err
	}

	*o = FileSvcDownloadFileRequest(varFileSvcDownloadFileRequest)

	return err
}

type NullableFileSvcDownloadFileRequest struct {
	value *FileSvcDownloadFileRequest
	isSet bool
}

func (v NullableFileSvcDownloadFileRequest) Get() *FileSvcDownloadFileRequest {
	return v.value
}

func (v *NullableFileSvcDownloadFileRequest) Set(val *FileSvcDownloadFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSvcDownloadFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSvcDownloadFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSvcDownloadFileRequest(val *FileSvcDownloadFileRequest) *NullableFileSvcDownloadFileRequest {
	return &NullableFileSvcDownloadFileRequest{value: val, isSet: true}
}

func (v NullableFileSvcDownloadFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSvcDownloadFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


