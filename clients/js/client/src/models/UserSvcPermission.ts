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
 * @interface UserSvcPermission
 */
export interface UserSvcPermission {
    /**
     * 
     * @type {string}
     * @memberof UserSvcPermission
     */
    createdAt?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcPermission
     */
    description?: string;
    /**
     * eg. "user.viewer"
     * @type {string}
     * @memberof UserSvcPermission
     */
    id?: string;
    /**
     * eg. "User Viewer"
     * @type {string}
     * @memberof UserSvcPermission
     */
    name?: string;
    /**
     * Service who owns the permission
     * 
     * Uncertain if this aligns with the system's use of slugs.
     * Issue encountered: I renamed Docker Svc to Container Svc in two steps (by mistake).
     * The name/slug had already changed to "container-svc," but data was still being saved
     * in the "dockerSvcCredentials" table.
     * After renaming the tables as well, I hit a "cannot update unowned permission" error
     * because ownership relies on this field rather than the user slug. YMMV.
     * @type {string}
     * @memberof UserSvcPermission
     */
    ownerId?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcPermission
     */
    updatedAt?: string;
}

/**
 * Check if a given object implements the UserSvcPermission interface.
 */
export function instanceOfUserSvcPermission(value: object): value is UserSvcPermission {
    return true;
}

export function UserSvcPermissionFromJSON(json: any): UserSvcPermission {
    return UserSvcPermissionFromJSONTyped(json, false);
}

export function UserSvcPermissionFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcPermission {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'description': json['description'] == null ? undefined : json['description'],
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'ownerId': json['ownerId'] == null ? undefined : json['ownerId'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
    };
}

export function UserSvcPermissionToJSON(json: any): UserSvcPermission {
    return UserSvcPermissionToJSONTyped(json, false);
}

export function UserSvcPermissionToJSONTyped(value?: UserSvcPermission | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'createdAt': value['createdAt'],
        'description': value['description'],
        'id': value['id'],
        'name': value['name'],
        'ownerId': value['ownerId'],
        'updatedAt': value['updatedAt'],
    };
}

