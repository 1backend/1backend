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
 * Check if a given object implements the UserSvcChangePasswordRequest interface.
 */
export function instanceOfUserSvcChangePasswordRequest(value) {
    return true;
}
export function UserSvcChangePasswordRequestFromJSON(json) {
    return UserSvcChangePasswordRequestFromJSONTyped(json, false);
}
export function UserSvcChangePasswordRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'currentPassword': json['currentPassword'] == null ? undefined : json['currentPassword'],
        'newPassword': json['newPassword'] == null ? undefined : json['newPassword'],
        'slug': json['slug'] == null ? undefined : json['slug'],
    };
}
export function UserSvcChangePasswordRequestToJSON(json) {
    return UserSvcChangePasswordRequestToJSONTyped(json, false);
}
export function UserSvcChangePasswordRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'currentPassword': value['currentPassword'],
        'newPassword': value['newPassword'],
        'slug': value['slug'],
    };
}
