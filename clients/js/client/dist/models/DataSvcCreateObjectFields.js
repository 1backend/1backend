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
/**
 * Check if a given object implements the DataSvcCreateObjectFields interface.
 */
export function instanceOfDataSvcCreateObjectFields(value) {
    if (!('data' in value) || value['data'] === undefined)
        return false;
    if (!('table' in value) || value['table'] === undefined)
        return false;
    return true;
}
export function DataSvcCreateObjectFieldsFromJSON(json) {
    return DataSvcCreateObjectFieldsFromJSONTyped(json, false);
}
export function DataSvcCreateObjectFieldsFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'authors': json['authors'] == null ? undefined : json['authors'],
        'data': json['data'],
        'deleters': json['deleters'] == null ? undefined : json['deleters'],
        'id': json['id'] == null ? undefined : json['id'],
        'readers': json['readers'] == null ? undefined : json['readers'],
        'table': json['table'],
        'writers': json['writers'] == null ? undefined : json['writers'],
    };
}
export function DataSvcCreateObjectFieldsToJSON(json) {
    return DataSvcCreateObjectFieldsToJSONTyped(json, false);
}
export function DataSvcCreateObjectFieldsToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'authors': value['authors'],
        'data': value['data'],
        'deleters': value['deleters'],
        'id': value['id'],
        'readers': value['readers'],
        'table': value['table'],
        'writers': value['writers'],
    };
}
