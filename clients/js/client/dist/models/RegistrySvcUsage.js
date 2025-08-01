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
/**
 * Check if a given object implements the RegistrySvcUsage interface.
 */
export function instanceOfRegistrySvcUsage(value) {
    return true;
}
export function RegistrySvcUsageFromJSON(json) {
    return RegistrySvcUsageFromJSONTyped(json, false);
}
export function RegistrySvcUsageFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'percent': json['percent'] == null ? undefined : json['percent'],
        'total': json['total'] == null ? undefined : json['total'],
        'used': json['used'] == null ? undefined : json['used'],
    };
}
export function RegistrySvcUsageToJSON(json) {
    return RegistrySvcUsageToJSONTyped(json, false);
}
export function RegistrySvcUsageToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'percent': value['percent'],
        'total': value['total'],
        'used': value['used'],
    };
}
