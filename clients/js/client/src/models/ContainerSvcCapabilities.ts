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
 * @interface ContainerSvcCapabilities
 */
export interface ContainerSvcCapabilities {
    /**
     * GPUEnabled specifies whether GPU support is enabled for the container.
     * @type {boolean}
     * @memberof ContainerSvcCapabilities
     */
    gpuEnabled?: boolean;
}

/**
 * Check if a given object implements the ContainerSvcCapabilities interface.
 */
export function instanceOfContainerSvcCapabilities(value: object): value is ContainerSvcCapabilities {
    return true;
}

export function ContainerSvcCapabilitiesFromJSON(json: any): ContainerSvcCapabilities {
    return ContainerSvcCapabilitiesFromJSONTyped(json, false);
}

export function ContainerSvcCapabilitiesFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcCapabilities {
    if (json == null) {
        return json;
    }
    return {
        
        'gpuEnabled': json['gpuEnabled'] == null ? undefined : json['gpuEnabled'],
    };
}

export function ContainerSvcCapabilitiesToJSON(json: any): ContainerSvcCapabilities {
    return ContainerSvcCapabilitiesToJSONTyped(json, false);
}

export function ContainerSvcCapabilitiesToJSONTyped(value?: ContainerSvcCapabilities | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gpuEnabled': value['gpuEnabled'],
    };
}

