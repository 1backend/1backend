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
import type { PromptSvcLlamaCppParameters } from './PromptSvcLlamaCppParameters';
import {
    PromptSvcLlamaCppParametersFromJSON,
    PromptSvcLlamaCppParametersFromJSONTyped,
    PromptSvcLlamaCppParametersToJSON,
    PromptSvcLlamaCppParametersToJSONTyped,
} from './PromptSvcLlamaCppParameters';
import type { PromptSvcStableDiffusionParameters } from './PromptSvcStableDiffusionParameters';
import {
    PromptSvcStableDiffusionParametersFromJSON,
    PromptSvcStableDiffusionParametersFromJSONTyped,
    PromptSvcStableDiffusionParametersToJSON,
    PromptSvcStableDiffusionParametersToJSONTyped,
} from './PromptSvcStableDiffusionParameters';

/**
 * 
 * @export
 * @interface PromptSvcEngineParameters
 */
export interface PromptSvcEngineParameters {
    /**
     * 
     * @type {PromptSvcLlamaCppParameters}
     * @memberof PromptSvcEngineParameters
     */
    llamaCppParameters?: PromptSvcLlamaCppParameters;
    /**
     * 
     * @type {PromptSvcStableDiffusionParameters}
     * @memberof PromptSvcEngineParameters
     */
    stableDiffusion?: PromptSvcStableDiffusionParameters;
}

/**
 * Check if a given object implements the PromptSvcEngineParameters interface.
 */
export function instanceOfPromptSvcEngineParameters(value: object): value is PromptSvcEngineParameters {
    return true;
}

export function PromptSvcEngineParametersFromJSON(json: any): PromptSvcEngineParameters {
    return PromptSvcEngineParametersFromJSONTyped(json, false);
}

export function PromptSvcEngineParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcEngineParameters {
    if (json == null) {
        return json;
    }
    return {
        
        'llamaCppParameters': json['llamaCppParameters'] == null ? undefined : PromptSvcLlamaCppParametersFromJSON(json['llamaCppParameters']),
        'stableDiffusion': json['stableDiffusion'] == null ? undefined : PromptSvcStableDiffusionParametersFromJSON(json['stableDiffusion']),
    };
}

export function PromptSvcEngineParametersToJSON(json: any): PromptSvcEngineParameters {
    return PromptSvcEngineParametersToJSONTyped(json, false);
}

export function PromptSvcEngineParametersToJSONTyped(value?: PromptSvcEngineParameters | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'llamaCppParameters': PromptSvcLlamaCppParametersToJSON(value['llamaCppParameters']),
        'stableDiffusion': PromptSvcStableDiffusionParametersToJSON(value['stableDiffusion']),
    };
}

