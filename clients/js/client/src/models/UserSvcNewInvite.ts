/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
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
 * @interface UserSvcNewInvite
 */
export interface UserSvcNewInvite {
    /**
     * 
     * @type {string}
     * @memberof UserSvcNewInvite
     */
    contactId: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcNewInvite
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcNewInvite
     */
    roleId: string;
}

/**
 * Check if a given object implements the UserSvcNewInvite interface.
 */
export function instanceOfUserSvcNewInvite(value: object): value is UserSvcNewInvite {
    if (!('contactId' in value) || value['contactId'] === undefined) return false;
    if (!('roleId' in value) || value['roleId'] === undefined) return false;
    return true;
}

export function UserSvcNewInviteFromJSON(json: any): UserSvcNewInvite {
    return UserSvcNewInviteFromJSONTyped(json, false);
}

export function UserSvcNewInviteFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcNewInvite {
    if (json == null) {
        return json;
    }
    return {
        
        'contactId': json['contactId'],
        'id': json['id'] == null ? undefined : json['id'],
        'roleId': json['roleId'],
    };
}

export function UserSvcNewInviteToJSON(json: any): UserSvcNewInvite {
    return UserSvcNewInviteToJSONTyped(json, false);
}

export function UserSvcNewInviteToJSONTyped(value?: UserSvcNewInvite | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'contactId': value['contactId'],
        'id': value['id'],
        'roleId': value['roleId'],
    };
}

