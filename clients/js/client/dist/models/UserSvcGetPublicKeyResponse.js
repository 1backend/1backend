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
 * Check if a given object implements the UserSvcGetPublicKeyResponse interface.
 */
export function instanceOfUserSvcGetPublicKeyResponse(value) {
    if (!('publicKey' in value) || value['publicKey'] === undefined)
        return false;
    return true;
}
export function UserSvcGetPublicKeyResponseFromJSON(json) {
    return UserSvcGetPublicKeyResponseFromJSONTyped(json, false);
}
export function UserSvcGetPublicKeyResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'publicKey': json['publicKey'],
    };
}
export function UserSvcGetPublicKeyResponseToJSON(json) {
    return UserSvcGetPublicKeyResponseToJSONTyped(json, false);
}
export function UserSvcGetPublicKeyResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'publicKey': value['publicKey'],
    };
}
