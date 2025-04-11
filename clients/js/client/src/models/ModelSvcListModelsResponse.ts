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
import type { ModelSvcModel } from './ModelSvcModel';
import {
    ModelSvcModelFromJSON,
    ModelSvcModelFromJSONTyped,
    ModelSvcModelToJSON,
    ModelSvcModelToJSONTyped,
} from './ModelSvcModel';

/**
 * 
 * @export
 * @interface ModelSvcListModelsResponse
 */
export interface ModelSvcListModelsResponse {
    /**
     * 
     * @type {Array<ModelSvcModel>}
     * @memberof ModelSvcListModelsResponse
     */
    models?: Array<ModelSvcModel>;
}

/**
 * Check if a given object implements the ModelSvcListModelsResponse interface.
 */
export function instanceOfModelSvcListModelsResponse(value: object): value is ModelSvcListModelsResponse {
    return true;
}

export function ModelSvcListModelsResponseFromJSON(json: any): ModelSvcListModelsResponse {
    return ModelSvcListModelsResponseFromJSONTyped(json, false);
}

export function ModelSvcListModelsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcListModelsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'models': json['models'] == null ? undefined : ((json['models'] as Array<any>).map(ModelSvcModelFromJSON)),
    };
}

export function ModelSvcListModelsResponseToJSON(json: any): ModelSvcListModelsResponse {
    return ModelSvcListModelsResponseToJSONTyped(json, false);
}

export function ModelSvcListModelsResponseToJSONTyped(value?: ModelSvcListModelsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'models': value['models'] == null ? undefined : ((value['models'] as Array<any>).map(ModelSvcModelToJSON)),
    };
}

