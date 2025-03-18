/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ModelSvcContainer } from './ModelSvcContainer';
/**
 *
 * @export
 * @interface ModelSvcDefaultParameters
 */
export interface ModelSvcDefaultParameters {
    /**
     *
     * @type {ModelSvcContainer}
     * @memberof ModelSvcDefaultParameters
     */
    container?: ModelSvcContainer;
}
/**
 * Check if a given object implements the ModelSvcDefaultParameters interface.
 */
export declare function instanceOfModelSvcDefaultParameters(value: object): value is ModelSvcDefaultParameters;
export declare function ModelSvcDefaultParametersFromJSON(json: any): ModelSvcDefaultParameters;
export declare function ModelSvcDefaultParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSvcDefaultParameters;
export declare function ModelSvcDefaultParametersToJSON(json: any): ModelSvcDefaultParameters;
export declare function ModelSvcDefaultParametersToJSONTyped(value?: ModelSvcDefaultParameters | null, ignoreDiscriminator?: boolean): any;
