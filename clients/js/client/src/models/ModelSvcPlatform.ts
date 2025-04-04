/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { PromptSvcPromptType } from './PromptSvcPromptType';
import {
    PromptSvcPromptTypeFromJSON,
    PromptSvcPromptTypeFromJSONTyped,
    PromptSvcPromptTypeToJSON,
    PromptSvcPromptTypeToJSONTyped,
} from './PromptSvcPromptType';
import type { ModelSvcArchitectures } from './ModelSvcArchitectures';
import {
    ModelSvcArchitecturesFromJSON,
    ModelSvcArchitecturesFromJSONTyped,
    ModelSvcArchitecturesToJSON,
    ModelSvcArchitecturesToJSONTyped,
} from './ModelSvcArchitectures';

/**
 * 
 * @export
 * @interface ModelSvcPlatform
 */
export interface ModelSvcPlatform {
    /**
     * 
     * @type {ModelSvcArchitectures}
     * @memberof ModelSvcPlatform
     */
    architectures?: ModelSvcArchitectures;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcPlatform
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcPlatform
     */
    name?: string;
    /**
     * List of prompt types that the AI engine supports.
     * @type {Array<PromptSvcPromptType>}
     * @memberof ModelSvcPlatform
     */
    types?: Array<PromptSvcPromptType>;
    /**
     * 
     * @type {number}
     * @memberof ModelSvcPlatform
     */
    version?: number;
}

/**
 * Check if a given object implements the ModelSvcPlatform interface.
 */
export function instanceOfModelSvcPlatform(value: object): value is ModelSvcPlatform {
    return true;
}

export function ModelSvcPlatformFromJSON(json: any): ModelSvcPlatform {
    return ModelSvcPlatformFromJSONTyped(json, false);
}

export function ModelSvcPlatformFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcPlatform {
    if (json == null) {
        return json;
    }
    return {
        
        'architectures': json['architectures'] == null ? undefined : ModelSvcArchitecturesFromJSON(json['architectures']),
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'types': json['types'] == null ? undefined : ((json['types'] as Array<any>).map(PromptSvcPromptTypeFromJSON)),
        'version': json['version'] == null ? undefined : json['version'],
    };
}

export function ModelSvcPlatformToJSON(json: any): ModelSvcPlatform {
    return ModelSvcPlatformToJSONTyped(json, false);
}

export function ModelSvcPlatformToJSONTyped(value?: ModelSvcPlatform | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'architectures': ModelSvcArchitecturesToJSON(value['architectures']),
        'id': value['id'],
        'name': value['name'],
        'types': value['types'] == null ? undefined : ((value['types'] as Array<any>).map(PromptSvcPromptTypeToJSON)),
        'version': value['version'],
    };
}

