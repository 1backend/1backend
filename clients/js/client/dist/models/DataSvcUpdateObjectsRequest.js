/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { DatastoreFilterFromJSON, DatastoreFilterToJSON, } from './DatastoreFilter';
import { DataSvcObjectFromJSON, DataSvcObjectToJSON, } from './DataSvcObject';
/**
 * Check if a given object implements the DataSvcUpdateObjectsRequest interface.
 */
export function instanceOfDataSvcUpdateObjectsRequest(value) {
    return true;
}
export function DataSvcUpdateObjectsRequestFromJSON(json) {
    return DataSvcUpdateObjectsRequestFromJSONTyped(json, false);
}
export function DataSvcUpdateObjectsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'filters': json['filters'] == null ? undefined : (json['filters'].map(DatastoreFilterFromJSON)),
        'object': json['object'] == null ? undefined : DataSvcObjectFromJSON(json['object']),
        'table': json['table'] == null ? undefined : json['table'],
    };
}
export function DataSvcUpdateObjectsRequestToJSON(json) {
    return DataSvcUpdateObjectsRequestToJSONTyped(json, false);
}
export function DataSvcUpdateObjectsRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'filters': value['filters'] == null ? undefined : (value['filters'].map(DatastoreFilterToJSON)),
        'object': DataSvcObjectToJSON(value['object']),
        'table': value['table'],
    };
}
