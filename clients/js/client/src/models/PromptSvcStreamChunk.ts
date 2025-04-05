/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { PromptSvcStreamChunkType } from './PromptSvcStreamChunkType';
import {
    PromptSvcStreamChunkTypeFromJSON,
    PromptSvcStreamChunkTypeFromJSONTyped,
    PromptSvcStreamChunkTypeToJSON,
    PromptSvcStreamChunkTypeToJSONTyped,
} from './PromptSvcStreamChunkType';

/**
 * 
 * @export
 * @interface PromptSvcStreamChunk
 */
export interface PromptSvcStreamChunk {
    /**
     * MessageId is the ChatSvc Message id that the chunk is part of.
     * Might only be available for "done" chunks.
     * @type {string}
     * @memberof PromptSvcStreamChunk
     */
    messageId?: string;
    /**
     * TextChunk contains a part of the text output from the stream.
     * @type {string}
     * @memberof PromptSvcStreamChunk
     */
    text?: string;
    /**
     * Type indicates the type of the stream event (e.g., text, done).
     * @type {PromptSvcStreamChunkType}
     * @memberof PromptSvcStreamChunk
     */
    type?: PromptSvcStreamChunkType;
}



/**
 * Check if a given object implements the PromptSvcStreamChunk interface.
 */
export function instanceOfPromptSvcStreamChunk(value: object): value is PromptSvcStreamChunk {
    return true;
}

export function PromptSvcStreamChunkFromJSON(json: any): PromptSvcStreamChunk {
    return PromptSvcStreamChunkFromJSONTyped(json, false);
}

export function PromptSvcStreamChunkFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcStreamChunk {
    if (json == null) {
        return json;
    }
    return {
        
        'messageId': json['messageId'] == null ? undefined : json['messageId'],
        'text': json['text'] == null ? undefined : json['text'],
        'type': json['type'] == null ? undefined : PromptSvcStreamChunkTypeFromJSON(json['type']),
    };
}

export function PromptSvcStreamChunkToJSON(json: any): PromptSvcStreamChunk {
    return PromptSvcStreamChunkToJSONTyped(json, false);
}

export function PromptSvcStreamChunkToJSONTyped(value?: PromptSvcStreamChunk | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'messageId': value['messageId'],
        'text': value['text'],
        'type': PromptSvcStreamChunkTypeToJSON(value['type']),
    };
}

