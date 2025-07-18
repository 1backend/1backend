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
import { UserSvcUserInputFromJSON, UserSvcUserInputToJSON, } from './UserSvcUserInput';
import { UserSvcContactFromJSON, UserSvcContactToJSON, } from './UserSvcContact';
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
        'app': json['app'] == null ? undefined : json['app'],
        'contacts': json['contacts'] == null ? undefined : (json['contacts'].map(UserSvcContactFromJSON)),
        'password': json['password'] == null ? undefined : json['password'],
        'roleIds': json['roleIds'] == null ? undefined : json['roleIds'],
        'user': json['user'] == null ? undefined : UserSvcUserInputFromJSON(json['user']),
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
        'app': value['app'],
        'contacts': value['contacts'] == null ? undefined : (value['contacts'].map(UserSvcContactToJSON)),
        'password': value['password'],
        'roleIds': value['roleIds'],
        'user': UserSvcUserInputToJSON(value['user']),
    };
}
