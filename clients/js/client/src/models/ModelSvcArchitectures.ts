/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { ModelSvcCudaParameters } from './ModelSvcCudaParameters';
import {
    ModelSvcCudaParametersFromJSON,
    ModelSvcCudaParametersFromJSONTyped,
    ModelSvcCudaParametersToJSON,
    ModelSvcCudaParametersToJSONTyped,
} from './ModelSvcCudaParameters';
import type { ModelSvcDefaultParameters } from './ModelSvcDefaultParameters';
import {
    ModelSvcDefaultParametersFromJSON,
    ModelSvcDefaultParametersFromJSONTyped,
    ModelSvcDefaultParametersToJSON,
    ModelSvcDefaultParametersToJSONTyped,
} from './ModelSvcDefaultParameters';

/**
 * 
 * @export
 * @interface ModelSvcArchitectures
 */
export interface ModelSvcArchitectures {
    /**
     * CUDA-specific container parameters, if applicable.
     * @type {ModelSvcCudaParameters}
     * @memberof ModelSvcArchitectures
     */
    cuda?: ModelSvcCudaParameters;
    /**
     * Default container configuration for non-GPU environments.
     * @type {ModelSvcDefaultParameters}
     * @memberof ModelSvcArchitectures
     */
    _default?: ModelSvcDefaultParameters;
}

/**
 * Check if a given object implements the ModelSvcArchitectures interface.
 */
export function instanceOfModelSvcArchitectures(value: object): value is ModelSvcArchitectures {
    return true;
}

export function ModelSvcArchitecturesFromJSON(json: any): ModelSvcArchitectures {
    return ModelSvcArchitecturesFromJSONTyped(json, false);
}

export function ModelSvcArchitecturesFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcArchitectures {
    if (json == null) {
        return json;
    }
    return {
        
        'cuda': json['cuda'] == null ? undefined : ModelSvcCudaParametersFromJSON(json['cuda']),
        '_default': json['default'] == null ? undefined : ModelSvcDefaultParametersFromJSON(json['default']),
    };
}

export function ModelSvcArchitecturesToJSON(json: any): ModelSvcArchitectures {
    return ModelSvcArchitecturesToJSONTyped(json, false);
}

export function ModelSvcArchitecturesToJSONTyped(value?: ModelSvcArchitectures | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cuda': ModelSvcCudaParametersToJSON(value['cuda']),
        'default': ModelSvcDefaultParametersToJSON(value['_default']),
    };
}

