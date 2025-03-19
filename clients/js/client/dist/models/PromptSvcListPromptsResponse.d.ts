/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { PromptSvcPrompt } from './PromptSvcPrompt';
/**
 *
 * @export
 * @interface PromptSvcListPromptsResponse
 */
export interface PromptSvcListPromptsResponse {
    /**
     *
     * @type {object}
     * @memberof PromptSvcListPromptsResponse
     */
    after?: object;
    /**
     *
     * @type {number}
     * @memberof PromptSvcListPromptsResponse
     */
    count?: number;
    /**
     *
     * @type {Array<PromptSvcPrompt>}
     * @memberof PromptSvcListPromptsResponse
     */
    prompts?: Array<PromptSvcPrompt>;
}
/**
 * Check if a given object implements the PromptSvcListPromptsResponse interface.
 */
export declare function instanceOfPromptSvcListPromptsResponse(value: object): value is PromptSvcListPromptsResponse;
export declare function PromptSvcListPromptsResponseFromJSON(json: any): PromptSvcListPromptsResponse;
export declare function PromptSvcListPromptsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcListPromptsResponse;
export declare function PromptSvcListPromptsResponseToJSON(json: any): PromptSvcListPromptsResponse;
export declare function PromptSvcListPromptsResponseToJSONTyped(value?: PromptSvcListPromptsResponse | null, ignoreDiscriminator?: boolean): any;
