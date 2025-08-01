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
 * @interface ContainerSvcLog
 */
export interface ContainerSvcLog {
    /**
     * ContainerId is the raw underlying container ID.
     * Eg. Docker container id. Node local.
     * @type {string}
     * @memberof ContainerSvcLog
     */
    containerId?: string;
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcLog
     */
    content?: string;
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcLog
     */
    createdAt: string;
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcLog
     */
    id?: string;
    /**
     * Node Id
     * Please see the documentation for the envar OB_NODE_ID
     * @type {string}
     * @memberof ContainerSvcLog
     */
    nodeId?: string;
}

/**
 * Check if a given object implements the ContainerSvcLog interface.
 */
export function instanceOfContainerSvcLog(value: object): value is ContainerSvcLog {
    if (!('createdAt' in value) || value['createdAt'] === undefined) return false;
    return true;
}

export function ContainerSvcLogFromJSON(json: any): ContainerSvcLog {
    return ContainerSvcLogFromJSONTyped(json, false);
}

export function ContainerSvcLogFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcLog {
    if (json == null) {
        return json;
    }
    return {
        
        'containerId': json['containerId'] == null ? undefined : json['containerId'],
        'content': json['content'] == null ? undefined : json['content'],
        'createdAt': json['createdAt'],
        'id': json['id'] == null ? undefined : json['id'],
        'nodeId': json['nodeId'] == null ? undefined : json['nodeId'],
    };
}

export function ContainerSvcLogToJSON(json: any): ContainerSvcLog {
    return ContainerSvcLogToJSONTyped(json, false);
}

export function ContainerSvcLogToJSONTyped(value?: ContainerSvcLog | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'containerId': value['containerId'],
        'content': value['content'],
        'createdAt': value['createdAt'],
        'id': value['id'],
        'nodeId': value['nodeId'],
    };
}

