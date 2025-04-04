/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.31
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PromptSvcListPromptsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptSvcListPromptsRequest{}

// PromptSvcListPromptsRequest struct for PromptSvcListPromptsRequest
type PromptSvcListPromptsRequest struct {
	Query *DatastoreQuery `json:"query,omitempty"`
}

// NewPromptSvcListPromptsRequest instantiates a new PromptSvcListPromptsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptSvcListPromptsRequest() *PromptSvcListPromptsRequest {
	this := PromptSvcListPromptsRequest{}
	return &this
}

// NewPromptSvcListPromptsRequestWithDefaults instantiates a new PromptSvcListPromptsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptSvcListPromptsRequestWithDefaults() *PromptSvcListPromptsRequest {
	this := PromptSvcListPromptsRequest{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *PromptSvcListPromptsRequest) GetQuery() DatastoreQuery {
	if o == nil || IsNil(o.Query) {
		var ret DatastoreQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptSvcListPromptsRequest) GetQueryOk() (*DatastoreQuery, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *PromptSvcListPromptsRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given DatastoreQuery and assigns it to the Query field.
func (o *PromptSvcListPromptsRequest) SetQuery(v DatastoreQuery) {
	o.Query = &v
}

func (o PromptSvcListPromptsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptSvcListPromptsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return toSerialize, nil
}

type NullablePromptSvcListPromptsRequest struct {
	value *PromptSvcListPromptsRequest
	isSet bool
}

func (v NullablePromptSvcListPromptsRequest) Get() *PromptSvcListPromptsRequest {
	return v.value
}

func (v *NullablePromptSvcListPromptsRequest) Set(val *PromptSvcListPromptsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptSvcListPromptsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptSvcListPromptsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptSvcListPromptsRequest(val *PromptSvcListPromptsRequest) *NullablePromptSvcListPromptsRequest {
	return &NullablePromptSvcListPromptsRequest{value: val, isSet: true}
}

func (v NullablePromptSvcListPromptsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptSvcListPromptsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


