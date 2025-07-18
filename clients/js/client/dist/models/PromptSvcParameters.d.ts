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
import type { PromptSvcTextToImageParameters } from './PromptSvcTextToImageParameters';
import type { PromptSvcTextToTextParameters } from './PromptSvcTextToTextParameters';
/**
 *
 * @export
 * @interface PromptSvcParameters
 */
export interface PromptSvcParameters {
    /**
     *
     * @type {PromptSvcTextToImageParameters}
     * @memberof PromptSvcParameters
     */
    textToImage?: PromptSvcTextToImageParameters;
    /**
     *
     * @type {PromptSvcTextToTextParameters}
     * @memberof PromptSvcParameters
     */
    textToText?: PromptSvcTextToTextParameters;
}
/**
 * Check if a given object implements the PromptSvcParameters interface.
 */
export declare function instanceOfPromptSvcParameters(value: object): value is PromptSvcParameters;
export declare function PromptSvcParametersFromJSON(json: any): PromptSvcParameters;
export declare function PromptSvcParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcParameters;
export declare function PromptSvcParametersToJSON(json: any): PromptSvcParameters;
export declare function PromptSvcParametersToJSONTyped(value?: PromptSvcParameters | null, ignoreDiscriminator?: boolean): any;
