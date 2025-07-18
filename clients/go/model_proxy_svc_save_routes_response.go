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

// checks if the ProxySvcSaveRoutesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxySvcSaveRoutesResponse{}

// ProxySvcSaveRoutesResponse struct for ProxySvcSaveRoutesResponse
type ProxySvcSaveRoutesResponse struct {
	Routes []ProxySvcRoute `json:"routes,omitempty"`
}

// NewProxySvcSaveRoutesResponse instantiates a new ProxySvcSaveRoutesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxySvcSaveRoutesResponse() *ProxySvcSaveRoutesResponse {
	this := ProxySvcSaveRoutesResponse{}
	return &this
}

// NewProxySvcSaveRoutesResponseWithDefaults instantiates a new ProxySvcSaveRoutesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxySvcSaveRoutesResponseWithDefaults() *ProxySvcSaveRoutesResponse {
	this := ProxySvcSaveRoutesResponse{}
	return &this
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *ProxySvcSaveRoutesResponse) GetRoutes() []ProxySvcRoute {
	if o == nil || IsNil(o.Routes) {
		var ret []ProxySvcRoute
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySvcSaveRoutesResponse) GetRoutesOk() ([]ProxySvcRoute, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *ProxySvcSaveRoutesResponse) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []ProxySvcRoute and assigns it to the Routes field.
func (o *ProxySvcSaveRoutesResponse) SetRoutes(v []ProxySvcRoute) {
	o.Routes = v
}

func (o ProxySvcSaveRoutesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxySvcSaveRoutesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	return toSerialize, nil
}

type NullableProxySvcSaveRoutesResponse struct {
	value *ProxySvcSaveRoutesResponse
	isSet bool
}

func (v NullableProxySvcSaveRoutesResponse) Get() *ProxySvcSaveRoutesResponse {
	return v.value
}

func (v *NullableProxySvcSaveRoutesResponse) Set(val *ProxySvcSaveRoutesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProxySvcSaveRoutesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProxySvcSaveRoutesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxySvcSaveRoutesResponse(val *ProxySvcSaveRoutesResponse) *NullableProxySvcSaveRoutesResponse {
	return &NullableProxySvcSaveRoutesResponse{value: val, isSet: true}
}

func (v NullableProxySvcSaveRoutesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxySvcSaveRoutesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


