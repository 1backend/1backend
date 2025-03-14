/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { RegistrySvcProcess } from './RegistrySvcProcess';
import {
    RegistrySvcProcessFromJSON,
    RegistrySvcProcessFromJSONTyped,
    RegistrySvcProcessToJSON,
    RegistrySvcProcessToJSONTyped,
} from './RegistrySvcProcess';

/**
 * 
 * @export
 * @interface RegistrySvcGPU
 */
export interface RegistrySvcGPU {
    /**
     * 
     * @type {string}
     * @memberof RegistrySvcGPU
     */
    busId?: string;
    /**
     * 
     * @type {string}
     * @memberof RegistrySvcGPU
     */
    computeMode?: string;
    /**
     * 
     * @type {string}
     * @memberof RegistrySvcGPU
     */
    cudaVersion?: string;
    /**
     * 
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    gpuUtilization?: number;
    /**
     * Id Node.URL + IntraNodeId
     * @type {string}
     * @memberof RegistrySvcGPU
     */
    id?: string;
    /**
     * 
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    intraNodeId?: number;
    /**
     * 
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    memoryTotal?: number;
    /**
     * 
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    memoryUsage?: number;
    /**
     * 
     * @type {string}
     * @memberof RegistrySvcGPU
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof RegistrySvcGPU
     */
    performanceState?: string;
    /**
     * 
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    powerCap?: number;
    /**
     * 
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    powerUsage?: number;
    /**
     * 
     * @type {Array<RegistrySvcProcess>}
     * @memberof RegistrySvcGPU
     */
    processDetails?: Array<RegistrySvcProcess>;
    /**
     * 
     * @type {number}
     * @memberof RegistrySvcGPU
     */
    temperature?: number;
}

/**
 * Check if a given object implements the RegistrySvcGPU interface.
 */
export function instanceOfRegistrySvcGPU(value: object): value is RegistrySvcGPU {
    return true;
}

export function RegistrySvcGPUFromJSON(json: any): RegistrySvcGPU {
    return RegistrySvcGPUFromJSONTyped(json, false);
}

export function RegistrySvcGPUFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcGPU {
    if (json == null) {
        return json;
    }
    return {
        
        'busId': json['busId'] == null ? undefined : json['busId'],
        'computeMode': json['computeMode'] == null ? undefined : json['computeMode'],
        'cudaVersion': json['cudaVersion'] == null ? undefined : json['cudaVersion'],
        'gpuUtilization': json['gpuUtilization'] == null ? undefined : json['gpuUtilization'],
        'id': json['id'] == null ? undefined : json['id'],
        'intraNodeId': json['intraNodeId'] == null ? undefined : json['intraNodeId'],
        'memoryTotal': json['memoryTotal'] == null ? undefined : json['memoryTotal'],
        'memoryUsage': json['memoryUsage'] == null ? undefined : json['memoryUsage'],
        'name': json['name'] == null ? undefined : json['name'],
        'performanceState': json['performanceState'] == null ? undefined : json['performanceState'],
        'powerCap': json['powerCap'] == null ? undefined : json['powerCap'],
        'powerUsage': json['powerUsage'] == null ? undefined : json['powerUsage'],
        'processDetails': json['processDetails'] == null ? undefined : ((json['processDetails'] as Array<any>).map(RegistrySvcProcessFromJSON)),
        'temperature': json['temperature'] == null ? undefined : json['temperature'],
    };
}

export function RegistrySvcGPUToJSON(json: any): RegistrySvcGPU {
    return RegistrySvcGPUToJSONTyped(json, false);
}

export function RegistrySvcGPUToJSONTyped(value?: RegistrySvcGPU | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'busId': value['busId'],
        'computeMode': value['computeMode'],
        'cudaVersion': value['cudaVersion'],
        'gpuUtilization': value['gpuUtilization'],
        'id': value['id'],
        'intraNodeId': value['intraNodeId'],
        'memoryTotal': value['memoryTotal'],
        'memoryUsage': value['memoryUsage'],
        'name': value['name'],
        'performanceState': value['performanceState'],
        'powerCap': value['powerCap'],
        'powerUsage': value['powerUsage'],
        'processDetails': value['processDetails'] == null ? undefined : ((value['processDetails'] as Array<any>).map(RegistrySvcProcessToJSON)),
        'temperature': value['temperature'],
    };
}

