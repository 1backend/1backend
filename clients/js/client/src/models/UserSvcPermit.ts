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
 * @interface UserSvcPermit
 */
export interface UserSvcPermit {
    /**
     * App of the permit.
     * Use `*` to match all apps, such as when bootstrapping
     * in services.
     * @type {string}
     * @memberof UserSvcPermit
     */
    app?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcPermit
     */
    createdAt: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcPermit
     */
    deletedAt?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcPermit
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcPermit
     */
    permission: string;
    /**
     * Role IDs that have been permited the specified permission.
     * 
     * Originally, permits were designed for slugs to facilitate service-to-service calls.
     * Due to their convenience—especially with CLI and infrastructure-as-code support—they were later extended to roles.
     * @type {Array<string>}
     * @memberof UserSvcPermit
     */
    roles?: Array<string>;
    /**
     * Slugs that have been permited the specified permission.
     * @type {Array<string>}
     * @memberof UserSvcPermit
     */
    slugs?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof UserSvcPermit
     */
    updatedAt: string;
}

/**
 * Check if a given object implements the UserSvcPermit interface.
 */
export function instanceOfUserSvcPermit(value: object): value is UserSvcPermit {
    if (!('createdAt' in value) || value['createdAt'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('permission' in value) || value['permission'] === undefined) return false;
    if (!('updatedAt' in value) || value['updatedAt'] === undefined) return false;
    return true;
}

export function UserSvcPermitFromJSON(json: any): UserSvcPermit {
    return UserSvcPermitFromJSONTyped(json, false);
}

export function UserSvcPermitFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcPermit {
    if (json == null) {
        return json;
    }
    return {
        
        'app': json['app'] == null ? undefined : json['app'],
        'createdAt': json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'id': json['id'],
        'permission': json['permission'],
        'roles': json['roles'] == null ? undefined : json['roles'],
        'slugs': json['slugs'] == null ? undefined : json['slugs'],
        'updatedAt': json['updatedAt'],
    };
}

export function UserSvcPermitToJSON(json: any): UserSvcPermit {
    return UserSvcPermitToJSONTyped(json, false);
}

export function UserSvcPermitToJSONTyped(value?: UserSvcPermit | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'app': value['app'],
        'createdAt': value['createdAt'],
        'deletedAt': value['deletedAt'],
        'id': value['id'],
        'permission': value['permission'],
        'roles': value['roles'],
        'slugs': value['slugs'],
        'updatedAt': value['updatedAt'],
    };
}

