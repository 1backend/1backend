/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import * as runtime from '../runtime';
import type { PromptSvcListPromptsRequest, PromptSvcListPromptsResponse, PromptSvcPromptRequest, PromptSvcPromptResponse, PromptSvcRemovePromptRequest, PromptSvcTypesResponse } from '../models/index';
export interface ListPromptsRequest {
    body?: PromptSvcListPromptsRequest;
}
export interface PromptRequest {
    body: PromptSvcPromptRequest;
}
export interface PromptTypesRequest {
    body: object;
}
export interface RemovePromptRequest {
    body: PromptSvcRemovePromptRequest;
}
export interface SubscribeToPromptResponsesRequest {
    threadId: string;
}
/**
 *
 */
export declare class PromptSvcApi extends runtime.BaseAPI {
    /**
     * List prompts that satisfy a query.
     * List Prompts
     */
    listPromptsRaw(requestParameters: ListPromptsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PromptSvcListPromptsResponse>>;
    /**
     * List prompts that satisfy a query.
     * List Prompts
     */
    listPrompts(requestParameters?: ListPromptsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PromptSvcListPromptsResponse>;
    /**
     * Sends a prompt and waits for a response if sync is true. If sync is false, adds the prompt to the queue and returns immediately.  Prompts can be used for `text-to-text`, `text-to-image`, `image-to-image`, and other types of generation. If no model ID is specified, the default model will be used (see `Model Svc` for details). The default model may or may not support the requested generation type.  **Prompting Modes** - **High-Level Parameters**: Uses predefined parameters relevant to `text-to-image`, `image-to-image`, etc. This mode abstracts away the underlying engine (e.g., LLaMA, Stable Diffusion) and focuses on functionality. - **Engine-Specific Parameters**: Uses `engineParameters` to directly specify an AI engine, exposing all available parameters for fine-tuned control.  **Permissions Required:** `prompt-svc:prompt:create`
     * Prompt an AI
     */
    promptRaw(requestParameters: PromptRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PromptSvcPromptResponse>>;
    /**
     * Sends a prompt and waits for a response if sync is true. If sync is false, adds the prompt to the queue and returns immediately.  Prompts can be used for `text-to-text`, `text-to-image`, `image-to-image`, and other types of generation. If no model ID is specified, the default model will be used (see `Model Svc` for details). The default model may or may not support the requested generation type.  **Prompting Modes** - **High-Level Parameters**: Uses predefined parameters relevant to `text-to-image`, `image-to-image`, etc. This mode abstracts away the underlying engine (e.g., LLaMA, Stable Diffusion) and focuses on functionality. - **Engine-Specific Parameters**: Uses `engineParameters` to directly specify an AI engine, exposing all available parameters for fine-tuned control.  **Permissions Required:** `prompt-svc:prompt:create`
     * Prompt an AI
     */
    prompt(requestParameters: PromptRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PromptSvcPromptResponse>;
    /**
     * The only purpose of this \"endpoint\" is to export types otherwise not appearing in the API docs. This endpoint otherwise does nothing. Do not depend on this endpoint, only its types.
     * Prompt Types
     */
    promptTypesRaw(requestParameters: PromptTypesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PromptSvcTypesResponse>>;
    /**
     * The only purpose of this \"endpoint\" is to export types otherwise not appearing in the API docs. This endpoint otherwise does nothing. Do not depend on this endpoint, only its types.
     * Prompt Types
     */
    promptTypes(requestParameters: PromptTypesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PromptSvcTypesResponse>;
    /**
     * Remove a prompt by ID.
     * Remove Prompt
     */
    removePromptRaw(requestParameters: RemovePromptRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Remove a prompt by ID.
     * Remove Prompt
     */
    removePrompt(requestParameters: RemovePromptRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Subscribe to prompt responses by thread via Server-Sent Events (SSE). You can subscribe to threads before they are created. The streamed strings are of type `StreamChunk`, see the PromptTypes endpoint for more details.
     * Subscribe to Prompt Responses by Thread
     */
    subscribeToPromptResponsesRaw(requestParameters: SubscribeToPromptResponsesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<string>>;
    /**
     * Subscribe to prompt responses by thread via Server-Sent Events (SSE). You can subscribe to threads before they are created. The streamed strings are of type `StreamChunk`, see the PromptTypes endpoint for more details.
     * Subscribe to Prompt Responses by Thread
     */
    subscribeToPromptResponses(requestParameters: SubscribeToPromptResponsesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<string>;
}
