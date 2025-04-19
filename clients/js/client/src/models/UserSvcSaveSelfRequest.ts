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
 * @interface UserSvcSaveSelfRequest
 */
export interface UserSvcSaveSelfRequest {
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof UserSvcSaveSelfRequest
     */
    labels?: { [key: string]: string; };
    /**
     * 
     * @type {string}
     * @memberof UserSvcSaveSelfRequest
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcSaveSelfRequest
     */
    thumbnailFileId?: string;
}

/**
 * Check if a given object implements the UserSvcSaveSelfRequest interface.
 */
export function instanceOfUserSvcSaveSelfRequest(value: object): value is UserSvcSaveSelfRequest {
    return true;
}

export function UserSvcSaveSelfRequestFromJSON(json: any): UserSvcSaveSelfRequest {
    return UserSvcSaveSelfRequestFromJSONTyped(json, false);
}

export function UserSvcSaveSelfRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveSelfRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'labels': json['labels'] == null ? undefined : json['labels'],
        'name': json['name'] == null ? undefined : json['name'],
        'thumbnailFileId': json['thumbnailFileId'] == null ? undefined : json['thumbnailFileId'],
    };
}

export function UserSvcSaveSelfRequestToJSON(json: any): UserSvcSaveSelfRequest {
    return UserSvcSaveSelfRequestToJSONTyped(json, false);
}

export function UserSvcSaveSelfRequestToJSONTyped(value?: UserSvcSaveSelfRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'labels': value['labels'],
        'name': value['name'],
        'thumbnailFileId': value['thumbnailFileId'],
    };
}

