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
import type { ContainerSvcContainer } from './ContainerSvcContainer';
import {
    ContainerSvcContainerFromJSON,
    ContainerSvcContainerFromJSONTyped,
    ContainerSvcContainerToJSON,
    ContainerSvcContainerToJSONTyped,
} from './ContainerSvcContainer';

/**
 * 
 * @export
 * @interface ContainerSvcListContainersResponse
 */
export interface ContainerSvcListContainersResponse {
    /**
     * 
     * @type {Array<ContainerSvcContainer>}
     * @memberof ContainerSvcListContainersResponse
     */
    containers?: Array<ContainerSvcContainer>;
}

/**
 * Check if a given object implements the ContainerSvcListContainersResponse interface.
 */
export function instanceOfContainerSvcListContainersResponse(value: object): value is ContainerSvcListContainersResponse {
    return true;
}

export function ContainerSvcListContainersResponseFromJSON(json: any): ContainerSvcListContainersResponse {
    return ContainerSvcListContainersResponseFromJSONTyped(json, false);
}

export function ContainerSvcListContainersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcListContainersResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'containers': json['containers'] == null ? undefined : ((json['containers'] as Array<any>).map(ContainerSvcContainerFromJSON)),
    };
}

export function ContainerSvcListContainersResponseToJSON(json: any): ContainerSvcListContainersResponse {
    return ContainerSvcListContainersResponseToJSONTyped(json, false);
}

export function ContainerSvcListContainersResponseToJSONTyped(value?: ContainerSvcListContainersResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'containers': value['containers'] == null ? undefined : ((value['containers'] as Array<any>).map(ContainerSvcContainerToJSON)),
    };
}

