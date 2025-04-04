/*
1Backend

AI-native microservices platform.

API version: 0.3.0-rc.32
Contact: sales@singulatron.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DataSvcQueryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSvcQueryRequest{}

// DataSvcQueryRequest struct for DataSvcQueryRequest
type DataSvcQueryRequest struct {
	Query *DatastoreQuery `json:"query,omitempty"`
	Readers []string `json:"readers,omitempty"`
	Table *string `json:"table,omitempty"`
}

// NewDataSvcQueryRequest instantiates a new DataSvcQueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSvcQueryRequest() *DataSvcQueryRequest {
	this := DataSvcQueryRequest{}
	return &this
}

// NewDataSvcQueryRequestWithDefaults instantiates a new DataSvcQueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSvcQueryRequestWithDefaults() *DataSvcQueryRequest {
	this := DataSvcQueryRequest{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *DataSvcQueryRequest) GetQuery() DatastoreQuery {
	if o == nil || IsNil(o.Query) {
		var ret DatastoreQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcQueryRequest) GetQueryOk() (*DatastoreQuery, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *DataSvcQueryRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given DatastoreQuery and assigns it to the Query field.
func (o *DataSvcQueryRequest) SetQuery(v DatastoreQuery) {
	o.Query = &v
}

// GetReaders returns the Readers field value if set, zero value otherwise.
func (o *DataSvcQueryRequest) GetReaders() []string {
	if o == nil || IsNil(o.Readers) {
		var ret []string
		return ret
	}
	return o.Readers
}

// GetReadersOk returns a tuple with the Readers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcQueryRequest) GetReadersOk() ([]string, bool) {
	if o == nil || IsNil(o.Readers) {
		return nil, false
	}
	return o.Readers, true
}

// HasReaders returns a boolean if a field has been set.
func (o *DataSvcQueryRequest) HasReaders() bool {
	if o != nil && !IsNil(o.Readers) {
		return true
	}

	return false
}

// SetReaders gets a reference to the given []string and assigns it to the Readers field.
func (o *DataSvcQueryRequest) SetReaders(v []string) {
	o.Readers = v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *DataSvcQueryRequest) GetTable() string {
	if o == nil || IsNil(o.Table) {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSvcQueryRequest) GetTableOk() (*string, bool) {
	if o == nil || IsNil(o.Table) {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *DataSvcQueryRequest) HasTable() bool {
	if o != nil && !IsNil(o.Table) {
		return true
	}

	return false
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *DataSvcQueryRequest) SetTable(v string) {
	o.Table = &v
}

func (o DataSvcQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSvcQueryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Readers) {
		toSerialize["readers"] = o.Readers
	}
	if !IsNil(o.Table) {
		toSerialize["table"] = o.Table
	}
	return toSerialize, nil
}

type NullableDataSvcQueryRequest struct {
	value *DataSvcQueryRequest
	isSet bool
}

func (v NullableDataSvcQueryRequest) Get() *DataSvcQueryRequest {
	return v.value
}

func (v *NullableDataSvcQueryRequest) Set(val *DataSvcQueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSvcQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSvcQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSvcQueryRequest(val *DataSvcQueryRequest) *NullableDataSvcQueryRequest {
	return &NullableDataSvcQueryRequest{value: val, isSet: true}
}

func (v NullableDataSvcQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSvcQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


