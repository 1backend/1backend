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
import type { ModelSvcModel } from './ModelSvcModel';
import {
    ModelSvcModelFromJSON,
    ModelSvcModelFromJSONTyped,
    ModelSvcModelToJSON,
    ModelSvcModelToJSONTyped,
} from './ModelSvcModel';
import type { ModelSvcPlatform } from './ModelSvcPlatform';
import {
    ModelSvcPlatformFromJSON,
    ModelSvcPlatformFromJSONTyped,
    ModelSvcPlatformToJSON,
    ModelSvcPlatformToJSONTyped,
} from './ModelSvcPlatform';

/**
 * 
 * @export
 * @interface ModelSvcGetModelResponse
 */
export interface ModelSvcGetModelResponse {
    /**
     * 
     * @type {boolean}
     * @memberof ModelSvcGetModelResponse
     */
    _exists: boolean;
    /**
     * 
     * @type {ModelSvcModel}
     * @memberof ModelSvcGetModelResponse
     */
    model: ModelSvcModel;
    /**
     * 
     * @type {ModelSvcPlatform}
     * @memberof ModelSvcGetModelResponse
     */
    platform: ModelSvcPlatform;
}

/**
 * Check if a given object implements the ModelSvcGetModelResponse interface.
 */
export function instanceOfModelSvcGetModelResponse(value: object): value is ModelSvcGetModelResponse {
    if (!('_exists' in value) || value['_exists'] === undefined) return false;
    if (!('model' in value) || value['model'] === undefined) return false;
    if (!('platform' in value) || value['platform'] === undefined) return false;
    return true;
}

export function ModelSvcGetModelResponseFromJSON(json: any): ModelSvcGetModelResponse {
    return ModelSvcGetModelResponseFromJSONTyped(json, false);
}

export function ModelSvcGetModelResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcGetModelResponse {
    if (json == null) {
        return json;
    }
    return {
        
        '_exists': json['exists'],
        'model': ModelSvcModelFromJSON(json['model']),
        'platform': ModelSvcPlatformFromJSON(json['platform']),
    };
}

export function ModelSvcGetModelResponseToJSON(json: any): ModelSvcGetModelResponse {
    return ModelSvcGetModelResponseToJSONTyped(json, false);
}

export function ModelSvcGetModelResponseToJSONTyped(value?: ModelSvcGetModelResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'exists': value['_exists'],
        'model': ModelSvcModelToJSON(value['model']),
        'platform': ModelSvcPlatformToJSON(value['platform']),
    };
}

