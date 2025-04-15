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
import type { ChatSvcMessage } from './ChatSvcMessage';
import type { PromptSvcPrompt } from './PromptSvcPrompt';
/**
 *
 * @export
 * @interface PromptSvcPromptResponse
 */
export interface PromptSvcPromptResponse {
    /**
     * Prompt contains the details of the prompt that was just created by this request.
     * This includes the ID, prompt text, status, and other associated metadata.
     * @type {PromptSvcPrompt}
     * @memberof PromptSvcPromptResponse
     */
    prompt?: PromptSvcPrompt;
    /**
     * Response message contains the response text and files.
     * This field is populated only for synchronous prompts (`sync = true`).
     * For asynchronous prompts, the response will provided in the associated
     * message identified by the `responseMessageId` of the `promptSvc.prompt` object once the prompt completes.
     * @type {ChatSvcMessage}
     * @memberof PromptSvcPromptResponse
     */
    responseMessage?: ChatSvcMessage;
}
/**
 * Check if a given object implements the PromptSvcPromptResponse interface.
 */
export declare function instanceOfPromptSvcPromptResponse(value: object): value is PromptSvcPromptResponse;
export declare function PromptSvcPromptResponseFromJSON(json: any): PromptSvcPromptResponse;
export declare function PromptSvcPromptResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcPromptResponse;
export declare function PromptSvcPromptResponseToJSON(json: any): PromptSvcPromptResponse;
export declare function PromptSvcPromptResponseToJSONTyped(value?: PromptSvcPromptResponse | null, ignoreDiscriminator?: boolean): any;
