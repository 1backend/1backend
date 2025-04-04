/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface PromptSvcTextToTextParameters
 */
export interface PromptSvcTextToTextParameters {
    /**
     * Template of the prompt. Optional. If not present it's derived from ModelId.
     * @type {string}
     * @memberof PromptSvcTextToTextParameters
     */
    template?: string;
}
/**
 * Check if a given object implements the PromptSvcTextToTextParameters interface.
 */
export declare function instanceOfPromptSvcTextToTextParameters(value: object): value is PromptSvcTextToTextParameters;
export declare function PromptSvcTextToTextParametersFromJSON(json: any): PromptSvcTextToTextParameters;
export declare function PromptSvcTextToTextParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcTextToTextParameters;
export declare function PromptSvcTextToTextParametersToJSON(json: any): PromptSvcTextToTextParameters;
export declare function PromptSvcTextToTextParametersToJSONTyped(value?: PromptSvcTextToTextParameters | null, ignoreDiscriminator?: boolean): any;
