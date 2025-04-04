/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
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
 * @interface UserSvcInvite
 */
export interface UserSvcInvite {
    /**
     * 
     * @type {string}
     * @memberof UserSvcInvite
     */
    appliedAt?: string;
    /**
     * ContactId represents the recipient of the invite.
     * If the user is already registered, the role is assigned immediately;
     * otherwise, it is applied upon registration.
     * @type {string}
     * @memberof UserSvcInvite
     */
    contactId: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcInvite
     */
    createdAt: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcInvite
     */
    deletedAt?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcInvite
     */
    expiresAt?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcInvite
     */
    id: string;
    /**
     * OwnerIds specifies the users who created the invite.
     * If you create an invite that already exists for a given role and contact ID,
     * you get added to the list of owners.
     * @type {Array<string>}
     * @memberof UserSvcInvite
     */
    ownerIds: Array<string>;
    /**
     * RoleId specifies the role to be assigned to the ContactId.
     * Callers can only assign roles they own, identified by their service slug
     * (e.g., if "my-service" creates an invite, the role must be "my-service:admin").
     * Dynamic organization roles can also be assigned
     * (e.g., "user-svc:org:{%orgId}:admin" or "user-svc:org:{%orgId}:user"),
     * but in this case, the caller must be an admin of the target organization.
     * @type {string}
     * @memberof UserSvcInvite
     */
    roleId: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcInvite
     */
    updatedAt?: string;
}

/**
 * Check if a given object implements the UserSvcInvite interface.
 */
export function instanceOfUserSvcInvite(value: object): value is UserSvcInvite {
    if (!('contactId' in value) || value['contactId'] === undefined) return false;
    if (!('createdAt' in value) || value['createdAt'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('ownerIds' in value) || value['ownerIds'] === undefined) return false;
    if (!('roleId' in value) || value['roleId'] === undefined) return false;
    return true;
}

export function UserSvcInviteFromJSON(json: any): UserSvcInvite {
    return UserSvcInviteFromJSONTyped(json, false);
}

export function UserSvcInviteFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcInvite {
    if (json == null) {
        return json;
    }
    return {
        
        'appliedAt': json['appliedAt'] == null ? undefined : json['appliedAt'],
        'contactId': json['contactId'],
        'createdAt': json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'expiresAt': json['expiresAt'] == null ? undefined : json['expiresAt'],
        'id': json['id'],
        'ownerIds': json['ownerIds'],
        'roleId': json['roleId'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
    };
}

export function UserSvcInviteToJSON(json: any): UserSvcInvite {
    return UserSvcInviteToJSONTyped(json, false);
}

export function UserSvcInviteToJSONTyped(value?: UserSvcInvite | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'appliedAt': value['appliedAt'],
        'contactId': value['contactId'],
        'createdAt': value['createdAt'],
        'deletedAt': value['deletedAt'],
        'expiresAt': value['expiresAt'],
        'id': value['id'],
        'ownerIds': value['ownerIds'],
        'roleId': value['roleId'],
        'updatedAt': value['updatedAt'],
    };
}

