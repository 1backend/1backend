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
  ChatSvcAddMessageRequest,
  ChatSvcAddThreadRequest,
  ChatSvcAddThreadResponse,
  ChatSvcEventThreadUpdate,
  ChatSvcGetMessageResponse,
  ChatSvcGetMessagesResponse,
  ChatSvcGetThreadResponse,
  ChatSvcGetThreadsResponse,
  ChatSvcUpdateThreadRequest,
} from '../models/index';
import {
    ChatSvcAddMessageRequestFromJSON,
    ChatSvcAddMessageRequestToJSON,
    ChatSvcAddThreadRequestFromJSON,
    ChatSvcAddThreadRequestToJSON,
    ChatSvcAddThreadResponseFromJSON,
    ChatSvcAddThreadResponseToJSON,
    ChatSvcEventThreadUpdateFromJSON,
    ChatSvcEventThreadUpdateToJSON,
    ChatSvcGetMessageResponseFromJSON,
    ChatSvcGetMessageResponseToJSON,
    ChatSvcGetMessagesResponseFromJSON,
    ChatSvcGetMessagesResponseToJSON,
    ChatSvcGetThreadResponseFromJSON,
    ChatSvcGetThreadResponseToJSON,
    ChatSvcGetThreadsResponseFromJSON,
    ChatSvcGetThreadsResponseToJSON,
    ChatSvcUpdateThreadRequestFromJSON,
    ChatSvcUpdateThreadRequestToJSON,
} from '../models/index';

export interface AddMessageRequest {
    threadId: string;
    body: ChatSvcAddMessageRequest;
}

export interface AddThreadRequest {
    body: ChatSvcAddThreadRequest;
}

export interface DeleteMessageRequest {
    messageId: string;
}

export interface DeleteThreadRequest {
    threadId: string;
}

export interface GetMessageRequest {
    messageId: string;
}

export interface GetMessagesRequest {
    threadId: string;
}

export interface GetThreadRequest {
    threadId: string;
}

export interface GetThreadsRequest {
    body?: object;
}

export interface UpdateThreadRequest {
    threadId: string;
    body: ChatSvcUpdateThreadRequest;
}

/**
 * 
 */
export class ChatSvcApi extends runtime.BaseAPI {

    /**
     * Add a new message to a specific thread.
     * Add Message
     */
    async addMessageRaw(requestParameters: AddMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{ [key: string]: any; }>> {
        if (requestParameters['threadId'] == null) {
            throw new runtime.RequiredError(
                'threadId',
                'Required parameter "threadId" was null or undefined when calling addMessage().'
            );
        }

        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling addMessage().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/chat-svc/thread/{threadId}/message`.replace(`{${"threadId"}}`, encodeURIComponent(String(requestParameters['threadId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ChatSvcAddMessageRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Add a new message to a specific thread.
     * Add Message
     */
    async addMessage(requestParameters: AddMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{ [key: string]: any; }> {
        const response = await this.addMessageRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a new chat thread and add the requesting user to it. Requires the `chat-svc:thread:create` permission.
     * Add Thread
     */
    async addThreadRaw(requestParameters: AddThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcAddThreadResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling addThread().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/chat-svc/thread`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ChatSvcAddThreadRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcAddThreadResponseFromJSON(jsonValue));
    }

    /**
     * Create a new chat thread and add the requesting user to it. Requires the `chat-svc:thread:create` permission.
     * Add Thread
     */
    async addThread(requestParameters: AddThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcAddThreadResponse> {
        const response = await this.addThreadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete a specific message from a chat thread by its ID
     * Delete a Message
     */
    async deleteMessageRaw(requestParameters: DeleteMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{ [key: string]: any; }>> {
        if (requestParameters['messageId'] == null) {
            throw new runtime.RequiredError(
                'messageId',
                'Required parameter "messageId" was null or undefined when calling deleteMessage().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/chat-svc/message/{messageId}`.replace(`{${"messageId"}}`, encodeURIComponent(String(requestParameters['messageId']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Delete a specific message from a chat thread by its ID
     * Delete a Message
     */
    async deleteMessage(requestParameters: DeleteMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{ [key: string]: any; }> {
        const response = await this.deleteMessageRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete a specific chat thread by its ID
     * Delete a Thread
     */
    async deleteThreadRaw(requestParameters: DeleteThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{ [key: string]: any; }>> {
        if (requestParameters['threadId'] == null) {
            throw new runtime.RequiredError(
                'threadId',
                'Required parameter "threadId" was null or undefined when calling deleteThread().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/chat-svc/thread/{threadId}`.replace(`{${"threadId"}}`, encodeURIComponent(String(requestParameters['threadId']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Delete a specific chat thread by its ID
     * Delete a Thread
     */
    async deleteThread(requestParameters: DeleteThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{ [key: string]: any; }> {
        const response = await this.deleteThreadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Events is a dummy endpoint to display documentation about the events that this service emits.
     * Events
     */
    async eventsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcEventThreadUpdate>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/chat-svc/events`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcEventThreadUpdateFromJSON(jsonValue));
    }

    /**
     * Events is a dummy endpoint to display documentation about the events that this service emits.
     * Events
     */
    async events(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcEventThreadUpdate> {
        const response = await this.eventsRaw(initOverrides);
        return await response.value();
    }

    /**
     * Fetch information about a specific chat message by its ID
     * Get Message
     */
    async getMessageRaw(requestParameters: GetMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcGetMessageResponse>> {
        if (requestParameters['messageId'] == null) {
            throw new runtime.RequiredError(
                'messageId',
                'Required parameter "messageId" was null or undefined when calling getMessage().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/chat-svc/message/{messageId}`.replace(`{${"messageId"}}`, encodeURIComponent(String(requestParameters['messageId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcGetMessageResponseFromJSON(jsonValue));
    }

    /**
     * Fetch information about a specific chat message by its ID
     * Get Message
     */
    async getMessage(requestParameters: GetMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcGetMessageResponse> {
        const response = await this.getMessageRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Fetch messages (and associated assets) for a specific chat thread.
     * List Messages
     */
    async getMessagesRaw(requestParameters: GetMessagesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcGetMessagesResponse>> {
        if (requestParameters['threadId'] == null) {
            throw new runtime.RequiredError(
                'threadId',
                'Required parameter "threadId" was null or undefined when calling getMessages().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/chat-svc/thread/{threadId}/messages`.replace(`{${"threadId"}}`, encodeURIComponent(String(requestParameters['threadId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcGetMessagesResponseFromJSON(jsonValue));
    }

    /**
     * Fetch messages (and associated assets) for a specific chat thread.
     * List Messages
     */
    async getMessages(requestParameters: GetMessagesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcGetMessagesResponse> {
        const response = await this.getMessagesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Fetch information about a specific chat thread by its ID
     * Get Thread
     */
    async getThreadRaw(requestParameters: GetThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcGetThreadResponse>> {
        if (requestParameters['threadId'] == null) {
            throw new runtime.RequiredError(
                'threadId',
                'Required parameter "threadId" was null or undefined when calling getThread().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/chat-svc/thread/{threadId}`.replace(`{${"threadId"}}`, encodeURIComponent(String(requestParameters['threadId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcGetThreadResponseFromJSON(jsonValue));
    }

    /**
     * Fetch information about a specific chat thread by its ID
     * Get Thread
     */
    async getThread(requestParameters: GetThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcGetThreadResponse> {
        const response = await this.getThreadRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Fetch all chat threads associated with a specific user
     * Get Threads
     */
    async getThreadsRaw(requestParameters: GetThreadsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcGetThreadsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/chat-svc/threads`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters['body'] as any,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcGetThreadsResponseFromJSON(jsonValue));
    }

    /**
     * Fetch all chat threads associated with a specific user
     * Get Threads
     */
    async getThreads(requestParameters: GetThreadsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcGetThreadsResponse> {
        const response = await this.getThreadsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Modify the details of a specific chat thread
     * Update Thread
     */
    async updateThreadRaw(requestParameters: UpdateThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcAddThreadResponse>> {
        if (requestParameters['threadId'] == null) {
            throw new runtime.RequiredError(
                'threadId',
                'Required parameter "threadId" was null or undefined when calling updateThread().'
            );
        }

        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling updateThread().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/chat-svc/thread/{threadId}`.replace(`{${"threadId"}}`, encodeURIComponent(String(requestParameters['threadId']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: ChatSvcUpdateThreadRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcAddThreadResponseFromJSON(jsonValue));
    }

    /**
     * Modify the details of a specific chat thread
     * Update Thread
     */
    async updateThread(requestParameters: UpdateThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcAddThreadResponse> {
        const response = await this.updateThreadRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
