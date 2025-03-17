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
)

// checks if the DatastoreQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatastoreQuery{}

// DatastoreQuery struct for DatastoreQuery
type DatastoreQuery struct {
	// Count true means return the count of the dataset filtered by Filters without after or limit.
	Count *bool `json:"count,omitempty"`
	// Filters are filtering options of a query. It is advised to use It's advised to use helper functions in your respective client library such as filter constructors (`all`, `equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
	Filters []DatastoreFilter `json:"filters,omitempty"`
	// JSONAfter is used for cursor-based pagination, which is more effective in scalable and distributed environments compared to offset-based pagination.  JSONAfter is a JSON-encoded string due to limitations in Swaggo (e.g., []interface{} gets converted to []map[string]interface{}). There is no way to specify a type that results in an any/interface{} type in the `go -> openapi -> go` generation process. As a result, JSONAfter is a JSON-marshalled string representing an array, e.g., `[42]`.
	JsonAfter *string `json:"jsonAfter,omitempty"`
	// Limit the number of records in the result set.
	Limit *int32 `json:"limit,omitempty"`
	// OrderBys order the result set.
	OrderBys []DatastoreOrderBy `json:"orderBys,omitempty"`
}

// NewDatastoreQuery instantiates a new DatastoreQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatastoreQuery() *DatastoreQuery {
	this := DatastoreQuery{}
	return &this
}

// NewDatastoreQueryWithDefaults instantiates a new DatastoreQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatastoreQueryWithDefaults() *DatastoreQuery {
	this := DatastoreQuery{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DatastoreQuery) GetCount() bool {
	if o == nil || IsNil(o.Count) {
		var ret bool
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreQuery) GetCountOk() (*bool, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DatastoreQuery) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given bool and assigns it to the Count field.
func (o *DatastoreQuery) SetCount(v bool) {
	o.Count = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *DatastoreQuery) GetFilters() []DatastoreFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []DatastoreFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreQuery) GetFiltersOk() ([]DatastoreFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *DatastoreQuery) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []DatastoreFilter and assigns it to the Filters field.
func (o *DatastoreQuery) SetFilters(v []DatastoreFilter) {
	o.Filters = v
}

// GetJsonAfter returns the JsonAfter field value if set, zero value otherwise.
func (o *DatastoreQuery) GetJsonAfter() string {
	if o == nil || IsNil(o.JsonAfter) {
		var ret string
		return ret
	}
	return *o.JsonAfter
}

// GetJsonAfterOk returns a tuple with the JsonAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreQuery) GetJsonAfterOk() (*string, bool) {
	if o == nil || IsNil(o.JsonAfter) {
		return nil, false
	}
	return o.JsonAfter, true
}

// HasJsonAfter returns a boolean if a field has been set.
func (o *DatastoreQuery) HasJsonAfter() bool {
	if o != nil && !IsNil(o.JsonAfter) {
		return true
	}

	return false
}

// SetJsonAfter gets a reference to the given string and assigns it to the JsonAfter field.
func (o *DatastoreQuery) SetJsonAfter(v string) {
	o.JsonAfter = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *DatastoreQuery) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreQuery) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *DatastoreQuery) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *DatastoreQuery) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderBys returns the OrderBys field value if set, zero value otherwise.
func (o *DatastoreQuery) GetOrderBys() []DatastoreOrderBy {
	if o == nil || IsNil(o.OrderBys) {
		var ret []DatastoreOrderBy
		return ret
	}
	return o.OrderBys
}

// GetOrderBysOk returns a tuple with the OrderBys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreQuery) GetOrderBysOk() ([]DatastoreOrderBy, bool) {
	if o == nil || IsNil(o.OrderBys) {
		return nil, false
	}
	return o.OrderBys, true
}

// HasOrderBys returns a boolean if a field has been set.
func (o *DatastoreQuery) HasOrderBys() bool {
	if o != nil && !IsNil(o.OrderBys) {
		return true
	}

	return false
}

// SetOrderBys gets a reference to the given []DatastoreOrderBy and assigns it to the OrderBys field.
func (o *DatastoreQuery) SetOrderBys(v []DatastoreOrderBy) {
	o.OrderBys = v
}

func (o DatastoreQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatastoreQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.JsonAfter) {
		toSerialize["jsonAfter"] = o.JsonAfter
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.OrderBys) {
		toSerialize["orderBys"] = o.OrderBys
	}
	return toSerialize, nil
}

type NullableDatastoreQuery struct {
	value *DatastoreQuery
	isSet bool
}

func (v NullableDatastoreQuery) Get() *DatastoreQuery {
	return v.value
}

func (v *NullableDatastoreQuery) Set(val *DatastoreQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableDatastoreQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableDatastoreQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatastoreQuery(val *DatastoreQuery) *NullableDatastoreQuery {
	return &NullableDatastoreQuery{value: val, isSet: true}
}

func (v NullableDatastoreQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatastoreQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


