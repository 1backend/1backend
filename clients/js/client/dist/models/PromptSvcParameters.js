/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { PromptSvcTextToImageParametersFromJSON, PromptSvcTextToImageParametersToJSON, } from './PromptSvcTextToImageParameters';
import { PromptSvcTextToTextParametersFromJSON, PromptSvcTextToTextParametersToJSON, } from './PromptSvcTextToTextParameters';
/**
 * Check if a given object implements the PromptSvcParameters interface.
 */
export function instanceOfPromptSvcParameters(value) {
    return true;
}
export function PromptSvcParametersFromJSON(json) {
    return PromptSvcParametersFromJSONTyped(json, false);
}
export function PromptSvcParametersFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'textToImage': json['textToImage'] == null ? undefined : PromptSvcTextToImageParametersFromJSON(json['textToImage']),
        'textToText': json['textToText'] == null ? undefined : PromptSvcTextToTextParametersFromJSON(json['textToText']),
    };
}
export function PromptSvcParametersToJSON(json) {
    return PromptSvcParametersToJSONTyped(json, false);
}
export function PromptSvcParametersToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'textToImage': PromptSvcTextToImageParametersToJSON(value['textToImage']),
        'textToText': PromptSvcTextToTextParametersToJSON(value['textToText']),
    };
}
