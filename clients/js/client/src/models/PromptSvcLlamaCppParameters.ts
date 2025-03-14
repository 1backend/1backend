/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
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
 * @interface PromptSvcLlamaCppParameters
 */
export interface PromptSvcLlamaCppParameters {
    /**
     * Template of the prompt. Optional. If not present it's derived from ModelId.
     * @type {string}
     * @memberof PromptSvcLlamaCppParameters
     */
    template?: string;
}

/**
 * Check if a given object implements the PromptSvcLlamaCppParameters interface.
 */
export function instanceOfPromptSvcLlamaCppParameters(value: object): value is PromptSvcLlamaCppParameters {
    return true;
}

export function PromptSvcLlamaCppParametersFromJSON(json: any): PromptSvcLlamaCppParameters {
    return PromptSvcLlamaCppParametersFromJSONTyped(json, false);
}

export function PromptSvcLlamaCppParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcLlamaCppParameters {
    if (json == null) {
        return json;
    }
    return {
        
        'template': json['template'] == null ? undefined : json['template'],
    };
}

export function PromptSvcLlamaCppParametersToJSON(json: any): PromptSvcLlamaCppParameters {
    return PromptSvcLlamaCppParametersToJSONTyped(json, false);
}

export function PromptSvcLlamaCppParametersToJSONTyped(value?: PromptSvcLlamaCppParameters | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'template': value['template'],
    };
}

