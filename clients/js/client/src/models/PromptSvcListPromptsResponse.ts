/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { PromptSvcPrompt } from './PromptSvcPrompt';
import {
    PromptSvcPromptFromJSON,
    PromptSvcPromptFromJSONTyped,
    PromptSvcPromptToJSON,
    PromptSvcPromptToJSONTyped,
} from './PromptSvcPrompt';

/**
 * 
 * @export
 * @interface PromptSvcListPromptsResponse
 */
export interface PromptSvcListPromptsResponse {
    /**
     * 
     * @type {object}
     * @memberof PromptSvcListPromptsResponse
     */
    after?: object;
    /**
     * 
     * @type {number}
     * @memberof PromptSvcListPromptsResponse
     */
    count?: number;
    /**
     * 
     * @type {Array<PromptSvcPrompt>}
     * @memberof PromptSvcListPromptsResponse
     */
    prompts?: Array<PromptSvcPrompt>;
}

/**
 * Check if a given object implements the PromptSvcListPromptsResponse interface.
 */
export function instanceOfPromptSvcListPromptsResponse(value: object): value is PromptSvcListPromptsResponse {
    return true;
}

export function PromptSvcListPromptsResponseFromJSON(json: any): PromptSvcListPromptsResponse {
    return PromptSvcListPromptsResponseFromJSONTyped(json, false);
}

export function PromptSvcListPromptsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcListPromptsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'after': json['after'] == null ? undefined : json['after'],
        'count': json['count'] == null ? undefined : json['count'],
        'prompts': json['prompts'] == null ? undefined : ((json['prompts'] as Array<any>).map(PromptSvcPromptFromJSON)),
    };
}

export function PromptSvcListPromptsResponseToJSON(json: any): PromptSvcListPromptsResponse {
    return PromptSvcListPromptsResponseToJSONTyped(json, false);
}

export function PromptSvcListPromptsResponseToJSONTyped(value?: PromptSvcListPromptsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'after': value['after'],
        'count': value['count'],
        'prompts': value['prompts'] == null ? undefined : ((value['prompts'] as Array<any>).map(PromptSvcPromptToJSON)),
    };
}

