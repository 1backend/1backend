/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { DataSvcObjectFromJSON, DataSvcObjectToJSON, } from './DataSvcObject';
/**
 * Check if a given object implements the DataSvcUpsertObjectResponse interface.
 */
export function instanceOfDataSvcUpsertObjectResponse(value) {
    return true;
}
export function DataSvcUpsertObjectResponseFromJSON(json) {
    return DataSvcUpsertObjectResponseFromJSONTyped(json, false);
}
export function DataSvcUpsertObjectResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'object': json['object'] == null ? undefined : DataSvcObjectFromJSON(json['object']),
    };
}
export function DataSvcUpsertObjectResponseToJSON(json) {
    return DataSvcUpsertObjectResponseToJSONTyped(json, false);
}
export function DataSvcUpsertObjectResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'object': DataSvcObjectToJSON(value['object']),
    };
}
