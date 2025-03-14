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
import type { UserSvcRole } from './UserSvcRole';
import {
    UserSvcRoleFromJSON,
    UserSvcRoleFromJSONTyped,
    UserSvcRoleToJSON,
    UserSvcRoleToJSONTyped,
} from './UserSvcRole';

/**
 * 
 * @export
 * @interface UserSvcCreateRoleResponse
 */
export interface UserSvcCreateRoleResponse {
    /**
     * 
     * @type {UserSvcRole}
     * @memberof UserSvcCreateRoleResponse
     */
    role?: UserSvcRole;
}

/**
 * Check if a given object implements the UserSvcCreateRoleResponse interface.
 */
export function instanceOfUserSvcCreateRoleResponse(value: object): value is UserSvcCreateRoleResponse {
    return true;
}

export function UserSvcCreateRoleResponseFromJSON(json: any): UserSvcCreateRoleResponse {
    return UserSvcCreateRoleResponseFromJSONTyped(json, false);
}

export function UserSvcCreateRoleResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcCreateRoleResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'role': json['role'] == null ? undefined : UserSvcRoleFromJSON(json['role']),
    };
}

export function UserSvcCreateRoleResponseToJSON(json: any): UserSvcCreateRoleResponse {
    return UserSvcCreateRoleResponseToJSONTyped(json, false);
}

export function UserSvcCreateRoleResponseToJSONTyped(value?: UserSvcCreateRoleResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'role': UserSvcRoleToJSON(value['role']),
    };
}

