/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
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
 * @interface ContainerSvcDaemonInfoResponse
 */
export interface ContainerSvcDaemonInfoResponse {
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcDaemonInfoResponse
     */
    address?: string;
    /**
     * 
     * @type {boolean}
     * @memberof ContainerSvcDaemonInfoResponse
     */
    available: boolean;
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcDaemonInfoResponse
     */
    error?: string;
}

/**
 * Check if a given object implements the ContainerSvcDaemonInfoResponse interface.
 */
export function instanceOfContainerSvcDaemonInfoResponse(value: object): value is ContainerSvcDaemonInfoResponse {
    if (!('available' in value) || value['available'] === undefined) return false;
    return true;
}

export function ContainerSvcDaemonInfoResponseFromJSON(json: any): ContainerSvcDaemonInfoResponse {
    return ContainerSvcDaemonInfoResponseFromJSONTyped(json, false);
}

export function ContainerSvcDaemonInfoResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcDaemonInfoResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'address': json['address'] == null ? undefined : json['address'],
        'available': json['available'],
        'error': json['error'] == null ? undefined : json['error'],
    };
}

export function ContainerSvcDaemonInfoResponseToJSON(json: any): ContainerSvcDaemonInfoResponse {
    return ContainerSvcDaemonInfoResponseToJSONTyped(json, false);
}

export function ContainerSvcDaemonInfoResponseToJSONTyped(value?: ContainerSvcDaemonInfoResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'address': value['address'],
        'available': value['available'],
        'error': value['error'],
    };
}

