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
 * @interface PolicySvcCheckResponse
 */
export interface PolicySvcCheckResponse {
    /**
     * 
     * @type {boolean}
     * @memberof PolicySvcCheckResponse
     */
    allowed: boolean;
}

/**
 * Check if a given object implements the PolicySvcCheckResponse interface.
 */
export function instanceOfPolicySvcCheckResponse(value: object): value is PolicySvcCheckResponse {
    if (!('allowed' in value) || value['allowed'] === undefined) return false;
    return true;
}

export function PolicySvcCheckResponseFromJSON(json: any): PolicySvcCheckResponse {
    return PolicySvcCheckResponseFromJSONTyped(json, false);
}

export function PolicySvcCheckResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcCheckResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'allowed': json['allowed'],
    };
}

export function PolicySvcCheckResponseToJSON(json: any): PolicySvcCheckResponse {
    return PolicySvcCheckResponseToJSONTyped(json, false);
}

export function PolicySvcCheckResponseToJSONTyped(value?: PolicySvcCheckResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'allowed': value['allowed'],
    };
}

