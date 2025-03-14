/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
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
 * @interface PolicySvcBlocklistParameters
 */
export interface PolicySvcBlocklistParameters {
    /**
     * 
     * @type {Array<string>}
     * @memberof PolicySvcBlocklistParameters
     */
    blockedIPs?: Array<string>;
}

/**
 * Check if a given object implements the PolicySvcBlocklistParameters interface.
 */
export function instanceOfPolicySvcBlocklistParameters(value: object): value is PolicySvcBlocklistParameters {
    return true;
}

export function PolicySvcBlocklistParametersFromJSON(json: any): PolicySvcBlocklistParameters {
    return PolicySvcBlocklistParametersFromJSONTyped(json, false);
}

export function PolicySvcBlocklistParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcBlocklistParameters {
    if (json == null) {
        return json;
    }
    return {
        
        'blockedIPs': json['blockedIPs'] == null ? undefined : json['blockedIPs'],
    };
}

export function PolicySvcBlocklistParametersToJSON(json: any): PolicySvcBlocklistParameters {
    return PolicySvcBlocklistParametersToJSONTyped(json, false);
}

export function PolicySvcBlocklistParametersToJSONTyped(value?: PolicySvcBlocklistParameters | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'blockedIPs': value['blockedIPs'],
    };
}

