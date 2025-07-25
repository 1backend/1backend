/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
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
import { ModelSvcGetModelResponseFromJSON, ModelSvcListModelsResponseFromJSON, ModelSvcListPlatformsResponseFromJSON, ModelSvcStatusResponseFromJSON, } from '../models/index';
/**
 *
 */
export class ModelSvcApi extends runtime.BaseAPI {
    /**
     * Retrieves the status of the default model.  Requires the `model-svc:model:view` permission.
     * Get Default Model Status
     */
    getDefaultModelStatusRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/model-svc/default-model/status`;
            const response = yield this.request({
                path: urlPath,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ModelSvcStatusResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieves the status of the default model.  Requires the `model-svc:model:view` permission.
     * Get Default Model Status
     */
    getDefaultModelStatus(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.getDefaultModelStatusRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * Retrieves the details of a model by its ID.  the Requires `model.view` permission.
     * Get a Model
     */
    getModelRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['modelId'] == null) {
                throw new runtime.RequiredError('modelId', 'Required parameter "modelId" was null or undefined when calling getModel().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/model-svc/model/{modelId}`;
            urlPath = urlPath.replace(`{${"modelId"}}`, encodeURIComponent(String(requestParameters['modelId'])));
            const response = yield this.request({
                path: urlPath,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ModelSvcGetModelResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieves the details of a model by its ID.  the Requires `model.view` permission.
     * Get a Model
     */
    getModel(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.getModelRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Retrieves the status of a model by ID.  Requires the `model-svc:model:view` permission.
     * Get Model Status
     */
    getModelStatusRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['modelId'] == null) {
                throw new runtime.RequiredError('modelId', 'Required parameter "modelId" was null or undefined when calling getModelStatus().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/model-svc/model/{modelId}/status`;
            urlPath = urlPath.replace(`{${"modelId"}}`, encodeURIComponent(String(requestParameters['modelId'])));
            const response = yield this.request({
                path: urlPath,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ModelSvcStatusResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieves the status of a model by ID.  Requires the `model-svc:model:view` permission.
     * Get Model Status
     */
    getModelStatus(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.getModelStatusRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Retrieves a list of models.  Requires `model-svc:model:view` permission.
     * List Models
     */
    listModelsRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/model-svc/models`;
            const response = yield this.request({
                path: urlPath,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ModelSvcListModelsResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieves a list of models.  Requires `model-svc:model:view` permission.
     * List Models
     */
    listModels(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.listModelsRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * Retrieves a list of AI platforms. Eg. LlamaCpp, StableDiffusion etc.  Requires `model-svc:platform:view` permission.
     * List Platforms
     */
    listPlatformsRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/model-svc/platforms`;
            const response = yield this.request({
                path: urlPath,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ModelSvcListPlatformsResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieves a list of AI platforms. Eg. LlamaCpp, StableDiffusion etc.  Requires `model-svc:platform:view` permission.
     * List Platforms
     */
    listPlatforms(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.listPlatformsRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * Sets a model as the default model — when prompts are sent without a Model ID, the default model is used.
     * Make a Model Default
     */
    makeDefaultRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['modelId'] == null) {
                throw new runtime.RequiredError('modelId', 'Required parameter "modelId" was null or undefined when calling makeDefault().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/model-svc/model/{modelId}/make-default`;
            urlPath = urlPath.replace(`{${"modelId"}}`, encodeURIComponent(String(requestParameters['modelId'])));
            const response = yield this.request({
                path: urlPath,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Sets a model as the default model — when prompts are sent without a Model ID, the default model is used.
     * Make a Model Default
     */
    makeDefault(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.makeDefaultRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Starts The Default Model.  Requires the `model-svc:model:create` permission.
     * Start the Default Model
     */
    startDefaultModelRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/model-svc/default-model/start`;
            const response = yield this.request({
                path: urlPath,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Starts The Default Model.  Requires the `model-svc:model:create` permission.
     * Start the Default Model
     */
    startDefaultModel(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.startDefaultModelRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * Starts a model by ID
     * Start a Model
     */
    startModelRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['modelId'] == null) {
                throw new runtime.RequiredError('modelId', 'Required parameter "modelId" was null or undefined when calling startModel().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/model-svc/model/{modelId}/start`;
            urlPath = urlPath.replace(`{${"modelId"}}`, encodeURIComponent(String(requestParameters['modelId'])));
            const response = yield this.request({
                path: urlPath,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Starts a model by ID
     * Start a Model
     */
    startModel(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.startModelRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
}
