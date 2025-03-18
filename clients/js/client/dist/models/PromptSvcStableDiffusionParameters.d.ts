/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { StableDiffusionTxt2ImgRequest } from './StableDiffusionTxt2ImgRequest';
/**
 *
 * @export
 * @interface PromptSvcStableDiffusionParameters
 */
export interface PromptSvcStableDiffusionParameters {
    /**
     * Text to image parameters
     * @type {StableDiffusionTxt2ImgRequest}
     * @memberof PromptSvcStableDiffusionParameters
     */
    txt2Img?: StableDiffusionTxt2ImgRequest;
}
/**
 * Check if a given object implements the PromptSvcStableDiffusionParameters interface.
 */
export declare function instanceOfPromptSvcStableDiffusionParameters(value: object): value is PromptSvcStableDiffusionParameters;
export declare function PromptSvcStableDiffusionParametersFromJSON(json: any): PromptSvcStableDiffusionParameters;
export declare function PromptSvcStableDiffusionParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcStableDiffusionParameters;
export declare function PromptSvcStableDiffusionParametersToJSON(json: any): PromptSvcStableDiffusionParameters;
export declare function PromptSvcStableDiffusionParametersToJSONTyped(value?: PromptSvcStableDiffusionParameters | null, ignoreDiscriminator?: boolean): any;
