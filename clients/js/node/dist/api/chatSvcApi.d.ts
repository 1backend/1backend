/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import http from 'http';
import { ChatSvcAddMessageRequest } from '../model/chatSvcAddMessageRequest';
import { ChatSvcAddThreadRequest } from '../model/chatSvcAddThreadRequest';
import { ChatSvcAddThreadResponse } from '../model/chatSvcAddThreadResponse';
import { ChatSvcEventThreadUpdate } from '../model/chatSvcEventThreadUpdate';
import { ChatSvcGetMessageResponse } from '../model/chatSvcGetMessageResponse';
import { ChatSvcGetMessagesResponse } from '../model/chatSvcGetMessagesResponse';
import { ChatSvcGetThreadResponse } from '../model/chatSvcGetThreadResponse';
import { ChatSvcGetThreadsResponse } from '../model/chatSvcGetThreadsResponse';
import { ChatSvcUpdateThreadRequest } from '../model/chatSvcUpdateThreadRequest';
import { Authentication, Interceptor } from '../model/models';
import { ApiKeyAuth } from '../model/models';
export declare enum ChatSvcApiApiKeys {
    BearerAuth = 0
}
export declare class ChatSvcApi {
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
    setApiKey(key: ChatSvcApiApiKeys, value: string): void;
    addInterceptor(interceptor: Interceptor): void;
    /**
     * Add a new message to a specific thread.
     * @summary Add Message
     * @param threadId Thread ID
     * @param body Add Message Request
     */
    addMessage(threadId: string, body: ChatSvcAddMessageRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: {
            [key: string]: any;
        };
    }>;
    /**
     * Create a new chat thread and add the requesting user to it. Requires the `chat-svc:thread:create` permission.
     * @summary Add Thread
     * @param body Add Thread Request
     */
    addThread(body: ChatSvcAddThreadRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ChatSvcAddThreadResponse;
    }>;
    /**
     * Delete a specific message from a chat thread by its ID
     * @summary Delete a Message
     * @param messageId Message ID
     */
    deleteMessage(messageId: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: {
            [key: string]: any;
        };
    }>;
    /**
     * Delete a specific chat thread by its ID
     * @summary Delete a Thread
     * @param threadId Thread ID
     */
    deleteThread(threadId: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: {
            [key: string]: any;
        };
    }>;
    /**
     * Events is a dummy endpoint to display documentation about the events that this service emits.
     * @summary Events
     */
    events(options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ChatSvcEventThreadUpdate;
    }>;
    /**
     * Fetch information about a specific chat message by its ID
     * @summary Get Message
     * @param messageId Message ID
     */
    getMessage(messageId: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ChatSvcGetMessageResponse;
    }>;
    /**
     * Fetch messages (and associated assets) for a specific chat thread.
     * @summary List Messages
     * @param threadId Thread ID
     */
    getMessages(threadId: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ChatSvcGetMessagesResponse;
    }>;
    /**
     * Fetch information about a specific chat thread by its ID
     * @summary Get Thread
     * @param threadId Thread ID
     */
    getThread(threadId: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ChatSvcGetThreadResponse;
    }>;
    /**
     * Fetch all chat threads associated with a specific user
     * @summary Get Threads
     * @param body Get Threads Request
     */
    getThreads(body?: object, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ChatSvcGetThreadsResponse;
    }>;
    /**
     * Modify the details of a specific chat thread
     * @summary Update Thread
     * @param threadId Thread ID
     * @param body Update Thread Request
     */
    updateThread(threadId: string, body: ChatSvcUpdateThreadRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ChatSvcAddThreadResponse;
    }>;
}
