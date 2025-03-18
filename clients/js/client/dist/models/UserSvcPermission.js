/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcPermission interface.
 */
export function instanceOfUserSvcPermission(value) {
    return true;
}
export function UserSvcPermissionFromJSON(json) {
    return UserSvcPermissionFromJSONTyped(json, false);
}
export function UserSvcPermissionFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'description': json['description'] == null ? undefined : json['description'],
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'ownerId': json['ownerId'] == null ? undefined : json['ownerId'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
    };
}
export function UserSvcPermissionToJSON(json) {
    return UserSvcPermissionToJSONTyped(json, false);
}
export function UserSvcPermissionToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'createdAt': value['createdAt'],
        'description': value['description'],
        'id': value['id'],
        'name': value['name'],
        'ownerId': value['ownerId'],
        'updatedAt': value['updatedAt'],
    };
}
