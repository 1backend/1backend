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
 * @interface ProxySvcRoute
 */
export interface ProxySvcRoute {
    /**
     * 
     * @type {string}
     * @memberof ProxySvcRoute
     */
    createdAt?: string;
    /**
     * Id is the host itself, e.g. "test.localhost"
     * @type {string}
     * @memberof ProxySvcRoute
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof ProxySvcRoute
     */
    target?: string;
    /**
     * 
     * @type {string}
     * @memberof ProxySvcRoute
     */
    updatedAt?: string;
}

/**
 * Check if a given object implements the ProxySvcRoute interface.
 */
export function instanceOfProxySvcRoute(value: object): value is ProxySvcRoute {
    return true;
}

export function ProxySvcRouteFromJSON(json: any): ProxySvcRoute {
    return ProxySvcRouteFromJSONTyped(json, false);
}

export function ProxySvcRouteFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProxySvcRoute {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'id': json['id'] == null ? undefined : json['id'],
        'target': json['target'] == null ? undefined : json['target'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
    };
}

export function ProxySvcRouteToJSON(json: any): ProxySvcRoute {
    return ProxySvcRouteToJSONTyped(json, false);
}

export function ProxySvcRouteToJSONTyped(value?: ProxySvcRoute | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'createdAt': value['createdAt'],
        'id': value['id'],
        'target': value['target'],
        'updatedAt': value['updatedAt'],
    };
}

