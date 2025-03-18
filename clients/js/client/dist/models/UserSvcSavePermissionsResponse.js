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
import { UserSvcPermissionFromJSON, UserSvcPermissionToJSON, } from './UserSvcPermission';
/**
 * Check if a given object implements the UserSvcSavePermissionsResponse interface.
 */
export function instanceOfUserSvcSavePermissionsResponse(value) {
    return true;
}
export function UserSvcSavePermissionsResponseFromJSON(json) {
    return UserSvcSavePermissionsResponseFromJSONTyped(json, false);
}
export function UserSvcSavePermissionsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'permissions': json['permissions'] == null ? undefined : (json['permissions'].map(UserSvcPermissionFromJSON)),
    };
}
export function UserSvcSavePermissionsResponseToJSON(json) {
    return UserSvcSavePermissionsResponseToJSONTyped(json, false);
}
export function UserSvcSavePermissionsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'permissions': value['permissions'] == null ? undefined : (value['permissions'].map(UserSvcPermissionToJSON)),
    };
}
