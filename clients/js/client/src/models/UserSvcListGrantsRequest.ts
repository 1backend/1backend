/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
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
 * @interface UserSvcListGrantsRequest
 */
export interface UserSvcListGrantsRequest {
    /**
     * 
     * @type {string}
     * @memberof UserSvcListGrantsRequest
     */
    permissionId?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcListGrantsRequest
     */
    slug?: string;
}

/**
 * Check if a given object implements the UserSvcListGrantsRequest interface.
 */
export function instanceOfUserSvcListGrantsRequest(value: object): value is UserSvcListGrantsRequest {
    return true;
}

export function UserSvcListGrantsRequestFromJSON(json: any): UserSvcListGrantsRequest {
    return UserSvcListGrantsRequestFromJSONTyped(json, false);
}

export function UserSvcListGrantsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcListGrantsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'permissionId': json['permissionId'] == null ? undefined : json['permissionId'],
        'slug': json['slug'] == null ? undefined : json['slug'],
    };
}

export function UserSvcListGrantsRequestToJSON(json: any): UserSvcListGrantsRequest {
    return UserSvcListGrantsRequestToJSONTyped(json, false);
}

export function UserSvcListGrantsRequestToJSONTyped(value?: UserSvcListGrantsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'permissionId': value['permissionId'],
        'slug': value['slug'],
    };
}

