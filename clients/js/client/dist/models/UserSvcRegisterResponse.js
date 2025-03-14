/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcAuthTokenFromJSON, UserSvcAuthTokenToJSON, } from './UserSvcAuthToken';
/**
 * Check if a given object implements the UserSvcRegisterResponse interface.
 */
export function instanceOfUserSvcRegisterResponse(value) {
    return true;
}
export function UserSvcRegisterResponseFromJSON(json) {
    return UserSvcRegisterResponseFromJSONTyped(json, false);
}
export function UserSvcRegisterResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'token': json['token'] == null ? undefined : UserSvcAuthTokenFromJSON(json['token']),
    };
}
export function UserSvcRegisterResponseToJSON(json) {
    return UserSvcRegisterResponseToJSONTyped(json, false);
}
export function UserSvcRegisterResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'token': UserSvcAuthTokenToJSON(value['token']),
    };
}
