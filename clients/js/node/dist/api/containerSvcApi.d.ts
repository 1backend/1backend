/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import http from 'http';
import { ContainerSvcBuildImageRequest } from '../model/containerSvcBuildImageRequest';
import { ContainerSvcContainerIsRunningResponse } from '../model/containerSvcContainerIsRunningResponse';
import { ContainerSvcDaemonInfoResponse } from '../model/containerSvcDaemonInfoResponse';
import { ContainerSvcGetContainerSummaryResponse } from '../model/containerSvcGetContainerSummaryResponse';
import { ContainerSvcGetHostResponse } from '../model/containerSvcGetHostResponse';
import { ContainerSvcImagePullableResponse } from '../model/containerSvcImagePullableResponse';
import { ContainerSvcListContainersRequest } from '../model/containerSvcListContainersRequest';
import { ContainerSvcListContainersResponse } from '../model/containerSvcListContainersResponse';
import { ContainerSvcListLogsRequest } from '../model/containerSvcListLogsRequest';
import { ContainerSvcListLogsResponse } from '../model/containerSvcListLogsResponse';
import { ContainerSvcRunContainerRequest } from '../model/containerSvcRunContainerRequest';
import { ContainerSvcRunContainerResponse } from '../model/containerSvcRunContainerResponse';
import { ContainerSvcStopContainerRequest } from '../model/containerSvcStopContainerRequest';
import { Authentication, Interceptor } from '../model/models';
import { ApiKeyAuth } from '../model/models';
export declare enum ContainerSvcApiApiKeys {
    BearerAuth = 0
}
export declare class ContainerSvcApi {
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
    setApiKey(key: ContainerSvcApiApiKeys, value: string): void;
    addInterceptor(interceptor: Interceptor): void;
    /**
     * Builds a Docker image with the specified parameters.  Requires the `container-svc:image:build` permission.
     * @summary Build an Image
     * @param body Build Image Request
     */
    buildImage(body: ContainerSvcBuildImageRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
    /**
     * Retrieve detailed information about the availability and status of container daemons on the node.
     * @summary Get Container Daemon Information
     */
    containerDaemonInfo(options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ContainerSvcDaemonInfoResponse;
    }>;
    /**
     * Check if a Docker container is running, identified by hash or name.
     * @summary Check If a Container Is Running
     * @param hash Container Hash
     * @param name Container Name
     */
    containerIsRunning(hash?: string, name?: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ContainerSvcContainerIsRunningResponse;
    }>;
    /**
     * Get a summary of the Docker container identified by hash or name, limited to a specified number of lines.
     * @summary Get Container Summary
     * @param hash Container Hash
     * @param name Container Name
     * @param lines Number of Lines
     */
    containerSummary(hash?: string, name?: string, lines?: number, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ContainerSvcGetContainerSummaryResponse;
    }>;
    /**
     * Retrieve information about the Container host
     * @summary Get Container Host
     */
    getHost(options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ContainerSvcGetHostResponse;
    }>;
    /**
     * Check if an image exists on in the container registry and is pullable.
     * @summary Check if Container Image is Pullable
     * @param imageName Image name
     */
    imagePullable(imageName: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ContainerSvcImagePullableResponse;
    }>;
    /**
     * List Container logs.  Requires the `container-svc:log:view` permission.
     * @summary List Logs
     * @param body List Logs Request
     */
    listContainerLogs(body: ContainerSvcListLogsRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ContainerSvcListLogsResponse;
    }>;
    /**
     * List containers.  Requires the `container-svc:container:view` permission.
     * @summary List Containers
     * @param body List Containers Request
     */
    listContainers(body: ContainerSvcListContainersRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ContainerSvcListContainersResponse;
    }>;
    /**
     * Runs a Docker container with the specified parameters.  Requires the `container-svc:container:run` permission.
     * @summary Run a Container
     * @param body Run Container Request
     */
    runContainer(body: ContainerSvcRunContainerRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ContainerSvcRunContainerResponse;
    }>;
    /**
     * Stops a Docker container with the specified parameters.  Requires the `container-svc:container:stop` permission.
     * @summary Stop a Container
     * @param body Stop Container Request
     */
    stopContainer(body: ContainerSvcStopContainerRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
}
