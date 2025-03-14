/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { PromptSvcStreamChunkType } from './promptSvcStreamChunkType';
export declare class PromptSvcStreamChunk {
    /**
    * MessageId is the ChatSvc Message id that the chunk is part of. Might only be available for \"done\" chunks.
    */
    'messageId'?: string;
    /**
    * TextChunk contains a part of the text output from the stream.
    */
    'text'?: string;
    /**
    * Type indicates the type of the stream event (e.g., text, done).
    */
    'type'?: PromptSvcStreamChunkType;
    static discriminator: string | undefined;
    static attributeTypeMap: Array<{
        name: string;
        baseName: string;
        type: string;
    }>;
    static getAttributeTypeMap(): {
        name: string;
        baseName: string;
        type: string;
    }[];
}
export declare namespace PromptSvcStreamChunk {
}
