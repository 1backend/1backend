/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { DatastoreQueryFromJSON, DatastoreQueryToJSON, } from './DatastoreQuery';
/**
 * Check if a given object implements the PromptSvcListPromptsRequest interface.
 */
export function instanceOfPromptSvcListPromptsRequest(value) {
    return true;
}
export function PromptSvcListPromptsRequestFromJSON(json) {
    return PromptSvcListPromptsRequestFromJSONTyped(json, false);
}
export function PromptSvcListPromptsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'query': json['query'] == null ? undefined : DatastoreQueryFromJSON(json['query']),
    };
}
export function PromptSvcListPromptsRequestToJSON(json) {
    return PromptSvcListPromptsRequestToJSONTyped(json, false);
}
export function PromptSvcListPromptsRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'query': DatastoreQueryToJSON(value['query']),
    };
}
