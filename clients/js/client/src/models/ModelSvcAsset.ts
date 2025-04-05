/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
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
 * @interface ModelSvcAsset
 */
export interface ModelSvcAsset {
    /**
     * 
     * @type {string}
     * @memberof ModelSvcAsset
     */
    envVarKey: string;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcAsset
     */
    url: string;
}

/**
 * Check if a given object implements the ModelSvcAsset interface.
 */
export function instanceOfModelSvcAsset(value: object): value is ModelSvcAsset {
    if (!('envVarKey' in value) || value['envVarKey'] === undefined) return false;
    if (!('url' in value) || value['url'] === undefined) return false;
    return true;
}

export function ModelSvcAssetFromJSON(json: any): ModelSvcAsset {
    return ModelSvcAssetFromJSONTyped(json, false);
}

export function ModelSvcAssetFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcAsset {
    if (json == null) {
        return json;
    }
    return {
        
        'envVarKey': json['envVarKey'],
        'url': json['url'],
    };
}

export function ModelSvcAssetToJSON(json: any): ModelSvcAsset {
    return ModelSvcAssetToJSONTyped(json, false);
}

export function ModelSvcAssetToJSONTyped(value?: ModelSvcAsset | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'envVarKey': value['envVarKey'],
        'url': value['url'],
    };
}

