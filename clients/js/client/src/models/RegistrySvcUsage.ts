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
 * @interface RegistrySvcUsage
 */
export interface RegistrySvcUsage {
    /**
     * Usage percentage.
     * @type {number}
     * @memberof RegistrySvcUsage
     */
    percent?: number;
    /**
     * Total available amount (in bytes).
     * @type {number}
     * @memberof RegistrySvcUsage
     */
    total?: number;
    /**
     * Used amount (in bytes).
     * @type {number}
     * @memberof RegistrySvcUsage
     */
    used?: number;
}

/**
 * Check if a given object implements the RegistrySvcUsage interface.
 */
export function instanceOfRegistrySvcUsage(value: object): value is RegistrySvcUsage {
    return true;
}

export function RegistrySvcUsageFromJSON(json: any): RegistrySvcUsage {
    return RegistrySvcUsageFromJSONTyped(json, false);
}

export function RegistrySvcUsageFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcUsage {
    if (json == null) {
        return json;
    }
    return {
        
        'percent': json['percent'] == null ? undefined : json['percent'],
        'total': json['total'] == null ? undefined : json['total'],
        'used': json['used'] == null ? undefined : json['used'],
    };
}

export function RegistrySvcUsageToJSON(json: any): RegistrySvcUsage {
    return RegistrySvcUsageToJSONTyped(json, false);
}

export function RegistrySvcUsageToJSONTyped(value?: RegistrySvcUsage | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'percent': value['percent'],
        'total': value['total'],
        'used': value['used'],
    };
}

