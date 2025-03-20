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
import type { PromptSvcStreamChunkType } from './PromptSvcStreamChunkType';
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
export declare function instanceOfPromptSvcStreamChunk(value: object): value is PromptSvcStreamChunk;
export declare function PromptSvcStreamChunkFromJSON(json: any): PromptSvcStreamChunk;
export declare function PromptSvcStreamChunkFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcStreamChunk;
export declare function PromptSvcStreamChunkToJSON(json: any): PromptSvcStreamChunk;
export declare function PromptSvcStreamChunkToJSONTyped(value?: PromptSvcStreamChunk | null, ignoreDiscriminator?: boolean): any;
