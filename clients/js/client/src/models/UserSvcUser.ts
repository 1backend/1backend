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
import type { UserSvcContact } from './UserSvcContact';
import {
    UserSvcContactFromJSON,
    UserSvcContactFromJSONTyped,
    UserSvcContactToJSON,
    UserSvcContactToJSONTyped,
} from './UserSvcContact';

/**
 * 
 * @export
 * @interface UserSvcUser
 */
export interface UserSvcUser {
    /**
     * Contacts are used for login and identification purposes.
     * @type {Array<UserSvcContact>}
     * @memberof UserSvcUser
     */
    contacts?: Array<UserSvcContact>;
    /**
     * 
     * @type {string}
     * @memberof UserSvcUser
     */
    createdAt?: string;
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
    id?: string;
    /**
     * Full name of the organization.
     * @type {string}
     * @memberof UserSvcUser
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcUser
     */
    passwordHash?: string;
    /**
     * URL-friendly unique (inside the Singularon platform) identifier for the `user`.
     * @type {string}
     * @memberof UserSvcUser
     */
    slug?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcUser
     */
    updatedAt?: string;
}

/**
 * Check if a given object implements the UserSvcUser interface.
 */
export function instanceOfUserSvcUser(value: object): value is UserSvcUser {
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
        
        'contacts': json['contacts'] == null ? undefined : ((json['contacts'] as Array<any>).map(UserSvcContactFromJSON)),
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'passwordHash': json['passwordHash'] == null ? undefined : json['passwordHash'],
        'slug': json['slug'] == null ? undefined : json['slug'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
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
        
        'contacts': value['contacts'] == null ? undefined : ((value['contacts'] as Array<any>).map(UserSvcContactToJSON)),
        'createdAt': value['createdAt'],
        'deletedAt': value['deletedAt'],
        'id': value['id'],
        'name': value['name'],
        'passwordHash': value['passwordHash'],
        'slug': value['slug'],
        'updatedAt': value['updatedAt'],
    };
}

