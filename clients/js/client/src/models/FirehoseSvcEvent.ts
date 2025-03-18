/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
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
 * @interface FirehoseSvcEvent
 */
export interface FirehoseSvcEvent {
    /**
     * 
     * @type {object}
     * @memberof FirehoseSvcEvent
     */
    data?: object;
    /**
     * 
     * @type {string}
     * @memberof FirehoseSvcEvent
     */
    name?: string;
}

/**
 * Check if a given object implements the FirehoseSvcEvent interface.
 */
export function instanceOfFirehoseSvcEvent(value: object): value is FirehoseSvcEvent {
    return true;
}

export function FirehoseSvcEventFromJSON(json: any): FirehoseSvcEvent {
    return FirehoseSvcEventFromJSONTyped(json, false);
}

export function FirehoseSvcEventFromJSONTyped(json: any, ignoreDiscriminator: boolean): FirehoseSvcEvent {
    if (json == null) {
        return json;
    }
    return {
        
        'data': json['data'] == null ? undefined : json['data'],
        'name': json['name'] == null ? undefined : json['name'],
    };
}

export function FirehoseSvcEventToJSON(json: any): FirehoseSvcEvent {
    return FirehoseSvcEventToJSONTyped(json, false);
}

export function FirehoseSvcEventToJSONTyped(value?: FirehoseSvcEvent | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'data': value['data'],
        'name': value['name'],
    };
}

