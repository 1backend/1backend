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
import { PromptSvcLlamaCppParametersFromJSON, PromptSvcLlamaCppParametersToJSON, } from './PromptSvcLlamaCppParameters';
import { PromptSvcStableDiffusionParametersFromJSON, PromptSvcStableDiffusionParametersToJSON, } from './PromptSvcStableDiffusionParameters';
/**
 * Check if a given object implements the PromptSvcEngineParameters interface.
 */
export function instanceOfPromptSvcEngineParameters(value) {
    return true;
}
export function PromptSvcEngineParametersFromJSON(json) {
    return PromptSvcEngineParametersFromJSONTyped(json, false);
}
export function PromptSvcEngineParametersFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'llamaCppParameters': json['llamaCppParameters'] == null ? undefined : PromptSvcLlamaCppParametersFromJSON(json['llamaCppParameters']),
        'stableDiffusion': json['stableDiffusion'] == null ? undefined : PromptSvcStableDiffusionParametersFromJSON(json['stableDiffusion']),
    };
}
export function PromptSvcEngineParametersToJSON(json) {
    return PromptSvcEngineParametersToJSONTyped(json, false);
}
export function PromptSvcEngineParametersToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'llamaCppParameters': PromptSvcLlamaCppParametersToJSON(value['llamaCppParameters']),
        'stableDiffusion': PromptSvcStableDiffusionParametersToJSON(value['stableDiffusion']),
    };
}
