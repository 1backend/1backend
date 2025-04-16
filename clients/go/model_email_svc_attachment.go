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

// checks if the EmailSvcAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSvcAttachment{}

// EmailSvcAttachment struct for EmailSvcAttachment
type EmailSvcAttachment struct {
	// Base64-encoded file content. Use this for small files. Required for inline attachments (i.e., those not using File Svc, see FileId).
	Content *string `json:"content,omitempty"`
	// MIME type of the file (e.g., \"application/pdf\", \"image/png\") Required for inline attachments (i.e., those not using File Svc, see FileId).
	ContentType string `json:"contentType"`
	// A File Svc file ID. Requires the file to be uploaded separately. Recommended for mid to large-sized files. If this field is specified, all other fields are optional.
	FileId *string `json:"fileId,omitempty"`
	// File name for the attachment. Required for inline attachments (i.e., those not using File Svc, see FileId).
	Filename string `json:"filename"`
}

type _EmailSvcAttachment EmailSvcAttachment

// NewEmailSvcAttachment instantiates a new EmailSvcAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSvcAttachment(contentType string, filename string) *EmailSvcAttachment {
	this := EmailSvcAttachment{}
	this.ContentType = contentType
	this.Filename = filename
	return &this
}

// NewEmailSvcAttachmentWithDefaults instantiates a new EmailSvcAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSvcAttachmentWithDefaults() *EmailSvcAttachment {
	this := EmailSvcAttachment{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *EmailSvcAttachment) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSvcAttachment) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *EmailSvcAttachment) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *EmailSvcAttachment) SetContent(v string) {
	o.Content = &v
}

// GetContentType returns the ContentType field value
func (o *EmailSvcAttachment) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *EmailSvcAttachment) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *EmailSvcAttachment) SetContentType(v string) {
	o.ContentType = v
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *EmailSvcAttachment) GetFileId() string {
	if o == nil || IsNil(o.FileId) {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSvcAttachment) GetFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *EmailSvcAttachment) HasFileId() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *EmailSvcAttachment) SetFileId(v string) {
	o.FileId = &v
}

// GetFilename returns the Filename field value
func (o *EmailSvcAttachment) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *EmailSvcAttachment) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *EmailSvcAttachment) SetFilename(v string) {
	o.Filename = v
}

func (o EmailSvcAttachment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSvcAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	toSerialize["contentType"] = o.ContentType
	if !IsNil(o.FileId) {
		toSerialize["fileId"] = o.FileId
	}
	toSerialize["filename"] = o.Filename
	return toSerialize, nil
}

func (o *EmailSvcAttachment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contentType",
		"filename",
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

	varEmailSvcAttachment := _EmailSvcAttachment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailSvcAttachment)

	if err != nil {
		return err
	}

	*o = EmailSvcAttachment(varEmailSvcAttachment)

	return err
}

type NullableEmailSvcAttachment struct {
	value *EmailSvcAttachment
	isSet bool
}

func (v NullableEmailSvcAttachment) Get() *EmailSvcAttachment {
	return v.value
}

func (v *NullableEmailSvcAttachment) Set(val *EmailSvcAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSvcAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSvcAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSvcAttachment(val *EmailSvcAttachment) *NullableEmailSvcAttachment {
	return &NullableEmailSvcAttachment{value: val, isSet: true}
}

func (v NullableEmailSvcAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSvcAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


