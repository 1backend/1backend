/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
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
export var FileSvcApiApiKeys;
(function (FileSvcApiApiKeys) {
    FileSvcApiApiKeys[FileSvcApiApiKeys["BearerAuth"] = 0] = "BearerAuth";
})(FileSvcApiApiKeys || (FileSvcApiApiKeys = {}));
export class FileSvcApi {
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
        this.authentications[FileSvcApiApiKeys[key]].apiKey = value;
    }
    addInterceptor(interceptor) {
        this.interceptors.push(interceptor);
    }
    /**
     * Start or resume the download for a specified URL.  Requires the `file-svc:download:create` permission.
     * @summary Download a File
     * @param body Download Request
     */
    downloadFile(body_1) {
        return __awaiter(this, arguments, void 0, function* (body, options = { headers: {} }) {
            const localVarPath = this.basePath + '/file-svc/download';
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
                throw new Error('Required parameter body was null or undefined when calling downloadFile.');
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
                body: ObjectSerializer.serialize(body, "FileSvcDownloadFileRequest")
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
     * Get a download by URL.  Requires the `file-svc:download:view` permission.
     * @summary Get a Download
     * @param url url
     */
    getDownload(url_1) {
        return __awaiter(this, arguments, void 0, function* (url, options = { headers: {} }) {
            const localVarPath = this.basePath + '/file-svc/download/{url}'
                .replace('{' + 'url' + '}', encodeURIComponent(String(url)));
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
            // verify required parameter 'url' is not null or undefined
            if (url === null || url === undefined) {
                throw new Error('Required parameter url was null or undefined when calling getDownload.');
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
                                body = ObjectSerializer.deserialize(body, "FileSvcGetDownloadResponse");
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
     * List download details.  Requires the `file-svc:download:view` permission.
     * @summary List Downloads
     */
    listDownloads() {
        return __awaiter(this, arguments, void 0, function* (options = { headers: {} }) {
            const localVarPath = this.basePath + '/file-svc/downloads';
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
                                body = ObjectSerializer.deserialize(body, "FileSvcDownloadsResponse");
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
     * List the uploaded files.  Requires the `file-svc:upload:view` permission.
     * @summary List Uploads
     * @param body List Uploads Request
     */
    listUploads(body_1) {
        return __awaiter(this, arguments, void 0, function* (body, options = { headers: {} }) {
            const localVarPath = this.basePath + '/file-svc/uploads';
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
                body: ObjectSerializer.serialize(body, "FileSvcListUploadsRequest")
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
                                body = ObjectSerializer.deserialize(body, "FileSvcListUploadsResponse");
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
     * Pause a download that is currently in progress.  Requires the `file-svc:download:edit` permission.
     * @summary Pause a Download
     * @param url Download URL
     */
    pauseDownload(url_1) {
        return __awaiter(this, arguments, void 0, function* (url, options = { headers: {} }) {
            const localVarPath = this.basePath + '/file-svc/download/{url}/pause'
                .replace('{' + 'url' + '}', encodeURIComponent(String(url)));
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
            // verify required parameter 'url' is not null or undefined
            if (url === null || url === undefined) {
                throw new Error('Required parameter url was null or undefined when calling pauseDownload.');
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
     * Serves a previously downloaded file based on its URL.
     * @summary Serve a Downloaded file
     * @param url URL of the file. Even after downloading, the file is still referenced by its original internet URL.
     */
    serveDownload(url_1) {
        return __awaiter(this, arguments, void 0, function* (url, options = { headers: {} }) {
            const localVarPath = this.basePath + '/file-svc/serve/download/{url}'
                .replace('{' + 'url' + '}', encodeURIComponent(String(url)));
            let localVarQueryParameters = {};
            let localVarHeaderParams = Object.assign({}, this._defaultHeaders);
            const produces = ['application/octet-stream'];
            // give precedence to 'application/json'
            if (produces.indexOf('application/json') >= 0) {
                localVarHeaderParams.Accept = 'application/json';
            }
            else {
                localVarHeaderParams.Accept = produces.join(',');
            }
            let localVarFormParams = {};
            // verify required parameter 'url' is not null or undefined
            if (url === null || url === undefined) {
                throw new Error('Required parameter url was null or undefined when calling serveDownload.');
            }
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            let localVarRequestOptions = {
                method: 'GET',
                qs: localVarQueryParameters,
                headers: localVarHeaderParams,
                uri: localVarPath,
                useQuerystring: this._useQuerystring,
                encoding: null,
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
                                body = ObjectSerializer.deserialize(body, "Buffer");
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
     * Retrieves and serves a previously uploaded file using its File ID. Note: The `ID` and `FileID` fields of an upload are different. - `FileID` is a unique identifier for the file itself. - `ID` is a unique identifier for a specific replica of the file. Since 1Backend is a distributed system, files can be replicated across multiple nodes. This means each uploaded file may have multiple records with the same `FileID` but different `ID`s.
     * @summary Serve an Uploaded File
     * @param fileId Upload ID
     */
    serveUpload(fileId_1) {
        return __awaiter(this, arguments, void 0, function* (fileId, options = { headers: {} }) {
            const localVarPath = this.basePath + '/file-svc/serve/upload/{fileId}'
                .replace('{' + 'fileId' + '}', encodeURIComponent(String(fileId)));
            let localVarQueryParameters = {};
            let localVarHeaderParams = Object.assign({}, this._defaultHeaders);
            const produces = ['application/octet-stream'];
            // give precedence to 'application/json'
            if (produces.indexOf('application/json') >= 0) {
                localVarHeaderParams.Accept = 'application/json';
            }
            else {
                localVarHeaderParams.Accept = produces.join(',');
            }
            let localVarFormParams = {};
            // verify required parameter 'fileId' is not null or undefined
            if (fileId === null || fileId === undefined) {
                throw new Error('Required parameter fileId was null or undefined when calling serveUpload.');
            }
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            let localVarRequestOptions = {
                method: 'GET',
                qs: localVarQueryParameters,
                headers: localVarHeaderParams,
                uri: localVarPath,
                useQuerystring: this._useQuerystring,
                encoding: null,
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
                                body = ObjectSerializer.deserialize(body, "Buffer");
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
     * Uploads a file to the server. Currently if using the clients only one file can be uploaded at a time due to this bug https://github.com/OpenAPITools/openapi-generator/issues/11341 Once that is fixed we should have an `PUT /file-svc/uploads`/uploadFiles (note the plural) endpoints. In reality the endpoint \"unofficially\" supports multiple files. YMMV.  Requires the `file-svc:upload:create` permission.
     * @summary Upload a File
     * @param file File to upload
     */
    uploadFile(file_1) {
        return __awaiter(this, arguments, void 0, function* (file, options = { headers: {} }) {
            const localVarPath = this.basePath + '/file-svc/upload';
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
            // verify required parameter 'file' is not null or undefined
            if (file === null || file === undefined) {
                throw new Error('Required parameter file was null or undefined when calling uploadFile.');
            }
            Object.assign(localVarHeaderParams, options.headers);
            let localVarUseFormData = false;
            if (file !== undefined) {
                localVarFormParams['file'] = file;
            }
            localVarUseFormData = true;
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
                                body = ObjectSerializer.deserialize(body, "FileSvcUploadFileResponse");
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
