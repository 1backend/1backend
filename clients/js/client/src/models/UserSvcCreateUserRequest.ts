/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { UserSvcUser } from './UserSvcUser';
import {
    UserSvcUserFromJSON,
    UserSvcUserFromJSONTyped,
    UserSvcUserToJSON,
    UserSvcUserToJSONTyped,
} from './UserSvcUser';

/**
 * 
 * @export
 * @interface UserSvcCreateUserRequest
 */
export interface UserSvcCreateUserRequest {
    /**
     * 
     * @type {string}
     * @memberof UserSvcCreateUserRequest
     */
    password?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof UserSvcCreateUserRequest
     */
    roleIds?: Array<string>;
    /**
     * 
     * @type {UserSvcUser}
     * @memberof UserSvcCreateUserRequest
     */
    user?: UserSvcUser;
}

/**
 * Check if a given object implements the UserSvcCreateUserRequest interface.
 */
export function instanceOfUserSvcCreateUserRequest(value: object): value is UserSvcCreateUserRequest {
    return true;
}

export function UserSvcCreateUserRequestFromJSON(json: any): UserSvcCreateUserRequest {
    return UserSvcCreateUserRequestFromJSONTyped(json, false);
}

export function UserSvcCreateUserRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcCreateUserRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'password': json['password'] == null ? undefined : json['password'],
        'roleIds': json['roleIds'] == null ? undefined : json['roleIds'],
        'user': json['user'] == null ? undefined : UserSvcUserFromJSON(json['user']),
    };
}

export function UserSvcCreateUserRequestToJSON(json: any): UserSvcCreateUserRequest {
    return UserSvcCreateUserRequestToJSONTyped(json, false);
}

export function UserSvcCreateUserRequestToJSONTyped(value?: UserSvcCreateUserRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'password': value['password'],
        'roleIds': value['roleIds'],
        'user': UserSvcUserToJSON(value['user']),
    };
}

