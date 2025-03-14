/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
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
import type { ContainerSvcPortMapping } from './ContainerSvcPortMapping';
import {
    ContainerSvcPortMappingFromJSON,
    ContainerSvcPortMappingFromJSONTyped,
    ContainerSvcPortMappingToJSON,
    ContainerSvcPortMappingToJSONTyped,
} from './ContainerSvcPortMapping';

/**
 * 
 * @export
 * @interface ContainerSvcRunContainerResponse
 */
export interface ContainerSvcRunContainerResponse {
    /**
     * Ports is returned here as host ports might get mapped dynamically.
     * @type {Array<ContainerSvcPortMapping>}
     * @memberof ContainerSvcRunContainerResponse
     */
    ports?: Array<ContainerSvcPortMapping>;
    /**
     * 
     * @type {boolean}
     * @memberof ContainerSvcRunContainerResponse
     */
    started?: boolean;
}

/**
 * Check if a given object implements the ContainerSvcRunContainerResponse interface.
 */
export function instanceOfContainerSvcRunContainerResponse(value: object): value is ContainerSvcRunContainerResponse {
    return true;
}

export function ContainerSvcRunContainerResponseFromJSON(json: any): ContainerSvcRunContainerResponse {
    return ContainerSvcRunContainerResponseFromJSONTyped(json, false);
}

export function ContainerSvcRunContainerResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcRunContainerResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'ports': json['ports'] == null ? undefined : ((json['ports'] as Array<any>).map(ContainerSvcPortMappingFromJSON)),
        'started': json['started'] == null ? undefined : json['started'],
    };
}

export function ContainerSvcRunContainerResponseToJSON(json: any): ContainerSvcRunContainerResponse {
    return ContainerSvcRunContainerResponseToJSONTyped(json, false);
}

export function ContainerSvcRunContainerResponseToJSONTyped(value?: ContainerSvcRunContainerResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'ports': value['ports'] == null ? undefined : ((value['ports'] as Array<any>).map(ContainerSvcPortMappingToJSON)),
        'started': value['started'],
    };
}

