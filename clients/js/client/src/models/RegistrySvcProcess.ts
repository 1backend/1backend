/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
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
 * @interface RegistrySvcProcess
 */
export interface RegistrySvcProcess {
    /**
     * 
     * @type {number}
     * @memberof RegistrySvcProcess
     */
    memoryUsage?: number;
    /**
     * 
     * @type {number}
     * @memberof RegistrySvcProcess
     */
    pid?: number;
    /**
     * 
     * @type {string}
     * @memberof RegistrySvcProcess
     */
    processName?: string;
}

/**
 * Check if a given object implements the RegistrySvcProcess interface.
 */
export function instanceOfRegistrySvcProcess(value: object): value is RegistrySvcProcess {
    return true;
}

export function RegistrySvcProcessFromJSON(json: any): RegistrySvcProcess {
    return RegistrySvcProcessFromJSONTyped(json, false);
}

export function RegistrySvcProcessFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcProcess {
    if (json == null) {
        return json;
    }
    return {
        
        'memoryUsage': json['memoryUsage'] == null ? undefined : json['memoryUsage'],
        'pid': json['pid'] == null ? undefined : json['pid'],
        'processName': json['processName'] == null ? undefined : json['processName'],
    };
}

export function RegistrySvcProcessToJSON(json: any): RegistrySvcProcess {
    return RegistrySvcProcessToJSONTyped(json, false);
}

export function RegistrySvcProcessToJSONTyped(value?: RegistrySvcProcess | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'memoryUsage': value['memoryUsage'],
        'pid': value['pid'],
        'processName': value['processName'],
    };
}

