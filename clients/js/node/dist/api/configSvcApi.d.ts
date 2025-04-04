/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import http from 'http';
import { ConfigSvcGetConfigResponse } from '../model/configSvcGetConfigResponse';
import { ConfigSvcSaveConfigRequest } from '../model/configSvcSaveConfigRequest';
import { Authentication, Interceptor } from '../model/models';
import { ApiKeyAuth } from '../model/models';
export declare enum ConfigSvcApiApiKeys {
    BearerAuth = 0
}
export declare class ConfigSvcApi {
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
    setApiKey(key: ConfigSvcApiApiKeys, value: string): void;
    addInterceptor(interceptor: Interceptor): void;
    /**
     * Fetch the current configuration from the server
     * @summary Get Config
     * @param namespace Namespace
     */
    getConfig(namespace?: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: ConfigSvcGetConfigResponse;
    }>;
    /**
     * Save the provided configuration to the server
     * @summary Save Config
     * @param body Save Config Request
     */
    saveConfig(body: ConfigSvcSaveConfigRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
}
