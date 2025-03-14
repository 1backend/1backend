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
/**
 * Check if a given object implements the ModelSvcModelStatus interface.
 */
export function instanceOfModelSvcModelStatus(value) {
    if (!('address' in value) || value['address'] === undefined)
        return false;
    if (!('assetsReady' in value) || value['assetsReady'] === undefined)
        return false;
    if (!('running' in value) || value['running'] === undefined)
        return false;
    return true;
}
export function ModelSvcModelStatusFromJSON(json) {
    return ModelSvcModelStatusFromJSONTyped(json, false);
}
export function ModelSvcModelStatusFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'address': json['address'],
        'assetsReady': json['assetsReady'],
        'running': json['running'],
    };
}
export function ModelSvcModelStatusToJSON(json) {
    return ModelSvcModelStatusToJSONTyped(json, false);
}
export function ModelSvcModelStatusToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'address': value['address'],
        'assetsReady': value['assetsReady'],
        'running': value['running'],
    };
}
