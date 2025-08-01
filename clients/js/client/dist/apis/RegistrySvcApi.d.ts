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
import type { RegistrySvcListDefinitionsResponse, RegistrySvcListInstancesResponse, RegistrySvcListNodesRequest, RegistrySvcListNodesResponse, RegistrySvcNodeSelfResponse, RegistrySvcRegisterInstanceRequest, RegistrySvcSaveDefinitionRequest } from '../models/index';
export interface DeleteDefinitionRequest {
    id: string;
}
export interface DeleteNodeRequest {
    url: string;
}
export interface ListInstancesRequest {
    scheme?: string;
    ip?: string;
    deploymentId?: string;
    host?: string;
    ip2?: string;
    id?: string;
    slug?: string;
}
export interface ListNodesRequest {
    body?: RegistrySvcListNodesRequest;
}
export interface RegisterInstanceRequest {
    body: RegistrySvcRegisterInstanceRequest;
}
export interface RemoveInstanceRequest {
    id: string;
}
export interface SaveDefinitionRequest {
    body: RegistrySvcSaveDefinitionRequest;
}
export interface SelfNodeRequest {
    body?: object;
}
/**
 *
 */
export declare class RegistrySvcApi extends runtime.BaseAPI {
    /**
     * Deletes a registered definition by ID.
     * Delete Definition
     */
    deleteDefinitionRaw(requestParameters: DeleteDefinitionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>>;
    /**
     * Deletes a registered definition by ID.
     * Delete Definition
     */
    deleteDefinition(requestParameters: DeleteDefinitionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void>;
    /**
     * Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it\'s still present in the database.
     * Delete Node
     */
    deleteNodeRaw(requestParameters: DeleteNodeRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>>;
    /**
     * Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it\'s still present in the database.
     * Delete Node
     */
    deleteNode(requestParameters: DeleteNodeRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void>;
    /**
     * This endpoint is used to test the server\'s response to a GET request. It echoes back the query parameters as a JSON object.
     * Echo the query parameters in the response body.
     */
    echoGetRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{
        [key: string]: any;
    }>>;
    /**
     * This endpoint is used to test the server\'s response to a GET request. It echoes back the query parameters as a JSON object.
     * Echo the query parameters in the response body.
     */
    echoGet(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{
        [key: string]: any;
    }>;
    /**
     * This endpoint is used to test the server\'s response to a request. It simply echoes back the request body as a JSON response.
     * Echo the request body in the response body.
     */
    echoPostRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{
        [key: string]: any;
    }>>;
    /**
     * This endpoint is used to test the server\'s response to a request. It simply echoes back the request body as a JSON response.
     * Echo the request body in the response body.
     */
    echoPost(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{
        [key: string]: any;
    }>;
    /**
     * This endpoint is used to test the server\'s response to a request. It simply echoes back the request body as a JSON response.
     * Echo the request body in the response body.
     */
    echoPutRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{
        [key: string]: any;
    }>>;
    /**
     * This endpoint is used to test the server\'s response to a request. It simply echoes back the request body as a JSON response.
     * Echo the request body in the response body.
     */
    echoPut(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{
        [key: string]: any;
    }>;
    /**
     * Retrieves a list of all definitions or filters them by specific criteria.
     * List Definitions
     */
    listDefinitionsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<RegistrySvcListDefinitionsResponse>>;
    /**
     * Retrieves a list of all definitions or filters them by specific criteria.
     * List Definitions
     */
    listDefinitions(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<RegistrySvcListDefinitionsResponse>;
    /**
     * Retrieves a list of all instances or filters them by specific criteria (e.g., host, IP).
     * List Service Instances
     */
    listInstancesRaw(requestParameters: ListInstancesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<RegistrySvcListInstancesResponse>>;
    /**
     * Retrieves a list of all instances or filters them by specific criteria (e.g., host, IP).
     * List Service Instances
     */
    listInstances(requestParameters?: ListInstancesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<RegistrySvcListInstancesResponse>;
    /**
     * Retrieve a list of nodes.
     * List Nodes
     */
    listNodesRaw(requestParameters: ListNodesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<RegistrySvcListNodesResponse>>;
    /**
     * Retrieve a list of nodes.
     * List Nodes
     */
    listNodes(requestParameters?: ListNodesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<RegistrySvcListNodesResponse>;
    /**
     * Registers an instance. Idempotent.
     * Register Instance
     */
    registerInstanceRaw(requestParameters: RegisterInstanceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Registers an instance. Idempotent.
     * Register Instance
     */
    registerInstance(requestParameters: RegisterInstanceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Removes a registered instance by ID.
     * Remove Instance
     */
    removeInstanceRaw(requestParameters: RemoveInstanceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>>;
    /**
     * Removes a registered instance by ID.
     * Remove Instance
     */
    removeInstance(requestParameters: RemoveInstanceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void>;
    /**
     * Registers a new definition, associating an definition address with a slug acquired from the bearer token.
     * Register a Definition
     */
    saveDefinitionRaw(requestParameters: SaveDefinitionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Registers a new definition, associating an definition address with a slug acquired from the bearer token.
     * Register a Definition
     */
    saveDefinition(requestParameters: SaveDefinitionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Show the local node.
     * View Self Node
     */
    selfNodeRaw(requestParameters: SelfNodeRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<RegistrySvcNodeSelfResponse>>;
    /**
     * Show the local node.
     * View Self Node
     */
    selfNode(requestParameters?: SelfNodeRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<RegistrySvcNodeSelfResponse>;
}
