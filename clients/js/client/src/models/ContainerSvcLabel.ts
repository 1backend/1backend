/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
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
 * @interface ContainerSvcLabel
 */
export interface ContainerSvcLabel {
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcLabel
     */
    key: string;
    /**
     * 
     * @type {string}
     * @memberof ContainerSvcLabel
     */
    value: string;
}

/**
 * Check if a given object implements the ContainerSvcLabel interface.
 */
export function instanceOfContainerSvcLabel(value: object): value is ContainerSvcLabel {
    if (!('key' in value) || value['key'] === undefined) return false;
    if (!('value' in value) || value['value'] === undefined) return false;
    return true;
}

export function ContainerSvcLabelFromJSON(json: any): ContainerSvcLabel {
    return ContainerSvcLabelFromJSONTyped(json, false);
}

export function ContainerSvcLabelFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContainerSvcLabel {
    if (json == null) {
        return json;
    }
    return {
        
        'key': json['key'],
        'value': json['value'],
    };
}

export function ContainerSvcLabelToJSON(json: any): ContainerSvcLabel {
    return ContainerSvcLabelToJSONTyped(json, false);
}

export function ContainerSvcLabelToJSONTyped(value?: ContainerSvcLabel | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'key': value['key'],
        'value': value['value'],
    };
}

