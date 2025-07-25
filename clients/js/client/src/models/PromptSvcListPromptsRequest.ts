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

import { mapValues } from '../runtime';
import type { DatastoreQuery } from './DatastoreQuery';
import {
    DatastoreQueryFromJSON,
    DatastoreQueryFromJSONTyped,
    DatastoreQueryToJSON,
    DatastoreQueryToJSONTyped,
} from './DatastoreQuery';

/**
 * 
 * @export
 * @interface PromptSvcListPromptsRequest
 */
export interface PromptSvcListPromptsRequest {
    /**
     * 
     * @type {DatastoreQuery}
     * @memberof PromptSvcListPromptsRequest
     */
    query?: DatastoreQuery;
}

/**
 * Check if a given object implements the PromptSvcListPromptsRequest interface.
 */
export function instanceOfPromptSvcListPromptsRequest(value: object): value is PromptSvcListPromptsRequest {
    return true;
}

export function PromptSvcListPromptsRequestFromJSON(json: any): PromptSvcListPromptsRequest {
    return PromptSvcListPromptsRequestFromJSONTyped(json, false);
}

export function PromptSvcListPromptsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcListPromptsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'query': json['query'] == null ? undefined : DatastoreQueryFromJSON(json['query']),
    };
}

export function PromptSvcListPromptsRequestToJSON(json: any): PromptSvcListPromptsRequest {
    return PromptSvcListPromptsRequestToJSONTyped(json, false);
}

export function PromptSvcListPromptsRequestToJSONTyped(value?: PromptSvcListPromptsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'query': DatastoreQueryToJSON(value['query']),
    };
}

