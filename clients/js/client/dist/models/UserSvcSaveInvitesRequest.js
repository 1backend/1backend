/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcNewInviteFromJSON, UserSvcNewInviteToJSON, } from './UserSvcNewInvite';
/**
 * Check if a given object implements the UserSvcSaveInvitesRequest interface.
 */
export function instanceOfUserSvcSaveInvitesRequest(value) {
    if (!('invites' in value) || value['invites'] === undefined)
        return false;
    return true;
}
export function UserSvcSaveInvitesRequestFromJSON(json) {
    return UserSvcSaveInvitesRequestFromJSONTyped(json, false);
}
export function UserSvcSaveInvitesRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'invites': (json['invites'].map(UserSvcNewInviteFromJSON)),
    };
}
export function UserSvcSaveInvitesRequestToJSON(json) {
    return UserSvcSaveInvitesRequestToJSONTyped(json, false);
}
export function UserSvcSaveInvitesRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'invites': (value['invites'].map(UserSvcNewInviteToJSON)),
    };
}
