/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
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
 * @interface RegistrySvcNodeSelfResponse
 */
export interface RegistrySvcNodeSelfResponse {
    /**
     * 
     * @type {RegistrySvcNode}
     * @memberof RegistrySvcNodeSelfResponse
     */
    node: RegistrySvcNode;
}

/**
 * Check if a given object implements the RegistrySvcNodeSelfResponse interface.
 */
export function instanceOfRegistrySvcNodeSelfResponse(value: object): value is RegistrySvcNodeSelfResponse {
    if (!('node' in value) || value['node'] === undefined) return false;
    return true;
}

export function RegistrySvcNodeSelfResponseFromJSON(json: any): RegistrySvcNodeSelfResponse {
    return RegistrySvcNodeSelfResponseFromJSONTyped(json, false);
}

export function RegistrySvcNodeSelfResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcNodeSelfResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'node': RegistrySvcNodeFromJSON(json['node']),
    };
}

export function RegistrySvcNodeSelfResponseToJSON(json: any): RegistrySvcNodeSelfResponse {
    return RegistrySvcNodeSelfResponseToJSONTyped(json, false);
}

export function RegistrySvcNodeSelfResponseToJSONTyped(value?: RegistrySvcNodeSelfResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'node': RegistrySvcNodeToJSON(value['node']),
    };
}

