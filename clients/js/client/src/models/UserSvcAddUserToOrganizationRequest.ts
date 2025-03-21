/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
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
 * @interface UserSvcAddUserToOrganizationRequest
 */
export interface UserSvcAddUserToOrganizationRequest {
    /**
     * 
     * @type {string}
     * @memberof UserSvcAddUserToOrganizationRequest
     */
    userId?: string;
}

/**
 * Check if a given object implements the UserSvcAddUserToOrganizationRequest interface.
 */
export function instanceOfUserSvcAddUserToOrganizationRequest(value: object): value is UserSvcAddUserToOrganizationRequest {
    return true;
}

export function UserSvcAddUserToOrganizationRequestFromJSON(json: any): UserSvcAddUserToOrganizationRequest {
    return UserSvcAddUserToOrganizationRequestFromJSONTyped(json, false);
}

export function UserSvcAddUserToOrganizationRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcAddUserToOrganizationRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}

export function UserSvcAddUserToOrganizationRequestToJSON(json: any): UserSvcAddUserToOrganizationRequest {
    return UserSvcAddUserToOrganizationRequestToJSONTyped(json, false);
}

export function UserSvcAddUserToOrganizationRequestToJSONTyped(value?: UserSvcAddUserToOrganizationRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'userId': value['userId'],
    };
}

