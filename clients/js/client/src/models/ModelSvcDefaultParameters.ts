/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { ModelSvcContainer } from './ModelSvcContainer';
import {
    ModelSvcContainerFromJSON,
    ModelSvcContainerFromJSONTyped,
    ModelSvcContainerToJSON,
    ModelSvcContainerToJSONTyped,
} from './ModelSvcContainer';

/**
 * 
 * @export
 * @interface ModelSvcDefaultParameters
 */
export interface ModelSvcDefaultParameters {
    /**
     * 
     * @type {ModelSvcContainer}
     * @memberof ModelSvcDefaultParameters
     */
    container?: ModelSvcContainer;
}

/**
 * Check if a given object implements the ModelSvcDefaultParameters interface.
 */
export function instanceOfModelSvcDefaultParameters(value: object): value is ModelSvcDefaultParameters {
    return true;
}

export function ModelSvcDefaultParametersFromJSON(json: any): ModelSvcDefaultParameters {
    return ModelSvcDefaultParametersFromJSONTyped(json, false);
}

export function ModelSvcDefaultParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcDefaultParameters {
    if (json == null) {
        return json;
    }
    return {
        
        'container': json['container'] == null ? undefined : ModelSvcContainerFromJSON(json['container']),
    };
}

export function ModelSvcDefaultParametersToJSON(json: any): ModelSvcDefaultParameters {
    return ModelSvcDefaultParametersToJSONTyped(json, false);
}

export function ModelSvcDefaultParametersToJSONTyped(value?: ModelSvcDefaultParameters | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'container': ModelSvcContainerToJSON(value['container']),
    };
}

