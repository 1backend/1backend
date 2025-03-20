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
 * @interface ContainerSvcPortMapping
 */
export interface ContainerSvcPortMapping {
    /**
     * 
     * @type {number}
     * @memberof ContainerSvcPortMapping
     */
    host: number;
    /**
     * 
     * @type {number}
     * @memberof ContainerSvcPortMapping
     */
    internal: number;
}

/**
 * Check if a given object implements the ContainerSvcPortMapping interface.
 */
export function instanceOfContainerSvcPortMapping(value: object): value is ContainerSvcPortMapping {
    if (!('host' in value) || value['host'] === undefined) return false;
    if (!('internal' in value) || value['internal'] === undefined) return false;
    return true;
}

export function ContainerSvcPortMappingFromJSON(json: any): ContainerSvcPortMapping {
    return ContainerSvcPortMappingFromJSONTyped(json, false);
}

export function ContainerSvcPortMappingFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcPortMapping {
    if (json == null) {
        return json;
    }
    return {
        
        'host': json['host'],
        'internal': json['internal'],
    };
}

export function ContainerSvcPortMappingToJSON(json: any): ContainerSvcPortMapping {
    return ContainerSvcPortMappingToJSONTyped(json, false);
}

export function ContainerSvcPortMappingToJSONTyped(value?: ContainerSvcPortMapping | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'host': value['host'],
        'internal': value['internal'],
    };
}

