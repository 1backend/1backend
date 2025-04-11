/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
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
 * @interface DataSvcErrorResponse
 */
export interface DataSvcErrorResponse {
    /**
     * 
     * @type {string}
     * @memberof DataSvcErrorResponse
     */
    error?: string;
}

/**
 * Check if a given object implements the DataSvcErrorResponse interface.
 */
export function instanceOfDataSvcErrorResponse(value: object): value is DataSvcErrorResponse {
    return true;
}

export function DataSvcErrorResponseFromJSON(json: any): DataSvcErrorResponse {
    return DataSvcErrorResponseFromJSONTyped(json, false);
}

export function DataSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcErrorResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'error': json['error'] == null ? undefined : json['error'],
    };
}

export function DataSvcErrorResponseToJSON(json: any): DataSvcErrorResponse {
    return DataSvcErrorResponseToJSONTyped(json, false);
}

export function DataSvcErrorResponseToJSONTyped(value?: DataSvcErrorResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'error': value['error'],
    };
}

