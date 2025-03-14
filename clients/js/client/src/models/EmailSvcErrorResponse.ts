/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
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
 * @interface EmailSvcErrorResponse
 */
export interface EmailSvcErrorResponse {
    /**
     * 
     * @type {string}
     * @memberof EmailSvcErrorResponse
     */
    error?: string;
}

/**
 * Check if a given object implements the EmailSvcErrorResponse interface.
 */
export function instanceOfEmailSvcErrorResponse(value: object): value is EmailSvcErrorResponse {
    return true;
}

export function EmailSvcErrorResponseFromJSON(json: any): EmailSvcErrorResponse {
    return EmailSvcErrorResponseFromJSONTyped(json, false);
}

export function EmailSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmailSvcErrorResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'error': json['error'] == null ? undefined : json['error'],
    };
}

export function EmailSvcErrorResponseToJSON(json: any): EmailSvcErrorResponse {
    return EmailSvcErrorResponseToJSONTyped(json, false);
}

export function EmailSvcErrorResponseToJSONTyped(value?: EmailSvcErrorResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'error': value['error'],
    };
}

