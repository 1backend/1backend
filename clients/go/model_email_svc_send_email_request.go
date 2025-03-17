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

// checks if the EmailSvcSendEmailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSvcSendEmailRequest{}

// EmailSvcSendEmailRequest struct for EmailSvcSendEmailRequest
type EmailSvcSendEmailRequest struct {
	// List of file attachments (optional)
	Attachments []EmailSvcFile `json:"attachments,omitempty"`
	// List of BCC recipient email addresses (optional)
	Bcc []string `json:"bcc,omitempty"`
	// Email body content (plain text or HTML)
	Body string `json:"body"`
	// List of CC recipient email addresses (optional)
	Cc []string `json:"cc,omitempty"`
	// Content type: \"text/plain\" or \"text/html\"
	ContentType string `json:"contentType"`
	// Timestamp of email creation
	CreatedAt string `json:"createdAt"`
	// Unique identifier
	Id *string `json:"id,omitempty"`
	// Email subject line
	Subject string `json:"subject"`
	// List of recipient email addresses
	To []string `json:"to"`
}

type _EmailSvcSendEmailRequest EmailSvcSendEmailRequest

// NewEmailSvcSendEmailRequest instantiates a new EmailSvcSendEmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSvcSendEmailRequest(body string, contentType string, createdAt string, subject string, to []string) *EmailSvcSendEmailRequest {
	this := EmailSvcSendEmailRequest{}
	this.Body = body
	this.ContentType = contentType
	this.CreatedAt = createdAt
	this.Subject = subject
	this.To = to
	return &this
}

// NewEmailSvcSendEmailRequestWithDefaults instantiates a new EmailSvcSendEmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSvcSendEmailRequestWithDefaults() *EmailSvcSendEmailRequest {
	this := EmailSvcSendEmailRequest{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *EmailSvcSendEmailRequest) GetAttachments() []EmailSvcFile {
	if o == nil || IsNil(o.Attachments) {
		var ret []EmailSvcFile
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSvcSendEmailRequest) GetAttachmentsOk() ([]EmailSvcFile, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *EmailSvcSendEmailRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []EmailSvcFile and assigns it to the Attachments field.
func (o *EmailSvcSendEmailRequest) SetAttachments(v []EmailSvcFile) {
	o.Attachments = v
}

// GetBcc returns the Bcc field value if set, zero value otherwise.
func (o *EmailSvcSendEmailRequest) GetBcc() []string {
	if o == nil || IsNil(o.Bcc) {
		var ret []string
		return ret
	}
	return o.Bcc
}

// GetBccOk returns a tuple with the Bcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSvcSendEmailRequest) GetBccOk() ([]string, bool) {
	if o == nil || IsNil(o.Bcc) {
		return nil, false
	}
	return o.Bcc, true
}

// HasBcc returns a boolean if a field has been set.
func (o *EmailSvcSendEmailRequest) HasBcc() bool {
	if o != nil && !IsNil(o.Bcc) {
		return true
	}

	return false
}

// SetBcc gets a reference to the given []string and assigns it to the Bcc field.
func (o *EmailSvcSendEmailRequest) SetBcc(v []string) {
	o.Bcc = v
}

// GetBody returns the Body field value
func (o *EmailSvcSendEmailRequest) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *EmailSvcSendEmailRequest) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *EmailSvcSendEmailRequest) SetBody(v string) {
	o.Body = v
}

// GetCc returns the Cc field value if set, zero value otherwise.
func (o *EmailSvcSendEmailRequest) GetCc() []string {
	if o == nil || IsNil(o.Cc) {
		var ret []string
		return ret
	}
	return o.Cc
}

// GetCcOk returns a tuple with the Cc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSvcSendEmailRequest) GetCcOk() ([]string, bool) {
	if o == nil || IsNil(o.Cc) {
		return nil, false
	}
	return o.Cc, true
}

// HasCc returns a boolean if a field has been set.
func (o *EmailSvcSendEmailRequest) HasCc() bool {
	if o != nil && !IsNil(o.Cc) {
		return true
	}

	return false
}

// SetCc gets a reference to the given []string and assigns it to the Cc field.
func (o *EmailSvcSendEmailRequest) SetCc(v []string) {
	o.Cc = v
}

// GetContentType returns the ContentType field value
func (o *EmailSvcSendEmailRequest) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *EmailSvcSendEmailRequest) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *EmailSvcSendEmailRequest) SetContentType(v string) {
	o.ContentType = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EmailSvcSendEmailRequest) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EmailSvcSendEmailRequest) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EmailSvcSendEmailRequest) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailSvcSendEmailRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSvcSendEmailRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailSvcSendEmailRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailSvcSendEmailRequest) SetId(v string) {
	o.Id = &v
}

// GetSubject returns the Subject field value
func (o *EmailSvcSendEmailRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *EmailSvcSendEmailRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *EmailSvcSendEmailRequest) SetSubject(v string) {
	o.Subject = v
}

// GetTo returns the To field value
func (o *EmailSvcSendEmailRequest) GetTo() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *EmailSvcSendEmailRequest) GetToOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.To, true
}

// SetTo sets field value
func (o *EmailSvcSendEmailRequest) SetTo(v []string) {
	o.To = v
}

func (o EmailSvcSendEmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSvcSendEmailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Bcc) {
		toSerialize["bcc"] = o.Bcc
	}
	toSerialize["body"] = o.Body
	if !IsNil(o.Cc) {
		toSerialize["cc"] = o.Cc
	}
	toSerialize["contentType"] = o.ContentType
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["subject"] = o.Subject
	toSerialize["to"] = o.To
	return toSerialize, nil
}

func (o *EmailSvcSendEmailRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"body",
		"contentType",
		"createdAt",
		"subject",
		"to",
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

	varEmailSvcSendEmailRequest := _EmailSvcSendEmailRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailSvcSendEmailRequest)

	if err != nil {
		return err
	}

	*o = EmailSvcSendEmailRequest(varEmailSvcSendEmailRequest)

	return err
}

type NullableEmailSvcSendEmailRequest struct {
	value *EmailSvcSendEmailRequest
	isSet bool
}

func (v NullableEmailSvcSendEmailRequest) Get() *EmailSvcSendEmailRequest {
	return v.value
}

func (v *NullableEmailSvcSendEmailRequest) Set(val *EmailSvcSendEmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSvcSendEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSvcSendEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSvcSendEmailRequest(val *EmailSvcSendEmailRequest) *NullableEmailSvcSendEmailRequest {
	return &NullableEmailSvcSendEmailRequest{value: val, isSet: true}
}

func (v NullableEmailSvcSendEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSvcSendEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


