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

// checks if the FileSvcUpload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSvcUpload{}

// FileSvcUpload struct for FileSvcUpload
type FileSvcUpload struct {
	CreatedAt string `json:"createdAt"`
	// Logical file ID spanning all replicas
	FileId string `json:"fileId"`
	// Filename is the original name of the file
	FileName string `json:"fileName"`
	// FilePath is the full node local path of the file
	FilePath string `json:"filePath"`
	FileSize int64 `json:"fileSize"`
	// Unique ID for this replica
	Id string `json:"id"`
	// ID of the node storing this replica
	NodeId string `json:"nodeId"`
	UpdatedAt string `json:"updatedAt"`
	UserId string `json:"userId"`
}

type _FileSvcUpload FileSvcUpload

// NewFileSvcUpload instantiates a new FileSvcUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSvcUpload(createdAt string, fileId string, fileName string, filePath string, fileSize int64, id string, nodeId string, updatedAt string, userId string) *FileSvcUpload {
	this := FileSvcUpload{}
	this.CreatedAt = createdAt
	this.FileId = fileId
	this.FileName = fileName
	this.FilePath = filePath
	this.FileSize = fileSize
	this.Id = id
	this.NodeId = nodeId
	this.UpdatedAt = updatedAt
	this.UserId = userId
	return &this
}

// NewFileSvcUploadWithDefaults instantiates a new FileSvcUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSvcUploadWithDefaults() *FileSvcUpload {
	this := FileSvcUpload{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *FileSvcUpload) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FileSvcUpload) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetFileId returns the FileId field value
func (o *FileSvcUpload) GetFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *FileSvcUpload) SetFileId(v string) {
	o.FileId = v
}

// GetFileName returns the FileName field value
func (o *FileSvcUpload) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *FileSvcUpload) SetFileName(v string) {
	o.FileName = v
}

// GetFilePath returns the FilePath field value
func (o *FileSvcUpload) GetFilePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilePath, true
}

// SetFilePath sets field value
func (o *FileSvcUpload) SetFilePath(v string) {
	o.FilePath = v
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

// GetId returns the Id field value
func (o *FileSvcUpload) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileSvcUpload) SetId(v string) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *FileSvcUpload) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *FileSvcUpload) SetNodeId(v string) {
	o.NodeId = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *FileSvcUpload) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *FileSvcUpload) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUserId returns the UserId field value
func (o *FileSvcUpload) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FileSvcUpload) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *FileSvcUpload) SetUserId(v string) {
	o.UserId = v
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
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["fileId"] = o.FileId
	toSerialize["fileName"] = o.FileName
	toSerialize["filePath"] = o.FilePath
	toSerialize["fileSize"] = o.FileSize
	toSerialize["id"] = o.Id
	toSerialize["nodeId"] = o.NodeId
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *FileSvcUpload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"fileId",
		"fileName",
		"filePath",
		"fileSize",
		"id",
		"nodeId",
		"updatedAt",
		"userId",
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


