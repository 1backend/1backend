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
/**
 * Check if a given object implements the DataSvcObject interface.
 */
export function instanceOfDataSvcObject(value) {
    if (!('data' in value) || value['data'] === undefined)
        return false;
    if (!('table' in value) || value['table'] === undefined)
        return false;
    return true;
}
export function DataSvcObjectFromJSON(json) {
    return DataSvcObjectFromJSONTyped(json, false);
}
export function DataSvcObjectFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'authors': json['authors'] == null ? undefined : json['authors'],
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'data': json['data'],
        'deleters': json['deleters'] == null ? undefined : json['deleters'],
        'id': json['id'] == null ? undefined : json['id'],
        'readers': json['readers'] == null ? undefined : json['readers'],
        'table': json['table'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
        'writers': json['writers'] == null ? undefined : json['writers'],
    };
}
export function DataSvcObjectToJSON(json) {
    return DataSvcObjectToJSONTyped(json, false);
}
export function DataSvcObjectToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'authors': value['authors'],
        'createdAt': value['createdAt'],
        'data': value['data'],
        'deleters': value['deleters'],
        'id': value['id'],
        'readers': value['readers'],
        'table': value['table'],
        'updatedAt': value['updatedAt'],
        'writers': value['writers'],
    };
}
