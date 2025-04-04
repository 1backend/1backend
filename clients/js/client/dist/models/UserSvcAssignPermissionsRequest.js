/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcPermissionLinkFromJSON, UserSvcPermissionLinkToJSON, } from './UserSvcPermissionLink';
/**
 * Check if a given object implements the UserSvcAssignPermissionsRequest interface.
 */
export function instanceOfUserSvcAssignPermissionsRequest(value) {
    return true;
}
export function UserSvcAssignPermissionsRequestFromJSON(json) {
    return UserSvcAssignPermissionsRequestFromJSONTyped(json, false);
}
export function UserSvcAssignPermissionsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'permissionLinks': json['permissionLinks'] == null ? undefined : (json['permissionLinks'].map(UserSvcPermissionLinkFromJSON)),
    };
}
export function UserSvcAssignPermissionsRequestToJSON(json) {
    return UserSvcAssignPermissionsRequestToJSONTyped(json, false);
}
export function UserSvcAssignPermissionsRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'permissionLinks': value['permissionLinks'] == null ? undefined : (value['permissionLinks'].map(UserSvcPermissionLinkToJSON)),
    };
}
