/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
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
import type { UserSvcAuthToken } from './UserSvcAuthToken';
import {
    UserSvcAuthTokenFromJSON,
    UserSvcAuthTokenFromJSONTyped,
    UserSvcAuthTokenToJSON,
    UserSvcAuthTokenToJSONTyped,
} from './UserSvcAuthToken';

/**
 * 
 * @export
 * @interface UserSvcLoginResponse
 */
export interface UserSvcLoginResponse {
    /**
     * 
     * @type {UserSvcAuthToken}
     * @memberof UserSvcLoginResponse
     */
    token?: UserSvcAuthToken;
}

/**
 * Check if a given object implements the UserSvcLoginResponse interface.
 */
export function instanceOfUserSvcLoginResponse(value: object): value is UserSvcLoginResponse {
    return true;
}

export function UserSvcLoginResponseFromJSON(json: any): UserSvcLoginResponse {
    return UserSvcLoginResponseFromJSONTyped(json, false);
}

export function UserSvcLoginResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcLoginResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'token': json['token'] == null ? undefined : UserSvcAuthTokenFromJSON(json['token']),
    };
}

export function UserSvcLoginResponseToJSON(json: any): UserSvcLoginResponse {
    return UserSvcLoginResponseToJSONTyped(json, false);
}

export function UserSvcLoginResponseToJSONTyped(value?: UserSvcLoginResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'token': UserSvcAuthTokenToJSON(value['token']),
    };
}

