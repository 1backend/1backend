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
import { PromptSvcStreamChunkFromJSON, PromptSvcStreamChunkToJSON, } from './PromptSvcStreamChunk';
/**
 * Check if a given object implements the PromptSvcTypesResponse interface.
 */
export function instanceOfPromptSvcTypesResponse(value) {
    return true;
}
export function PromptSvcTypesResponseFromJSON(json) {
    return PromptSvcTypesResponseFromJSONTyped(json, false);
}
export function PromptSvcTypesResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'chunk': json['chunk'] == null ? undefined : PromptSvcStreamChunkFromJSON(json['chunk']),
    };
}
export function PromptSvcTypesResponseToJSON(json) {
    return PromptSvcTypesResponseToJSONTyped(json, false);
}
export function PromptSvcTypesResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'chunk': PromptSvcStreamChunkToJSON(value['chunk']),
    };
}
