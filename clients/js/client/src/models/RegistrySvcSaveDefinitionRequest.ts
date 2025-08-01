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
import type { RegistrySvcDefinition } from './RegistrySvcDefinition';
import {
    RegistrySvcDefinitionFromJSON,
    RegistrySvcDefinitionFromJSONTyped,
    RegistrySvcDefinitionToJSON,
    RegistrySvcDefinitionToJSONTyped,
} from './RegistrySvcDefinition';

/**
 * 
 * @export
 * @interface RegistrySvcSaveDefinitionRequest
 */
export interface RegistrySvcSaveDefinitionRequest {
    /**
     * 
     * @type {RegistrySvcDefinition}
     * @memberof RegistrySvcSaveDefinitionRequest
     */
    definition?: RegistrySvcDefinition;
}

/**
 * Check if a given object implements the RegistrySvcSaveDefinitionRequest interface.
 */
export function instanceOfRegistrySvcSaveDefinitionRequest(value: object): value is RegistrySvcSaveDefinitionRequest {
    return true;
}

export function RegistrySvcSaveDefinitionRequestFromJSON(json: any): RegistrySvcSaveDefinitionRequest {
    return RegistrySvcSaveDefinitionRequestFromJSONTyped(json, false);
}

export function RegistrySvcSaveDefinitionRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcSaveDefinitionRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'definition': json['definition'] == null ? undefined : RegistrySvcDefinitionFromJSON(json['definition']),
    };
}

export function RegistrySvcSaveDefinitionRequestToJSON(json: any): RegistrySvcSaveDefinitionRequest {
    return RegistrySvcSaveDefinitionRequestToJSONTyped(json, false);
}

export function RegistrySvcSaveDefinitionRequestToJSONTyped(value?: RegistrySvcSaveDefinitionRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'definition': RegistrySvcDefinitionToJSON(value['definition']),
    };
}

