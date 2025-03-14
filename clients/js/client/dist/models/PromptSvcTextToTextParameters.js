/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
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
 * Check if a given object implements the PromptSvcTextToTextParameters interface.
 */
export function instanceOfPromptSvcTextToTextParameters(value) {
    return true;
}
export function PromptSvcTextToTextParametersFromJSON(json) {
    return PromptSvcTextToTextParametersFromJSONTyped(json, false);
}
export function PromptSvcTextToTextParametersFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'template': json['template'] == null ? undefined : json['template'],
    };
}
export function PromptSvcTextToTextParametersToJSON(json) {
    return PromptSvcTextToTextParametersToJSONTyped(json, false);
}
export function PromptSvcTextToTextParametersToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'template': value['template'],
    };
}
