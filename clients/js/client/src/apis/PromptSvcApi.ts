/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  PromptSvcErrorResponse,
  PromptSvcListPromptsRequest,
  PromptSvcListPromptsResponse,
  PromptSvcPromptRequest,
  PromptSvcPromptResponse,
  PromptSvcRemovePromptRequest,
  PromptSvcTypesResponse,
} from '../models/index';
import {
    PromptSvcErrorResponseFromJSON,
    PromptSvcErrorResponseToJSON,
    PromptSvcListPromptsRequestFromJSON,
    PromptSvcListPromptsRequestToJSON,
    PromptSvcListPromptsResponseFromJSON,
    PromptSvcListPromptsResponseToJSON,
    PromptSvcPromptRequestFromJSON,
    PromptSvcPromptRequestToJSON,
    PromptSvcPromptResponseFromJSON,
    PromptSvcPromptResponseToJSON,
    PromptSvcRemovePromptRequestFromJSON,
    PromptSvcRemovePromptRequestToJSON,
    PromptSvcTypesResponseFromJSON,
    PromptSvcTypesResponseToJSON,
} from '../models/index';

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
export class PromptSvcApi extends runtime.BaseAPI {

    /**
     * List prompts that satisfy a query.
     * List Prompts
     */
    async listPromptsRaw(requestParameters: ListPromptsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PromptSvcListPromptsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/prompt-svc/prompts`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: PromptSvcListPromptsRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PromptSvcListPromptsResponseFromJSON(jsonValue));
    }

    /**
     * List prompts that satisfy a query.
     * List Prompts
     */
    async listPrompts(requestParameters: ListPromptsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PromptSvcListPromptsResponse> {
        const response = await this.listPromptsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Sends a prompt and waits for a response if sync is true. If sync is false, adds the prompt to the queue and returns immediately.  Prompts can be used for `text-to-text`, `text-to-image`, `image-to-image`, and other types of generation. If no model ID is specified, the default model will be used (see `Model Svc` for details). The default model may or may not support the requested generation type.  **Prompting Modes** - **High-Level Parameters**: Uses predefined parameters relevant to `text-to-image`, `image-to-image`, etc. This mode abstracts away the underlying engine (e.g., LLaMA, Stable Diffusion) and focuses on functionality. - **Engine-Specific Parameters**: Uses `engineParameters` to directly specify an AI engine, exposing all available parameters for fine-tuned control.  **Permissions Required:** `prompt-svc:prompt:create`
     * Prompt an AI
     */
    async promptRaw(requestParameters: PromptRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PromptSvcPromptResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling prompt().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/prompt-svc/prompt`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: PromptSvcPromptRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PromptSvcPromptResponseFromJSON(jsonValue));
    }

    /**
     * Sends a prompt and waits for a response if sync is true. If sync is false, adds the prompt to the queue and returns immediately.  Prompts can be used for `text-to-text`, `text-to-image`, `image-to-image`, and other types of generation. If no model ID is specified, the default model will be used (see `Model Svc` for details). The default model may or may not support the requested generation type.  **Prompting Modes** - **High-Level Parameters**: Uses predefined parameters relevant to `text-to-image`, `image-to-image`, etc. This mode abstracts away the underlying engine (e.g., LLaMA, Stable Diffusion) and focuses on functionality. - **Engine-Specific Parameters**: Uses `engineParameters` to directly specify an AI engine, exposing all available parameters for fine-tuned control.  **Permissions Required:** `prompt-svc:prompt:create`
     * Prompt an AI
     */
    async prompt(requestParameters: PromptRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PromptSvcPromptResponse> {
        const response = await this.promptRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * The only purpose of this \"endpoint\" is to export types otherwise not appearing in the API docs. This endpoint otherwise does nothing. Do not depend on this endpoint, only its types.
     * Prompt Types
     */
    async promptTypesRaw(requestParameters: PromptTypesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PromptSvcTypesResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling promptTypes().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/prompt-svc/types`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters['body'] as any,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PromptSvcTypesResponseFromJSON(jsonValue));
    }

    /**
     * The only purpose of this \"endpoint\" is to export types otherwise not appearing in the API docs. This endpoint otherwise does nothing. Do not depend on this endpoint, only its types.
     * Prompt Types
     */
    async promptTypes(requestParameters: PromptTypesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PromptSvcTypesResponse> {
        const response = await this.promptTypesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Remove a prompt by ID.
     * Remove Prompt
     */
    async removePromptRaw(requestParameters: RemovePromptRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling removePrompt().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/prompt-svc/remove`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: PromptSvcRemovePromptRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Remove a prompt by ID.
     * Remove Prompt
     */
    async removePrompt(requestParameters: RemovePromptRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.removePromptRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Subscribe to prompt responses by thread via Server-Sent Events (SSE). You can subscribe to threads before they are created. The streamed strings are of type `StreamChunk`, see the PromptTypes endpoint for more details.
     * Subscribe to Prompt Responses by Thread
     */
    async subscribeToPromptResponsesRaw(requestParameters: SubscribeToPromptResponsesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<string>> {
        if (requestParameters['threadId'] == null) {
            throw new runtime.RequiredError(
                'threadId',
                'Required parameter "threadId" was null or undefined when calling subscribeToPromptResponses().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/prompt-svc/prompts/{threadId}/responses/subscribe`.replace(`{${"threadId"}}`, encodeURIComponent(String(requestParameters['threadId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        if (this.isJsonMime(response.headers.get('content-type'))) {
            return new runtime.JSONApiResponse<string>(response);
        } else {
            return new runtime.TextApiResponse(response) as any;
        }
    }

    /**
     * Subscribe to prompt responses by thread via Server-Sent Events (SSE). You can subscribe to threads before they are created. The streamed strings are of type `StreamChunk`, see the PromptTypes endpoint for more details.
     * Subscribe to Prompt Responses by Thread
     */
    async subscribeToPromptResponses(requestParameters: SubscribeToPromptResponsesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<string> {
        const response = await this.subscribeToPromptResponsesRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
