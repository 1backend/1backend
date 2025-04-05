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
/**
 * 
 * @export
 * @interface SourceSvcErrorResponse
 */
export interface SourceSvcErrorResponse {
    /**
     * 
     * @type {string}
     * @memberof SourceSvcErrorResponse
     */
    error?: string;
}

/**
 * Check if a given object implements the SourceSvcErrorResponse interface.
 */
export function instanceOfSourceSvcErrorResponse(value: object): value is SourceSvcErrorResponse {
    return true;
}

export function SourceSvcErrorResponseFromJSON(json: any): SourceSvcErrorResponse {
    return SourceSvcErrorResponseFromJSONTyped(json, false);
}

export function SourceSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): SourceSvcErrorResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'error': json['error'] == null ? undefined : json['error'],
    };
}

export function SourceSvcErrorResponseToJSON(json: any): SourceSvcErrorResponse {
    return SourceSvcErrorResponseToJSONTyped(json, false);
}

export function SourceSvcErrorResponseToJSONTyped(value?: SourceSvcErrorResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'error': value['error'],
    };
}

