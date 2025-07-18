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
import * as runtime from '../runtime';
import type { DeploySvcDeleteDeploymentRequest, DeploySvcListDeploymentsResponse, DeploySvcSaveDeploymentRequest } from '../models/index';
export interface DeleteDeploymentRequest {
    body?: DeploySvcDeleteDeploymentRequest;
}
export interface ListDeploymentsRequest {
    body?: object;
}
export interface SaveDeploymentRequest {
    body?: DeploySvcSaveDeploymentRequest;
}
/**
 *
 */
export declare class DeploySvcApi extends runtime.BaseAPI {
    /**
     * Delete a deployment.
     * Delete Deployment
     */
    deleteDeploymentRaw(requestParameters: DeleteDeploymentRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Delete a deployment.
     * Delete Deployment
     */
    deleteDeployment(requestParameters?: DeleteDeploymentRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Retrieve a list of deployments.
     * List Deployments
     */
    listDeploymentsRaw(requestParameters: ListDeploymentsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DeploySvcListDeploymentsResponse>>;
    /**
     * Retrieve a list of deployments.
     * List Deployments
     */
    listDeployments(requestParameters?: ListDeploymentsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DeploySvcListDeploymentsResponse>;
    /**
     * Save a deployment.
     * Save Deployment
     */
    saveDeploymentRaw(requestParameters: SaveDeploymentRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Save a deployment.
     * Save Deployment
     */
    saveDeployment(requestParameters?: SaveDeploymentRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
}
