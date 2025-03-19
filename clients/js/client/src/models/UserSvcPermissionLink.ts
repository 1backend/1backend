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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface UserSvcPermissionLink
 */
export interface UserSvcPermissionLink {
    /**
     * 
     * @type {string}
     * @memberof UserSvcPermissionLink
     */
    permissionId?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcPermissionLink
     */
    roleId?: string;
}

/**
 * Check if a given object implements the UserSvcPermissionLink interface.
 */
export function instanceOfUserSvcPermissionLink(value: object): value is UserSvcPermissionLink {
    return true;
}

export function UserSvcPermissionLinkFromJSON(json: any): UserSvcPermissionLink {
    return UserSvcPermissionLinkFromJSONTyped(json, false);
}

export function UserSvcPermissionLinkFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcPermissionLink {
    if (json == null) {
        return json;
    }
    return {
        
        'permissionId': json['permissionId'] == null ? undefined : json['permissionId'],
        'roleId': json['roleId'] == null ? undefined : json['roleId'],
    };
}

export function UserSvcPermissionLinkToJSON(json: any): UserSvcPermissionLink {
    return UserSvcPermissionLinkToJSONTyped(json, false);
}

export function UserSvcPermissionLinkToJSONTyped(value?: UserSvcPermissionLink | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'permissionId': value['permissionId'],
        'roleId': value['roleId'],
    };
}

