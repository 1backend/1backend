/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
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
 * @interface UserSvcSaveUserRequest
 */
export interface UserSvcSaveUserRequest {
    /**
     * 
     * @type {string}
     * @memberof UserSvcSaveUserRequest
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcSaveUserRequest
     */
    thumbnailFileId?: string;
}

/**
 * Check if a given object implements the UserSvcSaveUserRequest interface.
 */
export function instanceOfUserSvcSaveUserRequest(value: object): value is UserSvcSaveUserRequest {
    return true;
}

export function UserSvcSaveUserRequestFromJSON(json: any): UserSvcSaveUserRequest {
    return UserSvcSaveUserRequestFromJSONTyped(json, false);
}

export function UserSvcSaveUserRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveUserRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'name': json['name'] == null ? undefined : json['name'],
        'thumbnailFileId': json['thumbnailFileId'] == null ? undefined : json['thumbnailFileId'],
    };
}

export function UserSvcSaveUserRequestToJSON(json: any): UserSvcSaveUserRequest {
    return UserSvcSaveUserRequestToJSONTyped(json, false);
}

export function UserSvcSaveUserRequestToJSONTyped(value?: UserSvcSaveUserRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'name': value['name'],
        'thumbnailFileId': value['thumbnailFileId'],
    };
}

