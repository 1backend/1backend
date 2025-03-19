/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
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
 * @interface PromptSvcTextToTextParameters
 */
export interface PromptSvcTextToTextParameters {
    /**
     * Template of the prompt. Optional. If not present it's derived from ModelId.
     * @type {string}
     * @memberof PromptSvcTextToTextParameters
     */
    template?: string;
}

/**
 * Check if a given object implements the PromptSvcTextToTextParameters interface.
 */
export function instanceOfPromptSvcTextToTextParameters(value: object): value is PromptSvcTextToTextParameters {
    return true;
}

export function PromptSvcTextToTextParametersFromJSON(json: any): PromptSvcTextToTextParameters {
    return PromptSvcTextToTextParametersFromJSONTyped(json, false);
}

export function PromptSvcTextToTextParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcTextToTextParameters {
    if (json == null) {
        return json;
    }
    return {
        
        'template': json['template'] == null ? undefined : json['template'],
    };
}

export function PromptSvcTextToTextParametersToJSON(json: any): PromptSvcTextToTextParameters {
    return PromptSvcTextToTextParametersToJSONTyped(json, false);
}

export function PromptSvcTextToTextParametersToJSONTyped(value?: PromptSvcTextToTextParameters | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'template': value['template'],
    };
}

