/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcIsAuthorizedRequest interface.
 */
export function instanceOfUserSvcIsAuthorizedRequest(value) {
    return true;
}
export function UserSvcIsAuthorizedRequestFromJSON(json) {
    return UserSvcIsAuthorizedRequestFromJSONTyped(json, false);
}
export function UserSvcIsAuthorizedRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'contactsGranted': json['contactsGranted'] == null ? undefined : json['contactsGranted'],
        'grantedSlugs': json['grantedSlugs'] == null ? undefined : json['grantedSlugs'],
    };
}
export function UserSvcIsAuthorizedRequestToJSON(json) {
    return UserSvcIsAuthorizedRequestToJSONTyped(json, false);
}
export function UserSvcIsAuthorizedRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'contactsGranted': value['contactsGranted'],
        'grantedSlugs': value['grantedSlugs'],
    };
}
