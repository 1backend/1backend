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
 * @interface UserSvcSaveOrganizationRequest
 */
export interface UserSvcSaveOrganizationRequest {
    /**
     * 
     * @type {string}
     * @memberof UserSvcSaveOrganizationRequest
     */
    id?: string;
    /**
     * Full name of the organization.
     * @type {string}
     * @memberof UserSvcSaveOrganizationRequest
     */
    name?: string;
    /**
     * URL-friendly unique (inside the Singularon platform) identifier for the `organization`.
     * @type {string}
     * @memberof UserSvcSaveOrganizationRequest
     */
    slug: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcSaveOrganizationRequest
     */
    thumbnailFileId?: string;
}

/**
 * Check if a given object implements the UserSvcSaveOrganizationRequest interface.
 */
export function instanceOfUserSvcSaveOrganizationRequest(value: object): value is UserSvcSaveOrganizationRequest {
    if (!('slug' in value) || value['slug'] === undefined) return false;
    return true;
}

export function UserSvcSaveOrganizationRequestFromJSON(json: any): UserSvcSaveOrganizationRequest {
    return UserSvcSaveOrganizationRequestFromJSONTyped(json, false);
}

export function UserSvcSaveOrganizationRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveOrganizationRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'slug': json['slug'],
        'thumbnailFileId': json['thumbnailFileId'] == null ? undefined : json['thumbnailFileId'],
    };
}

export function UserSvcSaveOrganizationRequestToJSON(json: any): UserSvcSaveOrganizationRequest {
    return UserSvcSaveOrganizationRequestToJSONTyped(json, false);
}

export function UserSvcSaveOrganizationRequestToJSONTyped(value?: UserSvcSaveOrganizationRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'name': value['name'],
        'slug': value['slug'],
        'thumbnailFileId': value['thumbnailFileId'],
    };
}

