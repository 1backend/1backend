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
 * @interface UserSvcHasPermissionResponse
 */
export interface UserSvcHasPermissionResponse {
    /**
     * 
     * @type {boolean}
     * @memberof UserSvcHasPermissionResponse
     */
    authorized?: boolean;
    /**
     * 
     * @type {UserSvcUser}
     * @memberof UserSvcHasPermissionResponse
     */
    user?: UserSvcUser;
}

/**
 * Check if a given object implements the UserSvcHasPermissionResponse interface.
 */
export function instanceOfUserSvcHasPermissionResponse(value: object): value is UserSvcHasPermissionResponse {
    return true;
}

export function UserSvcHasPermissionResponseFromJSON(json: any): UserSvcHasPermissionResponse {
    return UserSvcHasPermissionResponseFromJSONTyped(json, false);
}

export function UserSvcHasPermissionResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcHasPermissionResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'authorized': json['authorized'] == null ? undefined : json['authorized'],
        'user': json['user'] == null ? undefined : UserSvcUserFromJSON(json['user']),
    };
}

export function UserSvcHasPermissionResponseToJSON(json: any): UserSvcHasPermissionResponse {
    return UserSvcHasPermissionResponseToJSONTyped(json, false);
}

export function UserSvcHasPermissionResponseToJSONTyped(value?: UserSvcHasPermissionResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'authorized': value['authorized'],
        'user': UserSvcUserToJSON(value['user']),
    };
}

