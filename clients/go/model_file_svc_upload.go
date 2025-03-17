/*
1Backend

A unified backend for your AI applications—microservices-based and built to scale.

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

// checks if the FileSvcUpload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSvcUpload{}

// FileSvcUpload struct for FileSvcUpload
type FileSvcUpload struct {
	CreatedAt *string `json:"createdAt,omitempty"`
	// Logical file ID spanning all replicas
	FileId *string `json:"fileId,omitempty"`
	// Filename is the original name of the file
	FileName *string `json:"fileName,omitempty"`
	// FilePath is the full node local path of the file
	FilePath *string `json:"filePath,omitempty"`
	FileSize int64 `json:"fileSize"`
	// Unique ID for this replica
	Id *string `json:"id,omitempty"`
	// ID of the node storing this replica
	NodeId *string `json:"nodeId,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	UserId *string `json:"userId,omitempty"`
}

type _FileSvcUpload FileSvcUpload

// NewFileSvcUpload instantiates a new FileSvcUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSvcUpload(fileSize int64) *FileSvcUpload {
	this := FileSvcUpload{}
	this.FileSize = fileSize
	return &this
}

// NewFileSvcUploadWithDefaults instantiates a new FileSvcUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSvcUploadWithDefaults() *FileSvcUpload {
	this := FileSvcUpload{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FileSvcUpload) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FileSvcUpload) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *FileSvcUpload) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *FileSvcUpload) GetFileId() string {
	if o == nil || IsNil(o.FileId) {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *FileSvcUpload) HasFileId() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *FileSvcUpload) SetFileId(v string) {
	o.FileId = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *FileSvcUpload) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *FileSvcUpload) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *FileSvcUpload) SetFileName(v string) {
	o.FileName = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *FileSvcUpload) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *FileSvcUpload) HasFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *FileSvcUpload) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFileSize returns the FileSize field value
func (o *FileSvcUpload) GetFileSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetFileSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSize, true
}

// SetFileSize sets field value
func (o *FileSvcUpload) SetFileSize(v int64) {
	o.FileSize = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileSvcUpload) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileSvcUpload) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FileSvcUpload) SetId(v string) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *FileSvcUpload) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *FileSvcUpload) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *FileSvcUpload) SetNodeId(v string) {
	o.NodeId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FileSvcUpload) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FileSvcUpload) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *FileSvcUpload) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FileSvcUpload) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FileSvcUpload) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *FileSvcUpload) SetUserId(v string) {
	o.UserId = &v
}

func (o FileSvcUpload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSvcUpload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.FileId) {
		toSerialize["fileId"] = o.FileId
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.FilePath) {
		toSerialize["filePath"] = o.FilePath
	}
	toSerialize["fileSize"] = o.FileSize
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NodeId) {
		toSerialize["nodeId"] = o.NodeId
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

func (o *FileSvcUpload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fileSize",
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

	varFileSvcUpload := _FileSvcUpload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileSvcUpload)

	if err != nil {
		return err
	}

	*o = FileSvcUpload(varFileSvcUpload)

	return err
}

type NullableFileSvcUpload struct {
	value *FileSvcUpload
	isSet bool
}

func (v NullableFileSvcUpload) Get() *FileSvcUpload {
	return v.value
}

func (v *NullableFileSvcUpload) Set(val *FileSvcUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSvcUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSvcUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSvcUpload(val *FileSvcUpload) *NullableFileSvcUpload {
	return &NullableFileSvcUpload{value: val, isSet: true}
}

func (v NullableFileSvcUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSvcUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


