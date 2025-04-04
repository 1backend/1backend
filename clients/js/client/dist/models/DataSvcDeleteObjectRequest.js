/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { DatastoreFilterFromJSON, DatastoreFilterToJSON, } from './DatastoreFilter';
/**
 * Check if a given object implements the DataSvcDeleteObjectRequest interface.
 */
export function instanceOfDataSvcDeleteObjectRequest(value) {
    return true;
}
export function DataSvcDeleteObjectRequestFromJSON(json) {
    return DataSvcDeleteObjectRequestFromJSONTyped(json, false);
}
export function DataSvcDeleteObjectRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'filters': json['filters'] == null ? undefined : (json['filters'].map(DatastoreFilterFromJSON)),
        'table': json['table'] == null ? undefined : json['table'],
    };
}
export function DataSvcDeleteObjectRequestToJSON(json) {
    return DataSvcDeleteObjectRequestToJSONTyped(json, false);
}
export function DataSvcDeleteObjectRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'filters': value['filters'] == null ? undefined : (value['filters'].map(DatastoreFilterToJSON)),
        'table': value['table'],
    };
}
