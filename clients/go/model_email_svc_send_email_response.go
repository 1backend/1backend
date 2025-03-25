/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.30
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the EmailSvcSendEmailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSvcSendEmailResponse{}

// EmailSvcSendEmailResponse struct for EmailSvcSendEmailResponse
type EmailSvcSendEmailResponse struct {
	// Unique identifier for the sent email
	EmailId *string `json:"emailId,omitempty"`
	// Status of the email send operation (\"sent\", \"queued\", etc.)
	Status *string `json:"status,omitempty"`
}

// NewEmailSvcSendEmailResponse instantiates a new EmailSvcSendEmailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSvcSendEmailResponse() *EmailSvcSendEmailResponse {
	this := EmailSvcSendEmailResponse{}
	return &this
}

// NewEmailSvcSendEmailResponseWithDefaults instantiates a new EmailSvcSendEmailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSvcSendEmailResponseWithDefaults() *EmailSvcSendEmailResponse {
	this := EmailSvcSendEmailResponse{}
	return &this
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *EmailSvcSendEmailResponse) GetEmailId() string {
	if o == nil || IsNil(o.EmailId) {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSvcSendEmailResponse) GetEmailIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmailId) {
		return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *EmailSvcSendEmailResponse) HasEmailId() bool {
	if o != nil && !IsNil(o.EmailId) {
		return true
	}

	return false
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *EmailSvcSendEmailResponse) SetEmailId(v string) {
	o.EmailId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EmailSvcSendEmailResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSvcSendEmailResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EmailSvcSendEmailResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EmailSvcSendEmailResponse) SetStatus(v string) {
	o.Status = &v
}

func (o EmailSvcSendEmailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSvcSendEmailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailId) {
		toSerialize["emailId"] = o.EmailId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableEmailSvcSendEmailResponse struct {
	value *EmailSvcSendEmailResponse
	isSet bool
}

func (v NullableEmailSvcSendEmailResponse) Get() *EmailSvcSendEmailResponse {
	return v.value
}

func (v *NullableEmailSvcSendEmailResponse) Set(val *EmailSvcSendEmailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSvcSendEmailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSvcSendEmailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSvcSendEmailResponse(val *EmailSvcSendEmailResponse) *NullableEmailSvcSendEmailResponse {
	return &NullableEmailSvcSendEmailResponse{value: val, isSet: true}
}

func (v NullableEmailSvcSendEmailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSvcSendEmailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


