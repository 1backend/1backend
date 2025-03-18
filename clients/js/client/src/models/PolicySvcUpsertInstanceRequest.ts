/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { PolicySvcInstance } from './PolicySvcInstance';
import {
    PolicySvcInstanceFromJSON,
    PolicySvcInstanceFromJSONTyped,
    PolicySvcInstanceToJSON,
    PolicySvcInstanceToJSONTyped,
} from './PolicySvcInstance';

/**
 * 
 * @export
 * @interface PolicySvcUpsertInstanceRequest
 */
export interface PolicySvcUpsertInstanceRequest {
    /**
     * 
     * @type {PolicySvcInstance}
     * @memberof PolicySvcUpsertInstanceRequest
     */
    instance?: PolicySvcInstance;
}

/**
 * Check if a given object implements the PolicySvcUpsertInstanceRequest interface.
 */
export function instanceOfPolicySvcUpsertInstanceRequest(value: object): value is PolicySvcUpsertInstanceRequest {
    return true;
}

export function PolicySvcUpsertInstanceRequestFromJSON(json: any): PolicySvcUpsertInstanceRequest {
    return PolicySvcUpsertInstanceRequestFromJSONTyped(json, false);
}

export function PolicySvcUpsertInstanceRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcUpsertInstanceRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'instance': json['instance'] == null ? undefined : PolicySvcInstanceFromJSON(json['instance']),
    };
}

export function PolicySvcUpsertInstanceRequestToJSON(json: any): PolicySvcUpsertInstanceRequest {
    return PolicySvcUpsertInstanceRequestToJSONTyped(json, false);
}

export function PolicySvcUpsertInstanceRequestToJSONTyped(value?: PolicySvcUpsertInstanceRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'instance': PolicySvcInstanceToJSON(value['instance']),
    };
}

