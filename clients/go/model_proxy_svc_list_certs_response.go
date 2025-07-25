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

// checks if the ProxySvcListCertsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxySvcListCertsResponse{}

// ProxySvcListCertsResponse struct for ProxySvcListCertsResponse
type ProxySvcListCertsResponse struct {
	Certs []ProxySvcCert `json:"certs,omitempty"`
}

// NewProxySvcListCertsResponse instantiates a new ProxySvcListCertsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxySvcListCertsResponse() *ProxySvcListCertsResponse {
	this := ProxySvcListCertsResponse{}
	return &this
}

// NewProxySvcListCertsResponseWithDefaults instantiates a new ProxySvcListCertsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxySvcListCertsResponseWithDefaults() *ProxySvcListCertsResponse {
	this := ProxySvcListCertsResponse{}
	return &this
}

// GetCerts returns the Certs field value if set, zero value otherwise.
func (o *ProxySvcListCertsResponse) GetCerts() []ProxySvcCert {
	if o == nil || IsNil(o.Certs) {
		var ret []ProxySvcCert
		return ret
	}
	return o.Certs
}

// GetCertsOk returns a tuple with the Certs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcListCertsResponse) GetCertsOk() ([]ProxySvcCert, bool) {
	if o == nil || IsNil(o.Certs) {
		return nil, false
	}
	return o.Certs, true
}

// HasCerts returns a boolean if a field has been set.
func (o *ProxySvcListCertsResponse) HasCerts() bool {
	if o != nil && !IsNil(o.Certs) {
		return true
	}

	return false
}

// SetCerts gets a reference to the given []ProxySvcCert and assigns it to the Certs field.
func (o *ProxySvcListCertsResponse) SetCerts(v []ProxySvcCert) {
	o.Certs = v
}

func (o ProxySvcListCertsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxySvcListCertsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Certs) {
		toSerialize["certs"] = o.Certs
	}
	return toSerialize, nil
}

type NullableProxySvcListCertsResponse struct {
	value *ProxySvcListCertsResponse
	isSet bool
}

func (v NullableProxySvcListCertsResponse) Get() *ProxySvcListCertsResponse {
	return v.value
}

func (v *NullableProxySvcListCertsResponse) Set(val *ProxySvcListCertsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProxySvcListCertsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProxySvcListCertsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxySvcListCertsResponse(val *ProxySvcListCertsResponse) *NullableProxySvcListCertsResponse {
	return &NullableProxySvcListCertsResponse{value: val, isSet: true}
}

func (v NullableProxySvcListCertsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxySvcListCertsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


