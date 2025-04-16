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

// checks if the PolicySvcCheckRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicySvcCheckRequest{}

// PolicySvcCheckRequest struct for PolicySvcCheckRequest
type PolicySvcCheckRequest struct {
	Endpoint *string `json:"endpoint,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Method *string `json:"method,omitempty"`
	UserId *string `json:"userId,omitempty"`
}

// NewPolicySvcCheckRequest instantiates a new PolicySvcCheckRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicySvcCheckRequest() *PolicySvcCheckRequest {
	this := PolicySvcCheckRequest{}
	return &this
}

// NewPolicySvcCheckRequestWithDefaults instantiates a new PolicySvcCheckRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicySvcCheckRequestWithDefaults() *PolicySvcCheckRequest {
	this := PolicySvcCheckRequest{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *PolicySvcCheckRequest) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySvcCheckRequest) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *PolicySvcCheckRequest) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *PolicySvcCheckRequest) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *PolicySvcCheckRequest) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySvcCheckRequest) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *PolicySvcCheckRequest) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *PolicySvcCheckRequest) SetIp(v string) {
	o.Ip = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PolicySvcCheckRequest) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySvcCheckRequest) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PolicySvcCheckRequest) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PolicySvcCheckRequest) SetMethod(v string) {
	o.Method = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PolicySvcCheckRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySvcCheckRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PolicySvcCheckRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PolicySvcCheckRequest) SetUserId(v string) {
	o.UserId = &v
}

func (o PolicySvcCheckRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicySvcCheckRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullablePolicySvcCheckRequest struct {
	value *PolicySvcCheckRequest
	isSet bool
}

func (v NullablePolicySvcCheckRequest) Get() *PolicySvcCheckRequest {
	return v.value
}

func (v *NullablePolicySvcCheckRequest) Set(val *PolicySvcCheckRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySvcCheckRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySvcCheckRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySvcCheckRequest(val *PolicySvcCheckRequest) *NullablePolicySvcCheckRequest {
	return &NullablePolicySvcCheckRequest{value: val, isSet: true}
}

func (v NullablePolicySvcCheckRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySvcCheckRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


