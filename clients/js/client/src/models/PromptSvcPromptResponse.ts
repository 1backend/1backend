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

import { mapValues } from '../runtime';
import type { ChatSvcMessage } from './ChatSvcMessage';
import {
    ChatSvcMessageFromJSON,
    ChatSvcMessageFromJSONTyped,
    ChatSvcMessageToJSON,
    ChatSvcMessageToJSONTyped,
} from './ChatSvcMessage';
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
 * @interface PromptSvcPromptResponse
 */
export interface PromptSvcPromptResponse {
    /**
     * Prompt contains the details of the prompt that was just created by this request.
     * This includes the ID, prompt text, status, and other associated metadata.
     * @type {PromptSvcPrompt}
     * @memberof PromptSvcPromptResponse
     */
    prompt?: PromptSvcPrompt;
    /**
     * Response message contains the response text and files.
     * This field is populated only for synchronous prompts (`sync = true`).
     * For asynchronous prompts, the response will provided in the associated
     * message identified by the `responseMessageId` of the `promptSvc.prompt` object once the prompt completes.
     * @type {ChatSvcMessage}
     * @memberof PromptSvcPromptResponse
     */
    responseMessage?: ChatSvcMessage;
}

/**
 * Check if a given object implements the PromptSvcPromptResponse interface.
 */
export function instanceOfPromptSvcPromptResponse(value: object): value is PromptSvcPromptResponse {
    return true;
}

export function PromptSvcPromptResponseFromJSON(json: any): PromptSvcPromptResponse {
    return PromptSvcPromptResponseFromJSONTyped(json, false);
}

export function PromptSvcPromptResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcPromptResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'prompt': json['prompt'] == null ? undefined : PromptSvcPromptFromJSON(json['prompt']),
        'responseMessage': json['responseMessage'] == null ? undefined : ChatSvcMessageFromJSON(json['responseMessage']),
    };
}

export function PromptSvcPromptResponseToJSON(json: any): PromptSvcPromptResponse {
    return PromptSvcPromptResponseToJSONTyped(json, false);
}

export function PromptSvcPromptResponseToJSONTyped(value?: PromptSvcPromptResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'prompt': PromptSvcPromptToJSON(value['prompt']),
        'responseMessage': ChatSvcMessageToJSON(value['responseMessage']),
    };
}

