/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
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
 * @interface ContainerSvcListLogsRequest
 */
export interface ContainerSvcListLogsRequest {
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcListLogsRequest
     */
    containerId?: string;
    /**
     * 
     * @type {number}
     * @memberof ContainerSvcListLogsRequest
     */
    limit?: number;
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcListLogsRequest
     */
    nodeId?: string;
}

/**
 * Check if a given object implements the ContainerSvcListLogsRequest interface.
 */
export function instanceOfContainerSvcListLogsRequest(value: object): value is ContainerSvcListLogsRequest {
    return true;
}

export function ContainerSvcListLogsRequestFromJSON(json: any): ContainerSvcListLogsRequest {
    return ContainerSvcListLogsRequestFromJSONTyped(json, false);
}

export function ContainerSvcListLogsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcListLogsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'containerId': json['containerId'] == null ? undefined : json['containerId'],
        'limit': json['limit'] == null ? undefined : json['limit'],
        'nodeId': json['nodeId'] == null ? undefined : json['nodeId'],
    };
}

export function ContainerSvcListLogsRequestToJSON(json: any): ContainerSvcListLogsRequest {
    return ContainerSvcListLogsRequestToJSONTyped(json, false);
}

export function ContainerSvcListLogsRequestToJSONTyped(value?: ContainerSvcListLogsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'containerId': value['containerId'],
        'limit': value['limit'],
        'nodeId': value['nodeId'],
    };
}

