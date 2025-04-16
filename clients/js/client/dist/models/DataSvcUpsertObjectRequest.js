/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { DataSvcCreateObjectFieldsFromJSON, DataSvcCreateObjectFieldsToJSON, } from './DataSvcCreateObjectFields';
/**
 * Check if a given object implements the DataSvcUpsertObjectRequest interface.
 */
export function instanceOfDataSvcUpsertObjectRequest(value) {
    return true;
}
export function DataSvcUpsertObjectRequestFromJSON(json) {
    return DataSvcUpsertObjectRequestFromJSONTyped(json, false);
}
export function DataSvcUpsertObjectRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'object': json['object'] == null ? undefined : DataSvcCreateObjectFieldsFromJSON(json['object']),
    };
}
export function DataSvcUpsertObjectRequestToJSON(json) {
    return DataSvcUpsertObjectRequestToJSONTyped(json, false);
}
export function DataSvcUpsertObjectRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'object': DataSvcCreateObjectFieldsToJSON(value['object']),
    };
}
