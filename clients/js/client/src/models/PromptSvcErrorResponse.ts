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
/**
 * 
 * @export
 * @interface PromptSvcErrorResponse
 */
export interface PromptSvcErrorResponse {
    /**
     * 
     * @type {string}
     * @memberof PromptSvcErrorResponse
     */
    error?: string;
}

/**
 * Check if a given object implements the PromptSvcErrorResponse interface.
 */
export function instanceOfPromptSvcErrorResponse(value: object): value is PromptSvcErrorResponse {
    return true;
}

export function PromptSvcErrorResponseFromJSON(json: any): PromptSvcErrorResponse {
    return PromptSvcErrorResponseFromJSONTyped(json, false);
}

export function PromptSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcErrorResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'error': json['error'] == null ? undefined : json['error'],
    };
}

export function PromptSvcErrorResponseToJSON(json: any): PromptSvcErrorResponse {
    return PromptSvcErrorResponseToJSONTyped(json, false);
}

export function PromptSvcErrorResponseToJSONTyped(value?: PromptSvcErrorResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'error': value['error'],
    };
}

