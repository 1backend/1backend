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
 * Check if a given object implements the UserSvcLoginRequest interface.
 */
export function instanceOfUserSvcLoginRequest(value) {
    return true;
}
export function UserSvcLoginRequestFromJSON(json) {
    return UserSvcLoginRequestFromJSONTyped(json, false);
}
export function UserSvcLoginRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'contact': json['contact'] == null ? undefined : json['contact'],
        'password': json['password'] == null ? undefined : json['password'],
        'slug': json['slug'] == null ? undefined : json['slug'],
    };
}
export function UserSvcLoginRequestToJSON(json) {
    return UserSvcLoginRequestToJSONTyped(json, false);
}
export function UserSvcLoginRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'contact': value['contact'],
        'password': value['password'],
        'slug': value['slug'],
    };
}
