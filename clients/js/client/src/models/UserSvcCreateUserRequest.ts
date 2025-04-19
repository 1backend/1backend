/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { UserSvcContact } from './UserSvcContact';
import {
    UserSvcContactFromJSON,
    UserSvcContactFromJSONTyped,
    UserSvcContactToJSON,
    UserSvcContactToJSONTyped,
} from './UserSvcContact';
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
     * @type {Array<UserSvcContact>}
     * @memberof UserSvcCreateUserRequest
     */
    contacts?: Array<UserSvcContact>;
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
        
        'contacts': json['contacts'] == null ? undefined : ((json['contacts'] as Array<any>).map(UserSvcContactFromJSON)),
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
        
        'contacts': value['contacts'] == null ? undefined : ((value['contacts'] as Array<any>).map(UserSvcContactToJSON)),
        'password': value['password'],
        'roleIds': value['roleIds'],
        'user': UserSvcUserToJSON(value['user']),
    };
}

