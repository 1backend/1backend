/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
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
 * @interface ModelSvcCudaParameters
 */
export interface ModelSvcCudaParameters {
    /**
     * Container configuration related to CUDA usage.
     * @type {ModelSvcContainer}
     * @memberof ModelSvcCudaParameters
     */
    container?: ModelSvcContainer;
    /**
     * Level of precision for selecting the CUDA version when resolving the container image.
     * - 2 -> Use "major.minor" (e.g., "12.2")
     * - 3 -> Use "major.minor.patch" (e.g., "12.2.0")
     * @type {number}
     * @memberof ModelSvcCudaParameters
     */
    cudaVersionPrecision?: number;
    /**
     * Default CUDA version to use (e.g., "12.2" or "12.2.0").
     * @type {string}
     * @memberof ModelSvcCudaParameters
     */
    defaultCudaVersion?: string;
    /**
     * Default cuDNN version to use alongside CUDA.
     * @type {string}
     * @memberof ModelSvcCudaParameters
     */
    defaultCudnnVersion?: string;
}

/**
 * Check if a given object implements the ModelSvcCudaParameters interface.
 */
export function instanceOfModelSvcCudaParameters(value: object): value is ModelSvcCudaParameters {
    return true;
}

export function ModelSvcCudaParametersFromJSON(json: any): ModelSvcCudaParameters {
    return ModelSvcCudaParametersFromJSONTyped(json, false);
}

export function ModelSvcCudaParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcCudaParameters {
    if (json == null) {
        return json;
    }
    return {
        
        'container': json['container'] == null ? undefined : ModelSvcContainerFromJSON(json['container']),
        'cudaVersionPrecision': json['cudaVersionPrecision'] == null ? undefined : json['cudaVersionPrecision'],
        'defaultCudaVersion': json['defaultCudaVersion'] == null ? undefined : json['defaultCudaVersion'],
        'defaultCudnnVersion': json['defaultCudnnVersion'] == null ? undefined : json['defaultCudnnVersion'],
    };
}

export function ModelSvcCudaParametersToJSON(json: any): ModelSvcCudaParameters {
    return ModelSvcCudaParametersToJSONTyped(json, false);
}

export function ModelSvcCudaParametersToJSONTyped(value?: ModelSvcCudaParameters | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'container': ModelSvcContainerToJSON(value['container']),
        'cudaVersionPrecision': value['cudaVersionPrecision'],
        'defaultCudaVersion': value['defaultCudaVersion'],
        'defaultCudnnVersion': value['defaultCudnnVersion'],
    };
}

