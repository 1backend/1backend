/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
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
 * @interface UserSvcChangePasswordRequest
 */
export interface UserSvcChangePasswordRequest {
    /**
     * 
     * @type {string}
     * @memberof UserSvcChangePasswordRequest
     */
    currentPassword: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcChangePasswordRequest
     */
    newPassword: string;
}

/**
 * Check if a given object implements the UserSvcChangePasswordRequest interface.
 */
export function instanceOfUserSvcChangePasswordRequest(value: object): value is UserSvcChangePasswordRequest {
    if (!('currentPassword' in value) || value['currentPassword'] === undefined) return false;
    if (!('newPassword' in value) || value['newPassword'] === undefined) return false;
    return true;
}

export function UserSvcChangePasswordRequestFromJSON(json: any): UserSvcChangePasswordRequest {
    return UserSvcChangePasswordRequestFromJSONTyped(json, false);
}

export function UserSvcChangePasswordRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcChangePasswordRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'currentPassword': json['currentPassword'],
        'newPassword': json['newPassword'],
    };
}

export function UserSvcChangePasswordRequestToJSON(json: any): UserSvcChangePasswordRequest {
    return UserSvcChangePasswordRequestToJSONTyped(json, false);
}

export function UserSvcChangePasswordRequestToJSONTyped(value?: UserSvcChangePasswordRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'currentPassword': value['currentPassword'],
        'newPassword': value['newPassword'],
    };
}

