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
)

// checks if the ProxySvcListCertsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxySvcListCertsRequest{}

// ProxySvcListCertsRequest struct for ProxySvcListCertsRequest
type ProxySvcListCertsRequest struct {
	Ids []string `json:"ids,omitempty"`
}

// NewProxySvcListCertsRequest instantiates a new ProxySvcListCertsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxySvcListCertsRequest() *ProxySvcListCertsRequest {
	this := ProxySvcListCertsRequest{}
	return &this
}

// NewProxySvcListCertsRequestWithDefaults instantiates a new ProxySvcListCertsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxySvcListCertsRequestWithDefaults() *ProxySvcListCertsRequest {
	this := ProxySvcListCertsRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ProxySvcListCertsRequest) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcListCertsRequest) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ProxySvcListCertsRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *ProxySvcListCertsRequest) SetIds(v []string) {
	o.Ids = v
}

func (o ProxySvcListCertsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxySvcListCertsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableProxySvcListCertsRequest struct {
	value *ProxySvcListCertsRequest
	isSet bool
}

func (v NullableProxySvcListCertsRequest) Get() *ProxySvcListCertsRequest {
	return v.value
}

func (v *NullableProxySvcListCertsRequest) Set(val *ProxySvcListCertsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProxySvcListCertsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProxySvcListCertsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxySvcListCertsRequest(val *ProxySvcListCertsRequest) *NullableProxySvcListCertsRequest {
	return &NullableProxySvcListCertsRequest{value: val, isSet: true}
}

func (v NullableProxySvcListCertsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxySvcListCertsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


