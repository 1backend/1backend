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
/**
 * 
 * @export
 * @interface ContainerSvcErrorResponse
 */
export interface ContainerSvcErrorResponse {
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcErrorResponse
     */
    error?: string;
}

/**
 * Check if a given object implements the ContainerSvcErrorResponse interface.
 */
export function instanceOfContainerSvcErrorResponse(value: object): value is ContainerSvcErrorResponse {
    return true;
}

export function ContainerSvcErrorResponseFromJSON(json: any): ContainerSvcErrorResponse {
    return ContainerSvcErrorResponseFromJSONTyped(json, false);
}

export function ContainerSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcErrorResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'error': json['error'] == null ? undefined : json['error'],
    };
}

export function ContainerSvcErrorResponseToJSON(json: any): ContainerSvcErrorResponse {
    return ContainerSvcErrorResponseToJSONTyped(json, false);
}

export function ContainerSvcErrorResponseToJSONTyped(value?: ContainerSvcErrorResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'error': value['error'],
    };
}

