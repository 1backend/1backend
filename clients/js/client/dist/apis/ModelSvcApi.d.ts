/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import * as runtime from '../runtime';
import type { ModelSvcGetModelResponse, ModelSvcListModelsResponse, ModelSvcListPlatformsResponse, ModelSvcStatusResponse } from '../models/index';
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
export declare class ModelSvcApi extends runtime.BaseAPI {
    /**
     * Retrieves the status of the default model.  Requires the `model-svc:model:view` permission.
     * Get Default Model Status
     */
    getDefaultModelStatusRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelSvcStatusResponse>>;
    /**
     * Retrieves the status of the default model.  Requires the `model-svc:model:view` permission.
     * Get Default Model Status
     */
    getDefaultModelStatus(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelSvcStatusResponse>;
    /**
     * Retrieves the details of a model by its ID.  the Requires `model.view` permission.
     * Get a Model
     */
    getModelRaw(requestParameters: GetModelRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelSvcGetModelResponse>>;
    /**
     * Retrieves the details of a model by its ID.  the Requires `model.view` permission.
     * Get a Model
     */
    getModel(requestParameters: GetModelRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelSvcGetModelResponse>;
    /**
     * Retrieves the status of a model by ID.  Requires the `model-svc:model:view` permission.
     * Get Model Status
     */
    getModelStatusRaw(requestParameters: GetModelStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelSvcStatusResponse>>;
    /**
     * Retrieves the status of a model by ID.  Requires the `model-svc:model:view` permission.
     * Get Model Status
     */
    getModelStatus(requestParameters: GetModelStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelSvcStatusResponse>;
    /**
     * Retrieves a list of models.  Requires `model-svc:model:view` permission.
     * List Models
     */
    listModelsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelSvcListModelsResponse>>;
    /**
     * Retrieves a list of models.  Requires `model-svc:model:view` permission.
     * List Models
     */
    listModels(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelSvcListModelsResponse>;
    /**
     * Retrieves a list of AI platforms. Eg. LlamaCpp, StableDiffusion etc.  Requires `model-svc:platform:view` permission.
     * List Platforms
     */
    listPlatformsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelSvcListPlatformsResponse>>;
    /**
     * Retrieves a list of AI platforms. Eg. LlamaCpp, StableDiffusion etc.  Requires `model-svc:platform:view` permission.
     * List Platforms
     */
    listPlatforms(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelSvcListPlatformsResponse>;
    /**
     * Sets a model as the default model — when prompts are sent without a Model ID, the default model is used.
     * Make a Model Default
     */
    makeDefaultRaw(requestParameters: MakeDefaultRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Sets a model as the default model — when prompts are sent without a Model ID, the default model is used.
     * Make a Model Default
     */
    makeDefault(requestParameters: MakeDefaultRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Starts The Default Model.  Requires the `model-svc:model:create` permission.
     * Start the Default Model
     */
    startDefaultModelRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Starts The Default Model.  Requires the `model-svc:model:create` permission.
     * Start the Default Model
     */
    startDefaultModel(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Starts a model by ID
     * Start a Model
     */
    startModelRaw(requestParameters: StartModelRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Starts a model by ID
     * Start a Model
     */
    startModel(requestParameters: StartModelRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
}
