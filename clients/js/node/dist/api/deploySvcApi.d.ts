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
import http from 'http';
import { DeploySvcDeleteDeploymentRequest } from '../model/deploySvcDeleteDeploymentRequest';
import { DeploySvcListDeploymentsResponse } from '../model/deploySvcListDeploymentsResponse';
import { DeploySvcSaveDeploymentRequest } from '../model/deploySvcSaveDeploymentRequest';
import { Authentication, Interceptor } from '../model/models';
import { ApiKeyAuth } from '../model/models';
export declare enum DeploySvcApiApiKeys {
    BearerAuth = 0
}
export declare class DeploySvcApi {
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
    setApiKey(key: DeploySvcApiApiKeys, value: string): void;
    addInterceptor(interceptor: Interceptor): void;
    /**
     * Delete a deployment.
     * @summary Delete Deployment
     * @param body Delete Deploys Request
     */
    deleteDeployment(body?: DeploySvcDeleteDeploymentRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
    /**
     * Retrieve a list of deployments.
     * @summary List Deployments
     * @param body List Deploys Request
     */
    listDeployments(body?: object, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: DeploySvcListDeploymentsResponse;
    }>;
    /**
     * Save a deployment.
     * @summary Save Deployment
     * @param body Save Deploys Request
     */
    saveDeployment(body?: DeploySvcSaveDeploymentRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
}
