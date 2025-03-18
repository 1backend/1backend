/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
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
 * @interface RegistrySvcListDefinitionsResponse
 */
export interface RegistrySvcListDefinitionsResponse {
    /**
     * 
     * @type {Array<RegistrySvcDefinition>}
     * @memberof RegistrySvcListDefinitionsResponse
     */
    definitions?: Array<RegistrySvcDefinition>;
}

/**
 * Check if a given object implements the RegistrySvcListDefinitionsResponse interface.
 */
export function instanceOfRegistrySvcListDefinitionsResponse(value: object): value is RegistrySvcListDefinitionsResponse {
    return true;
}

export function RegistrySvcListDefinitionsResponseFromJSON(json: any): RegistrySvcListDefinitionsResponse {
    return RegistrySvcListDefinitionsResponseFromJSONTyped(json, false);
}

export function RegistrySvcListDefinitionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcListDefinitionsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'definitions': json['definitions'] == null ? undefined : ((json['definitions'] as Array<any>).map(RegistrySvcDefinitionFromJSON)),
    };
}

export function RegistrySvcListDefinitionsResponseToJSON(json: any): RegistrySvcListDefinitionsResponse {
    return RegistrySvcListDefinitionsResponseToJSONTyped(json, false);
}

export function RegistrySvcListDefinitionsResponseToJSONTyped(value?: RegistrySvcListDefinitionsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'definitions': value['definitions'] == null ? undefined : ((value['definitions'] as Array<any>).map(RegistrySvcDefinitionToJSON)),
    };
}

