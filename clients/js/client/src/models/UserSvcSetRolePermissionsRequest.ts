/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
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
 * @interface UserSvcSetRolePermissionsRequest
 */
export interface UserSvcSetRolePermissionsRequest {
    /**
     * 
     * @type {Array<string>}
     * @memberof UserSvcSetRolePermissionsRequest
     */
    permissionIds?: Array<string>;
}

/**
 * Check if a given object implements the UserSvcSetRolePermissionsRequest interface.
 */
export function instanceOfUserSvcSetRolePermissionsRequest(value: object): value is UserSvcSetRolePermissionsRequest {
    return true;
}

export function UserSvcSetRolePermissionsRequestFromJSON(json: any): UserSvcSetRolePermissionsRequest {
    return UserSvcSetRolePermissionsRequestFromJSONTyped(json, false);
}

export function UserSvcSetRolePermissionsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSetRolePermissionsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'permissionIds': json['permissionIds'] == null ? undefined : json['permissionIds'],
    };
}

export function UserSvcSetRolePermissionsRequestToJSON(json: any): UserSvcSetRolePermissionsRequest {
    return UserSvcSetRolePermissionsRequestToJSONTyped(json, false);
}

export function UserSvcSetRolePermissionsRequestToJSONTyped(value?: UserSvcSetRolePermissionsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'permissionIds': value['permissionIds'],
    };
}

