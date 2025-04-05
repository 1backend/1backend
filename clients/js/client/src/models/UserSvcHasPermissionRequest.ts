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
 * @interface UserSvcHasPermissionRequest
 */
export interface UserSvcHasPermissionRequest {
    /**
     * 
     * @type {Array<string>}
     * @memberof UserSvcHasPermissionRequest
     */
    contactsGranted?: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof UserSvcHasPermissionRequest
     */
    grantedSlugs?: Array<string>;
}

/**
 * Check if a given object implements the UserSvcHasPermissionRequest interface.
 */
export function instanceOfUserSvcHasPermissionRequest(value: object): value is UserSvcHasPermissionRequest {
    return true;
}

export function UserSvcHasPermissionRequestFromJSON(json: any): UserSvcHasPermissionRequest {
    return UserSvcHasPermissionRequestFromJSONTyped(json, false);
}

export function UserSvcHasPermissionRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcHasPermissionRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'contactsGranted': json['contactsGranted'] == null ? undefined : json['contactsGranted'],
        'grantedSlugs': json['grantedSlugs'] == null ? undefined : json['grantedSlugs'],
    };
}

export function UserSvcHasPermissionRequestToJSON(json: any): UserSvcHasPermissionRequest {
    return UserSvcHasPermissionRequestToJSONTyped(json, false);
}

export function UserSvcHasPermissionRequestToJSONTyped(value?: UserSvcHasPermissionRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'contactsGranted': value['contactsGranted'],
        'grantedSlugs': value['grantedSlugs'],
    };
}

