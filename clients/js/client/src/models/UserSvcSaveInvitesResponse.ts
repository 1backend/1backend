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
import type { UserSvcInvite } from './UserSvcInvite';
import {
    UserSvcInviteFromJSON,
    UserSvcInviteFromJSONTyped,
    UserSvcInviteToJSON,
    UserSvcInviteToJSONTyped,
} from './UserSvcInvite';

/**
 * 
 * @export
 * @interface UserSvcSaveInvitesResponse
 */
export interface UserSvcSaveInvitesResponse {
    /**
     * 
     * @type {Array<UserSvcInvite>}
     * @memberof UserSvcSaveInvitesResponse
     */
    invites: Array<UserSvcInvite>;
}

/**
 * Check if a given object implements the UserSvcSaveInvitesResponse interface.
 */
export function instanceOfUserSvcSaveInvitesResponse(value: object): value is UserSvcSaveInvitesResponse {
    if (!('invites' in value) || value['invites'] === undefined) return false;
    return true;
}

export function UserSvcSaveInvitesResponseFromJSON(json: any): UserSvcSaveInvitesResponse {
    return UserSvcSaveInvitesResponseFromJSONTyped(json, false);
}

export function UserSvcSaveInvitesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveInvitesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'invites': ((json['invites'] as Array<any>).map(UserSvcInviteFromJSON)),
    };
}

export function UserSvcSaveInvitesResponseToJSON(json: any): UserSvcSaveInvitesResponse {
    return UserSvcSaveInvitesResponseToJSONTyped(json, false);
}

export function UserSvcSaveInvitesResponseToJSONTyped(value?: UserSvcSaveInvitesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'invites': ((value['invites'] as Array<any>).map(UserSvcInviteToJSON)),
    };
}

