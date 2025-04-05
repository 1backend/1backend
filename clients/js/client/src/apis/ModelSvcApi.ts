/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  ModelSvcErrorResponse,
  ModelSvcGetModelResponse,
  ModelSvcListModelsResponse,
  ModelSvcListPlatformsResponse,
  ModelSvcStatusResponse,
} from '../models/index';
import {
    ModelSvcErrorResponseFromJSON,
    ModelSvcErrorResponseToJSON,
    ModelSvcGetModelResponseFromJSON,
    ModelSvcGetModelResponseToJSON,
    ModelSvcListModelsResponseFromJSON,
    ModelSvcListModelsResponseToJSON,
    ModelSvcListPlatformsResponseFromJSON,
    ModelSvcListPlatformsResponseToJSON,
    ModelSvcStatusResponseFromJSON,
    ModelSvcStatusResponseToJSON,
} from '../models/index';

export interface GetModelRequest {
    modelId: string;
}

export interface GetModelStatusRequest {
    modelId: string;
}

export interface MakeDefaultRequest {
    modelId: string;
}

export interface StartModelRequest {
    modelId: string;
}

/**
 * 
 */
export class ModelSvcApi extends runtime.BaseAPI {

    /**
     * Retrieves the status of the default model.  Requires the `model-svc:model:view` permission.
     * Get Default Model Status
     */
    async getDefaultModelStatusRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelSvcStatusResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/model-svc/default-model/status`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ModelSvcStatusResponseFromJSON(jsonValue));
    }

    /**
     * Retrieves the status of the default model.  Requires the `model-svc:model:view` permission.
     * Get Default Model Status
     */
    async getDefaultModelStatus(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelSvcStatusResponse> {
        const response = await this.getDefaultModelStatusRaw(initOverrides);
        return await response.value();
    }

    /**
     * Retrieves the details of a model by its ID.  the Requires `model.view` permission.
     * Get a Model
     */
    async getModelRaw(requestParameters: GetModelRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelSvcGetModelResponse>> {
        if (requestParameters['modelId'] == null) {
            throw new runtime.RequiredError(
                'modelId',
                'Required parameter "modelId" was null or undefined when calling getModel().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/model-svc/model/{modelId}`.replace(`{${"modelId"}}`, encodeURIComponent(String(requestParameters['modelId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ModelSvcGetModelResponseFromJSON(jsonValue));
    }

    /**
     * Retrieves the details of a model by its ID.  the Requires `model.view` permission.
     * Get a Model
     */
    async getModel(requestParameters: GetModelRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelSvcGetModelResponse> {
        const response = await this.getModelRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieves the status of a model by ID.  Requires the `model-svc:model:view` permission.
     * Get Model Status
     */
    async getModelStatusRaw(requestParameters: GetModelStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelSvcStatusResponse>> {
        if (requestParameters['modelId'] == null) {
            throw new runtime.RequiredError(
                'modelId',
                'Required parameter "modelId" was null or undefined when calling getModelStatus().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/model-svc/model/{modelId}/status`.replace(`{${"modelId"}}`, encodeURIComponent(String(requestParameters['modelId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ModelSvcStatusResponseFromJSON(jsonValue));
    }

    /**
     * Retrieves the status of a model by ID.  Requires the `model-svc:model:view` permission.
     * Get Model Status
     */
    async getModelStatus(requestParameters: GetModelStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelSvcStatusResponse> {
        const response = await this.getModelStatusRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieves a list of models.  Requires `model-svc:model:view` permission.
     * List Models
     */
    async listModelsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelSvcListModelsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/model-svc/models`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ModelSvcListModelsResponseFromJSON(jsonValue));
    }

    /**
     * Retrieves a list of models.  Requires `model-svc:model:view` permission.
     * List Models
     */
    async listModels(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelSvcListModelsResponse> {
        const response = await this.listModelsRaw(initOverrides);
        return await response.value();
    }

    /**
     * Retrieves a list of AI platforms. Eg. LlamaCpp, StableDiffusion etc.  Requires `model-svc:platform:view` permission.
     * List Platforms
     */
    async listPlatformsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelSvcListPlatformsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/model-svc/platforms`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ModelSvcListPlatformsResponseFromJSON(jsonValue));
    }

    /**
     * Retrieves a list of AI platforms. Eg. LlamaCpp, StableDiffusion etc.  Requires `model-svc:platform:view` permission.
     * List Platforms
     */
    async listPlatforms(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelSvcListPlatformsResponse> {
        const response = await this.listPlatformsRaw(initOverrides);
        return await response.value();
    }

    /**
     * Sets a model as the default model — when prompts are sent without a Model ID, the default model is used.
     * Make a Model Default
     */
    async makeDefaultRaw(requestParameters: MakeDefaultRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['modelId'] == null) {
            throw new runtime.RequiredError(
                'modelId',
                'Required parameter "modelId" was null or undefined when calling makeDefault().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/model-svc/model/{modelId}/make-default`.replace(`{${"modelId"}}`, encodeURIComponent(String(requestParameters['modelId']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Sets a model as the default model — when prompts are sent without a Model ID, the default model is used.
     * Make a Model Default
     */
    async makeDefault(requestParameters: MakeDefaultRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.makeDefaultRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Starts The Default Model.  Requires the `model-svc:model:create` permission.
     * Start the Default Model
     */
    async startDefaultModelRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/model-svc/default-model/start`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Starts The Default Model.  Requires the `model-svc:model:create` permission.
     * Start the Default Model
     */
    async startDefaultModel(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.startDefaultModelRaw(initOverrides);
        return await response.value();
    }

    /**
     * Starts a model by ID
     * Start a Model
     */
    async startModelRaw(requestParameters: StartModelRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['modelId'] == null) {
            throw new runtime.RequiredError(
                'modelId',
                'Required parameter "modelId" was null or undefined when calling startModel().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/model-svc/model/{modelId}/start`.replace(`{${"modelId"}}`, encodeURIComponent(String(requestParameters['modelId']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Starts a model by ID
     * Start a Model
     */
    async startModel(requestParameters: StartModelRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.startModelRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
