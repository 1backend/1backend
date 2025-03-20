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
import type { PromptSvcLlamaCppParameters } from './PromptSvcLlamaCppParameters';
import type { PromptSvcStableDiffusionParameters } from './PromptSvcStableDiffusionParameters';
/**
 *
 * @export
 * @interface PromptSvcEngineParameters
 */
export interface PromptSvcEngineParameters {
    /**
     *
     * @type {PromptSvcLlamaCppParameters}
     * @memberof PromptSvcEngineParameters
     */
    llamaCppParameters?: PromptSvcLlamaCppParameters;
    /**
     *
     * @type {PromptSvcStableDiffusionParameters}
     * @memberof PromptSvcEngineParameters
     */
    stableDiffusion?: PromptSvcStableDiffusionParameters;
}
/**
 * Check if a given object implements the PromptSvcEngineParameters interface.
 */
export declare function instanceOfPromptSvcEngineParameters(value: object): value is PromptSvcEngineParameters;
export declare function PromptSvcEngineParametersFromJSON(json: any): PromptSvcEngineParameters;
export declare function PromptSvcEngineParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcEngineParameters;
export declare function PromptSvcEngineParametersToJSON(json: any): PromptSvcEngineParameters;
export declare function PromptSvcEngineParametersToJSONTyped(value?: PromptSvcEngineParameters | null, ignoreDiscriminator?: boolean): any;
