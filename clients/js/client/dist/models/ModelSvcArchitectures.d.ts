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
import type { ModelSvcCudaParameters } from './ModelSvcCudaParameters';
import type { ModelSvcDefaultParameters } from './ModelSvcDefaultParameters';
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
export declare function instanceOfModelSvcArchitectures(value: object): value is ModelSvcArchitectures;
export declare function ModelSvcArchitecturesFromJSON(json: any): ModelSvcArchitectures;
export declare function ModelSvcArchitecturesFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcArchitectures;
export declare function ModelSvcArchitecturesToJSON(json: any): ModelSvcArchitectures;
export declare function ModelSvcArchitecturesToJSONTyped(value?: ModelSvcArchitectures | null, ignoreDiscriminator?: boolean): any;
