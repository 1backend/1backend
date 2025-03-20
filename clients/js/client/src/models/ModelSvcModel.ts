/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { ModelSvcAsset } from './ModelSvcAsset';
import {
    ModelSvcAssetFromJSON,
    ModelSvcAssetFromJSONTyped,
    ModelSvcAssetToJSON,
    ModelSvcAssetToJSONTyped,
} from './ModelSvcAsset';

/**
 * 
 * @export
 * @interface ModelSvcModel
 */
export interface ModelSvcModel {
    /**
     * 
     * @type {Array<ModelSvcAsset>}
     * @memberof ModelSvcModel
     */
    assets?: Array<ModelSvcAsset>;
    /**
     * 
     * @type {number}
     * @memberof ModelSvcModel
     */
    bits?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcModel
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcModel
     */
    extension?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcModel
     */
    flavour?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcModel
     */
    fullName?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcModel
     */
    id: string;
    /**
     * 
     * @type {number}
     * @memberof ModelSvcModel
     */
    maxBits?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelSvcModel
     */
    maxRam?: number;
    /**
     * 
     * @type {Array<string>}
     * @memberof ModelSvcModel
     */
    mirrors?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcModel
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcModel
     */
    parameters?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcModel
     */
    platformId: string;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcModel
     */
    promptTemplate?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcModel
     */
    quality?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcModel
     */
    quantComment?: string;
    /**
     * 
     * @type {number}
     * @memberof ModelSvcModel
     */
    size?: number;
    /**
     * 
     * @type {Array<string>}
     * @memberof ModelSvcModel
     */
    tags?: Array<string>;
    /**
     * 
     * @type {boolean}
     * @memberof ModelSvcModel
     */
    uncensored?: boolean;
    /**
     * 
     * @type {string}
     * @memberof ModelSvcModel
     */
    version?: string;
}

/**
 * Check if a given object implements the ModelSvcModel interface.
 */
export function instanceOfModelSvcModel(value: object): value is ModelSvcModel {
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('name' in value) || value['name'] === undefined) return false;
    if (!('platformId' in value) || value['platformId'] === undefined) return false;
    return true;
}

export function ModelSvcModelFromJSON(json: any): ModelSvcModel {
    return ModelSvcModelFromJSONTyped(json, false);
}

export function ModelSvcModelFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcModel {
    if (json == null) {
        return json;
    }
    return {
        
        'assets': json['assets'] == null ? undefined : ((json['assets'] as Array<any>).map(ModelSvcAssetFromJSON)),
        'bits': json['bits'] == null ? undefined : json['bits'],
        'description': json['description'] == null ? undefined : json['description'],
        'extension': json['extension'] == null ? undefined : json['extension'],
        'flavour': json['flavour'] == null ? undefined : json['flavour'],
        'fullName': json['fullName'] == null ? undefined : json['fullName'],
        'id': json['id'],
        'maxBits': json['maxBits'] == null ? undefined : json['maxBits'],
        'maxRam': json['maxRam'] == null ? undefined : json['maxRam'],
        'mirrors': json['mirrors'] == null ? undefined : json['mirrors'],
        'name': json['name'],
        'parameters': json['parameters'] == null ? undefined : json['parameters'],
        'platformId': json['platformId'],
        'promptTemplate': json['promptTemplate'] == null ? undefined : json['promptTemplate'],
        'quality': json['quality'] == null ? undefined : json['quality'],
        'quantComment': json['quantComment'] == null ? undefined : json['quantComment'],
        'size': json['size'] == null ? undefined : json['size'],
        'tags': json['tags'] == null ? undefined : json['tags'],
        'uncensored': json['uncensored'] == null ? undefined : json['uncensored'],
        'version': json['version'] == null ? undefined : json['version'],
    };
}

export function ModelSvcModelToJSON(json: any): ModelSvcModel {
    return ModelSvcModelToJSONTyped(json, false);
}

export function ModelSvcModelToJSONTyped(value?: ModelSvcModel | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'assets': value['assets'] == null ? undefined : ((value['assets'] as Array<any>).map(ModelSvcAssetToJSON)),
        'bits': value['bits'],
        'description': value['description'],
        'extension': value['extension'],
        'flavour': value['flavour'],
        'fullName': value['fullName'],
        'id': value['id'],
        'maxBits': value['maxBits'],
        'maxRam': value['maxRam'],
        'mirrors': value['mirrors'],
        'name': value['name'],
        'parameters': value['parameters'],
        'platformId': value['platformId'],
        'promptTemplate': value['promptTemplate'],
        'quality': value['quality'],
        'quantComment': value['quantComment'],
        'size': value['size'],
        'tags': value['tags'],
        'uncensored': value['uncensored'],
        'version': value['version'],
    };
}

