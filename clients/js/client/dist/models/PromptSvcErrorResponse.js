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
/**
 * Check if a given object implements the PromptSvcErrorResponse interface.
 */
export function instanceOfPromptSvcErrorResponse(value) {
    return true;
}
export function PromptSvcErrorResponseFromJSON(json) {
    return PromptSvcErrorResponseFromJSONTyped(json, false);
}
export function PromptSvcErrorResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'error': json['error'] == null ? undefined : json['error'],
    };
}
export function PromptSvcErrorResponseToJSON(json) {
    return PromptSvcErrorResponseToJSONTyped(json, false);
}
export function PromptSvcErrorResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'error': value['error'],
    };
}
