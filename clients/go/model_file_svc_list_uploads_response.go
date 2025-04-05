/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.33
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FileSvcListUploadsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSvcListUploadsResponse{}

// FileSvcListUploadsResponse struct for FileSvcListUploadsResponse
type FileSvcListUploadsResponse struct {
	Uploads []FileSvcUpload `json:"uploads,omitempty"`
}

// NewFileSvcListUploadsResponse instantiates a new FileSvcListUploadsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSvcListUploadsResponse() *FileSvcListUploadsResponse {
	this := FileSvcListUploadsResponse{}
	return &this
}

// NewFileSvcListUploadsResponseWithDefaults instantiates a new FileSvcListUploadsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSvcListUploadsResponseWithDefaults() *FileSvcListUploadsResponse {
	this := FileSvcListUploadsResponse{}
	return &this
}

// GetUploads returns the Uploads field value if set, zero value otherwise.
func (o *FileSvcListUploadsResponse) GetUploads() []FileSvcUpload {
	if o == nil || IsNil(o.Uploads) {
		var ret []FileSvcUpload
		return ret
	}
	return o.Uploads
}

// GetUploadsOk returns a tuple with the Uploads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcListUploadsResponse) GetUploadsOk() ([]FileSvcUpload, bool) {
	if o == nil || IsNil(o.Uploads) {
		return nil, false
	}
	return o.Uploads, true
}

// HasUploads returns a boolean if a field has been set.
func (o *FileSvcListUploadsResponse) HasUploads() bool {
	if o != nil && !IsNil(o.Uploads) {
		return true
	}

	return false
}

// SetUploads gets a reference to the given []FileSvcUpload and assigns it to the Uploads field.
func (o *FileSvcListUploadsResponse) SetUploads(v []FileSvcUpload) {
	o.Uploads = v
}

func (o FileSvcListUploadsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSvcListUploadsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uploads) {
		toSerialize["uploads"] = o.Uploads
	}
	return toSerialize, nil
}

type NullableFileSvcListUploadsResponse struct {
	value *FileSvcListUploadsResponse
	isSet bool
}

func (v NullableFileSvcListUploadsResponse) Get() *FileSvcListUploadsResponse {
	return v.value
}

func (v *NullableFileSvcListUploadsResponse) Set(val *FileSvcListUploadsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSvcListUploadsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSvcListUploadsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSvcListUploadsResponse(val *FileSvcListUploadsResponse) *NullableFileSvcListUploadsResponse {
	return &NullableFileSvcListUploadsResponse{value: val, isSet: true}
}

func (v NullableFileSvcListUploadsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSvcListUploadsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


