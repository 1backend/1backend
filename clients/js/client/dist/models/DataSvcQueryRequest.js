/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { DatastoreQueryFromJSON, DatastoreQueryToJSON, } from './DatastoreQuery';
/**
 * Check if a given object implements the DataSvcQueryRequest interface.
 */
export function instanceOfDataSvcQueryRequest(value) {
    return true;
}
export function DataSvcQueryRequestFromJSON(json) {
    return DataSvcQueryRequestFromJSONTyped(json, false);
}
export function DataSvcQueryRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'query': json['query'] == null ? undefined : DatastoreQueryFromJSON(json['query']),
        'readers': json['readers'] == null ? undefined : json['readers'],
        'table': json['table'] == null ? undefined : json['table'],
    };
}
export function DataSvcQueryRequestToJSON(json) {
    return DataSvcQueryRequestToJSONTyped(json, false);
}
export function DataSvcQueryRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'query': DatastoreQueryToJSON(value['query']),
        'readers': value['readers'],
        'table': value['table'],
    };
}
