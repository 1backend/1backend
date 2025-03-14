/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
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
 * @interface ContainerSvcResources
 */
export interface ContainerSvcResources {
    /**
     * CPU cores allocated to the container (e.g., 0.5 = 500m, 2 = 2 cores).
     * @type {number}
     * @memberof ContainerSvcResources
     */
    cpu?: number;
    /**
     * Disk space allocated to the container in megabytes.
     * @type {number}
     * @memberof ContainerSvcResources
     */
    diskMB?: number;
    /**
     * Memory allocated to the container in megabytes.
     * @type {number}
     * @memberof ContainerSvcResources
     */
    memoryMB?: number;
}

/**
 * Check if a given object implements the ContainerSvcResources interface.
 */
export function instanceOfContainerSvcResources(value: object): value is ContainerSvcResources {
    return true;
}

export function ContainerSvcResourcesFromJSON(json: any): ContainerSvcResources {
    return ContainerSvcResourcesFromJSONTyped(json, false);
}

export function ContainerSvcResourcesFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcResources {
    if (json == null) {
        return json;
    }
    return {
        
        'cpu': json['cpu'] == null ? undefined : json['cpu'],
        'diskMB': json['diskMB'] == null ? undefined : json['diskMB'],
        'memoryMB': json['memoryMB'] == null ? undefined : json['memoryMB'],
    };
}

export function ContainerSvcResourcesToJSON(json: any): ContainerSvcResources {
    return ContainerSvcResourcesToJSONTyped(json, false);
}

export function ContainerSvcResourcesToJSONTyped(value?: ContainerSvcResources | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cpu': value['cpu'],
        'diskMB': value['diskMB'],
        'memoryMB': value['memoryMB'],
    };
}

