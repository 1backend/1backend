/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import http from 'http';
import { ModelSvcGetModelResponse } from '../model/modelSvcGetModelResponse';
import { ModelSvcListModelsResponse } from '../model/modelSvcListModelsResponse';
import { ModelSvcListPlatformsResponse } from '../model/modelSvcListPlatformsResponse';
import { ModelSvcStatusResponse } from '../model/modelSvcStatusResponse';
import { Authentication, Interceptor } from '../model/models';
import { ApiKeyAuth } from '../model/models';
export declare enum ModelSvcApiApiKeys {
    BearerAuth = 0
}
export declare class ModelSvcApi {
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
    setApiKey(key: ModelSvcApiApiKeys, value: string): void;
    addInterceptor(interceptor: Interceptor): void;
    /**
     * Retrieves the status of the default model.  Requires the `model-svc:model:view` permission.
     * @summary Get Default Model Status
     */
    getDefaultModelStatus(options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ModelSvcStatusResponse;
    }>;
    /**
     * Retrieves the details of a model by its ID.  the Requires `model.view` permission.
     * @summary Get a Model
     * @param modelId Model ID
     */
    getModel(modelId: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ModelSvcGetModelResponse;
    }>;
    /**
     * Retrieves the status of a model by ID.  Requires the `model-svc:model:view` permission.
     * @summary Get Model Status
     * @param modelId Model ID
     */
    getModelStatus(modelId: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ModelSvcStatusResponse;
    }>;
    /**
     * Retrieves a list of models.  Requires `model-svc:model:view` permission.
     * @summary List Models
     */
    listModels(options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ModelSvcListModelsResponse;
    }>;
    /**
     * Retrieves a list of AI platforms. Eg. LlamaCpp, StableDiffusion etc.  Requires `model-svc:platform:view` permission.
     * @summary List Platforms
     */
    listPlatforms(options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ModelSvcListPlatformsResponse;
    }>;
    /**
     * Sets a model as the default model — when prompts are sent without a Model ID, the default model is used.
     * @summary Make a Model Default
     * @param modelId Model ID
     */
    makeDefault(modelId: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
    /**
     * Starts The Default Model.  Requires the `model-svc:model:create` permission.
     * @summary Start the Default Model
     */
    startDefaultModel(options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
    /**
     * Starts a model by ID
     * @summary Start a Model
     * @param modelId Model ID
     */
    startModel(modelId: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
}
