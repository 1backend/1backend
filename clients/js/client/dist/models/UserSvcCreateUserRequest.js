/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcContactFromJSON, UserSvcContactToJSON, } from './UserSvcContact';
import { UserSvcUserFromJSON, UserSvcUserToJSON, } from './UserSvcUser';
/**
 * Check if a given object implements the UserSvcCreateUserRequest interface.
 */
export function instanceOfUserSvcCreateUserRequest(value) {
    return true;
}
export function UserSvcCreateUserRequestFromJSON(json) {
    return UserSvcCreateUserRequestFromJSONTyped(json, false);
}
export function UserSvcCreateUserRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'contacts': json['contacts'] == null ? undefined : (json['contacts'].map(UserSvcContactFromJSON)),
        'password': json['password'] == null ? undefined : json['password'],
        'roleIds': json['roleIds'] == null ? undefined : json['roleIds'],
        'user': json['user'] == null ? undefined : UserSvcUserFromJSON(json['user']),
    };
}
export function UserSvcCreateUserRequestToJSON(json) {
    return UserSvcCreateUserRequestToJSONTyped(json, false);
}
export function UserSvcCreateUserRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'contacts': value['contacts'] == null ? undefined : (value['contacts'].map(UserSvcContactToJSON)),
        'password': value['password'],
        'roleIds': value['roleIds'],
        'user': UserSvcUserToJSON(value['user']),
    };
}
