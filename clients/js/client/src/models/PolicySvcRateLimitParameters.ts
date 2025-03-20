/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { PolicySvcScope } from './PolicySvcScope';
import {
    PolicySvcScopeFromJSON,
    PolicySvcScopeFromJSONTyped,
    PolicySvcScopeToJSON,
    PolicySvcScopeToJSONTyped,
} from './PolicySvcScope';
import type { PolicySvcEntity } from './PolicySvcEntity';
import {
    PolicySvcEntityFromJSON,
    PolicySvcEntityFromJSONTyped,
    PolicySvcEntityToJSON,
    PolicySvcEntityToJSONTyped,
} from './PolicySvcEntity';

/**
 * 
 * @export
 * @interface PolicySvcRateLimitParameters
 */
export interface PolicySvcRateLimitParameters {
    /**
     * 
     * @type {PolicySvcEntity}
     * @memberof PolicySvcRateLimitParameters
     */
    entity?: PolicySvcEntity;
    /**
     * 
     * @type {number}
     * @memberof PolicySvcRateLimitParameters
     */
    maxRequests?: number;
    /**
     * 
     * @type {PolicySvcScope}
     * @memberof PolicySvcRateLimitParameters
     */
    scope?: PolicySvcScope;
    /**
     * 
     * @type {string}
     * @memberof PolicySvcRateLimitParameters
     */
    timeWindow?: string;
}



/**
 * Check if a given object implements the PolicySvcRateLimitParameters interface.
 */
export function instanceOfPolicySvcRateLimitParameters(value: object): value is PolicySvcRateLimitParameters {
    return true;
}

export function PolicySvcRateLimitParametersFromJSON(json: any): PolicySvcRateLimitParameters {
    return PolicySvcRateLimitParametersFromJSONTyped(json, false);
}

export function PolicySvcRateLimitParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcRateLimitParameters {
    if (json == null) {
        return json;
    }
    return {
        
        'entity': json['entity'] == null ? undefined : PolicySvcEntityFromJSON(json['entity']),
        'maxRequests': json['maxRequests'] == null ? undefined : json['maxRequests'],
        'scope': json['scope'] == null ? undefined : PolicySvcScopeFromJSON(json['scope']),
        'timeWindow': json['timeWindow'] == null ? undefined : json['timeWindow'],
    };
}

export function PolicySvcRateLimitParametersToJSON(json: any): PolicySvcRateLimitParameters {
    return PolicySvcRateLimitParametersToJSONTyped(json, false);
}

export function PolicySvcRateLimitParametersToJSONTyped(value?: PolicySvcRateLimitParameters | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'entity': PolicySvcEntityToJSON(value['entity']),
        'maxRequests': value['maxRequests'],
        'scope': PolicySvcScopeToJSON(value['scope']),
        'timeWindow': value['timeWindow'],
    };
}

