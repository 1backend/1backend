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
/**
 * 
 * @export
 * @interface ContainerSvcContainerIsRunningResponse
 */
export interface ContainerSvcContainerIsRunningResponse {
    /**
     * 
     * @type {boolean}
     * @memberof ContainerSvcContainerIsRunningResponse
     */
    isRunning: boolean;
}

/**
 * Check if a given object implements the ContainerSvcContainerIsRunningResponse interface.
 */
export function instanceOfContainerSvcContainerIsRunningResponse(value: object): value is ContainerSvcContainerIsRunningResponse {
    if (!('isRunning' in value) || value['isRunning'] === undefined) return false;
    return true;
}

export function ContainerSvcContainerIsRunningResponseFromJSON(json: any): ContainerSvcContainerIsRunningResponse {
    return ContainerSvcContainerIsRunningResponseFromJSONTyped(json, false);
}

export function ContainerSvcContainerIsRunningResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcContainerIsRunningResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'isRunning': json['isRunning'],
    };
}

export function ContainerSvcContainerIsRunningResponseToJSON(json: any): ContainerSvcContainerIsRunningResponse {
    return ContainerSvcContainerIsRunningResponseToJSONTyped(json, false);
}

export function ContainerSvcContainerIsRunningResponseToJSONTyped(value?: ContainerSvcContainerIsRunningResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'isRunning': value['isRunning'],
    };
}

