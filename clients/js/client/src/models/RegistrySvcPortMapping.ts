/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
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
 * @interface RegistrySvcPortMapping
 */
export interface RegistrySvcPortMapping {
    /**
     * 
     * @type {number}
     * @memberof RegistrySvcPortMapping
     */
    host: number;
    /**
     * 
     * @type {number}
     * @memberof RegistrySvcPortMapping
     */
    internal: number;
}

/**
 * Check if a given object implements the RegistrySvcPortMapping interface.
 */
export function instanceOfRegistrySvcPortMapping(value: object): value is RegistrySvcPortMapping {
    if (!('host' in value) || value['host'] === undefined) return false;
    if (!('internal' in value) || value['internal'] === undefined) return false;
    return true;
}

export function RegistrySvcPortMappingFromJSON(json: any): RegistrySvcPortMapping {
    return RegistrySvcPortMappingFromJSONTyped(json, false);
}

export function RegistrySvcPortMappingFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcPortMapping {
    if (json == null) {
        return json;
    }
    return {
        
        'host': json['host'],
        'internal': json['internal'],
    };
}

export function RegistrySvcPortMappingToJSON(json: any): RegistrySvcPortMapping {
    return RegistrySvcPortMappingToJSONTyped(json, false);
}

export function RegistrySvcPortMappingToJSONTyped(value?: RegistrySvcPortMapping | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'host': value['host'],
        'internal': value['internal'],
    };
}

