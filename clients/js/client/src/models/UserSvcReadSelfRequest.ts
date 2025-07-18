/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
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
 * @interface UserSvcReadSelfRequest
 */
export interface UserSvcReadSelfRequest {
    /**
     * 
     * @type {boolean}
     * @memberof UserSvcReadSelfRequest
     */
    countTokens?: boolean;
}

/**
 * Check if a given object implements the UserSvcReadSelfRequest interface.
 */
export function instanceOfUserSvcReadSelfRequest(value: object): value is UserSvcReadSelfRequest {
    return true;
}

export function UserSvcReadSelfRequestFromJSON(json: any): UserSvcReadSelfRequest {
    return UserSvcReadSelfRequestFromJSONTyped(json, false);
}

export function UserSvcReadSelfRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcReadSelfRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'countTokens': json['countTokens'] == null ? undefined : json['countTokens'],
    };
}

export function UserSvcReadSelfRequestToJSON(json: any): UserSvcReadSelfRequest {
    return UserSvcReadSelfRequestToJSONTyped(json, false);
}

export function UserSvcReadSelfRequestToJSONTyped(value?: UserSvcReadSelfRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'countTokens': value['countTokens'],
    };
}

