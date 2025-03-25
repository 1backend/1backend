/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
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
 * @interface ContainerSvcGetHostResponse
 */
export interface ContainerSvcGetHostResponse {
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcGetHostResponse
     */
    host: string;
}

/**
 * Check if a given object implements the ContainerSvcGetHostResponse interface.
 */
export function instanceOfContainerSvcGetHostResponse(value: object): value is ContainerSvcGetHostResponse {
    if (!('host' in value) || value['host'] === undefined) return false;
    return true;
}

export function ContainerSvcGetHostResponseFromJSON(json: any): ContainerSvcGetHostResponse {
    return ContainerSvcGetHostResponseFromJSONTyped(json, false);
}

export function ContainerSvcGetHostResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcGetHostResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'host': json['host'],
    };
}

export function ContainerSvcGetHostResponseToJSON(json: any): ContainerSvcGetHostResponse {
    return ContainerSvcGetHostResponseToJSONTyped(json, false);
}

export function ContainerSvcGetHostResponseToJSONTyped(value?: ContainerSvcGetHostResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'host': value['host'],
    };
}

