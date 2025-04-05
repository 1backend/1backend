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
/**
 * Check if a given object implements the UserSvcListInvitesRequest interface.
 */
export function instanceOfUserSvcListInvitesRequest(value) {
    return true;
}
export function UserSvcListInvitesRequestFromJSON(json) {
    return UserSvcListInvitesRequestFromJSONTyped(json, false);
}
export function UserSvcListInvitesRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'contactId': json['contactId'] == null ? undefined : json['contactId'],
        'roleId': json['roleId'] == null ? undefined : json['roleId'],
    };
}
export function UserSvcListInvitesRequestToJSON(json) {
    return UserSvcListInvitesRequestToJSONTyped(json, false);
}
export function UserSvcListInvitesRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'contactId': value['contactId'],
        'roleId': value['roleId'],
    };
}
