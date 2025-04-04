/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcUserFromJSON, UserSvcUserToJSON, } from './UserSvcUser';
/**
 * Check if a given object implements the UserSvcIsAuthorizedResponse interface.
 */
export function instanceOfUserSvcIsAuthorizedResponse(value) {
    return true;
}
export function UserSvcIsAuthorizedResponseFromJSON(json) {
    return UserSvcIsAuthorizedResponseFromJSONTyped(json, false);
}
export function UserSvcIsAuthorizedResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'authorized': json['authorized'] == null ? undefined : json['authorized'],
        'user': json['user'] == null ? undefined : UserSvcUserFromJSON(json['user']),
    };
}
export function UserSvcIsAuthorizedResponseToJSON(json) {
    return UserSvcIsAuthorizedResponseToJSONTyped(json, false);
}
export function UserSvcIsAuthorizedResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'authorized': value['authorized'],
        'user': UserSvcUserToJSON(value['user']),
    };
}
