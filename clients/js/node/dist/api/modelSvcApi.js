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
let defaultBasePath = 'http://localhost:11337';
// ===============================================
// This file is autogenerated - Please do not edit
// ===============================================
export var ModelSvcApiApiKeys;
(function (ModelSvcApiApiKeys) {
    ModelSvcApiApiKeys[ModelSvcApiApiKeys["BearerAuth"] = 0] = "BearerAuth";
})(ModelSvcApiApiKeys || (ModelSvcApiApiKeys = {}));
export class ModelSvcApi {
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
        this.authentications[ModelSvcApiApiKeys[key]].apiKey = value;
    }
    addInterceptor(interceptor) {
        this.interceptors.push(interceptor);
    }
    /**
     * Retrieves the status of the default model.  Requires the `model-svc:model:view` permission.
     * @summary Get Default Model Status
     */
    getDefaultModelStatus() {
        return __awaiter(this, arguments, void 0, function* (options = { headers: {} }) {
            const localVarPath = this.basePath + '/model-svc/default-model/status';
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
                                body = ObjectSerializer.deserialize(body, "ModelSvcStatusResponse");
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
     * Retrieves the details of a model by its ID.  the Requires `model.view` permission.
     * @summary Get a Model
     * @param modelId Model ID
     */
    getModel(modelId_1) {
        return __awaiter(this, arguments, void 0, function* (modelId, options = { headers: {} }) {
            const localVarPath = this.basePath + '/model-svc/model/{modelId}'
                .replace('{' + 'modelId' + '}', encodeURIComponent(String(modelId)));
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
            // verify required parameter 'modelId' is not null or undefined
            if (modelId === null || modelId === undefined) {
                throw new Error('Required parameter modelId was null or undefined when calling getModel.');
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
                                body = ObjectSerializer.deserialize(body, "ModelSvcGetModelResponse");
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
     * Retrieves the status of a model by ID.  Requires the `model-svc:model:view` permission.
     * @summary Get Model Status
     * @param modelId Model ID
     */
    getModelStatus(modelId_1) {
        return __awaiter(this, arguments, void 0, function* (modelId, options = { headers: {} }) {
            const localVarPath = this.basePath + '/model-svc/model/{modelId}/status'
                .replace('{' + 'modelId' + '}', encodeURIComponent(String(modelId)));
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
            // verify required parameter 'modelId' is not null or undefined
            if (modelId === null || modelId === undefined) {
                throw new Error('Required parameter modelId was null or undefined when calling getModelStatus.');
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
                                body = ObjectSerializer.deserialize(body, "ModelSvcStatusResponse");
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
     * Retrieves a list of models.  Requires `model-svc:model:view` permission.
     * @summary List Models
     */
    listModels() {
        return __awaiter(this, arguments, void 0, function* (options = { headers: {} }) {
            const localVarPath = this.basePath + '/model-svc/models';
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
                                body = ObjectSerializer.deserialize(body, "ModelSvcListModelsResponse");
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
     * Retrieves a list of AI platforms. Eg. LlamaCpp, StableDiffusion etc.  Requires `model-svc:platform:view` permission.
     * @summary List Platforms
     */
    listPlatforms() {
        return __awaiter(this, arguments, void 0, function* (options = { headers: {} }) {
            const localVarPath = this.basePath + '/model-svc/platforms';
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
                                body = ObjectSerializer.deserialize(body, "ModelSvcListPlatformsResponse");
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
     * Sets a model as the default model — when prompts are sent without a Model ID, the default model is used.
     * @summary Make a Model Default
     * @param modelId Model ID
     */
    makeDefault(modelId_1) {
        return __awaiter(this, arguments, void 0, function* (modelId, options = { headers: {} }) {
            const localVarPath = this.basePath + '/model-svc/model/{modelId}/make-default'
                .replace('{' + 'modelId' + '}', encodeURIComponent(String(modelId)));
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
            // verify required parameter 'modelId' is not null or undefined
            if (modelId === null || modelId === undefined) {
                throw new Error('Required parameter modelId was null or undefined when calling makeDefault.');
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
                                body = ObjectSerializer.deserialize(body, "object");
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
     * Starts The Default Model.  Requires the `model-svc:model:create` permission.
     * @summary Start the Default Model
     */
    startDefaultModel() {
        return __awaiter(this, arguments, void 0, function* (options = { headers: {} }) {
            const localVarPath = this.basePath + '/model-svc/default-model/start';
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
                method: 'PUT',
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
                                body = ObjectSerializer.deserialize(body, "object");
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
     * Starts a model by ID
     * @summary Start a Model
     * @param modelId Model ID
     */
    startModel(modelId_1) {
        return __awaiter(this, arguments, void 0, function* (modelId, options = { headers: {} }) {
            const localVarPath = this.basePath + '/model-svc/model/{modelId}/start'
                .replace('{' + 'modelId' + '}', encodeURIComponent(String(modelId)));
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
            // verify required parameter 'modelId' is not null or undefined
            if (modelId === null || modelId === undefined) {
                throw new Error('Required parameter modelId was null or undefined when calling startModel.');
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
                                body = ObjectSerializer.deserialize(body, "object");
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
