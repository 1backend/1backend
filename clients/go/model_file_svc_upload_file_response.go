/*
1Backend

A unified backend development platform for microservices-based AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FileSvcUploadFileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSvcUploadFileResponse{}

// FileSvcUploadFileResponse struct for FileSvcUploadFileResponse
type FileSvcUploadFileResponse struct {
	Upload *FileSvcUpload `json:"upload,omitempty"`
}

// NewFileSvcUploadFileResponse instantiates a new FileSvcUploadFileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSvcUploadFileResponse() *FileSvcUploadFileResponse {
	this := FileSvcUploadFileResponse{}
	return &this
}

// NewFileSvcUploadFileResponseWithDefaults instantiates a new FileSvcUploadFileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSvcUploadFileResponseWithDefaults() *FileSvcUploadFileResponse {
	this := FileSvcUploadFileResponse{}
	return &this
}

// GetUpload returns the Upload field value if set, zero value otherwise.
func (o *FileSvcUploadFileResponse) GetUpload() FileSvcUpload {
	if o == nil || IsNil(o.Upload) {
		var ret FileSvcUpload
		return ret
	}
	return *o.Upload
}

// GetUploadOk returns a tuple with the Upload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcUploadFileResponse) GetUploadOk() (*FileSvcUpload, bool) {
	if o == nil || IsNil(o.Upload) {
		return nil, false
	}
	return o.Upload, true
}

// HasUpload returns a boolean if a field has been set.
func (o *FileSvcUploadFileResponse) HasUpload() bool {
	if o != nil && !IsNil(o.Upload) {
		return true
	}

	return false
}

// SetUpload gets a reference to the given FileSvcUpload and assigns it to the Upload field.
func (o *FileSvcUploadFileResponse) SetUpload(v FileSvcUpload) {
	o.Upload = &v
}

func (o FileSvcUploadFileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSvcUploadFileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Upload) {
		toSerialize["upload"] = o.Upload
	}
	return toSerialize, nil
}

type NullableFileSvcUploadFileResponse struct {
	value *FileSvcUploadFileResponse
	isSet bool
}

func (v NullableFileSvcUploadFileResponse) Get() *FileSvcUploadFileResponse {
	return v.value
}

func (v *NullableFileSvcUploadFileResponse) Set(val *FileSvcUploadFileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSvcUploadFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSvcUploadFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSvcUploadFileResponse(val *FileSvcUploadFileResponse) *NullableFileSvcUploadFileResponse {
	return &NullableFileSvcUploadFileResponse{value: val, isSet: true}
}

func (v NullableFileSvcUploadFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSvcUploadFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


