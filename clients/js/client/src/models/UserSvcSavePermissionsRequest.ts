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
import type { UserSvcPermission } from './UserSvcPermission';
import {
    UserSvcPermissionFromJSON,
    UserSvcPermissionFromJSONTyped,
    UserSvcPermissionToJSON,
    UserSvcPermissionToJSONTyped,
} from './UserSvcPermission';

/**
 * 
 * @export
 * @interface UserSvcSavePermissionsRequest
 */
export interface UserSvcSavePermissionsRequest {
    /**
     * 
     * @type {Array<UserSvcPermission>}
     * @memberof UserSvcSavePermissionsRequest
     */
    permissions?: Array<UserSvcPermission>;
}

/**
 * Check if a given object implements the UserSvcSavePermissionsRequest interface.
 */
export function instanceOfUserSvcSavePermissionsRequest(value: object): value is UserSvcSavePermissionsRequest {
    return true;
}

export function UserSvcSavePermissionsRequestFromJSON(json: any): UserSvcSavePermissionsRequest {
    return UserSvcSavePermissionsRequestFromJSONTyped(json, false);
}

export function UserSvcSavePermissionsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSavePermissionsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'permissions': json['permissions'] == null ? undefined : ((json['permissions'] as Array<any>).map(UserSvcPermissionFromJSON)),
    };
}

export function UserSvcSavePermissionsRequestToJSON(json: any): UserSvcSavePermissionsRequest {
    return UserSvcSavePermissionsRequestToJSONTyped(json, false);
}

export function UserSvcSavePermissionsRequestToJSONTyped(value?: UserSvcSavePermissionsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'permissions': value['permissions'] == null ? undefined : ((value['permissions'] as Array<any>).map(UserSvcPermissionToJSON)),
    };
}

