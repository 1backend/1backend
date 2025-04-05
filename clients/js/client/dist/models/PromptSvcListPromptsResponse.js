/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { PromptSvcPromptFromJSON, PromptSvcPromptToJSON, } from './PromptSvcPrompt';
/**
 * Check if a given object implements the PromptSvcListPromptsResponse interface.
 */
export function instanceOfPromptSvcListPromptsResponse(value) {
    return true;
}
export function PromptSvcListPromptsResponseFromJSON(json) {
    return PromptSvcListPromptsResponseFromJSONTyped(json, false);
}
export function PromptSvcListPromptsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'after': json['after'] == null ? undefined : json['after'],
        'count': json['count'] == null ? undefined : json['count'],
        'prompts': json['prompts'] == null ? undefined : (json['prompts'].map(PromptSvcPromptFromJSON)),
    };
}
export function PromptSvcListPromptsResponseToJSON(json) {
    return PromptSvcListPromptsResponseToJSONTyped(json, false);
}
export function PromptSvcListPromptsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'after': value['after'],
        'count': value['count'],
        'prompts': value['prompts'] == null ? undefined : (value['prompts'].map(PromptSvcPromptToJSON)),
    };
}
