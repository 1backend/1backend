/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { DatastoreFilterFromJSON, DatastoreFilterToJSON, } from './DatastoreFilter';
import { DatastoreOrderByFromJSON, DatastoreOrderByToJSON, } from './DatastoreOrderBy';
/**
 * Check if a given object implements the DatastoreQuery interface.
 */
export function instanceOfDatastoreQuery(value) {
    return true;
}
export function DatastoreQueryFromJSON(json) {
    return DatastoreQueryFromJSONTyped(json, false);
}
export function DatastoreQueryFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'count': json['count'] == null ? undefined : json['count'],
        'filters': json['filters'] == null ? undefined : (json['filters'].map(DatastoreFilterFromJSON)),
        'jsonAfter': json['jsonAfter'] == null ? undefined : json['jsonAfter'],
        'limit': json['limit'] == null ? undefined : json['limit'],
        'orderBys': json['orderBys'] == null ? undefined : (json['orderBys'].map(DatastoreOrderByFromJSON)),
    };
}
export function DatastoreQueryToJSON(json) {
    return DatastoreQueryToJSONTyped(json, false);
}
export function DatastoreQueryToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'count': value['count'],
        'filters': value['filters'] == null ? undefined : (value['filters'].map(DatastoreFilterToJSON)),
        'jsonAfter': value['jsonAfter'],
        'limit': value['limit'],
        'orderBys': value['orderBys'] == null ? undefined : (value['orderBys'].map(DatastoreOrderByToJSON)),
    };
}
