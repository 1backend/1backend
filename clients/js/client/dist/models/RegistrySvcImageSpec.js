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
 * Check if a given object implements the RegistrySvcImageSpec interface.
 */
export function instanceOfRegistrySvcImageSpec(value) {
    if (!('internalPorts' in value) || value['internalPorts'] === undefined)
        return false;
    if (!('name' in value) || value['name'] === undefined)
        return false;
    return true;
}
export function RegistrySvcImageSpecFromJSON(json) {
    return RegistrySvcImageSpecFromJSONTyped(json, false);
}
export function RegistrySvcImageSpecFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'internalPorts': json['internalPorts'],
        'name': json['name'],
    };
}
export function RegistrySvcImageSpecToJSON(json) {
    return RegistrySvcImageSpecToJSONTyped(json, false);
}
export function RegistrySvcImageSpecToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'internalPorts': value['internalPorts'],
        'name': value['name'],
    };
}
