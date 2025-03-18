/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
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
 * Check if a given object implements the UserSvcCreateRoleResponse interface.
 */
export function instanceOfUserSvcCreateRoleResponse(value) {
    return true;
}
export function UserSvcCreateRoleResponseFromJSON(json) {
    return UserSvcCreateRoleResponseFromJSONTyped(json, false);
}
export function UserSvcCreateRoleResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'role': json['role'] == null ? undefined : UserSvcRoleFromJSON(json['role']),
    };
}
export function UserSvcCreateRoleResponseToJSON(json) {
    return UserSvcCreateRoleResponseToJSONTyped(json, false);
}
export function UserSvcCreateRoleResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'role': UserSvcRoleToJSON(value['role']),
    };
}
