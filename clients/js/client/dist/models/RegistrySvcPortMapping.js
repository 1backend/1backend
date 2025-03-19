/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the RegistrySvcPortMapping interface.
 */
export function instanceOfRegistrySvcPortMapping(value) {
    if (!('host' in value) || value['host'] === undefined)
        return false;
    if (!('internal' in value) || value['internal'] === undefined)
        return false;
    return true;
}
export function RegistrySvcPortMappingFromJSON(json) {
    return RegistrySvcPortMappingFromJSONTyped(json, false);
}
export function RegistrySvcPortMappingFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'host': json['host'],
        'internal': json['internal'],
    };
}
export function RegistrySvcPortMappingToJSON(json) {
    return RegistrySvcPortMappingToJSONTyped(json, false);
}
export function RegistrySvcPortMappingToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'host': value['host'],
        'internal': value['internal'],
    };
}
