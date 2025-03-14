/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
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
import { ChatSvcAddMessageRequestToJSON, ChatSvcAddThreadRequestToJSON, ChatSvcAddThreadResponseFromJSON, ChatSvcEventThreadUpdateFromJSON, ChatSvcGetMessageResponseFromJSON, ChatSvcGetMessagesResponseFromJSON, ChatSvcGetThreadResponseFromJSON, ChatSvcGetThreadsResponseFromJSON, ChatSvcUpdateThreadRequestToJSON, } from '../models/index';
/**
 *
 */
export class ChatSvcApi extends runtime.BaseAPI {
    /**
     * Add a new message to a specific thread.
     * Add Message
     */
    addMessageRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['threadId'] == null) {
                throw new runtime.RequiredError('threadId', 'Required parameter "threadId" was null or undefined when calling addMessage().');
            }
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling addMessage().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/chat-svc/thread/{threadId}/message`.replace(`{${"threadId"}}`, encodeURIComponent(String(requestParameters['threadId']))),
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: ChatSvcAddMessageRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Add a new message to a specific thread.
     * Add Message
     */
    addMessage(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.addMessageRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Create a new chat thread and add the requesting user to it. Requires the `chat-svc:thread:create` permission.
     * Add Thread
     */
    addThreadRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling addThread().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/chat-svc/thread`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: ChatSvcAddThreadRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcAddThreadResponseFromJSON(jsonValue));
        });
    }
    /**
     * Create a new chat thread and add the requesting user to it. Requires the `chat-svc:thread:create` permission.
     * Add Thread
     */
    addThread(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.addThreadRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Delete a specific message from a chat thread by its ID
     * Delete a Message
     */
    deleteMessageRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['messageId'] == null) {
                throw new runtime.RequiredError('messageId', 'Required parameter "messageId" was null or undefined when calling deleteMessage().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/chat-svc/message/{messageId}`.replace(`{${"messageId"}}`, encodeURIComponent(String(requestParameters['messageId']))),
                method: 'DELETE',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Delete a specific message from a chat thread by its ID
     * Delete a Message
     */
    deleteMessage(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.deleteMessageRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Delete a specific chat thread by its ID
     * Delete a Thread
     */
    deleteThreadRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['threadId'] == null) {
                throw new runtime.RequiredError('threadId', 'Required parameter "threadId" was null or undefined when calling deleteThread().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/chat-svc/thread/{threadId}`.replace(`{${"threadId"}}`, encodeURIComponent(String(requestParameters['threadId']))),
                method: 'DELETE',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Delete a specific chat thread by its ID
     * Delete a Thread
     */
    deleteThread(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.deleteThreadRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Events is a dummy endpoint to display documentation about the events that this service emits.
     * Events
     */
    eventsRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            const response = yield this.request({
                path: `/chat-svc/events`,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcEventThreadUpdateFromJSON(jsonValue));
        });
    }
    /**
     * Events is a dummy endpoint to display documentation about the events that this service emits.
     * Events
     */
    events(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.eventsRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * Fetch information about a specific chat message by its ID
     * Get Message
     */
    getMessageRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['messageId'] == null) {
                throw new runtime.RequiredError('messageId', 'Required parameter "messageId" was null or undefined when calling getMessage().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/chat-svc/message/{messageId}`.replace(`{${"messageId"}}`, encodeURIComponent(String(requestParameters['messageId']))),
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcGetMessageResponseFromJSON(jsonValue));
        });
    }
    /**
     * Fetch information about a specific chat message by its ID
     * Get Message
     */
    getMessage(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.getMessageRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Fetch messages (and associated assets) for a specific chat thread.
     * List Messages
     */
    getMessagesRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['threadId'] == null) {
                throw new runtime.RequiredError('threadId', 'Required parameter "threadId" was null or undefined when calling getMessages().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/chat-svc/thread/{threadId}/messages`.replace(`{${"threadId"}}`, encodeURIComponent(String(requestParameters['threadId']))),
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcGetMessagesResponseFromJSON(jsonValue));
        });
    }
    /**
     * Fetch messages (and associated assets) for a specific chat thread.
     * List Messages
     */
    getMessages(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.getMessagesRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Fetch information about a specific chat thread by its ID
     * Get Thread
     */
    getThreadRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['threadId'] == null) {
                throw new runtime.RequiredError('threadId', 'Required parameter "threadId" was null or undefined when calling getThread().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/chat-svc/thread/{threadId}`.replace(`{${"threadId"}}`, encodeURIComponent(String(requestParameters['threadId']))),
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcGetThreadResponseFromJSON(jsonValue));
        });
    }
    /**
     * Fetch information about a specific chat thread by its ID
     * Get Thread
     */
    getThread(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.getThreadRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Fetch all chat threads associated with a specific user
     * Get Threads
     */
    getThreadsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/chat-svc/threads`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: requestParameters['body'],
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcGetThreadsResponseFromJSON(jsonValue));
        });
    }
    /**
     * Fetch all chat threads associated with a specific user
     * Get Threads
     */
    getThreads() {
        return __awaiter(this, arguments, void 0, function* (requestParameters = {}, initOverrides) {
            const response = yield this.getThreadsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Modify the details of a specific chat thread
     * Update Thread
     */
    updateThreadRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['threadId'] == null) {
                throw new runtime.RequiredError('threadId', 'Required parameter "threadId" was null or undefined when calling updateThread().');
            }
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling updateThread().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/chat-svc/thread/{threadId}`.replace(`{${"threadId"}}`, encodeURIComponent(String(requestParameters['threadId']))),
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: ChatSvcUpdateThreadRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ChatSvcAddThreadResponseFromJSON(jsonValue));
        });
    }
    /**
     * Modify the details of a specific chat thread
     * Update Thread
     */
    updateThread(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.updateThreadRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
}
