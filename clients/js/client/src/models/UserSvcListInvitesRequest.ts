/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
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
 * @interface UserSvcListInvitesRequest
 */
export interface UserSvcListInvitesRequest {
    /**
     * 
     * @type {string}
     * @memberof UserSvcListInvitesRequest
     */
    contactId?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcListInvitesRequest
     */
    roleId?: string;
}

/**
 * Check if a given object implements the UserSvcListInvitesRequest interface.
 */
export function instanceOfUserSvcListInvitesRequest(value: object): value is UserSvcListInvitesRequest {
    return true;
}

export function UserSvcListInvitesRequestFromJSON(json: any): UserSvcListInvitesRequest {
    return UserSvcListInvitesRequestFromJSONTyped(json, false);
}

export function UserSvcListInvitesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcListInvitesRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'contactId': json['contactId'] == null ? undefined : json['contactId'],
        'roleId': json['roleId'] == null ? undefined : json['roleId'],
    };
}

export function UserSvcListInvitesRequestToJSON(json: any): UserSvcListInvitesRequest {
    return UserSvcListInvitesRequestToJSONTyped(json, false);
}

export function UserSvcListInvitesRequestToJSONTyped(value?: UserSvcListInvitesRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'contactId': value['contactId'],
        'roleId': value['roleId'],
    };
}

