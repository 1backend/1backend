/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import http from 'http';
import { PromptSvcListPromptsRequest } from '../model/promptSvcListPromptsRequest';
import { PromptSvcListPromptsResponse } from '../model/promptSvcListPromptsResponse';
import { PromptSvcPromptRequest } from '../model/promptSvcPromptRequest';
import { PromptSvcPromptResponse } from '../model/promptSvcPromptResponse';
import { PromptSvcRemovePromptRequest } from '../model/promptSvcRemovePromptRequest';
import { PromptSvcTypesResponse } from '../model/promptSvcTypesResponse';
import { Authentication, Interceptor } from '../model/models';
import { ApiKeyAuth } from '../model/models';
export declare enum PromptSvcApiApiKeys {
    BearerAuth = 0
}
export declare class PromptSvcApi {
    protected _basePath: string;
    protected _defaultHeaders: any;
    protected _useQuerystring: boolean;
    protected authentications: {
        default: Authentication;
        BearerAuth: ApiKeyAuth;
    };
    protected interceptors: Interceptor[];
    constructor(basePath?: string);
    set useQuerystring(value: boolean);
    set basePath(basePath: string);
    set defaultHeaders(defaultHeaders: any);
    get defaultHeaders(): any;
    get basePath(): string;
    setDefaultAuthentication(auth: Authentication): void;
    setApiKey(key: PromptSvcApiApiKeys, value: string): void;
    addInterceptor(interceptor: Interceptor): void;
    /**
     * List prompts that satisfy a query.
     * @summary List Prompts
     * @param body List Prompts Request
     */
    listPrompts(body?: PromptSvcListPromptsRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: PromptSvcListPromptsResponse;
    }>;
    /**
     * Sends a prompt and waits for a response if sync is true. If sync is false, adds the prompt to the queue and returns immediately.  Prompts can be used for `text-to-text`, `text-to-image`, `image-to-image`, and other types of generation. If no model ID is specified, the default model will be used (see `Model Svc` for details). The default model may or may not support the requested generation type.  **Prompting Modes** - **High-Level Parameters**: Uses predefined parameters relevant to `text-to-image`, `image-to-image`, etc. This mode abstracts away the underlying engine (e.g., LLaMA, Stable Diffusion) and focuses on functionality. - **Engine-Specific Parameters**: Uses `engineParameters` to directly specify an AI engine, exposing all available parameters for fine-tuned control.  **Permissions Required:** `prompt-svc:prompt:create`
     * @summary Prompt an AI
     * @param body Add Prompt Request
     */
    prompt(body: PromptSvcPromptRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: PromptSvcPromptResponse;
    }>;
    /**
     * The only purpose of this \"endpoint\" is to export types otherwise not appearing in the API docs. This endpoint otherwise does nothing. Do not depend on this endpoint, only its types.
     * @summary Prompt Types
     * @param body Types Request
     */
    promptTypes(body: object, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: PromptSvcTypesResponse;
    }>;
    /**
     * Remove a prompt by ID.
     * @summary Remove Prompt
     * @param body Remove Prompt Request
     */
    removePrompt(body: PromptSvcRemovePromptRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
    /**
     * Subscribe to prompt responses by thread via Server-Sent Events (SSE). You can subscribe to threads before they are created. The streamed strings are of type `StreamChunk`, see the PromptTypes endpoint for more details.
     * @summary Subscribe to Prompt Responses by Thread
     * @param threadId Thread ID
     */
    subscribeToPromptResponses(threadId: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: string;
    }>;
}
