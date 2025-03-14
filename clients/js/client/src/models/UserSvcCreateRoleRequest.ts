/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
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
 * @interface UserSvcCreateRoleRequest
 */
export interface UserSvcCreateRoleRequest {
    /**
     * 
     * @type {string}
     * @memberof UserSvcCreateRoleRequest
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcCreateRoleRequest
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcCreateRoleRequest
     */
    name?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof UserSvcCreateRoleRequest
     */
    permissionIds?: Array<string>;
}

/**
 * Check if a given object implements the UserSvcCreateRoleRequest interface.
 */
export function instanceOfUserSvcCreateRoleRequest(value: object): value is UserSvcCreateRoleRequest {
    if (!('id' in value) || value['id'] === undefined) return false;
    return true;
}

export function UserSvcCreateRoleRequestFromJSON(json: any): UserSvcCreateRoleRequest {
    return UserSvcCreateRoleRequestFromJSONTyped(json, false);
}

export function UserSvcCreateRoleRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcCreateRoleRequest {
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

export function UserSvcCreateRoleRequestToJSON(json: any): UserSvcCreateRoleRequest {
    return UserSvcCreateRoleRequestToJSONTyped(json, false);
}

export function UserSvcCreateRoleRequestToJSONTyped(value?: UserSvcCreateRoleRequest | null, ignoreDiscriminator: boolean = false): any {
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

