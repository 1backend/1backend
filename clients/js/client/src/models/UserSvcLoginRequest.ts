/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
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
 * @interface UserSvcLoginRequest
 */
export interface UserSvcLoginRequest {
    /**
     * 
     * @type {string}
     * @memberof UserSvcLoginRequest
     */
    contact?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcLoginRequest
     */
    password?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcLoginRequest
     */
    slug?: string;
}

/**
 * Check if a given object implements the UserSvcLoginRequest interface.
 */
export function instanceOfUserSvcLoginRequest(value: object): value is UserSvcLoginRequest {
    return true;
}

export function UserSvcLoginRequestFromJSON(json: any): UserSvcLoginRequest {
    return UserSvcLoginRequestFromJSONTyped(json, false);
}

export function UserSvcLoginRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcLoginRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'contact': json['contact'] == null ? undefined : json['contact'],
        'password': json['password'] == null ? undefined : json['password'],
        'slug': json['slug'] == null ? undefined : json['slug'],
    };
}

export function UserSvcLoginRequestToJSON(json: any): UserSvcLoginRequest {
    return UserSvcLoginRequestToJSONTyped(json, false);
}

export function UserSvcLoginRequestToJSONTyped(value?: UserSvcLoginRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'contact': value['contact'],
        'password': value['password'],
        'slug': value['slug'],
    };
}

