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
 * @interface UserSvcReadUserByTokenResponse
 */
export interface UserSvcReadUserByTokenResponse {
    /**
     * 
     * @type {string}
     * @memberof UserSvcReadUserByTokenResponse
     */
    activeOrganizationId?: string;
    /**
     * 
     * @type {Array<UserSvcOrganization>}
     * @memberof UserSvcReadUserByTokenResponse
     */
    organizations?: Array<UserSvcOrganization>;
    /**
     * 
     * @type {UserSvcUser}
     * @memberof UserSvcReadUserByTokenResponse
     */
    user?: UserSvcUser;
}

/**
 * Check if a given object implements the UserSvcReadUserByTokenResponse interface.
 */
export function instanceOfUserSvcReadUserByTokenResponse(value: object): value is UserSvcReadUserByTokenResponse {
    return true;
}

export function UserSvcReadUserByTokenResponseFromJSON(json: any): UserSvcReadUserByTokenResponse {
    return UserSvcReadUserByTokenResponseFromJSONTyped(json, false);
}

export function UserSvcReadUserByTokenResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcReadUserByTokenResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'activeOrganizationId': json['activeOrganizationId'] == null ? undefined : json['activeOrganizationId'],
        'organizations': json['organizations'] == null ? undefined : ((json['organizations'] as Array<any>).map(UserSvcOrganizationFromJSON)),
        'user': json['user'] == null ? undefined : UserSvcUserFromJSON(json['user']),
    };
}

export function UserSvcReadUserByTokenResponseToJSON(json: any): UserSvcReadUserByTokenResponse {
    return UserSvcReadUserByTokenResponseToJSONTyped(json, false);
}

export function UserSvcReadUserByTokenResponseToJSONTyped(value?: UserSvcReadUserByTokenResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'activeOrganizationId': value['activeOrganizationId'],
        'organizations': value['organizations'] == null ? undefined : ((value['organizations'] as Array<any>).map(UserSvcOrganizationToJSON)),
        'user': UserSvcUserToJSON(value['user']),
    };
}

