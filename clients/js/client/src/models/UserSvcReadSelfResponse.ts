/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { UserSvcOrganization } from './UserSvcOrganization';
import {
    UserSvcOrganizationFromJSON,
    UserSvcOrganizationFromJSONTyped,
    UserSvcOrganizationToJSON,
    UserSvcOrganizationToJSONTyped,
} from './UserSvcOrganization';
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
 * @interface UserSvcReadSelfResponse
 */
export interface UserSvcReadSelfResponse {
    /**
     * Active organization of the caller user, if it has any.
     * @type {string}
     * @memberof UserSvcReadSelfResponse
     */
    activeOrganizationId?: string;
    /**
     * Organizations of the caller user.
     * @type {Array<UserSvcOrganization>}
     * @memberof UserSvcReadSelfResponse
     */
    organizations?: Array<UserSvcOrganization>;
    /**
     * Roles the token has that made this request.
     * @type {Array<string>}
     * @memberof UserSvcReadSelfResponse
     */
    roles?: Array<string>;
    /**
     * The user who made the request.
     * @type {UserSvcUser}
     * @memberof UserSvcReadSelfResponse
     */
    user: UserSvcUser;
}

/**
 * Check if a given object implements the UserSvcReadSelfResponse interface.
 */
export function instanceOfUserSvcReadSelfResponse(value: object): value is UserSvcReadSelfResponse {
    if (!('user' in value) || value['user'] === undefined) return false;
    return true;
}

export function UserSvcReadSelfResponseFromJSON(json: any): UserSvcReadSelfResponse {
    return UserSvcReadSelfResponseFromJSONTyped(json, false);
}

export function UserSvcReadSelfResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcReadSelfResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'activeOrganizationId': json['activeOrganizationId'] == null ? undefined : json['activeOrganizationId'],
        'organizations': json['organizations'] == null ? undefined : ((json['organizations'] as Array<any>).map(UserSvcOrganizationFromJSON)),
        'roles': json['roles'] == null ? undefined : json['roles'],
        'user': UserSvcUserFromJSON(json['user']),
    };
}

export function UserSvcReadSelfResponseToJSON(json: any): UserSvcReadSelfResponse {
    return UserSvcReadSelfResponseToJSONTyped(json, false);
}

export function UserSvcReadSelfResponseToJSONTyped(value?: UserSvcReadSelfResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'activeOrganizationId': value['activeOrganizationId'],
        'organizations': value['organizations'] == null ? undefined : ((value['organizations'] as Array<any>).map(UserSvcOrganizationToJSON)),
        'roles': value['roles'],
        'user': UserSvcUserToJSON(value['user']),
    };
}

