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
/**
 * Check if a given object implements the PromptSvcLlamaCppParameters interface.
 */
export function instanceOfPromptSvcLlamaCppParameters(value) {
    return true;
}
export function PromptSvcLlamaCppParametersFromJSON(json) {
    return PromptSvcLlamaCppParametersFromJSONTyped(json, false);
}
export function PromptSvcLlamaCppParametersFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'template': json['template'] == null ? undefined : json['template'],
    };
}
export function PromptSvcLlamaCppParametersToJSON(json) {
    return PromptSvcLlamaCppParametersToJSONTyped(json, false);
}
export function PromptSvcLlamaCppParametersToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'template': value['template'],
    };
}
