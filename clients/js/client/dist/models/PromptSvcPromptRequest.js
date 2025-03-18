/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { PromptSvcParametersFromJSON, PromptSvcParametersToJSON, } from './PromptSvcParameters';
import { PromptSvcEngineParametersFromJSON, PromptSvcEngineParametersToJSON, } from './PromptSvcEngineParameters';
/**
 * Check if a given object implements the PromptSvcPromptRequest interface.
 */
export function instanceOfPromptSvcPromptRequest(value) {
    if (!('prompt' in value) || value['prompt'] === undefined)
        return false;
    return true;
}
export function PromptSvcPromptRequestFromJSON(json) {
    return PromptSvcPromptRequestFromJSONTyped(json, false);
}
export function PromptSvcPromptRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'engineParameters': json['engineParameters'] == null ? undefined : PromptSvcEngineParametersFromJSON(json['engineParameters']),
        'id': json['id'] == null ? undefined : json['id'],
        'maxRetries': json['maxRetries'] == null ? undefined : json['maxRetries'],
        'modelId': json['modelId'] == null ? undefined : json['modelId'],
        'parameters': json['parameters'] == null ? undefined : PromptSvcParametersFromJSON(json['parameters']),
        'prompt': json['prompt'],
        'sync': json['sync'] == null ? undefined : json['sync'],
        'threadId': json['threadId'] == null ? undefined : json['threadId'],
    };
}
export function PromptSvcPromptRequestToJSON(json) {
    return PromptSvcPromptRequestToJSONTyped(json, false);
}
export function PromptSvcPromptRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'engineParameters': PromptSvcEngineParametersToJSON(value['engineParameters']),
        'id': value['id'],
        'maxRetries': value['maxRetries'],
        'modelId': value['modelId'],
        'parameters': PromptSvcParametersToJSON(value['parameters']),
        'prompt': value['prompt'],
        'sync': value['sync'],
        'threadId': value['threadId'],
    };
}
