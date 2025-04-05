/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { UserSvcNewInvite } from './UserSvcNewInvite';
import {
    UserSvcNewInviteFromJSON,
    UserSvcNewInviteFromJSONTyped,
    UserSvcNewInviteToJSON,
    UserSvcNewInviteToJSONTyped,
} from './UserSvcNewInvite';

/**
 * 
 * @export
 * @interface UserSvcSaveInvitesRequest
 */
export interface UserSvcSaveInvitesRequest {
    /**
     * 
     * @type {Array<UserSvcNewInvite>}
     * @memberof UserSvcSaveInvitesRequest
     */
    invites: Array<UserSvcNewInvite>;
}

/**
 * Check if a given object implements the UserSvcSaveInvitesRequest interface.
 */
export function instanceOfUserSvcSaveInvitesRequest(value: object): value is UserSvcSaveInvitesRequest {
    if (!('invites' in value) || value['invites'] === undefined) return false;
    return true;
}

export function UserSvcSaveInvitesRequestFromJSON(json: any): UserSvcSaveInvitesRequest {
    return UserSvcSaveInvitesRequestFromJSONTyped(json, false);
}

export function UserSvcSaveInvitesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveInvitesRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'invites': ((json['invites'] as Array<any>).map(UserSvcNewInviteFromJSON)),
    };
}

export function UserSvcSaveInvitesRequestToJSON(json: any): UserSvcSaveInvitesRequest {
    return UserSvcSaveInvitesRequestToJSONTyped(json, false);
}

export function UserSvcSaveInvitesRequestToJSONTyped(value?: UserSvcSaveInvitesRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'invites': ((value['invites'] as Array<any>).map(UserSvcNewInviteToJSON)),
    };
}

