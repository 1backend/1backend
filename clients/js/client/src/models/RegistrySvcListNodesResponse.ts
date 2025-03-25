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
import type { RegistrySvcNode } from './RegistrySvcNode';
import {
    RegistrySvcNodeFromJSON,
    RegistrySvcNodeFromJSONTyped,
    RegistrySvcNodeToJSON,
    RegistrySvcNodeToJSONTyped,
} from './RegistrySvcNode';

/**
 * 
 * @export
 * @interface RegistrySvcListNodesResponse
 */
export interface RegistrySvcListNodesResponse {
    /**
     * 
     * @type {Array<RegistrySvcNode>}
     * @memberof RegistrySvcListNodesResponse
     */
    nodes?: Array<RegistrySvcNode>;
}

/**
 * Check if a given object implements the RegistrySvcListNodesResponse interface.
 */
export function instanceOfRegistrySvcListNodesResponse(value: object): value is RegistrySvcListNodesResponse {
    return true;
}

export function RegistrySvcListNodesResponseFromJSON(json: any): RegistrySvcListNodesResponse {
    return RegistrySvcListNodesResponseFromJSONTyped(json, false);
}

export function RegistrySvcListNodesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcListNodesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'nodes': json['nodes'] == null ? undefined : ((json['nodes'] as Array<any>).map(RegistrySvcNodeFromJSON)),
    };
}

export function RegistrySvcListNodesResponseToJSON(json: any): RegistrySvcListNodesResponse {
    return RegistrySvcListNodesResponseToJSONTyped(json, false);
}

export function RegistrySvcListNodesResponseToJSONTyped(value?: RegistrySvcListNodesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'nodes': value['nodes'] == null ? undefined : ((value['nodes'] as Array<any>).map(RegistrySvcNodeToJSON)),
    };
}

