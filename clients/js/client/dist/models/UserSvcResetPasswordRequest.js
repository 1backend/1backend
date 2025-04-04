/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcResetPasswordRequest interface.
 */
export function instanceOfUserSvcResetPasswordRequest(value) {
    return true;
}
export function UserSvcResetPasswordRequestFromJSON(json) {
    return UserSvcResetPasswordRequestFromJSONTyped(json, false);
}
export function UserSvcResetPasswordRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'newPassword': json['newPassword'] == null ? undefined : json['newPassword'],
        'slug': json['slug'] == null ? undefined : json['slug'],
    };
}
export function UserSvcResetPasswordRequestToJSON(json) {
    return UserSvcResetPasswordRequestToJSONTyped(json, false);
}
export function UserSvcResetPasswordRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'newPassword': value['newPassword'],
        'slug': value['slug'],
    };
}
