/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
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
 * @interface DeploySvcErrorResponse
 */
export interface DeploySvcErrorResponse {
    /**
     * 
     * @type {string}
     * @memberof DeploySvcErrorResponse
     */
    error?: string;
}

/**
 * Check if a given object implements the DeploySvcErrorResponse interface.
 */
export function instanceOfDeploySvcErrorResponse(value: object): value is DeploySvcErrorResponse {
    return true;
}

export function DeploySvcErrorResponseFromJSON(json: any): DeploySvcErrorResponse {
    return DeploySvcErrorResponseFromJSONTyped(json, false);
}

export function DeploySvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcErrorResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'error': json['error'] == null ? undefined : json['error'],
    };
}

export function DeploySvcErrorResponseToJSON(json: any): DeploySvcErrorResponse {
    return DeploySvcErrorResponseToJSONTyped(json, false);
}

export function DeploySvcErrorResponseToJSONTyped(value?: DeploySvcErrorResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'error': value['error'],
    };
}

