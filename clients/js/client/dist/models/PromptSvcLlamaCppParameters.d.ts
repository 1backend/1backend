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
/**
 *
 * @export
 * @interface PromptSvcLlamaCppParameters
 */
export interface PromptSvcLlamaCppParameters {
    /**
     * Template of the prompt. Optional. If not present it's derived from ModelId.
     * @type {string}
     * @memberof PromptSvcLlamaCppParameters
     */
    template?: string;
}
/**
 * Check if a given object implements the PromptSvcLlamaCppParameters interface.
 */
export declare function instanceOfPromptSvcLlamaCppParameters(value: object): value is PromptSvcLlamaCppParameters;
export declare function PromptSvcLlamaCppParametersFromJSON(json: any): PromptSvcLlamaCppParameters;
export declare function PromptSvcLlamaCppParametersFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcLlamaCppParameters;
export declare function PromptSvcLlamaCppParametersToJSON(json: any): PromptSvcLlamaCppParameters;
export declare function PromptSvcLlamaCppParametersToJSONTyped(value?: PromptSvcLlamaCppParameters | null, ignoreDiscriminator?: boolean): any;
