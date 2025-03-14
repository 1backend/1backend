/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcContactFromJSON, UserSvcContactToJSON, } from './UserSvcContact';
/**
 * Check if a given object implements the UserSvcRegisterRequest interface.
 */
export function instanceOfUserSvcRegisterRequest(value) {
    return true;
}
export function UserSvcRegisterRequestFromJSON(json) {
    return UserSvcRegisterRequestFromJSONTyped(json, false);
}
export function UserSvcRegisterRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'contact': json['contact'] == null ? undefined : UserSvcContactFromJSON(json['contact']),
        'name': json['name'] == null ? undefined : json['name'],
        'password': json['password'] == null ? undefined : json['password'],
        'slug': json['slug'] == null ? undefined : json['slug'],
    };
}
export function UserSvcRegisterRequestToJSON(json) {
    return UserSvcRegisterRequestToJSONTyped(json, false);
}
export function UserSvcRegisterRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'contact': UserSvcContactToJSON(value['contact']),
        'name': value['name'],
        'password': value['password'],
        'slug': value['slug'],
    };
}
