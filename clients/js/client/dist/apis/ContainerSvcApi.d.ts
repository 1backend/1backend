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
import type { ContainerSvcBuildImageRequest, ContainerSvcContainerIsRunningResponse, ContainerSvcDaemonInfoResponse, ContainerSvcGetContainerSummaryResponse, ContainerSvcGetHostResponse, ContainerSvcImagePullableResponse, ContainerSvcListContainersRequest, ContainerSvcListContainersResponse, ContainerSvcListLogsRequest, ContainerSvcListLogsResponse, ContainerSvcRunContainerRequest, ContainerSvcRunContainerResponse, ContainerSvcStopContainerRequest } from '../models/index';
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
export declare class ContainerSvcApi extends runtime.BaseAPI {
    /**
     * Builds a Docker image with the specified parameters.  Requires the `container-svc:image:build` permission.
     * Build an Image
     */
    buildImageRaw(requestParameters: BuildImageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Builds a Docker image with the specified parameters.  Requires the `container-svc:image:build` permission.
     * Build an Image
     */
    buildImage(requestParameters: BuildImageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Retrieve detailed information about the availability and status of container daemons on the node.
     * Get Container Daemon Information
     */
    containerDaemonInfoRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcDaemonInfoResponse>>;
    /**
     * Retrieve detailed information about the availability and status of container daemons on the node.
     * Get Container Daemon Information
     */
    containerDaemonInfo(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcDaemonInfoResponse>;
    /**
     * Check if a Docker container is running, identified by hash or name.
     * Check If a Container Is Running
     */
    containerIsRunningRaw(requestParameters: ContainerIsRunningRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcContainerIsRunningResponse>>;
    /**
     * Check if a Docker container is running, identified by hash or name.
     * Check If a Container Is Running
     */
    containerIsRunning(requestParameters?: ContainerIsRunningRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcContainerIsRunningResponse>;
    /**
     * Get a summary of the Docker container identified by hash or name, limited to a specified number of lines.
     * Get Container Summary
     */
    containerSummaryRaw(requestParameters: ContainerSummaryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcGetContainerSummaryResponse>>;
    /**
     * Get a summary of the Docker container identified by hash or name, limited to a specified number of lines.
     * Get Container Summary
     */
    containerSummary(requestParameters?: ContainerSummaryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcGetContainerSummaryResponse>;
    /**
     * Retrieve information about the Container host
     * Get Container Host
     */
    getHostRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcGetHostResponse>>;
    /**
     * Retrieve information about the Container host
     * Get Container Host
     */
    getHost(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcGetHostResponse>;
    /**
     * Check if an image exists on in the container registry and is pullable.
     * Check if Container Image is Pullable
     */
    imagePullableRaw(requestParameters: ImagePullableRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcImagePullableResponse>>;
    /**
     * Check if an image exists on in the container registry and is pullable.
     * Check if Container Image is Pullable
     */
    imagePullable(requestParameters: ImagePullableRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcImagePullableResponse>;
    /**
     * List Container logs.  Requires the `container-svc:log:view` permission.
     * List Logs
     */
    listContainerLogsRaw(requestParameters: ListContainerLogsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcListLogsResponse>>;
    /**
     * List Container logs.  Requires the `container-svc:log:view` permission.
     * List Logs
     */
    listContainerLogs(requestParameters: ListContainerLogsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcListLogsResponse>;
    /**
     * List containers.  Requires the `container-svc:container:view` permission.
     * List Containers
     */
    listContainersRaw(requestParameters: ListContainersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcListContainersResponse>>;
    /**
     * List containers.  Requires the `container-svc:container:view` permission.
     * List Containers
     */
    listContainers(requestParameters: ListContainersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcListContainersResponse>;
    /**
     * Runs a Docker container with the specified parameters.  Requires the `container-svc:container:run` permission.
     * Run a Container
     */
    runContainerRaw(requestParameters: RunContainerRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ContainerSvcRunContainerResponse>>;
    /**
     * Runs a Docker container with the specified parameters.  Requires the `container-svc:container:run` permission.
     * Run a Container
     */
    runContainer(requestParameters: RunContainerRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ContainerSvcRunContainerResponse>;
    /**
     * Stops a Docker container with the specified parameters.  Requires the `container-svc:container:stop` permission.
     * Stop a Container
     */
    stopContainerRaw(requestParameters: StopContainerRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Stops a Docker container with the specified parameters.  Requires the `container-svc:container:stop` permission.
     * Stop a Container
     */
    stopContainer(requestParameters: StopContainerRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
}
