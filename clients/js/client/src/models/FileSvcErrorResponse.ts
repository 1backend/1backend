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
/**
 * 
 * @export
 * @interface FileSvcErrorResponse
 */
export interface FileSvcErrorResponse {
    /**
     * 
     * @type {string}
     * @memberof FileSvcErrorResponse
     */
    error?: string;
}

/**
 * Check if a given object implements the FileSvcErrorResponse interface.
 */
export function instanceOfFileSvcErrorResponse(value: object): value is FileSvcErrorResponse {
    return true;
}

export function FileSvcErrorResponseFromJSON(json: any): FileSvcErrorResponse {
    return FileSvcErrorResponseFromJSONTyped(json, false);
}

export function FileSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcErrorResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'error': json['error'] == null ? undefined : json['error'],
    };
}

export function FileSvcErrorResponseToJSON(json: any): FileSvcErrorResponse {
    return FileSvcErrorResponseToJSONTyped(json, false);
}

export function FileSvcErrorResponseToJSONTyped(value?: FileSvcErrorResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'error': value['error'],
    };
}

