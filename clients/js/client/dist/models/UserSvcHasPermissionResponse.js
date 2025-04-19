/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcUserFromJSON, UserSvcUserToJSON, } from './UserSvcUser';
/**
 * Check if a given object implements the UserSvcHasPermissionResponse interface.
 */
export function instanceOfUserSvcHasPermissionResponse(value) {
    return true;
}
export function UserSvcHasPermissionResponseFromJSON(json) {
    return UserSvcHasPermissionResponseFromJSONTyped(json, false);
}
export function UserSvcHasPermissionResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'authorized': json['authorized'] == null ? undefined : json['authorized'],
        'user': json['user'] == null ? undefined : UserSvcUserFromJSON(json['user']),
    };
}
export function UserSvcHasPermissionResponseToJSON(json) {
    return UserSvcHasPermissionResponseToJSONTyped(json, false);
}
export function UserSvcHasPermissionResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'authorized': value['authorized'],
        'user': UserSvcUserToJSON(value['user']),
    };
}
