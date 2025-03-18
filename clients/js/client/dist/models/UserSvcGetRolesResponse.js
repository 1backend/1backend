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
import { UserSvcRoleFromJSON, UserSvcRoleToJSON, } from './UserSvcRole';
/**
 * Check if a given object implements the UserSvcGetRolesResponse interface.
 */
export function instanceOfUserSvcGetRolesResponse(value) {
    return true;
}
export function UserSvcGetRolesResponseFromJSON(json) {
    return UserSvcGetRolesResponseFromJSONTyped(json, false);
}
export function UserSvcGetRolesResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'roles': json['roles'] == null ? undefined : (json['roles'].map(UserSvcRoleFromJSON)),
    };
}
export function UserSvcGetRolesResponseToJSON(json) {
    return UserSvcGetRolesResponseToJSONTyped(json, false);
}
export function UserSvcGetRolesResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'roles': value['roles'] == null ? undefined : (value['roles'].map(UserSvcRoleToJSON)),
    };
}
