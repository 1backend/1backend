/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
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
  ContainerSvcBuildImageRequest,
  ContainerSvcContainerIsRunningResponse,
  ContainerSvcDaemonInfoResponse,
  ContainerSvcErrorResponse,
  ContainerSvcGetContainerSummaryResponse,
  ContainerSvcGetHostResponse,
  ContainerSvcImagePullableResponse,
  ContainerSvcListContainersRequest,
  ContainerSvcListContainersResponse,
  ContainerSvcListLogsRequest,
  ContainerSvcListLogsResponse,
  ContainerSvcRunContainerRequest,
  ContainerSvcRunContainerResponse,
  ContainerSvcStopContainerRequest,
} from '../models/index';
import {
    ContainerSvcBuildImageRequestFromJSON,
    ContainerSvcBuildImageRequestToJSON,
    ContainerSvcContainerIsRunningResponseFromJSON,
    ContainerSvcContainerIsRunningResponseToJSON,
    ContainerSvcDaemonInfoResponseFromJSON,
    ContainerSvcDaemonInfoResponseToJSON,
    ContainerSvcErrorResponseFromJSON,
    ContainerSvcErrorResponseToJSON,
    ContainerSvcGetContainerSummaryResponseFromJSON,
    ContainerSvcGetContainerSummaryResponseToJSON,
    ContainerSvcGetHostResponseFromJSON,
    ContainerSvcGetHostResponseToJSON,
    ContainerSvcImagePullableResponseFromJSON,
    ContainerSvcImagePullableResponseToJSON,
    ContainerSvcListContainersRequestFromJSON,
    ContainerSvcListContainersRequestToJSON,
    ContainerSvcListContainersResponseFromJSON,
    ContainerSvcListContainersResponseToJSON,
    ContainerSvcListLogsRequestFromJSON,
    ContainerSvcListLogsRequestToJSON,
    ContainerSvcListLogsResponseFromJSON,
    ContainerSvcListLogsResponseToJSON,
    ContainerSvcRunContainerRequestFromJSON,
    ContainerSvcRunContainerRequestToJSON,
    ContainerSvcRunContainerResponseFromJSON,
    ContainerSvcRunContainerResponseToJSON,
    ContainerSvcStopContainerRequestFromJSON,
    ContainerSvcStopContainerRequestToJSON,
} from '../models/index';

export interface BuildImageRequest {
    body: ContainerSvcBuildImageRequest;
}

export interface ContainerIsRunningRequest {
    hash?: string;
    name?: string;
}

export interface ContainerSummaryRequest {
    hash?: string;
    name?: string;
    lines?: number;
}

export interface ImagePullableRequest {
    imageName: string;
}

export interface ListContainerLogsRequest {
    body: ContainerSvcListLogsRequest;
}

export interface ListContainersRequest {
    body: ContainerSvcListContainersRequest;
}

export interface RunContainerRequest {
    body: ContainerSvcRunContainerRequest;
}

export interface StopContainerRequest {
    body: ContainerSvcStopContainerRequest;
}

/**
 * 
 */
export class ContainerSvcApi extends runtime.BaseAPI {

    /**
     * Builds a Docker image with the specified parameters.  Requires the `container-svc:image:build` permission.
     * Build an Image
     */
    async buildImageRaw(requestParameters: BuildImageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling buildImage().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/container-svc/image`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: ContainerSvcBuildImageRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Builds a Docker image with the specified parameters.  Requires the `container-svc:image:build` permission.
     * Build an Image
     */
    async buildImage(requestParameters: BuildImageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.buildImageRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieve detailed information about the availability and status of container daemons on the node.
     * Get Container Daemon Information
     */
    async containerDaemonInfoRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcDaemonInfoResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/container-svc/daemon/info`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcDaemonInfoResponseFromJSON(jsonValue));
    }

    /**
     * Retrieve detailed information about the availability and status of container daemons on the node.
     * Get Container Daemon Information
     */
    async containerDaemonInfo(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcDaemonInfoResponse> {
        const response = await this.containerDaemonInfoRaw(initOverrides);
        return await response.value();
    }

    /**
     * Check if a Docker container is running, identified by hash or name.
     * Check If a Container Is Running
     */
    async containerIsRunningRaw(requestParameters: ContainerIsRunningRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcContainerIsRunningResponse>> {
        const queryParameters: any = {};

        if (requestParameters['hash'] != null) {
            queryParameters['hash'] = requestParameters['hash'];
        }

        if (requestParameters['name'] != null) {
            queryParameters['name'] = requestParameters['name'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/container-svc/container/is-running`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcContainerIsRunningResponseFromJSON(jsonValue));
    }

    /**
     * Check if a Docker container is running, identified by hash or name.
     * Check If a Container Is Running
     */
    async containerIsRunning(requestParameters: ContainerIsRunningRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcContainerIsRunningResponse> {
        const response = await this.containerIsRunningRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get a summary of the Docker container identified by hash or name, limited to a specified number of lines.
     * Get Container Summary
     */
    async containerSummaryRaw(requestParameters: ContainerSummaryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcGetContainerSummaryResponse>> {
        const queryParameters: any = {};

        if (requestParameters['hash'] != null) {
            queryParameters['hash'] = requestParameters['hash'];
        }

        if (requestParameters['name'] != null) {
            queryParameters['name'] = requestParameters['name'];
        }

        if (requestParameters['lines'] != null) {
            queryParameters['lines'] = requestParameters['lines'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/container-svc/container/summary`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcGetContainerSummaryResponseFromJSON(jsonValue));
    }

    /**
     * Get a summary of the Docker container identified by hash or name, limited to a specified number of lines.
     * Get Container Summary
     */
    async containerSummary(requestParameters: ContainerSummaryRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcGetContainerSummaryResponse> {
        const response = await this.containerSummaryRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieve information about the Container host
     * Get Container Host
     */
    async getHostRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcGetHostResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/container-svc/host`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcGetHostResponseFromJSON(jsonValue));
    }

    /**
     * Retrieve information about the Container host
     * Get Container Host
     */
    async getHost(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcGetHostResponse> {
        const response = await this.getHostRaw(initOverrides);
        return await response.value();
    }

    /**
     * Check if an image exists on in the container registry and is pullable.
     * Check if Container Image is Pullable
     */
    async imagePullableRaw(requestParameters: ImagePullableRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcImagePullableResponse>> {
        if (requestParameters['imageName'] == null) {
            throw new runtime.RequiredError(
                'imageName',
                'Required parameter "imageName" was null or undefined when calling imagePullable().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/container-svc/image/{imageName}/pullable`.replace(`{${"imageName"}}`, encodeURIComponent(String(requestParameters['imageName']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcImagePullableResponseFromJSON(jsonValue));
    }

    /**
     * Check if an image exists on in the container registry and is pullable.
     * Check if Container Image is Pullable
     */
    async imagePullable(requestParameters: ImagePullableRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcImagePullableResponse> {
        const response = await this.imagePullableRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List Container logs.  Requires the `container-svc:log:view` permission.
     * List Logs
     */
    async listContainerLogsRaw(requestParameters: ListContainerLogsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcListLogsResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling listContainerLogs().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/container-svc/logs`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ContainerSvcListLogsRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcListLogsResponseFromJSON(jsonValue));
    }

    /**
     * List Container logs.  Requires the `container-svc:log:view` permission.
     * List Logs
     */
    async listContainerLogs(requestParameters: ListContainerLogsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcListLogsResponse> {
        const response = await this.listContainerLogsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List containers.  Requires the `container-svc:container:view` permission.
     * List Containers
     */
    async listContainersRaw(requestParameters: ListContainersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcListContainersResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling listContainers().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/container-svc/containers`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ContainerSvcListContainersRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcListContainersResponseFromJSON(jsonValue));
    }

    /**
     * List containers.  Requires the `container-svc:container:view` permission.
     * List Containers
     */
    async listContainers(requestParameters: ListContainersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcListContainersResponse> {
        const response = await this.listContainersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Runs a Docker container with the specified parameters.  Requires the `container-svc:container:run` permission.
     * Run a Container
     */
    async runContainerRaw(requestParameters: RunContainerRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcRunContainerResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling runContainer().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/container-svc/container`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: ContainerSvcRunContainerRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcRunContainerResponseFromJSON(jsonValue));
    }

    /**
     * Runs a Docker container with the specified parameters.  Requires the `container-svc:container:run` permission.
     * Run a Container
     */
    async runContainer(requestParameters: RunContainerRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcRunContainerResponse> {
        const response = await this.runContainerRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Stops a Docker container with the specified parameters.  Requires the `container-svc:container:stop` permission.
     * Stop a Container
     */
    async stopContainerRaw(requestParameters: StopContainerRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling stopContainer().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/container-svc/container/stop`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: ContainerSvcStopContainerRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Stops a Docker container with the specified parameters.  Requires the `container-svc:container:stop` permission.
     * Stop a Container
     */
    async stopContainer(requestParameters: StopContainerRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.stopContainerRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
