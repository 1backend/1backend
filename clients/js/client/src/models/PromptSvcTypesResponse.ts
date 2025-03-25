/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { PromptSvcStreamChunk } from './PromptSvcStreamChunk';
import {
    PromptSvcStreamChunkFromJSON,
    PromptSvcStreamChunkFromJSONTyped,
    PromptSvcStreamChunkToJSON,
    PromptSvcStreamChunkToJSONTyped,
} from './PromptSvcStreamChunk';

/**
 * 
 * @export
 * @interface PromptSvcTypesResponse
 */
export interface PromptSvcTypesResponse {
    /**
     * 
     * @type {PromptSvcStreamChunk}
     * @memberof PromptSvcTypesResponse
     */
    chunk?: PromptSvcStreamChunk;
}

/**
 * Check if a given object implements the PromptSvcTypesResponse interface.
 */
export function instanceOfPromptSvcTypesResponse(value: object): value is PromptSvcTypesResponse {
    return true;
}

export function PromptSvcTypesResponseFromJSON(json: any): PromptSvcTypesResponse {
    return PromptSvcTypesResponseFromJSONTyped(json, false);
}

export function PromptSvcTypesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcTypesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'chunk': json['chunk'] == null ? undefined : PromptSvcStreamChunkFromJSON(json['chunk']),
    };
}

export function PromptSvcTypesResponseToJSON(json: any): PromptSvcTypesResponse {
    return PromptSvcTypesResponseToJSONTyped(json, false);
}

export function PromptSvcTypesResponseToJSONTyped(value?: PromptSvcTypesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chunk': PromptSvcStreamChunkToJSON(value['chunk']),
    };
}

