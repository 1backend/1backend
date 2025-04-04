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


/**
 * 
 * @export
 */
export const PromptSvcStreamChunkType = {
    ChunkTypeProgress: 'progress',
    ChunkTypeDone: 'done'
} as const;
export type PromptSvcStreamChunkType = typeof PromptSvcStreamChunkType[keyof typeof PromptSvcStreamChunkType];


export function instanceOfPromptSvcStreamChunkType(value: any): boolean {
    for (const key in PromptSvcStreamChunkType) {
        if (Object.prototype.hasOwnProperty.call(PromptSvcStreamChunkType, key)) {
            if (PromptSvcStreamChunkType[key as keyof typeof PromptSvcStreamChunkType] === value) {
                return true;
            }
        }
    }
    return false;
}

export function PromptSvcStreamChunkTypeFromJSON(json: any): PromptSvcStreamChunkType {
    return PromptSvcStreamChunkTypeFromJSONTyped(json, false);
}

export function PromptSvcStreamChunkTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcStreamChunkType {
    return json as PromptSvcStreamChunkType;
}

export function PromptSvcStreamChunkTypeToJSON(value?: PromptSvcStreamChunkType | null): any {
    return value as any;
}

export function PromptSvcStreamChunkTypeToJSONTyped(value: any, ignoreDiscriminator: boolean): PromptSvcStreamChunkType {
    return value as PromptSvcStreamChunkType;
}

