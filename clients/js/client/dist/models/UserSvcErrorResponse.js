/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcErrorResponse interface.
 */
export function instanceOfUserSvcErrorResponse(value) {
    return true;
}
export function UserSvcErrorResponseFromJSON(json) {
    return UserSvcErrorResponseFromJSONTyped(json, false);
}
export function UserSvcErrorResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'error': json['error'] == null ? undefined : json['error'],
    };
}
export function UserSvcErrorResponseToJSON(json) {
    return UserSvcErrorResponseToJSONTyped(json, false);
}
export function UserSvcErrorResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'error': value['error'],
    };
}
