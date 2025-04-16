/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
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
 * @interface UserSvcResetPasswordRequest
 */
export interface UserSvcResetPasswordRequest {
    /**
     * 
     * @type {string}
     * @memberof UserSvcResetPasswordRequest
     */
    newPassword: string;
}

/**
 * Check if a given object implements the UserSvcResetPasswordRequest interface.
 */
export function instanceOfUserSvcResetPasswordRequest(value: object): value is UserSvcResetPasswordRequest {
    if (!('newPassword' in value) || value['newPassword'] === undefined) return false;
    return true;
}

export function UserSvcResetPasswordRequestFromJSON(json: any): UserSvcResetPasswordRequest {
    return UserSvcResetPasswordRequestFromJSONTyped(json, false);
}

export function UserSvcResetPasswordRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcResetPasswordRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'newPassword': json['newPassword'],
    };
}

export function UserSvcResetPasswordRequestToJSON(json: any): UserSvcResetPasswordRequest {
    return UserSvcResetPasswordRequestToJSONTyped(json, false);
}

export function UserSvcResetPasswordRequestToJSONTyped(value?: UserSvcResetPasswordRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'newPassword': value['newPassword'],
    };
}

