/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { PromptSvcParameters } from './PromptSvcParameters';
import type { PromptSvcEngineParameters } from './PromptSvcEngineParameters';
/**
 *
 * @export
 * @interface PromptSvcPromptRequest
 */
export interface PromptSvcPromptRequest {
    /**
     * AI engine/platform (eg. Llama, Stable Diffusion) specific parameters
     * @type {PromptSvcEngineParameters}
     * @memberof PromptSvcPromptRequest
     */
    engineParameters?: PromptSvcEngineParameters;
    /**
     * Id is the unique ID of the prompt.
     * @type {string}
     * @memberof PromptSvcPromptRequest
     */
    id?: string;
    /**
     * MaxRetries specified how many times the system should retry a prompt when it keeps erroring.
     * @type {number}
     * @memberof PromptSvcPromptRequest
     */
    maxRetries?: number;
    /**
     * ModelId is just the 1Backend internal ID of the model.
     * @type {string}
     * @memberof PromptSvcPromptRequest
     */
    modelId?: string;
    /**
     * AI engine/platform (eg. Llama, Stable Diffusion) agnostic parameters.
     * Use these high level parameters when you don't care about the actual engine, only
     * the functionality (eg. text to image, image to image) it provides.
     * @type {PromptSvcParameters}
     * @memberof PromptSvcPromptRequest
     */
    parameters?: PromptSvcParameters;
    /**
     * Prompt is the message itself eg. "What's a banana?
     * @type {string}
     * @memberof PromptSvcPromptRequest
     */
    prompt: string;
    /**
     * Sync drives whether prompt add request should wait and hang until
     * the prompt is done executing. By default the prompt just gets put on a queue
     * and the client will just subscribe to a Thread Stream.
     * For quick and dirty scripting however it's often times easier to do things synchronously.
     * In those cases set Sync to true.
     * @type {boolean}
     * @memberof PromptSvcPromptRequest
     */
    sync?: boolean;
    /**
     * ThreadId is the ID of the thread a prompt belongs to.
     * Clients subscribe to Thread Streams to see the answer to a prompt,
     * or set `prompt.sync` to true for a blocking answer.
     * @type {string}
     * @memberof PromptSvcPromptRequest
     */
    threadId?: string;
}
/**
 * Check if a given object implements the PromptSvcPromptRequest interface.
 */
export declare function instanceOfPromptSvcPromptRequest(value: object): value is PromptSvcPromptRequest;
export declare function PromptSvcPromptRequestFromJSON(json: any): PromptSvcPromptRequest;
export declare function PromptSvcPromptRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcPromptRequest;
export declare function PromptSvcPromptRequestToJSON(json: any): PromptSvcPromptRequest;
export declare function PromptSvcPromptRequestToJSONTyped(value?: PromptSvcPromptRequest | null, ignoreDiscriminator?: boolean): any;
