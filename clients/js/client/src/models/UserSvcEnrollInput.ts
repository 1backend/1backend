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
 * @interface UserSvcEnrollInput
 */
export interface UserSvcEnrollInput {
    /**
     * ContactId is the the recipient of the enroll.
     * If the user is already registered, the role is assigned immediately;
     * otherwise, it is applied upon registration.
     * @type {string}
     * @memberof UserSvcEnrollInput
     */
    contactId?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcEnrollInput
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcEnrollInput
     */
    role: string;
    /**
     * UserId is the recipient of the enroll.
     * If the user is already registered, the role is assigned immediately;
     * otherwise, it is applied upon registration.
     * @type {string}
     * @memberof UserSvcEnrollInput
     */
    userId?: string;
}

/**
 * Check if a given object implements the UserSvcEnrollInput interface.
 */
export function instanceOfUserSvcEnrollInput(value: object): value is UserSvcEnrollInput {
    if (!('role' in value) || value['role'] === undefined) return false;
    return true;
}

export function UserSvcEnrollInputFromJSON(json: any): UserSvcEnrollInput {
    return UserSvcEnrollInputFromJSONTyped(json, false);
}

export function UserSvcEnrollInputFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcEnrollInput {
    if (json == null) {
        return json;
    }
    return {
        
        'contactId': json['contactId'] == null ? undefined : json['contactId'],
        'id': json['id'] == null ? undefined : json['id'],
        'role': json['role'],
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}

export function UserSvcEnrollInputToJSON(json: any): UserSvcEnrollInput {
    return UserSvcEnrollInputToJSONTyped(json, false);
}

export function UserSvcEnrollInputToJSONTyped(value?: UserSvcEnrollInput | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'contactId': value['contactId'],
        'id': value['id'],
        'role': value['role'],
        'userId': value['userId'],
    };
}

