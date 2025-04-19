/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the RegistrySvcErrorResponse interface.
 */
export function instanceOfRegistrySvcErrorResponse(value) {
    return true;
}
export function RegistrySvcErrorResponseFromJSON(json) {
    return RegistrySvcErrorResponseFromJSONTyped(json, false);
}
export function RegistrySvcErrorResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'error': json['error'] == null ? undefined : json['error'],
    };
}
export function RegistrySvcErrorResponseToJSON(json) {
    return RegistrySvcErrorResponseToJSONTyped(json, false);
}
export function RegistrySvcErrorResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'error': value['error'],
    };
}
