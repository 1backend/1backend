/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcAuthToken interface.
 */
export function instanceOfUserSvcAuthToken(value) {
    if (!('token' in value) || value['token'] === undefined)
        return false;
    if (!('userId' in value) || value['userId'] === undefined)
        return false;
    return true;
}
export function UserSvcAuthTokenFromJSON(json) {
    return UserSvcAuthTokenFromJSONTyped(json, false);
}
export function UserSvcAuthTokenFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'active': json['active'] == null ? undefined : json['active'],
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'id': json['id'] == null ? undefined : json['id'],
        'token': json['token'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
        'userId': json['userId'],
    };
}
export function UserSvcAuthTokenToJSON(json) {
    return UserSvcAuthTokenToJSONTyped(json, false);
}
export function UserSvcAuthTokenToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'active': value['active'],
        'createdAt': value['createdAt'],
        'deletedAt': value['deletedAt'],
        'id': value['id'],
        'token': value['token'],
        'updatedAt': value['updatedAt'],
        'userId': value['userId'],
    };
}
