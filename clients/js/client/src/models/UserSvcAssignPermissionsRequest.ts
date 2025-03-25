/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { UserSvcPermissionLink } from './UserSvcPermissionLink';
import {
    UserSvcPermissionLinkFromJSON,
    UserSvcPermissionLinkFromJSONTyped,
    UserSvcPermissionLinkToJSON,
    UserSvcPermissionLinkToJSONTyped,
} from './UserSvcPermissionLink';

/**
 * 
 * @export
 * @interface UserSvcAssignPermissionsRequest
 */
export interface UserSvcAssignPermissionsRequest {
    /**
     * 
     * @type {Array<UserSvcPermissionLink>}
     * @memberof UserSvcAssignPermissionsRequest
     */
    permissionLinks?: Array<UserSvcPermissionLink>;
}

/**
 * Check if a given object implements the UserSvcAssignPermissionsRequest interface.
 */
export function instanceOfUserSvcAssignPermissionsRequest(value: object): value is UserSvcAssignPermissionsRequest {
    return true;
}

export function UserSvcAssignPermissionsRequestFromJSON(json: any): UserSvcAssignPermissionsRequest {
    return UserSvcAssignPermissionsRequestFromJSONTyped(json, false);
}

export function UserSvcAssignPermissionsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcAssignPermissionsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'permissionLinks': json['permissionLinks'] == null ? undefined : ((json['permissionLinks'] as Array<any>).map(UserSvcPermissionLinkFromJSON)),
    };
}

export function UserSvcAssignPermissionsRequestToJSON(json: any): UserSvcAssignPermissionsRequest {
    return UserSvcAssignPermissionsRequestToJSONTyped(json, false);
}

export function UserSvcAssignPermissionsRequestToJSONTyped(value?: UserSvcAssignPermissionsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'permissionLinks': value['permissionLinks'] == null ? undefined : ((value['permissionLinks'] as Array<any>).map(UserSvcPermissionLinkToJSON)),
    };
}

