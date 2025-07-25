/* tslint:disable */
/* eslint-disable */
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
import { RegistrySvcListDefinitionsResponseFromJSON, RegistrySvcListInstancesResponseFromJSON, RegistrySvcListNodesRequestToJSON, RegistrySvcListNodesResponseFromJSON, RegistrySvcNodeSelfResponseFromJSON, RegistrySvcRegisterInstanceRequestToJSON, RegistrySvcSaveDefinitionRequestToJSON, } from '../models/index';
/**
 *
 */
export class RegistrySvcApi extends runtime.BaseAPI {
    /**
     * Deletes a registered definition by ID.
     * Delete Definition
     */
    deleteDefinitionRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['id'] == null) {
                throw new runtime.RequiredError('id', 'Required parameter "id" was null or undefined when calling deleteDefinition().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/registry-svc/definition/{id}`;
            urlPath = urlPath.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id'])));
            const response = yield this.request({
                path: urlPath,
                method: 'DELETE',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.VoidApiResponse(response);
        });
    }
    /**
     * Deletes a registered definition by ID.
     * Delete Definition
     */
    deleteDefinition(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            yield this.deleteDefinitionRaw(requestParameters, initOverrides);
        });
    }
    /**
     * Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it\'s still present in the database.
     * Delete Node
     */
    deleteNodeRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['url'] == null) {
                throw new runtime.RequiredError('url', 'Required parameter "url" was null or undefined when calling deleteNode().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/registry-svc/node/{url}`;
            urlPath = urlPath.replace(`{${"url"}}`, encodeURIComponent(String(requestParameters['url'])));
            const response = yield this.request({
                path: urlPath,
                method: 'DELETE',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.VoidApiResponse(response);
        });
    }
    /**
     * Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it\'s still present in the database.
     * Delete Node
     */
    deleteNode(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            yield this.deleteNodeRaw(requestParameters, initOverrides);
        });
    }
    /**
     * This endpoint is used to test the server\'s response to a GET request. It echoes back the query parameters as a JSON object.
     * Echo the query parameters in the response body.
     */
    echoGetRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/registry-svc/echo`;
            const response = yield this.request({
                path: urlPath,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * This endpoint is used to test the server\'s response to a GET request. It echoes back the query parameters as a JSON object.
     * Echo the query parameters in the response body.
     */
    echoGet(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.echoGetRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * This endpoint is used to test the server\'s response to a request. It simply echoes back the request body as a JSON response.
     * Echo the request body in the response body.
     */
    echoPostRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/registry-svc/echo`;
            const response = yield this.request({
                path: urlPath,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * This endpoint is used to test the server\'s response to a request. It simply echoes back the request body as a JSON response.
     * Echo the request body in the response body.
     */
    echoPost(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.echoPostRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * This endpoint is used to test the server\'s response to a request. It simply echoes back the request body as a JSON response.
     * Echo the request body in the response body.
     */
    echoPutRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/registry-svc/echo`;
            const response = yield this.request({
                path: urlPath,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * This endpoint is used to test the server\'s response to a request. It simply echoes back the request body as a JSON response.
     * Echo the request body in the response body.
     */
    echoPut(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.echoPutRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * Retrieves a list of all definitions or filters them by specific criteria.
     * List Definitions
     */
    listDefinitionsRaw(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/registry-svc/definitions`;
            const response = yield this.request({
                path: urlPath,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => RegistrySvcListDefinitionsResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieves a list of all definitions or filters them by specific criteria.
     * List Definitions
     */
    listDefinitions(initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.listDefinitionsRaw(initOverrides);
            return yield response.value();
        });
    }
    /**
     * Retrieves a list of all instances or filters them by specific criteria (e.g., host, IP).
     * List Service Instances
     */
    listInstancesRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
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
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/registry-svc/instances`;
            const response = yield this.request({
                path: urlPath,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => RegistrySvcListInstancesResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieves a list of all instances or filters them by specific criteria (e.g., host, IP).
     * List Service Instances
     */
    listInstances() {
        return __awaiter(this, arguments, void 0, function* (requestParameters = {}, initOverrides) {
            const response = yield this.listInstancesRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Retrieve a list of nodes.
     * List Nodes
     */
    listNodesRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/registry-svc/nodes`;
            const response = yield this.request({
                path: urlPath,
                method: 'POST',
                headers: headerParameters,
                query: queryParameters,
                body: RegistrySvcListNodesRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => RegistrySvcListNodesResponseFromJSON(jsonValue));
        });
    }
    /**
     * Retrieve a list of nodes.
     * List Nodes
     */
    listNodes() {
        return __awaiter(this, arguments, void 0, function* (requestParameters = {}, initOverrides) {
            const response = yield this.listNodesRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Registers an instance. Idempotent.
     * Register Instance
     */
    registerInstanceRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling registerInstance().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/registry-svc/instance`;
            const response = yield this.request({
                path: urlPath,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: RegistrySvcRegisterInstanceRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Registers an instance. Idempotent.
     * Register Instance
     */
    registerInstance(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.registerInstanceRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Removes a registered instance by ID.
     * Remove Instance
     */
    removeInstanceRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['id'] == null) {
                throw new runtime.RequiredError('id', 'Required parameter "id" was null or undefined when calling removeInstance().');
            }
            const queryParameters = {};
            const headerParameters = {};
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/registry-svc/instance/{id}`;
            urlPath = urlPath.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id'])));
            const response = yield this.request({
                path: urlPath,
                method: 'DELETE',
                headers: headerParameters,
                query: queryParameters,
            }, initOverrides);
            return new runtime.VoidApiResponse(response);
        });
    }
    /**
     * Removes a registered instance by ID.
     * Remove Instance
     */
    removeInstance(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            yield this.removeInstanceRaw(requestParameters, initOverrides);
        });
    }
    /**
     * Registers a new definition, associating an definition address with a slug acquired from the bearer token.
     * Register a Definition
     */
    saveDefinitionRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            if (requestParameters['body'] == null) {
                throw new runtime.RequiredError('body', 'Required parameter "body" was null or undefined when calling saveDefinition().');
            }
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/registry-svc/definition`;
            const response = yield this.request({
                path: urlPath,
                method: 'PUT',
                headers: headerParameters,
                query: queryParameters,
                body: RegistrySvcSaveDefinitionRequestToJSON(requestParameters['body']),
            }, initOverrides);
            return new runtime.JSONApiResponse(response);
        });
    }
    /**
     * Registers a new definition, associating an definition address with a slug acquired from the bearer token.
     * Register a Definition
     */
    saveDefinition(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield this.saveDefinitionRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
    /**
     * Show the local node.
     * View Self Node
     */
    selfNodeRaw(requestParameters, initOverrides) {
        return __awaiter(this, void 0, void 0, function* () {
            const queryParameters = {};
            const headerParameters = {};
            headerParameters['Content-Type'] = 'application/json';
            if (this.configuration && this.configuration.apiKey) {
                headerParameters["Authorization"] = yield this.configuration.apiKey("Authorization"); // BearerAuth authentication
            }
            let urlPath = `/registry-svc/node/self`;
            const response = yield this.request({
                path: urlPath,
                method: 'GET',
                headers: headerParameters,
                query: queryParameters,
                body: requestParameters['body'],
            }, initOverrides);
            return new runtime.JSONApiResponse(response, (jsonValue) => RegistrySvcNodeSelfResponseFromJSON(jsonValue));
        });
    }
    /**
     * Show the local node.
     * View Self Node
     */
    selfNode() {
        return __awaiter(this, arguments, void 0, function* (requestParameters = {}, initOverrides) {
            const response = yield this.selfNodeRaw(requestParameters, initOverrides);
            return yield response.value();
        });
    }
}
