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
 * @interface ModelSvcListPlatformsResponse
 */
export interface ModelSvcListPlatformsResponse {
    /**
     * 
     * @type {Array<ModelSvcPlatform>}
     * @memberof ModelSvcListPlatformsResponse
     */
    platforms: Array<ModelSvcPlatform>;
}

/**
 * Check if a given object implements the ModelSvcListPlatformsResponse interface.
 */
export function instanceOfModelSvcListPlatformsResponse(value: object): value is ModelSvcListPlatformsResponse {
    if (!('platforms' in value) || value['platforms'] === undefined) return false;
    return true;
}

export function ModelSvcListPlatformsResponseFromJSON(json: any): ModelSvcListPlatformsResponse {
    return ModelSvcListPlatformsResponseFromJSONTyped(json, false);
}

export function ModelSvcListPlatformsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcListPlatformsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'platforms': ((json['platforms'] as Array<any>).map(ModelSvcPlatformFromJSON)),
    };
}

export function ModelSvcListPlatformsResponseToJSON(json: any): ModelSvcListPlatformsResponse {
    return ModelSvcListPlatformsResponseToJSONTyped(json, false);
}

export function ModelSvcListPlatformsResponseToJSONTyped(value?: ModelSvcListPlatformsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'platforms': ((value['platforms'] as Array<any>).map(ModelSvcPlatformToJSON)),
    };
}

