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
 * @interface UserSvcRegisterResponse
 */
export interface UserSvcRegisterResponse {
    /**
     * 
     * @type {UserSvcAuthToken}
     * @memberof UserSvcRegisterResponse
     */
    token?: UserSvcAuthToken;
}

/**
 * Check if a given object implements the UserSvcRegisterResponse interface.
 */
export function instanceOfUserSvcRegisterResponse(value: object): value is UserSvcRegisterResponse {
    return true;
}

export function UserSvcRegisterResponseFromJSON(json: any): UserSvcRegisterResponse {
    return UserSvcRegisterResponseFromJSONTyped(json, false);
}

export function UserSvcRegisterResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcRegisterResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'token': json['token'] == null ? undefined : UserSvcAuthTokenFromJSON(json['token']),
    };
}

export function UserSvcRegisterResponseToJSON(json: any): UserSvcRegisterResponse {
    return UserSvcRegisterResponseToJSONTyped(json, false);
}

export function UserSvcRegisterResponseToJSONTyped(value?: UserSvcRegisterResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'token': UserSvcAuthTokenToJSON(value['token']),
    };
}

