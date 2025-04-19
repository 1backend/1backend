/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
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
 * @interface UserSvcListPermitsRequest
 */
export interface UserSvcListPermitsRequest {
    /**
     * 
     * @type {string}
     * @memberof UserSvcListPermitsRequest
     */
    permission?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcListPermitsRequest
     */
    slug?: string;
}

/**
 * Check if a given object implements the UserSvcListPermitsRequest interface.
 */
export function instanceOfUserSvcListPermitsRequest(value: object): value is UserSvcListPermitsRequest {
    return true;
}

export function UserSvcListPermitsRequestFromJSON(json: any): UserSvcListPermitsRequest {
    return UserSvcListPermitsRequestFromJSONTyped(json, false);
}

export function UserSvcListPermitsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcListPermitsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'permission': json['permission'] == null ? undefined : json['permission'],
        'slug': json['slug'] == null ? undefined : json['slug'],
    };
}

export function UserSvcListPermitsRequestToJSON(json: any): UserSvcListPermitsRequest {
    return UserSvcListPermitsRequestToJSONTyped(json, false);
}

export function UserSvcListPermitsRequestToJSONTyped(value?: UserSvcListPermitsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'permission': value['permission'],
        'slug': value['slug'],
    };
}

