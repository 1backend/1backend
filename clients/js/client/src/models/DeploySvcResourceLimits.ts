/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
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
 * @interface DeploySvcResourceLimits
 */
export interface DeploySvcResourceLimits {
    /**
     * CPU limit, e.g., "500m" for 0.5 cores
     * @type {string}
     * @memberof DeploySvcResourceLimits
     */
    cpu?: string;
    /**
     * Memory limit, e.g., "128Mi"
     * @type {string}
     * @memberof DeploySvcResourceLimits
     */
    memory?: string;
    /**
     * Optional: GPU VRAM requirement, e.g., "48GB"
     * @type {string}
     * @memberof DeploySvcResourceLimits
     */
    vram?: string;
}

/**
 * Check if a given object implements the DeploySvcResourceLimits interface.
 */
export function instanceOfDeploySvcResourceLimits(value: object): value is DeploySvcResourceLimits {
    return true;
}

export function DeploySvcResourceLimitsFromJSON(json: any): DeploySvcResourceLimits {
    return DeploySvcResourceLimitsFromJSONTyped(json, false);
}

export function DeploySvcResourceLimitsFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcResourceLimits {
    if (json == null) {
        return json;
    }
    return {
        
        'cpu': json['cpu'] == null ? undefined : json['cpu'],
        'memory': json['memory'] == null ? undefined : json['memory'],
        'vram': json['vram'] == null ? undefined : json['vram'],
    };
}

export function DeploySvcResourceLimitsToJSON(json: any): DeploySvcResourceLimits {
    return DeploySvcResourceLimitsToJSONTyped(json, false);
}

export function DeploySvcResourceLimitsToJSONTyped(value?: DeploySvcResourceLimits | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cpu': value['cpu'],
        'memory': value['memory'],
        'vram': value['vram'],
    };
}

