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
/**
 * Check if a given object implements the EmailSvcErrorResponse interface.
 */
export function instanceOfEmailSvcErrorResponse(value) {
    return true;
}
export function EmailSvcErrorResponseFromJSON(json) {
    return EmailSvcErrorResponseFromJSONTyped(json, false);
}
export function EmailSvcErrorResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'error': json['error'] == null ? undefined : json['error'],
    };
}
export function EmailSvcErrorResponseToJSON(json) {
    return EmailSvcErrorResponseToJSONTyped(json, false);
}
export function EmailSvcErrorResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'error': value['error'],
    };
}
