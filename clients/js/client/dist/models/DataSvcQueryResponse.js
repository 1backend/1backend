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
import { DataSvcObjectFromJSON, DataSvcObjectToJSON, } from './DataSvcObject';
/**
 * Check if a given object implements the DataSvcQueryResponse interface.
 */
export function instanceOfDataSvcQueryResponse(value) {
    return true;
}
export function DataSvcQueryResponseFromJSON(json) {
    return DataSvcQueryResponseFromJSONTyped(json, false);
}
export function DataSvcQueryResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'objects': json['objects'] == null ? undefined : (json['objects'].map(DataSvcObjectFromJSON)),
    };
}
export function DataSvcQueryResponseToJSON(json) {
    return DataSvcQueryResponseToJSONTyped(json, false);
}
export function DataSvcQueryResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'objects': value['objects'] == null ? undefined : (value['objects'].map(DataSvcObjectToJSON)),
    };
}
