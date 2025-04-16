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
 * @interface UserSvcAuthToken
 */
export interface UserSvcAuthToken {
    /**
     * Active tokens contain the most up-to-date information.
     * When a user's role changes—due to role assignment, organization
     * creation/assignment, etc.—all existing tokens are marked inactive.
     * Active tokens are reused during login, while inactive tokens
     * are retained for historical reference.
     * @type {boolean}
     * @memberof UserSvcAuthToken
     */
    active?: boolean;
    /**
     * 
     * @type {string}
     * @memberof UserSvcAuthToken
     */
    createdAt?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcAuthToken
     */
    deletedAt?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcAuthToken
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcAuthToken
     */
    token: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcAuthToken
     */
    updatedAt?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcAuthToken
     */
    userId: string;
}

/**
 * Check if a given object implements the UserSvcAuthToken interface.
 */
export function instanceOfUserSvcAuthToken(value: object): value is UserSvcAuthToken {
    if (!('token' in value) || value['token'] === undefined) return false;
    if (!('userId' in value) || value['userId'] === undefined) return false;
    return true;
}

export function UserSvcAuthTokenFromJSON(json: any): UserSvcAuthToken {
    return UserSvcAuthTokenFromJSONTyped(json, false);
}

export function UserSvcAuthTokenFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcAuthToken {
    if (json == null) {
        return json;
    }
    return {
        
        'active': json['active'] == null ? undefined : json['active'],
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'id': json['id'] == null ? undefined : json['id'],
        'token': json['token'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
        'userId': json['userId'],
    };
}

export function UserSvcAuthTokenToJSON(json: any): UserSvcAuthToken {
    return UserSvcAuthTokenToJSONTyped(json, false);
}

export function UserSvcAuthTokenToJSONTyped(value?: UserSvcAuthToken | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'active': value['active'],
        'createdAt': value['createdAt'],
        'deletedAt': value['deletedAt'],
        'id': value['id'],
        'token': value['token'],
        'updatedAt': value['updatedAt'],
        'userId': value['userId'],
    };
}

