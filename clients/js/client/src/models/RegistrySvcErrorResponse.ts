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
 * @interface RegistrySvcErrorResponse
 */
export interface RegistrySvcErrorResponse {
    /**
     * 
     * @type {string}
     * @memberof RegistrySvcErrorResponse
     */
    error?: string;
}

/**
 * Check if a given object implements the RegistrySvcErrorResponse interface.
 */
export function instanceOfRegistrySvcErrorResponse(value: object): value is RegistrySvcErrorResponse {
    return true;
}

export function RegistrySvcErrorResponseFromJSON(json: any): RegistrySvcErrorResponse {
    return RegistrySvcErrorResponseFromJSONTyped(json, false);
}

export function RegistrySvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcErrorResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'error': json['error'] == null ? undefined : json['error'],
    };
}

export function RegistrySvcErrorResponseToJSON(json: any): RegistrySvcErrorResponse {
    return RegistrySvcErrorResponseToJSONTyped(json, false);
}

export function RegistrySvcErrorResponseToJSONTyped(value?: RegistrySvcErrorResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'error': value['error'],
    };
}

