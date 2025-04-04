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
import * as runtime from '../runtime';
import type { ChatSvcAddMessageRequest, ChatSvcAddThreadRequest, ChatSvcAddThreadResponse, ChatSvcEventThreadUpdate, ChatSvcGetMessageResponse, ChatSvcGetMessagesResponse, ChatSvcGetThreadResponse, ChatSvcGetThreadsResponse, ChatSvcUpdateThreadRequest } from '../models/index';
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
export declare class ChatSvcApi extends runtime.BaseAPI {
    /**
     * Add a new message to a specific thread.
     * Add Message
     */
    addMessageRaw(requestParameters: AddMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{
        [key: string]: any;
    }>>;
    /**
     * Add a new message to a specific thread.
     * Add Message
     */
    addMessage(requestParameters: AddMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{
        [key: string]: any;
    }>;
    /**
     * Create a new chat thread and add the requesting user to it. Requires the `chat-svc:thread:create` permission.
     * Add Thread
     */
    addThreadRaw(requestParameters: AddThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcAddThreadResponse>>;
    /**
     * Create a new chat thread and add the requesting user to it. Requires the `chat-svc:thread:create` permission.
     * Add Thread
     */
    addThread(requestParameters: AddThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcAddThreadResponse>;
    /**
     * Delete a specific message from a chat thread by its ID
     * Delete a Message
     */
    deleteMessageRaw(requestParameters: DeleteMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{
        [key: string]: any;
    }>>;
    /**
     * Delete a specific message from a chat thread by its ID
     * Delete a Message
     */
    deleteMessage(requestParameters: DeleteMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{
        [key: string]: any;
    }>;
    /**
     * Delete a specific chat thread by its ID
     * Delete a Thread
     */
    deleteThreadRaw(requestParameters: DeleteThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{
        [key: string]: any;
    }>>;
    /**
     * Delete a specific chat thread by its ID
     * Delete a Thread
     */
    deleteThread(requestParameters: DeleteThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{
        [key: string]: any;
    }>;
    /**
     * Events is a dummy endpoint to display documentation about the events that this service emits.
     * Events
     */
    eventsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcEventThreadUpdate>>;
    /**
     * Events is a dummy endpoint to display documentation about the events that this service emits.
     * Events
     */
    events(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcEventThreadUpdate>;
    /**
     * Fetch information about a specific chat message by its ID
     * Get Message
     */
    getMessageRaw(requestParameters: GetMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcGetMessageResponse>>;
    /**
     * Fetch information about a specific chat message by its ID
     * Get Message
     */
    getMessage(requestParameters: GetMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcGetMessageResponse>;
    /**
     * Fetch messages (and associated assets) for a specific chat thread.
     * List Messages
     */
    getMessagesRaw(requestParameters: GetMessagesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcGetMessagesResponse>>;
    /**
     * Fetch messages (and associated assets) for a specific chat thread.
     * List Messages
     */
    getMessages(requestParameters: GetMessagesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcGetMessagesResponse>;
    /**
     * Fetch information about a specific chat thread by its ID
     * Get Thread
     */
    getThreadRaw(requestParameters: GetThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcGetThreadResponse>>;
    /**
     * Fetch information about a specific chat thread by its ID
     * Get Thread
     */
    getThread(requestParameters: GetThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcGetThreadResponse>;
    /**
     * Fetch all chat threads associated with a specific user
     * Get Threads
     */
    getThreadsRaw(requestParameters: GetThreadsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcGetThreadsResponse>>;
    /**
     * Fetch all chat threads associated with a specific user
     * Get Threads
     */
    getThreads(requestParameters?: GetThreadsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcGetThreadsResponse>;
    /**
     * Modify the details of a specific chat thread
     * Update Thread
     */
    updateThreadRaw(requestParameters: UpdateThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatSvcAddThreadResponse>>;
    /**
     * Modify the details of a specific chat thread
     * Update Thread
     */
    updateThread(requestParameters: UpdateThreadRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatSvcAddThreadResponse>;
}
