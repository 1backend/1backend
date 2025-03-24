/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcNewInvite interface.
 */
export function instanceOfUserSvcNewInvite(value) {
    if (!('contactId' in value) || value['contactId'] === undefined)
        return false;
    if (!('roleId' in value) || value['roleId'] === undefined)
        return false;
    return true;
}
export function UserSvcNewInviteFromJSON(json) {
    return UserSvcNewInviteFromJSONTyped(json, false);
}
export function UserSvcNewInviteFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'contactId': json['contactId'],
        'id': json['id'] == null ? undefined : json['id'],
        'roleId': json['roleId'],
    };
}
export function UserSvcNewInviteToJSON(json) {
    return UserSvcNewInviteToJSONTyped(json, false);
}
export function UserSvcNewInviteToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'contactId': value['contactId'],
        'id': value['id'],
        'roleId': value['roleId'],
    };
}
