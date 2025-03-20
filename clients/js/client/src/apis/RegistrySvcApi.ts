/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  RegistrySvcErrorResponse,
  RegistrySvcListDefinitionsResponse,
  RegistrySvcListInstancesResponse,
  RegistrySvcListNodesRequest,
  RegistrySvcListNodesResponse,
  RegistrySvcNodeSelfResponse,
  RegistrySvcRegisterInstanceRequest,
  RegistrySvcSaveDefinitionRequest,
} from '../models/index';
import {
    RegistrySvcErrorResponseFromJSON,
    RegistrySvcErrorResponseToJSON,
    RegistrySvcListDefinitionsResponseFromJSON,
    RegistrySvcListDefinitionsResponseToJSON,
    RegistrySvcListInstancesResponseFromJSON,
    RegistrySvcListInstancesResponseToJSON,
    RegistrySvcListNodesRequestFromJSON,
    RegistrySvcListNodesRequestToJSON,
    RegistrySvcListNodesResponseFromJSON,
    RegistrySvcListNodesResponseToJSON,
    RegistrySvcNodeSelfResponseFromJSON,
    RegistrySvcNodeSelfResponseToJSON,
    RegistrySvcRegisterInstanceRequestFromJSON,
    RegistrySvcRegisterInstanceRequestToJSON,
    RegistrySvcSaveDefinitionRequestFromJSON,
    RegistrySvcSaveDefinitionRequestToJSON,
} from '../models/index';

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
export class RegistrySvcApi extends runtime.BaseAPI {

    /**
     * Deletes a registered definition by ID.
     * Delete Definition
     */
    async deleteDefinitionRaw(requestParameters: DeleteDefinitionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling deleteDefinition().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/registry-svc/definition/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Deletes a registered definition by ID.
     * Delete Definition
     */
    async deleteDefinition(requestParameters: DeleteDefinitionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteDefinitionRaw(requestParameters, initOverrides);
    }

    /**
     * Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it\'s still present in the database.
     * Delete Node
     */
    async deleteNodeRaw(requestParameters: DeleteNodeRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['url'] == null) {
            throw new runtime.RequiredError(
                'url',
                'Required parameter "url" was null or undefined when calling deleteNode().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/registry-svc/node/{url}`.replace(`{${"url"}}`, encodeURIComponent(String(requestParameters['url']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it\'s still present in the database.
     * Delete Node
     */
    async deleteNode(requestParameters: DeleteNodeRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteNodeRaw(requestParameters, initOverrides);
    }

    /**
     * Retrieves a list of all definitions or filters them by specific criteria.
     * List Definitions
     */
    async listDefinitionsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<RegistrySvcListDefinitionsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/registry-svc/definitions`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => RegistrySvcListDefinitionsResponseFromJSON(jsonValue));
    }

    /**
     * Retrieves a list of all definitions or filters them by specific criteria.
     * List Definitions
     */
    async listDefinitions(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<RegistrySvcListDefinitionsResponse> {
        const response = await this.listDefinitionsRaw(initOverrides);
        return await response.value();
    }

    /**
     * Retrieves a list of all instances or filters them by specific criteria (e.g., host, IP).
     * List Service Instances
     */
    async listInstancesRaw(requestParameters: ListInstancesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<RegistrySvcListInstancesResponse>> {
        const queryParameters: any = {};

        if (requestParameters['scheme'] != null) {
            queryParameters['scheme'] = requestParameters['scheme'];
        }

        if (requestParameters['ip'] != null) {
            queryParameters['ip'] = requestParameters['ip'];
        }

        if (requestParameters['deploymentId'] != null) {
            queryParameters['deploymentId'] = requestParameters['deploymentId'];
        }

        if (requestParameters['host'] != null) {
            queryParameters['host'] = requestParameters['host'];
        }

        if (requestParameters['ip2'] != null) {
            queryParameters['ip'] = requestParameters['ip2'];
        }

        if (requestParameters['id'] != null) {
            queryParameters['id'] = requestParameters['id'];
        }

        if (requestParameters['slug'] != null) {
            queryParameters['slug'] = requestParameters['slug'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/registry-svc/instances`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => RegistrySvcListInstancesResponseFromJSON(jsonValue));
    }

    /**
     * Retrieves a list of all instances or filters them by specific criteria (e.g., host, IP).
     * List Service Instances
     */
    async listInstances(requestParameters: ListInstancesRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<RegistrySvcListInstancesResponse> {
        const response = await this.listInstancesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieve a list of nodes.
     * List Nodes
     */
    async listNodesRaw(requestParameters: ListNodesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<RegistrySvcListNodesResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/registry-svc/nodes`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: RegistrySvcListNodesRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => RegistrySvcListNodesResponseFromJSON(jsonValue));
    }

    /**
     * Retrieve a list of nodes.
     * List Nodes
     */
    async listNodes(requestParameters: ListNodesRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<RegistrySvcListNodesResponse> {
        const response = await this.listNodesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Registers an instance. Idempotent.
     * Register Instance
     */
    async registerInstanceRaw(requestParameters: RegisterInstanceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling registerInstance().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/registry-svc/instance`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: RegistrySvcRegisterInstanceRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Registers an instance. Idempotent.
     * Register Instance
     */
    async registerInstance(requestParameters: RegisterInstanceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.registerInstanceRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Removes a registered instance by ID.
     * Remove Instance
     */
    async removeInstanceRaw(requestParameters: RemoveInstanceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling removeInstance().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/registry-svc/instance/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Removes a registered instance by ID.
     * Remove Instance
     */
    async removeInstance(requestParameters: RemoveInstanceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.removeInstanceRaw(requestParameters, initOverrides);
    }

    /**
     * Registers a new definition, associating an definition address with a slug acquired from the bearer token.
     * Register a Definition
     */
    async saveDefinitionRaw(requestParameters: SaveDefinitionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling saveDefinition().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/registry-svc/definition`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: RegistrySvcSaveDefinitionRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Registers a new definition, associating an definition address with a slug acquired from the bearer token.
     * Register a Definition
     */
    async saveDefinition(requestParameters: SaveDefinitionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.saveDefinitionRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Show the local node.
     * View Self Node
     */
    async selfNodeRaw(requestParameters: SelfNodeRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<RegistrySvcNodeSelfResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/registry-svc/node/self`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters['body'] as any,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => RegistrySvcNodeSelfResponseFromJSON(jsonValue));
    }

    /**
     * Show the local node.
     * View Self Node
     */
    async selfNode(requestParameters: SelfNodeRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<RegistrySvcNodeSelfResponse> {
        const response = await this.selfNodeRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
