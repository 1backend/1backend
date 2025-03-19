/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
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
import localVarRequest from 'request';
import { ObjectSerializer, VoidAuth } from '../model/models';
import { ApiKeyAuth } from '../model/models';
import { HttpError } from './apis';
let defaultBasePath = 'http://localhost:58231';
// ===============================================
// This file is autogenerated - Please do not edit
// ===============================================
export var ChatSvcApiApiKeys;
(function (ChatSvcApiApiKeys) {
    ChatSvcApiApiKeys[ChatSvcApiApiKeys["BearerAuth"] = 0] = "BearerAuth";
})(ChatSvcApiApiKeys || (ChatSvcApiApiKeys = {}));
export class ChatSvcApi {
    constructor(basePathOrUsername, password, basePath) {
        this._basePath = defaultBasePath;
        this._defaultHeaders = {};
        this._useQuerystring = false;
        this.authentications = {
            'default': new VoidAuth(),
            'BearerAuth': new ApiKeyAuth('header', 'Authorization'),
        };
        this.interceptors = [];
        if (password) {
            if (basePath) {
                this.basePath = basePath;
            }
        }
        else {
            if (basePathOrUsername) {
                this.basePath = basePathOrUsername;
            }
        }
    }
    set useQuerystring(value) {
        this._useQuerystring = value;
    }
    set basePath(basePath) {
        this._basePath = basePath;
    }
    set defaultHeaders(defaultHeaders) {
        this._defaultHeaders = defaultHeaders;
    }
    get defaultHeaders() {
        return this._defaultHeaders;
    }
    get basePath() {
        return this._basePath;
    }
    setDefaultAuthentication(auth) {
        this.authentications.default = auth;
    }
    setApiKey(key, value) {
        this.authentications[ChatSvcApiApiKeys[key]].apiKey = value;
    }
    addInterceptor(interceptor) {
        this.interceptors.push(interceptor);
    }
    /**
     * Add a new message to a specific thread.
     * @summary Add Message
     * @param threadId Thread ID
     * @param body Add Message Request
     */
    addMessage(threadId_1, body_1) {
        return __awaiter(this, arguments, void 0, function* (threadId, body, options = { headers: {} }) {
            const localVarPath = this.basePath + '/chat-svc/thread/{threadId}/message'
                .replace('{' + 'threadId' + '}', encodeURIComponent(String(threadId)));
            let localVarQueryParameters = {};
            let localVarHeaderParams = Object.assign({}, this._defaultHeaders);
            const produces = ['application/json'];
            // give precedence to 'application/json'
            if (produces.indexOf('application/json') >= 0) {
                localVarHeaderParams.Accept = 'application/json';
            }
            else {
                localVarHeaderParams.Accept = produces.join(',');
            }
            let localVarFormParams = {};
            // verify required parameter 'threadId' is not null or undefined
            if (threadId === null || threadId === undefined) {
                throw new Error('Required parameter threadId was null or undefined when calling addMessage.');
            }
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new Error('Required parameter body was null or undefined when calling addMessage.');
            }
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            let localVarRequestOptions = {
                method: 'POST',
                qs: localVarQueryParameters,
                headers: localVarHeaderParams,
                uri: localVarPath,
                useQuerystring: this._useQuerystring,
                json: true,
                body: ObjectSerializer.serialize(body, "ChatSvcAddMessageRequest")
            };
            let authenticationPromise = Promise.resolve();
            if (this.authentications.BearerAuth.apiKey) {
                authenticationPromise = authenticationPromise.then(() => this.authentications.BearerAuth.applyToRequest(localVarRequestOptions));
            }
            authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));
            let interceptorPromise = authenticationPromise;
            for (const interceptor of this.interceptors) {
                interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
            }
            return interceptorPromise.then(() => {
                if (Object.keys(localVarFormParams).length) {
                    if (localVarUseFormData) {
                        localVarRequestOptions.formData = localVarFormParams;
                    }
                    else {
                        localVarRequestOptions.form = localVarFormParams;
                    }
                }
                return new Promise((resolve, reject) => {
                    localVarRequest(localVarRequestOptions, (error, response, body) => {
                        if (error) {
                            reject(error);
                        }
                        else {
                            if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                                body = ObjectSerializer.deserialize(body, "{ [key: string]: any; }");
                                resolve({ response: response, body: body });
                            }
                            else {
                                reject(new HttpError(response, body, response.statusCode));
                            }
                        }
                    });
                });
            });
        });
    }
    /**
     * Create a new chat thread and add the requesting user to it. Requires the `chat-svc:thread:create` permission.
     * @summary Add Thread
     * @param body Add Thread Request
     */
    addThread(body_1) {
        return __awaiter(this, arguments, void 0, function* (body, options = { headers: {} }) {
            const localVarPath = this.basePath + '/chat-svc/thread';
            let localVarQueryParameters = {};
            let localVarHeaderParams = Object.assign({}, this._defaultHeaders);
            const produces = ['application/json'];
            // give precedence to 'application/json'
            if (produces.indexOf('application/json') >= 0) {
                localVarHeaderParams.Accept = 'application/json';
            }
            else {
                localVarHeaderParams.Accept = produces.join(',');
            }
            let localVarFormParams = {};
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new Error('Required parameter body was null or undefined when calling addThread.');
            }
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            let localVarRequestOptions = {
                method: 'POST',
                qs: localVarQueryParameters,
                headers: localVarHeaderParams,
                uri: localVarPath,
                useQuerystring: this._useQuerystring,
                json: true,
                body: ObjectSerializer.serialize(body, "ChatSvcAddThreadRequest")
            };
            let authenticationPromise = Promise.resolve();
            if (this.authentications.BearerAuth.apiKey) {
                authenticationPromise = authenticationPromise.then(() => this.authentications.BearerAuth.applyToRequest(localVarRequestOptions));
            }
            authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));
            let interceptorPromise = authenticationPromise;
            for (const interceptor of this.interceptors) {
                interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
            }
            return interceptorPromise.then(() => {
                if (Object.keys(localVarFormParams).length) {
                    if (localVarUseFormData) {
                        localVarRequestOptions.formData = localVarFormParams;
                    }
                    else {
                        localVarRequestOptions.form = localVarFormParams;
                    }
                }
                return new Promise((resolve, reject) => {
                    localVarRequest(localVarRequestOptions, (error, response, body) => {
                        if (error) {
                            reject(error);
                        }
                        else {
                            if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                                body = ObjectSerializer.deserialize(body, "ChatSvcAddThreadResponse");
                                resolve({ response: response, body: body });
                            }
                            else {
                                reject(new HttpError(response, body, response.statusCode));
                            }
                        }
                    });
                });
            });
        });
    }
    /**
     * Delete a specific message from a chat thread by its ID
     * @summary Delete a Message
     * @param messageId Message ID
     */
    deleteMessage(messageId_1) {
        return __awaiter(this, arguments, void 0, function* (messageId, options = { headers: {} }) {
            const localVarPath = this.basePath + '/chat-svc/message/{messageId}'
                .replace('{' + 'messageId' + '}', encodeURIComponent(String(messageId)));
            let localVarQueryParameters = {};
            let localVarHeaderParams = Object.assign({}, this._defaultHeaders);
            const produces = ['application/json'];
            // give precedence to 'application/json'
            if (produces.indexOf('application/json') >= 0) {
                localVarHeaderParams.Accept = 'application/json';
            }
            else {
                localVarHeaderParams.Accept = produces.join(',');
            }
            let localVarFormParams = {};
            // verify required parameter 'messageId' is not null or undefined
            if (messageId === null || messageId === undefined) {
                throw new Error('Required parameter messageId was null or undefined when calling deleteMessage.');
            }
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            let localVarRequestOptions = {
                method: 'DELETE',
                qs: localVarQueryParameters,
                headers: localVarHeaderParams,
                uri: localVarPath,
                useQuerystring: this._useQuerystring,
                json: true,
            };
            let authenticationPromise = Promise.resolve();
            if (this.authentications.BearerAuth.apiKey) {
                authenticationPromise = authenticationPromise.then(() => this.authentications.BearerAuth.applyToRequest(localVarRequestOptions));
            }
            authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));
            let interceptorPromise = authenticationPromise;
            for (const interceptor of this.interceptors) {
                interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
            }
            return interceptorPromise.then(() => {
                if (Object.keys(localVarFormParams).length) {
                    if (localVarUseFormData) {
                        localVarRequestOptions.formData = localVarFormParams;
                    }
                    else {
                        localVarRequestOptions.form = localVarFormParams;
                    }
                }
                return new Promise((resolve, reject) => {
                    localVarRequest(localVarRequestOptions, (error, response, body) => {
                        if (error) {
                            reject(error);
                        }
                        else {
                            if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                                body = ObjectSerializer.deserialize(body, "{ [key: string]: any; }");
                                resolve({ response: response, body: body });
                            }
                            else {
                                reject(new HttpError(response, body, response.statusCode));
                            }
                        }
                    });
                });
            });
        });
    }
    /**
     * Delete a specific chat thread by its ID
     * @summary Delete a Thread
     * @param threadId Thread ID
     */
    deleteThread(threadId_1) {
        return __awaiter(this, arguments, void 0, function* (threadId, options = { headers: {} }) {
            const localVarPath = this.basePath + '/chat-svc/thread/{threadId}'
                .replace('{' + 'threadId' + '}', encodeURIComponent(String(threadId)));
            let localVarQueryParameters = {};
            let localVarHeaderParams = Object.assign({}, this._defaultHeaders);
            const produces = ['application/json'];
            // give precedence to 'application/json'
            if (produces.indexOf('application/json') >= 0) {
                localVarHeaderParams.Accept = 'application/json';
            }
            else {
                localVarHeaderParams.Accept = produces.join(',');
            }
            let localVarFormParams = {};
            // verify required parameter 'threadId' is not null or undefined
            if (threadId === null || threadId === undefined) {
                throw new Error('Required parameter threadId was null or undefined when calling deleteThread.');
            }
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            let localVarRequestOptions = {
                method: 'DELETE',
                qs: localVarQueryParameters,
                headers: localVarHeaderParams,
                uri: localVarPath,
                useQuerystring: this._useQuerystring,
                json: true,
            };
            let authenticationPromise = Promise.resolve();
            if (this.authentications.BearerAuth.apiKey) {
                authenticationPromise = authenticationPromise.then(() => this.authentications.BearerAuth.applyToRequest(localVarRequestOptions));
            }
            authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));
            let interceptorPromise = authenticationPromise;
            for (const interceptor of this.interceptors) {
                interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
            }
            return interceptorPromise.then(() => {
                if (Object.keys(localVarFormParams).length) {
                    if (localVarUseFormData) {
                        localVarRequestOptions.formData = localVarFormParams;
                    }
                    else {
                        localVarRequestOptions.form = localVarFormParams;
                    }
                }
                return new Promise((resolve, reject) => {
                    localVarRequest(localVarRequestOptions, (error, response, body) => {
                        if (error) {
                            reject(error);
                        }
                        else {
                            if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                                body = ObjectSerializer.deserialize(body, "{ [key: string]: any; }");
                                resolve({ response: response, body: body });
                            }
                            else {
                                reject(new HttpError(response, body, response.statusCode));
                            }
                        }
                    });
                });
            });
        });
    }
    /**
     * Events is a dummy endpoint to display documentation about the events that this service emits.
     * @summary Events
     */
    events() {
        return __awaiter(this, arguments, void 0, function* (options = { headers: {} }) {
            const localVarPath = this.basePath + '/chat-svc/events';
            let localVarQueryParameters = {};
            let localVarHeaderParams = Object.assign({}, this._defaultHeaders);
            const produces = ['application/json'];
            // give precedence to 'application/json'
            if (produces.indexOf('application/json') >= 0) {
                localVarHeaderParams.Accept = 'application/json';
            }
            else {
                localVarHeaderParams.Accept = produces.join(',');
            }
            let localVarFormParams = {};
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            let localVarRequestOptions = {
                method: 'GET',
                qs: localVarQueryParameters,
                headers: localVarHeaderParams,
                uri: localVarPath,
                useQuerystring: this._useQuerystring,
                json: true,
            };
            let authenticationPromise = Promise.resolve();
            authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));
            let interceptorPromise = authenticationPromise;
            for (const interceptor of this.interceptors) {
                interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
            }
            return interceptorPromise.then(() => {
                if (Object.keys(localVarFormParams).length) {
                    if (localVarUseFormData) {
                        localVarRequestOptions.formData = localVarFormParams;
                    }
                    else {
                        localVarRequestOptions.form = localVarFormParams;
                    }
                }
                return new Promise((resolve, reject) => {
                    localVarRequest(localVarRequestOptions, (error, response, body) => {
                        if (error) {
                            reject(error);
                        }
                        else {
                            if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                                body = ObjectSerializer.deserialize(body, "ChatSvcEventThreadUpdate");
                                resolve({ response: response, body: body });
                            }
                            else {
                                reject(new HttpError(response, body, response.statusCode));
                            }
                        }
                    });
                });
            });
        });
    }
    /**
     * Fetch information about a specific chat message by its ID
     * @summary Get Message
     * @param messageId Message ID
     */
    getMessage(messageId_1) {
        return __awaiter(this, arguments, void 0, function* (messageId, options = { headers: {} }) {
            const localVarPath = this.basePath + '/chat-svc/message/{messageId}'
                .replace('{' + 'messageId' + '}', encodeURIComponent(String(messageId)));
            let localVarQueryParameters = {};
            let localVarHeaderParams = Object.assign({}, this._defaultHeaders);
            const produces = ['application/json'];
            // give precedence to 'application/json'
            if (produces.indexOf('application/json') >= 0) {
                localVarHeaderParams.Accept = 'application/json';
            }
            else {
                localVarHeaderParams.Accept = produces.join(',');
            }
            let localVarFormParams = {};
            // verify required parameter 'messageId' is not null or undefined
            if (messageId === null || messageId === undefined) {
                throw new Error('Required parameter messageId was null or undefined when calling getMessage.');
            }
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            let localVarRequestOptions = {
                method: 'GET',
                qs: localVarQueryParameters,
                headers: localVarHeaderParams,
                uri: localVarPath,
                useQuerystring: this._useQuerystring,
                json: true,
            };
            let authenticationPromise = Promise.resolve();
            if (this.authentications.BearerAuth.apiKey) {
                authenticationPromise = authenticationPromise.then(() => this.authentications.BearerAuth.applyToRequest(localVarRequestOptions));
            }
            authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));
            let interceptorPromise = authenticationPromise;
            for (const interceptor of this.interceptors) {
                interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
            }
            return interceptorPromise.then(() => {
                if (Object.keys(localVarFormParams).length) {
                    if (localVarUseFormData) {
                        localVarRequestOptions.formData = localVarFormParams;
                    }
                    else {
                        localVarRequestOptions.form = localVarFormParams;
                    }
                }
                return new Promise((resolve, reject) => {
                    localVarRequest(localVarRequestOptions, (error, response, body) => {
                        if (error) {
                            reject(error);
                        }
                        else {
                            if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                                body = ObjectSerializer.deserialize(body, "ChatSvcGetMessageResponse");
                                resolve({ response: response, body: body });
                            }
                            else {
                                reject(new HttpError(response, body, response.statusCode));
                            }
                        }
                    });
                });
            });
        });
    }
    /**
     * Fetch messages (and associated assets) for a specific chat thread.
     * @summary List Messages
     * @param threadId Thread ID
     */
    getMessages(threadId_1) {
        return __awaiter(this, arguments, void 0, function* (threadId, options = { headers: {} }) {
            const localVarPath = this.basePath + '/chat-svc/thread/{threadId}/messages'
                .replace('{' + 'threadId' + '}', encodeURIComponent(String(threadId)));
            let localVarQueryParameters = {};
            let localVarHeaderParams = Object.assign({}, this._defaultHeaders);
            const produces = ['application/json'];
            // give precedence to 'application/json'
            if (produces.indexOf('application/json') >= 0) {
                localVarHeaderParams.Accept = 'application/json';
            }
            else {
                localVarHeaderParams.Accept = produces.join(',');
            }
            let localVarFormParams = {};
            // verify required parameter 'threadId' is not null or undefined
            if (threadId === null || threadId === undefined) {
                throw new Error('Required parameter threadId was null or undefined when calling getMessages.');
            }
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            let localVarRequestOptions = {
                method: 'POST',
                qs: localVarQueryParameters,
                headers: localVarHeaderParams,
                uri: localVarPath,
                useQuerystring: this._useQuerystring,
                json: true,
            };
            let authenticationPromise = Promise.resolve();
            if (this.authentications.BearerAuth.apiKey) {
                authenticationPromise = authenticationPromise.then(() => this.authentications.BearerAuth.applyToRequest(localVarRequestOptions));
            }
            authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));
            let interceptorPromise = authenticationPromise;
            for (const interceptor of this.interceptors) {
                interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
            }
            return interceptorPromise.then(() => {
                if (Object.keys(localVarFormParams).length) {
                    if (localVarUseFormData) {
                        localVarRequestOptions.formData = localVarFormParams;
                    }
                    else {
                        localVarRequestOptions.form = localVarFormParams;
                    }
                }
                return new Promise((resolve, reject) => {
                    localVarRequest(localVarRequestOptions, (error, response, body) => {
                        if (error) {
                            reject(error);
                        }
                        else {
                            if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                                body = ObjectSerializer.deserialize(body, "ChatSvcGetMessagesResponse");
                                resolve({ response: response, body: body });
                            }
                            else {
                                reject(new HttpError(response, body, response.statusCode));
                            }
                        }
                    });
                });
            });
        });
    }
    /**
     * Fetch information about a specific chat thread by its ID
     * @summary Get Thread
     * @param threadId Thread ID
     */
    getThread(threadId_1) {
        return __awaiter(this, arguments, void 0, function* (threadId, options = { headers: {} }) {
            const localVarPath = this.basePath + '/chat-svc/thread/{threadId}'
                .replace('{' + 'threadId' + '}', encodeURIComponent(String(threadId)));
            let localVarQueryParameters = {};
            let localVarHeaderParams = Object.assign({}, this._defaultHeaders);
            const produces = ['application/json'];
            // give precedence to 'application/json'
            if (produces.indexOf('application/json') >= 0) {
                localVarHeaderParams.Accept = 'application/json';
            }
            else {
                localVarHeaderParams.Accept = produces.join(',');
            }
            let localVarFormParams = {};
            // verify required parameter 'threadId' is not null or undefined
            if (threadId === null || threadId === undefined) {
                throw new Error('Required parameter threadId was null or undefined when calling getThread.');
            }
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            let localVarRequestOptions = {
                method: 'GET',
                qs: localVarQueryParameters,
                headers: localVarHeaderParams,
                uri: localVarPath,
                useQuerystring: this._useQuerystring,
                json: true,
            };
            let authenticationPromise = Promise.resolve();
            if (this.authentications.BearerAuth.apiKey) {
                authenticationPromise = authenticationPromise.then(() => this.authentications.BearerAuth.applyToRequest(localVarRequestOptions));
            }
            authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));
            let interceptorPromise = authenticationPromise;
            for (const interceptor of this.interceptors) {
                interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
            }
            return interceptorPromise.then(() => {
                if (Object.keys(localVarFormParams).length) {
                    if (localVarUseFormData) {
                        localVarRequestOptions.formData = localVarFormParams;
                    }
                    else {
                        localVarRequestOptions.form = localVarFormParams;
                    }
                }
                return new Promise((resolve, reject) => {
                    localVarRequest(localVarRequestOptions, (error, response, body) => {
                        if (error) {
                            reject(error);
                        }
                        else {
                            if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                                body = ObjectSerializer.deserialize(body, "ChatSvcGetThreadResponse");
                                resolve({ response: response, body: body });
                            }
                            else {
                                reject(new HttpError(response, body, response.statusCode));
                            }
                        }
                    });
                });
            });
        });
    }
    /**
     * Fetch all chat threads associated with a specific user
     * @summary Get Threads
     * @param body Get Threads Request
     */
    getThreads(body_1) {
        return __awaiter(this, arguments, void 0, function* (body, options = { headers: {} }) {
            const localVarPath = this.basePath + '/chat-svc/threads';
            let localVarQueryParameters = {};
            let localVarHeaderParams = Object.assign({}, this._defaultHeaders);
            const produces = ['application/json'];
            // give precedence to 'application/json'
            if (produces.indexOf('application/json') >= 0) {
                localVarHeaderParams.Accept = 'application/json';
            }
            else {
                localVarHeaderParams.Accept = produces.join(',');
            }
            let localVarFormParams = {};
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            let localVarRequestOptions = {
                method: 'POST',
                qs: localVarQueryParameters,
                headers: localVarHeaderParams,
                uri: localVarPath,
                useQuerystring: this._useQuerystring,
                json: true,
                body: ObjectSerializer.serialize(body, "object")
            };
            let authenticationPromise = Promise.resolve();
            if (this.authentications.BearerAuth.apiKey) {
                authenticationPromise = authenticationPromise.then(() => this.authentications.BearerAuth.applyToRequest(localVarRequestOptions));
            }
            authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));
            let interceptorPromise = authenticationPromise;
            for (const interceptor of this.interceptors) {
                interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
            }
            return interceptorPromise.then(() => {
                if (Object.keys(localVarFormParams).length) {
                    if (localVarUseFormData) {
                        localVarRequestOptions.formData = localVarFormParams;
                    }
                    else {
                        localVarRequestOptions.form = localVarFormParams;
                    }
                }
                return new Promise((resolve, reject) => {
                    localVarRequest(localVarRequestOptions, (error, response, body) => {
                        if (error) {
                            reject(error);
                        }
                        else {
                            if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                                body = ObjectSerializer.deserialize(body, "ChatSvcGetThreadsResponse");
                                resolve({ response: response, body: body });
                            }
                            else {
                                reject(new HttpError(response, body, response.statusCode));
                            }
                        }
                    });
                });
            });
        });
    }
    /**
     * Modify the details of a specific chat thread
     * @summary Update Thread
     * @param threadId Thread ID
     * @param body Update Thread Request
     */
    updateThread(threadId_1, body_1) {
        return __awaiter(this, arguments, void 0, function* (threadId, body, options = { headers: {} }) {
            const localVarPath = this.basePath + '/chat-svc/thread/{threadId}'
                .replace('{' + 'threadId' + '}', encodeURIComponent(String(threadId)));
            let localVarQueryParameters = {};
            let localVarHeaderParams = Object.assign({}, this._defaultHeaders);
            const produces = ['application/json'];
            // give precedence to 'application/json'
            if (produces.indexOf('application/json') >= 0) {
                localVarHeaderParams.Accept = 'application/json';
            }
            else {
                localVarHeaderParams.Accept = produces.join(',');
            }
            let localVarFormParams = {};
            // verify required parameter 'threadId' is not null or undefined
            if (threadId === null || threadId === undefined) {
                throw new Error('Required parameter threadId was null or undefined when calling updateThread.');
            }
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new Error('Required parameter body was null or undefined when calling updateThread.');
            }
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            let localVarRequestOptions = {
                method: 'PUT',
                qs: localVarQueryParameters,
                headers: localVarHeaderParams,
                uri: localVarPath,
                useQuerystring: this._useQuerystring,
                json: true,
                body: ObjectSerializer.serialize(body, "ChatSvcUpdateThreadRequest")
            };
            let authenticationPromise = Promise.resolve();
            if (this.authentications.BearerAuth.apiKey) {
                authenticationPromise = authenticationPromise.then(() => this.authentications.BearerAuth.applyToRequest(localVarRequestOptions));
            }
            authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));
            let interceptorPromise = authenticationPromise;
            for (const interceptor of this.interceptors) {
                interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
            }
            return interceptorPromise.then(() => {
                if (Object.keys(localVarFormParams).length) {
                    if (localVarUseFormData) {
                        localVarRequestOptions.formData = localVarFormParams;
                    }
                    else {
                        localVarRequestOptions.form = localVarFormParams;
                    }
                }
                return new Promise((resolve, reject) => {
                    localVarRequest(localVarRequestOptions, (error, response, body) => {
                        if (error) {
                            reject(error);
                        }
                        else {
                            if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                                body = ObjectSerializer.deserialize(body, "ChatSvcAddThreadResponse");
                                resolve({ response: response, body: body });
                            }
                            else {
                                reject(new HttpError(response, body, response.statusCode));
                            }
                        }
                    });
                });
            });
        });
    }
}
