/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
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
 * @interface ContainerSvcStopContainerRequest
 */
export interface ContainerSvcStopContainerRequest {
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcStopContainerRequest
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcStopContainerRequest
     */
    name?: string;
}

/**
 * Check if a given object implements the ContainerSvcStopContainerRequest interface.
 */
export function instanceOfContainerSvcStopContainerRequest(value: object): value is ContainerSvcStopContainerRequest {
    return true;
}

export function ContainerSvcStopContainerRequestFromJSON(json: any): ContainerSvcStopContainerRequest {
    return ContainerSvcStopContainerRequestFromJSONTyped(json, false);
}

export function ContainerSvcStopContainerRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcStopContainerRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
    };
}

export function ContainerSvcStopContainerRequestToJSON(json: any): ContainerSvcStopContainerRequest {
    return ContainerSvcStopContainerRequestToJSONTyped(json, false);
}

export function ContainerSvcStopContainerRequestToJSONTyped(value?: ContainerSvcStopContainerRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'name': value['name'],
    };
}

