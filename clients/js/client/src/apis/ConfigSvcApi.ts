/* tslint:disable */
/* eslint-disable */
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


import * as runtime from '../runtime';
import type {
  ConfigSvcGetConfigResponse,
  ConfigSvcSaveConfigRequest,
} from '../models/index';
import {
    ConfigSvcGetConfigResponseFromJSON,
    ConfigSvcGetConfigResponseToJSON,
    ConfigSvcSaveConfigRequestFromJSON,
    ConfigSvcSaveConfigRequestToJSON,
} from '../models/index';

export interface GetConfigRequest {
    namespace?: string;
}

export interface SaveConfigRequest {
    body: ConfigSvcSaveConfigRequest;
}

/**
 * 
 */
export class ConfigSvcApi extends runtime.BaseAPI {

    /**
     * Fetch the current configuration from the server
     * Get Config
     */
    async getConfigRaw(requestParameters: GetConfigRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ConfigSvcGetConfigResponse>> {
        const queryParameters: any = {};

        if (requestParameters['namespace'] != null) {
            queryParameters['namespace'] = requestParameters['namespace'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/config-svc/config`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ConfigSvcGetConfigResponseFromJSON(jsonValue));
    }

    /**
     * Fetch the current configuration from the server
     * Get Config
     */
    async getConfig(requestParameters: GetConfigRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ConfigSvcGetConfigResponse> {
        const response = await this.getConfigRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Save the provided configuration to the server
     * Save Config
     */
    async saveConfigRaw(requestParameters: SaveConfigRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling saveConfig().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/config-svc/config`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: ConfigSvcSaveConfigRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Save the provided configuration to the server
     * Save Config
     */
    async saveConfig(requestParameters: SaveConfigRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.saveConfigRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
