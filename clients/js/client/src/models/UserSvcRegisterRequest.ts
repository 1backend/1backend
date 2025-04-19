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
 * @interface UserSvcRegisterRequest
 */
export interface UserSvcRegisterRequest {
    /**
     * 
     * @type {UserSvcContact}
     * @memberof UserSvcRegisterRequest
     */
    contact?: UserSvcContact;
    /**
     * 
     * @type {string}
     * @memberof UserSvcRegisterRequest
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcRegisterRequest
     */
    password?: string;
    /**
     * Slug is a URL-friendly unique (inside the 1Backend platform) identifier for the `user`.
     * Required due to its central role in the platform.
     * If your project has no use for a slug, just derive it from the email or similar.
     * @type {string}
     * @memberof UserSvcRegisterRequest
     */
    slug: string;
}

/**
 * Check if a given object implements the UserSvcRegisterRequest interface.
 */
export function instanceOfUserSvcRegisterRequest(value: object): value is UserSvcRegisterRequest {
    if (!('slug' in value) || value['slug'] === undefined) return false;
    return true;
}

export function UserSvcRegisterRequestFromJSON(json: any): UserSvcRegisterRequest {
    return UserSvcRegisterRequestFromJSONTyped(json, false);
}

export function UserSvcRegisterRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcRegisterRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'contact': json['contact'] == null ? undefined : UserSvcContactFromJSON(json['contact']),
        'name': json['name'] == null ? undefined : json['name'],
        'password': json['password'] == null ? undefined : json['password'],
        'slug': json['slug'],
    };
}

export function UserSvcRegisterRequestToJSON(json: any): UserSvcRegisterRequest {
    return UserSvcRegisterRequestToJSONTyped(json, false);
}

export function UserSvcRegisterRequestToJSONTyped(value?: UserSvcRegisterRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'contact': UserSvcContactToJSON(value['contact']),
        'name': value['name'],
        'password': value['password'],
        'slug': value['slug'],
    };
}

