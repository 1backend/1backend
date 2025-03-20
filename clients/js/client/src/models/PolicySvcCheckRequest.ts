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
/**
 * 
 * @export
 * @interface PolicySvcCheckRequest
 */
export interface PolicySvcCheckRequest {
    /**
     * 
     * @type {string}
     * @memberof PolicySvcCheckRequest
     */
    endpoint?: string;
    /**
     * 
     * @type {string}
     * @memberof PolicySvcCheckRequest
     */
    ip?: string;
    /**
     * 
     * @type {string}
     * @memberof PolicySvcCheckRequest
     */
    method?: string;
    /**
     * 
     * @type {string}
     * @memberof PolicySvcCheckRequest
     */
    userId?: string;
}

/**
 * Check if a given object implements the PolicySvcCheckRequest interface.
 */
export function instanceOfPolicySvcCheckRequest(value: object): value is PolicySvcCheckRequest {
    return true;
}

export function PolicySvcCheckRequestFromJSON(json: any): PolicySvcCheckRequest {
    return PolicySvcCheckRequestFromJSONTyped(json, false);
}

export function PolicySvcCheckRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcCheckRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'endpoint': json['endpoint'] == null ? undefined : json['endpoint'],
        'ip': json['ip'] == null ? undefined : json['ip'],
        'method': json['method'] == null ? undefined : json['method'],
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}

export function PolicySvcCheckRequestToJSON(json: any): PolicySvcCheckRequest {
    return PolicySvcCheckRequestToJSONTyped(json, false);
}

export function PolicySvcCheckRequestToJSONTyped(value?: PolicySvcCheckRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'endpoint': value['endpoint'],
        'ip': value['ip'],
        'method': value['method'],
        'userId': value['userId'],
    };
}

