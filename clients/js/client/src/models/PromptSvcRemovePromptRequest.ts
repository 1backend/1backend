/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface PromptSvcRemovePromptRequest
 */
export interface PromptSvcRemovePromptRequest {
    /**
     * 
     * @type {string}
     * @memberof PromptSvcRemovePromptRequest
     */
    promptId?: string;
}

/**
 * Check if a given object implements the PromptSvcRemovePromptRequest interface.
 */
export function instanceOfPromptSvcRemovePromptRequest(value: object): value is PromptSvcRemovePromptRequest {
    return true;
}

export function PromptSvcRemovePromptRequestFromJSON(json: any): PromptSvcRemovePromptRequest {
    return PromptSvcRemovePromptRequestFromJSONTyped(json, false);
}

export function PromptSvcRemovePromptRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcRemovePromptRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'promptId': json['promptId'] == null ? undefined : json['promptId'],
    };
}

export function PromptSvcRemovePromptRequestToJSON(json: any): PromptSvcRemovePromptRequest {
    return PromptSvcRemovePromptRequestToJSONTyped(json, false);
}

export function PromptSvcRemovePromptRequestToJSONTyped(value?: PromptSvcRemovePromptRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'promptId': value['promptId'],
    };
}

