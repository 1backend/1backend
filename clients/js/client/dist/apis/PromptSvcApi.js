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
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
import * as runtime from '../runtime';
import { PromptSvcListPromptsRequestToJSON, PromptSvcListPromptsResponseFromJSON, PromptSvcPromptRequestToJSON, PromptSvcPromptResponseFromJSON, PromptSvcRemovePromptRequestToJSON, PromptSvcTypesResponseFromJSON, } from '../models/index';
/**
 *
 */
export class PromptSvcApi extends runtime.BaseAPI {
    /**
     * List prompts that satisfy a query.
     * List Prompts
     */
    listPromptsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/prompt-svc/prompts`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: PromptSvcListPromptsRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => PromptSvcListPromptsResponseFromJSON(jsonValue));
        });
    }
    /**
     * List prompts that satisfy a query.
     * List Prompts
     */
    listPrompts() {
        return __awaiter(this, arguments, void 0, function* (requestParameters = {}, initOverrides) {
            const response = yield this.listPromptsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Sends a prompt and waits for a response if sync is true. If sync is false, adds the prompt to the queue and returns immediately.  Prompts can be used for `text-to-text`, `text-to-image`, `image-to-image`, and other types of generation. If no model ID is specified, the default model will be used (see `Model Svc` for details). The default model may or may not support the requested generation type.  **Prompting Modes** - **High-Level Parameters**: Uses predefined parameters relevant to `text-to-image`, `image-to-image`, etc. This mode abstracts away the underlying engine (e.g., LLaMA, Stable Diffusion) and focuses on functionality. - **Engine-Specific Parameters**: Uses `engineParameters` to directly specify an AI engine, exposing all available parameters for fine-tuned control.  **Permissions Required:** `prompt-svc:prompt:create`
     * Prompt an AI
     */
    promptRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling prompt().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/prompt-svc/prompt`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: PromptSvcPromptRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => PromptSvcPromptResponseFromJSON(jsonValue));
        });
    }
    /**
     * Sends a prompt and waits for a response if sync is true. If sync is false, adds the prompt to the queue and returns immediately.  Prompts can be used for `text-to-text`, `text-to-image`, `image-to-image`, and other types of generation. If no model ID is specified, the default model will be used (see `Model Svc` for details). The default model may or may not support the requested generation type.  **Prompting Modes** - **High-Level Parameters**: Uses predefined parameters relevant to `text-to-image`, `image-to-image`, etc. This mode abstracts away the underlying engine (e.g., LLaMA, Stable Diffusion) and focuses on functionality. - **Engine-Specific Parameters**: Uses `engineParameters` to directly specify an AI engine, exposing all available parameters for fine-tuned control.  **Permissions Required:** `prompt-svc:prompt:create`
     * Prompt an AI
     */
    prompt(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.promptRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * The only purpose of this \"endpoint\" is to export types otherwise not appearing in the API docs. This endpoint otherwise does nothing. Do not depend on this endpoint, only its types.
     * Prompt Types
     */
    promptTypesRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling promptTypes().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/prompt-svc/types`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: requestParameters['body'],
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => PromptSvcTypesResponseFromJSON(jsonValue));
        });
    }
    /**
     * The only purpose of this \"endpoint\" is to export types otherwise not appearing in the API docs. This endpoint otherwise does nothing. Do not depend on this endpoint, only its types.
     * Prompt Types
     */
    promptTypes(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.promptTypesRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Remove a prompt by ID.
     * Remove Prompt
     */
    removePromptRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling removePrompt().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/prompt-svc/remove`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: PromptSvcRemovePromptRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Remove a prompt by ID.
     * Remove Prompt
     */
    removePrompt(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.removePromptRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Subscribe to prompt responses by thread via Server-Sent Events (SSE). You can subscribe to threads before they are created. The streamed strings are of type `StreamChunk`, see the PromptTypes endpoint for more details.
     * Subscribe to Prompt Responses by Thread
     */
    subscribeToPromptResponsesRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['threadId'] == null) {
                throw new runtime.RequiredError('threadId', 'Required parameter "threadId" was null or undefined when calling subscribeToPromptResponses().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/prompt-svc/prompts/{threadId}/responses/subscribe`.replace(`{${"threadId"}}`, encodeURIComponent(String(requestParameters['threadId']))),
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            if (this.isJsonMime(response.headers.get('content-type'))) {
                return new runtime.JSONApiResponse(response);
            }
            else {
                return new runtime.TextApiResponse(response);
            }
        });
    }
    /**
     * Subscribe to prompt responses by thread via Server-Sent Events (SSE). You can subscribe to threads before they are created. The streamed strings are of type `StreamChunk`, see the PromptTypes endpoint for more details.
     * Subscribe to Prompt Responses by Thread
     */
    subscribeToPromptResponses(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.subscribeToPromptResponsesRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
}
