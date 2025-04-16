/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcEnrollFromJSON, UserSvcEnrollToJSON, } from './UserSvcEnroll';
/**
 * Check if a given object implements the UserSvcListEnrollsResponse interface.
 */
export function instanceOfUserSvcListEnrollsResponse(value) {
    if (!('enrolls' in value) || value['enrolls'] === undefined)
        return false;
    return true;
}
export function UserSvcListEnrollsResponseFromJSON(json) {
    return UserSvcListEnrollsResponseFromJSONTyped(json, false);
}
export function UserSvcListEnrollsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'enrolls': (json['enrolls'].map(UserSvcEnrollFromJSON)),
    };
}
export function UserSvcListEnrollsResponseToJSON(json) {
    return UserSvcListEnrollsResponseToJSONTyped(json, false);
}
export function UserSvcListEnrollsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'enrolls': (value['enrolls'].map(UserSvcEnrollToJSON)),
    };
}
