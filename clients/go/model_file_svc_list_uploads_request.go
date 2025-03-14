/*
OpenOrch

A language-agnostic microservices framework for building AI applications.

API version: 0.3.0-rc.29
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FileSvcListUploadsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSvcListUploadsRequest{}

// FileSvcListUploadsRequest struct for FileSvcListUploadsRequest
type FileSvcListUploadsRequest struct {
	// After time value
	After *string `json:"after,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	UserId *string `json:"userId,omitempty"`
}

// NewFileSvcListUploadsRequest instantiates a new FileSvcListUploadsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSvcListUploadsRequest() *FileSvcListUploadsRequest {
	this := FileSvcListUploadsRequest{}
	return &this
}

// NewFileSvcListUploadsRequestWithDefaults instantiates a new FileSvcListUploadsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSvcListUploadsRequestWithDefaults() *FileSvcListUploadsRequest {
	this := FileSvcListUploadsRequest{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *FileSvcListUploadsRequest) GetAfter() string {
	if o == nil || IsNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcListUploadsRequest) GetAfterOk() (*string, bool) {
	if o == nil || IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *FileSvcListUploadsRequest) HasAfter() bool {
	if o != nil && !IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *FileSvcListUploadsRequest) SetAfter(v string) {
	o.After = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *FileSvcListUploadsRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcListUploadsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *FileSvcListUploadsRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *FileSvcListUploadsRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FileSvcListUploadsRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSvcListUploadsRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FileSvcListUploadsRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *FileSvcListUploadsRequest) SetUserId(v string) {
	o.UserId = &v
}

func (o FileSvcListUploadsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSvcListUploadsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.After) {
		toSerialize["after"] = o.After
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableFileSvcListUploadsRequest struct {
	value *FileSvcListUploadsRequest
	isSet bool
}

func (v NullableFileSvcListUploadsRequest) Get() *FileSvcListUploadsRequest {
	return v.value
}

func (v *NullableFileSvcListUploadsRequest) Set(val *FileSvcListUploadsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSvcListUploadsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSvcListUploadsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSvcListUploadsRequest(val *FileSvcListUploadsRequest) *NullableFileSvcListUploadsRequest {
	return &NullableFileSvcListUploadsRequest{value: val, isSet: true}
}

func (v NullableFileSvcListUploadsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSvcListUploadsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


