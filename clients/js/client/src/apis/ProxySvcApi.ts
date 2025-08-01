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


import * as runtime from '../runtime';
import type {
  ProxySvcErrorResponse,
  ProxySvcListCertsRequest,
  ProxySvcListCertsResponse,
  ProxySvcListRoutesRequest,
  ProxySvcListRoutesResponse,
  ProxySvcSaveCertsRequest,
  ProxySvcSaveRoutesRequest,
  ProxySvcSaveRoutesResponse,
} from '../models/index';
import {
    ProxySvcErrorResponseFromJSON,
    ProxySvcErrorResponseToJSON,
    ProxySvcListCertsRequestFromJSON,
    ProxySvcListCertsRequestToJSON,
    ProxySvcListCertsResponseFromJSON,
    ProxySvcListCertsResponseToJSON,
    ProxySvcListRoutesRequestFromJSON,
    ProxySvcListRoutesRequestToJSON,
    ProxySvcListRoutesResponseFromJSON,
    ProxySvcListRoutesResponseToJSON,
    ProxySvcSaveCertsRequestFromJSON,
    ProxySvcSaveCertsRequestToJSON,
    ProxySvcSaveRoutesRequestFromJSON,
    ProxySvcSaveRoutesRequestToJSON,
    ProxySvcSaveRoutesResponseFromJSON,
    ProxySvcSaveRoutesResponseToJSON,
} from '../models/index';

export interface ListCertsRequest {
    body?: ProxySvcListCertsRequest;
}

export interface ListRoutesRequest {
    body?: ProxySvcListRoutesRequest;
}

export interface SaveCertsRequest {
    body: ProxySvcSaveCertsRequest;
}

export interface SaveRoutesRequest {
    body: ProxySvcSaveRoutesRequest;
}

/**
 * 
 */
export class ProxySvcApi extends runtime.BaseAPI {

    /**
     * List certs that the edge proxy will use to cert requests.
     * List Certs
     */
    async listCertsRaw(requestParameters: ListCertsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ProxySvcListCertsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }


        let urlPath = `/proxy-svc/certs`;

        const response = await this.request({
            path: urlPath,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ProxySvcListCertsRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ProxySvcListCertsResponseFromJSON(jsonValue));
    }

    /**
     * List certs that the edge proxy will use to cert requests.
     * List Certs
     */
    async listCerts(requestParameters: ListCertsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ProxySvcListCertsResponse> {
        const response = await this.listCertsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List routes that the edge proxy will use to route requests.
     * List Routes
     */
    async listRoutesRaw(requestParameters: ListRoutesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ProxySvcListRoutesResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }


        let urlPath = `/proxy-svc/routes`;

        const response = await this.request({
            path: urlPath,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ProxySvcListRoutesRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ProxySvcListRoutesResponseFromJSON(jsonValue));
    }

    /**
     * List routes that the edge proxy will use to route requests.
     * List Routes
     */
    async listRoutes(requestParameters: ListRoutesRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ProxySvcListRoutesResponse> {
        const response = await this.listRoutesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint only exist for testing purposes. Only callable by admins Certs should be saved by the Proxy Svc and its edge proxying functionality internally, not through this endpoint.
     * Save Certs
     */
    async saveCertsRaw(requestParameters: SaveCertsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling saveCerts().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }


        let urlPath = `/proxy-svc/certs`;

        const response = await this.request({
            path: urlPath,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: ProxySvcSaveCertsRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * This endpoint only exist for testing purposes. Only callable by admins Certs should be saved by the Proxy Svc and its edge proxying functionality internally, not through this endpoint.
     * Save Certs
     */
    async saveCerts(requestParameters: SaveCertsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.saveCertsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Save routes that the edge proxy will use to route requests.
     * Save Routes
     */
    async saveRoutesRaw(requestParameters: SaveRoutesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ProxySvcSaveRoutesResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling saveRoutes().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }


        let urlPath = `/proxy-svc/routes`;

        const response = await this.request({
            path: urlPath,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: ProxySvcSaveRoutesRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ProxySvcSaveRoutesResponseFromJSON(jsonValue));
    }

    /**
     * Save routes that the edge proxy will use to route requests.
     * Save Routes
     */
    async saveRoutes(requestParameters: SaveRoutesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ProxySvcSaveRoutesResponse> {
        const response = await this.saveRoutesRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
