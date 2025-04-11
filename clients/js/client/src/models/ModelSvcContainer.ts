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
import type { ModelSvcKeep } from './ModelSvcKeep';
import {
    ModelSvcKeepFromJSON,
    ModelSvcKeepFromJSONTyped,
    ModelSvcKeepToJSON,
    ModelSvcKeepToJSONTyped,
} from './ModelSvcKeep';
import type { ModelSvcEnvVar } from './ModelSvcEnvVar';
import {
    ModelSvcEnvVarFromJSON,
    ModelSvcEnvVarFromJSONTyped,
    ModelSvcEnvVarToJSON,
    ModelSvcEnvVarToJSONTyped,
} from './ModelSvcEnvVar';

/**
 * 
 * @export
 * @interface ModelSvcContainer
 */
export interface ModelSvcContainer {
    /**
     * Environment variables to be passed to the container (e.g., "DEVICES=all").
     * @type {Array<ModelSvcEnvVar>}
     * @memberof ModelSvcContainer
     */
    envars?: Array<ModelSvcEnvVar>;
    /**
     * Template for constructing the container image name.
     * @type {string}
     * @memberof ModelSvcContainer
     */
    imageTemplate?: string;
    /**
     * List of container paths that should persist across restarts.
     * @type {Array<ModelSvcKeep>}
     * @memberof ModelSvcContainer
     */
    keeps?: Array<ModelSvcKeep>;
    /**
     * Internal port exposed by the container.
     * @type {number}
     * @memberof ModelSvcContainer
     */
    port?: number;
}

/**
 * Check if a given object implements the ModelSvcContainer interface.
 */
export function instanceOfModelSvcContainer(value: object): value is ModelSvcContainer {
    return true;
}

export function ModelSvcContainerFromJSON(json: any): ModelSvcContainer {
    return ModelSvcContainerFromJSONTyped(json, false);
}

export function ModelSvcContainerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcContainer {
    if (json == null) {
        return json;
    }
    return {
        
        'envars': json['envars'] == null ? undefined : ((json['envars'] as Array<any>).map(ModelSvcEnvVarFromJSON)),
        'imageTemplate': json['imageTemplate'] == null ? undefined : json['imageTemplate'],
        'keeps': json['keeps'] == null ? undefined : ((json['keeps'] as Array<any>).map(ModelSvcKeepFromJSON)),
        'port': json['port'] == null ? undefined : json['port'],
    };
}

export function ModelSvcContainerToJSON(json: any): ModelSvcContainer {
    return ModelSvcContainerToJSONTyped(json, false);
}

export function ModelSvcContainerToJSONTyped(value?: ModelSvcContainer | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'envars': value['envars'] == null ? undefined : ((value['envars'] as Array<any>).map(ModelSvcEnvVarToJSON)),
        'imageTemplate': value['imageTemplate'],
        'keeps': value['keeps'] == null ? undefined : ((value['keeps'] as Array<any>).map(ModelSvcKeepToJSON)),
        'port': value['port'],
    };
}

