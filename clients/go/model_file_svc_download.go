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
)

// checks if the FileSvcDownload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSvcDownload{}

// FileSvcDownload struct for FileSvcDownload
type FileSvcDownload struct {
	CreatedAt *string `json:"createdAt,omitempty"`
	// DownloadedBytes exists to show the download progress in terms of the number of bytes already downloaded.
	DownloadedBytes *int64 `json:"downloadedBytes,omitempty"`
	Error *string `json:"error,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	FilePath *string `json:"filePath,omitempty"`
	// FileSize is the full final downloaded file size.
	FileSize *int64 `json:"fileSize,omitempty"`
	Id *string `json:"id,omitempty"`
	Progress *float32 `json:"progress,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewFileSvcDownload instantiates a new FileSvcDownload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSvcDownload() *FileSvcDownload {
	this := FileSvcDownload{}
	return &this
}

// NewFileSvcDownloadWithDefaults instantiates a new FileSvcDownload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSvcDownloadWithDefaults() *FileSvcDownload {
	this := FileSvcDownload{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FileSvcDownload) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownload) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FileSvcDownload) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *FileSvcDownload) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDownloadedBytes returns the DownloadedBytes field value if set, zero value otherwise.
func (o *FileSvcDownload) GetDownloadedBytes() int64 {
	if o == nil || IsNil(o.DownloadedBytes) {
		var ret int64
		return ret
	}
	return *o.DownloadedBytes
}

// GetDownloadedBytesOk returns a tuple with the DownloadedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownload) GetDownloadedBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DownloadedBytes) {
		return nil, false
	}
	return o.DownloadedBytes, true
}

// HasDownloadedBytes returns a boolean if a field has been set.
func (o *FileSvcDownload) HasDownloadedBytes() bool {
	if o != nil && !IsNil(o.DownloadedBytes) {
		return true
	}

	return false
}

// SetDownloadedBytes gets a reference to the given int64 and assigns it to the DownloadedBytes field.
func (o *FileSvcDownload) SetDownloadedBytes(v int64) {
	o.DownloadedBytes = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *FileSvcDownload) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownload) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *FileSvcDownload) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *FileSvcDownload) SetError(v string) {
	o.Error = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *FileSvcDownload) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownload) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *FileSvcDownload) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *FileSvcDownload) SetFileName(v string) {
	o.FileName = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *FileSvcDownload) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownload) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *FileSvcDownload) HasFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *FileSvcDownload) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *FileSvcDownload) GetFileSize() int64 {
	if o == nil || IsNil(o.FileSize) {
		var ret int64
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownload) GetFileSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *FileSvcDownload) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int64 and assigns it to the FileSize field.
func (o *FileSvcDownload) SetFileSize(v int64) {
	o.FileSize = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileSvcDownload) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownload) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileSvcDownload) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FileSvcDownload) SetId(v string) {
	o.Id = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *FileSvcDownload) GetProgress() float32 {
	if o == nil || IsNil(o.Progress) {
		var ret float32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownload) GetProgressOk() (*float32, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *FileSvcDownload) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given float32 and assigns it to the Progress field.
func (o *FileSvcDownload) SetProgress(v float32) {
	o.Progress = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FileSvcDownload) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownload) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FileSvcDownload) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FileSvcDownload) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FileSvcDownload) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownload) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FileSvcDownload) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *FileSvcDownload) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *FileSvcDownload) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcDownload) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *FileSvcDownload) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *FileSvcDownload) SetUrl(v string) {
	o.Url = &v
}

func (o FileSvcDownload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSvcDownload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.DownloadedBytes) {
		toSerialize["downloadedBytes"] = o.DownloadedBytes
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.FilePath) {
		toSerialize["filePath"] = o.FilePath
	}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableFileSvcDownload struct {
	value *FileSvcDownload
	isSet bool
}

func (v NullableFileSvcDownload) Get() *FileSvcDownload {
	return v.value
}

func (v *NullableFileSvcDownload) Set(val *FileSvcDownload) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSvcDownload) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSvcDownload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSvcDownload(val *FileSvcDownload) *NullableFileSvcDownload {
	return &NullableFileSvcDownload{value: val, isSet: true}
}

func (v NullableFileSvcDownload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSvcDownload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


