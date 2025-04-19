/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
import * as runtime from '../runtime';
import { ContainerSvcBuildImageRequestToJSON, ContainerSvcContainerIsRunningResponseFromJSON, ContainerSvcDaemonInfoResponseFromJSON, ContainerSvcGetContainerSummaryResponseFromJSON, ContainerSvcGetHostResponseFromJSON, ContainerSvcImagePullableResponseFromJSON, ContainerSvcListContainersRequestToJSON, ContainerSvcListContainersResponseFromJSON, ContainerSvcListLogsRequestToJSON, ContainerSvcListLogsResponseFromJSON, ContainerSvcRunContainerRequestToJSON, ContainerSvcRunContainerResponseFromJSON, ContainerSvcStopContainerRequestToJSON, } from '../models/index';
/**
 *
 */
export class ContainerSvcApi extends runtime.BaseAPI {
    /**
     * Builds a Docker image with the specified parameters.  Requires the `container-svc:image:build` permission.
     * Build an Image
     */
    buildImageRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling buildImage().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/container-svc/image`,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: ContainerSvcBuildImageRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Builds a Docker image with the specified parameters.  Requires the `container-svc:image:build` permission.
     * Build an Image
     */
    buildImage(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.buildImageRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Retrieve detailed information about the availability and status of container daemons on the node.
     * Get Container Daemon Information
     */
    containerDaemonInfoRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/container-svc/daemon/info`,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcDaemonInfoResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieve detailed information about the availability and status of container daemons on the node.
     * Get Container Daemon Information
     */
    containerDaemonInfo(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.containerDaemonInfoRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * Check if a Docker container is running, identified by hash or name.
     * Check If a Container Is Running
     */
    containerIsRunningRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            if (requestParameters['hash'] != null) {
                queryParameters['hash'] = requestParameters['hash'];
            }
            if (requestParameters['name'] != null) {
                queryParameters['name'] = requestParameters['name'];
            }
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/container-svc/container/is-running`,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcContainerIsRunningResponseFromJSON(jsonValue));
        });
    }
    /**
     * Check if a Docker container is running, identified by hash or name.
     * Check If a Container Is Running
     */
    containerIsRunning() {
        return __awaiter(this, arguments, void 0, function* (requestParameters = {}, initOverrides) {
            const response = yield this.containerIsRunningRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Get a summary of the Docker container identified by hash or name, limited to a specified number of lines.
     * Get Container Summary
     */
    containerSummaryRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            if (requestParameters['hash'] != null) {
                queryParameters['hash'] = requestParameters['hash'];
            }
            if (requestParameters['name'] != null) {
                queryParameters['name'] = requestParameters['name'];
            }
            if (requestParameters['lines'] != null) {
                queryParameters['lines'] = requestParameters['lines'];
            }
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/container-svc/container/summary`,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcGetContainerSummaryResponseFromJSON(jsonValue));
        });
    }
    /**
     * Get a summary of the Docker container identified by hash or name, limited to a specified number of lines.
     * Get Container Summary
     */
    containerSummary() {
        return __awaiter(this, arguments, void 0, function* (requestParameters = {}, initOverrides) {
            const response = yield this.containerSummaryRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Retrieve information about the Container host
     * Get Container Host
     */
    getHostRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/container-svc/host`,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcGetHostResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieve information about the Container host
     * Get Container Host
     */
    getHost(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.getHostRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * Check if an image exists on in the container registry and is pullable.
     * Check if Container Image is Pullable
     */
    imagePullableRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['imageName'] == null) {
                throw new runtime.RequiredError('imageName', 'Required parameter "imageName" was null or undefined when calling imagePullable().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/container-svc/image/{imageName}/pullable`.replace(`{${"imageName"}}`, encodeURIComponent(String(requestParameters['imageName']))),
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcImagePullableResponseFromJSON(jsonValue));
        });
    }
    /**
     * Check if an image exists on in the container registry and is pullable.
     * Check if Container Image is Pullable
     */
    imagePullable(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.imagePullableRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * List Container logs.  Requires the `container-svc:log:view` permission.
     * List Logs
     */
    listContainerLogsRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling listContainerLogs().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/container-svc/logs`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: ContainerSvcListLogsRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcListLogsResponseFromJSON(jsonValue));
        });
    }
    /**
     * List Container logs.  Requires the `container-svc:log:view` permission.
     * List Logs
     */
    listContainerLogs(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.listContainerLogsRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * List containers.  Requires the `container-svc:container:view` permission.
     * List Containers
     */
    listContainersRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling listContainers().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/container-svc/containers`,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: ContainerSvcListContainersRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcListContainersResponseFromJSON(jsonValue));
        });
    }
    /**
     * List containers.  Requires the `container-svc:container:view` permission.
     * List Containers
     */
    listContainers(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.listContainersRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Runs a Docker container with the specified parameters.  Requires the `container-svc:container:run` permission.
     * Run a Container
     */
    runContainerRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling runContainer().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/container-svc/container`,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: ContainerSvcRunContainerRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => ContainerSvcRunContainerResponseFromJSON(jsonValue));
        });
    }
    /**
     * Runs a Docker container with the specified parameters.  Requires the `container-svc:container:run` permission.
     * Run a Container
     */
    runContainer(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.runContainerRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Stops a Docker container with the specified parameters.  Requires the `container-svc:container:stop` permission.
     * Stop a Container
     */
    stopContainerRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling stopContainer().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            const response = yield this.request({
                path: `/container-svc/container/stop`,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: ContainerSvcStopContainerRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Stops a Docker container with the specified parameters.  Requires the `container-svc:container:stop` permission.
     * Stop a Container
     */
    stopContainer(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.stopContainerRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
}
