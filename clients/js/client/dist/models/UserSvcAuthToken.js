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
 * Check if a given object implements the UserSvcAuthToken interface.
 */
export function instanceOfUserSvcAuthToken(value) {
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
        'token': json['token'] == null ? undefined : json['token'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
        'userId': json['userId'] == null ? undefined : json['userId'],
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
