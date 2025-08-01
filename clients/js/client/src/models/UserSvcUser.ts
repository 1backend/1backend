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
 * @interface UserSvcUser
 */
export interface UserSvcUser {
    /**
     * 
     * @type {string}
     * @memberof UserSvcUser
     */
    createdAt: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcUser
     */
    deletedAt?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcUser
     */
    id: string;
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof UserSvcUser
     */
    labels?: { [key: string]: string; };
    /**
     * Full name of the user.
     * @type {string}
     * @memberof UserSvcUser
     */
    name?: string;
    /**
     * URL-friendly unique (inside the 1Backend platform) identifier for the `user`.
     * @type {string}
     * @memberof UserSvcUser
     */
    slug: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcUser
     */
    thumbnailFileId?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcUser
     */
    updatedAt: string;
}

/**
 * Check if a given object implements the UserSvcUser interface.
 */
export function instanceOfUserSvcUser(value: object): value is UserSvcUser {
    if (!('createdAt' in value) || value['createdAt'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('slug' in value) || value['slug'] === undefined) return false;
    if (!('updatedAt' in value) || value['updatedAt'] === undefined) return false;
    return true;
}

export function UserSvcUserFromJSON(json: any): UserSvcUser {
    return UserSvcUserFromJSONTyped(json, false);
}

export function UserSvcUserFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcUser {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'id': json['id'],
        'labels': json['labels'] == null ? undefined : json['labels'],
        'name': json['name'] == null ? undefined : json['name'],
        'slug': json['slug'],
        'thumbnailFileId': json['thumbnailFileId'] == null ? undefined : json['thumbnailFileId'],
        'updatedAt': json['updatedAt'],
    };
}

export function UserSvcUserToJSON(json: any): UserSvcUser {
    return UserSvcUserToJSONTyped(json, false);
}

export function UserSvcUserToJSONTyped(value?: UserSvcUser | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'createdAt': value['createdAt'],
        'deletedAt': value['deletedAt'],
        'id': value['id'],
        'labels': value['labels'],
        'name': value['name'],
        'slug': value['slug'],
        'thumbnailFileId': value['thumbnailFileId'],
        'updatedAt': value['updatedAt'],
    };
}

