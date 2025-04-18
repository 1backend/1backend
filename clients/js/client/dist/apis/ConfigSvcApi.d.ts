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
import * as runtime from '../runtime';
import type { ConfigSvcGetConfigResponse, ConfigSvcSaveConfigRequest } from '../models/index';
export interface GetConfigRequest {
    namespace?: string;
}
export interface SaveConfigRequest {
    body: ConfigSvcSaveConfigRequest;
}
/**
 *
 */
export declare class ConfigSvcApi extends runtime.BaseAPI {
    /**
     * Fetch the current configuration from the server
     * Get Config
     */
    getConfigRaw(requestParameters: GetConfigRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ConfigSvcGetConfigResponse>>;
    /**
     * Fetch the current configuration from the server
     * Get Config
     */
    getConfig(requestParameters?: GetConfigRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ConfigSvcGetConfigResponse>;
    /**
     * Save the provided configuration to the server
     * Save Config
     */
    saveConfigRaw(requestParameters: SaveConfigRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Save the provided configuration to the server
     * Save Config
     */
    saveConfig(requestParameters: SaveConfigRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
}
