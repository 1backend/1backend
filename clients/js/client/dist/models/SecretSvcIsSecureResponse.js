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
 * Check if a given object implements the SecretSvcIsSecureResponse interface.
 */
export function instanceOfSecretSvcIsSecureResponse(value) {
    if (!('isSecure' in value) || value['isSecure'] === undefined)
        return false;
    return true;
}
export function SecretSvcIsSecureResponseFromJSON(json) {
    return SecretSvcIsSecureResponseFromJSONTyped(json, false);
}
export function SecretSvcIsSecureResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'isSecure': json['isSecure'],
    };
}
export function SecretSvcIsSecureResponseToJSON(json) {
    return SecretSvcIsSecureResponseToJSONTyped(json, false);
}
export function SecretSvcIsSecureResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'isSecure': value['isSecure'],
    };
}
