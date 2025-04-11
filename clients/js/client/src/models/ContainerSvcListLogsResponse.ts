/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { ContainerSvcLog } from './ContainerSvcLog';
import {
    ContainerSvcLogFromJSON,
    ContainerSvcLogFromJSONTyped,
    ContainerSvcLogToJSON,
    ContainerSvcLogToJSONTyped,
} from './ContainerSvcLog';

/**
 * 
 * @export
 * @interface ContainerSvcListLogsResponse
 */
export interface ContainerSvcListLogsResponse {
    /**
     * 
     * @type {Array<ContainerSvcLog>}
     * @memberof ContainerSvcListLogsResponse
     */
    logs?: Array<ContainerSvcLog>;
}

/**
 * Check if a given object implements the ContainerSvcListLogsResponse interface.
 */
export function instanceOfContainerSvcListLogsResponse(value: object): value is ContainerSvcListLogsResponse {
    return true;
}

export function ContainerSvcListLogsResponseFromJSON(json: any): ContainerSvcListLogsResponse {
    return ContainerSvcListLogsResponseFromJSONTyped(json, false);
}

export function ContainerSvcListLogsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcListLogsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'logs': json['logs'] == null ? undefined : ((json['logs'] as Array<any>).map(ContainerSvcLogFromJSON)),
    };
}

export function ContainerSvcListLogsResponseToJSON(json: any): ContainerSvcListLogsResponse {
    return ContainerSvcListLogsResponseToJSONTyped(json, false);
}

export function ContainerSvcListLogsResponseToJSONTyped(value?: ContainerSvcListLogsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'logs': value['logs'] == null ? undefined : ((value['logs'] as Array<any>).map(ContainerSvcLogToJSON)),
    };
}

