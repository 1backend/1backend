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
 * Check if a given object implements the UserSvcCreateRoleRequest interface.
 */
export function instanceOfUserSvcCreateRoleRequest(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    return true;
}
export function UserSvcCreateRoleRequestFromJSON(json) {
    return UserSvcCreateRoleRequestFromJSONTyped(json, false);
}
export function UserSvcCreateRoleRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'description': json['description'] == null ? undefined : json['description'],
        'id': json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'permissionIds': json['permissionIds'] == null ? undefined : json['permissionIds'],
    };
}
export function UserSvcCreateRoleRequestToJSON(json) {
    return UserSvcCreateRoleRequestToJSONTyped(json, false);
}
export function UserSvcCreateRoleRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'description': value['description'],
        'id': value['id'],
        'name': value['name'],
        'permissionIds': value['permissionIds'],
    };
}
