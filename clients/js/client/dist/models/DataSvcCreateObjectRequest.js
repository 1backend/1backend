/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { DataSvcCreateObjectFieldsFromJSON, DataSvcCreateObjectFieldsToJSON, } from './DataSvcCreateObjectFields';
/**
 * Check if a given object implements the DataSvcCreateObjectRequest interface.
 */
export function instanceOfDataSvcCreateObjectRequest(value) {
    return true;
}
export function DataSvcCreateObjectRequestFromJSON(json) {
    return DataSvcCreateObjectRequestFromJSONTyped(json, false);
}
export function DataSvcCreateObjectRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'object': json['object'] == null ? undefined : DataSvcCreateObjectFieldsFromJSON(json['object']),
    };
}
export function DataSvcCreateObjectRequestToJSON(json) {
    return DataSvcCreateObjectRequestToJSONTyped(json, false);
}
export function DataSvcCreateObjectRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'object': DataSvcCreateObjectFieldsToJSON(value['object']),
    };
}
