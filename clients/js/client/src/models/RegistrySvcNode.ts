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
import type { RegistrySvcGPU } from './RegistrySvcGPU';
import {
    RegistrySvcGPUFromJSON,
    RegistrySvcGPUFromJSONTyped,
    RegistrySvcGPUToJSON,
    RegistrySvcGPUToJSONTyped,
} from './RegistrySvcGPU';
import type { RegistrySvcResourceUsage } from './RegistrySvcResourceUsage';
import {
    RegistrySvcResourceUsageFromJSON,
    RegistrySvcResourceUsageFromJSONTyped,
    RegistrySvcResourceUsageToJSON,
    RegistrySvcResourceUsageToJSONTyped,
} from './RegistrySvcResourceUsage';

/**
 * 
 * @export
 * @interface RegistrySvcNode
 */
export interface RegistrySvcNode {
    /**
     * The availability zone of the node
     * @type {string}
     * @memberof RegistrySvcNode
     */
    availabilityZone?: string;
    /**
     * List of GPUs available on the node
     * @type {Array<RegistrySvcGPU>}
     * @memberof RegistrySvcNode
     */
    gpus?: Array<RegistrySvcGPU>;
    /**
     * Required: ID of the instance
     * @type {string}
     * @memberof RegistrySvcNode
     */
    id: string;
    /**
     * Last time the instance gave a sign of life
     * @type {string}
     * @memberof RegistrySvcNode
     */
    lastHeartbeat?: string;
    /**
     * The region of the node
     * @type {string}
     * @memberof RegistrySvcNode
     */
    region?: string;
    /**
     * URL of the daemon running on the node.
     * If not configured defaults to hostname + default 1Backend server port.
     * @type {string}
     * @memberof RegistrySvcNode
     */
    url: string;
    /**
     * Resource usage metrics of the node.
     * @type {RegistrySvcResourceUsage}
     * @memberof RegistrySvcNode
     */
    usage?: RegistrySvcResourceUsage;
}

/**
 * Check if a given object implements the RegistrySvcNode interface.
 */
export function instanceOfRegistrySvcNode(value: object): value is RegistrySvcNode {
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('url' in value) || value['url'] === undefined) return false;
    return true;
}

export function RegistrySvcNodeFromJSON(json: any): RegistrySvcNode {
    return RegistrySvcNodeFromJSONTyped(json, false);
}

export function RegistrySvcNodeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcNode {
    if (json == null) {
        return json;
    }
    return {
        
        'availabilityZone': json['availabilityZone'] == null ? undefined : json['availabilityZone'],
        'gpus': json['gpus'] == null ? undefined : ((json['gpus'] as Array<any>).map(RegistrySvcGPUFromJSON)),
        'id': json['id'],
        'lastHeartbeat': json['lastHeartbeat'] == null ? undefined : json['lastHeartbeat'],
        'region': json['region'] == null ? undefined : json['region'],
        'url': json['url'],
        'usage': json['usage'] == null ? undefined : RegistrySvcResourceUsageFromJSON(json['usage']),
    };
}

export function RegistrySvcNodeToJSON(json: any): RegistrySvcNode {
    return RegistrySvcNodeToJSONTyped(json, false);
}

export function RegistrySvcNodeToJSONTyped(value?: RegistrySvcNode | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'availabilityZone': value['availabilityZone'],
        'gpus': value['gpus'] == null ? undefined : ((value['gpus'] as Array<any>).map(RegistrySvcGPUToJSON)),
        'id': value['id'],
        'lastHeartbeat': value['lastHeartbeat'],
        'region': value['region'],
        'url': value['url'],
        'usage': RegistrySvcResourceUsageToJSON(value['usage']),
    };
}

