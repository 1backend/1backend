/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
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
 * @interface DeploySvcAutoScalingConfig
 */
export interface DeploySvcAutoScalingConfig {
    /**
     * CPU usage threshold for scaling (as a percentage)
     * @type {number}
     * @memberof DeploySvcAutoScalingConfig
     */
    cpuThreshold?: number;
    /**
     * Maximum number of replicas to run
     * @type {number}
     * @memberof DeploySvcAutoScalingConfig
     */
    maxReplicas?: number;
    /**
     * Minimum number of replicas to run
     * @type {number}
     * @memberof DeploySvcAutoScalingConfig
     */
    minReplicas?: number;
}

/**
 * Check if a given object implements the DeploySvcAutoScalingConfig interface.
 */
export function instanceOfDeploySvcAutoScalingConfig(value: object): value is DeploySvcAutoScalingConfig {
    return true;
}

export function DeploySvcAutoScalingConfigFromJSON(json: any): DeploySvcAutoScalingConfig {
    return DeploySvcAutoScalingConfigFromJSONTyped(json, false);
}

export function DeploySvcAutoScalingConfigFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcAutoScalingConfig {
    if (json == null) {
        return json;
    }
    return {
        
        'cpuThreshold': json['cpuThreshold'] == null ? undefined : json['cpuThreshold'],
        'maxReplicas': json['maxReplicas'] == null ? undefined : json['maxReplicas'],
        'minReplicas': json['minReplicas'] == null ? undefined : json['minReplicas'],
    };
}

export function DeploySvcAutoScalingConfigToJSON(json: any): DeploySvcAutoScalingConfig {
    return DeploySvcAutoScalingConfigToJSONTyped(json, false);
}

export function DeploySvcAutoScalingConfigToJSONTyped(value?: DeploySvcAutoScalingConfig | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cpuThreshold': value['cpuThreshold'],
        'maxReplicas': value['maxReplicas'],
        'minReplicas': value['minReplicas'],
    };
}

