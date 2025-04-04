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
 * @interface UserSvcIsAuthorizedRequest
 */
export interface UserSvcIsAuthorizedRequest {
    /**
     * 
     * @type {Array<string>}
     * @memberof UserSvcIsAuthorizedRequest
     */
    contactsGranted?: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof UserSvcIsAuthorizedRequest
     */
    grantedSlugs?: Array<string>;
}

/**
 * Check if a given object implements the UserSvcIsAuthorizedRequest interface.
 */
export function instanceOfUserSvcIsAuthorizedRequest(value: object): value is UserSvcIsAuthorizedRequest {
    return true;
}

export function UserSvcIsAuthorizedRequestFromJSON(json: any): UserSvcIsAuthorizedRequest {
    return UserSvcIsAuthorizedRequestFromJSONTyped(json, false);
}

export function UserSvcIsAuthorizedRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcIsAuthorizedRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'contactsGranted': json['contactsGranted'] == null ? undefined : json['contactsGranted'],
        'grantedSlugs': json['grantedSlugs'] == null ? undefined : json['grantedSlugs'],
    };
}

export function UserSvcIsAuthorizedRequestToJSON(json: any): UserSvcIsAuthorizedRequest {
    return UserSvcIsAuthorizedRequestToJSONTyped(json, false);
}

export function UserSvcIsAuthorizedRequestToJSONTyped(value?: UserSvcIsAuthorizedRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'contactsGranted': value['contactsGranted'],
        'grantedSlugs': value['grantedSlugs'],
    };
}

