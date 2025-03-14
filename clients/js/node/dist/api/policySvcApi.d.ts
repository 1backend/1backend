/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import http from 'http';
import { PolicySvcCheckRequest } from '../model/policySvcCheckRequest';
import { PolicySvcCheckResponse } from '../model/policySvcCheckResponse';
import { PolicySvcUpsertInstanceRequest } from '../model/policySvcUpsertInstanceRequest';
import { Authentication, Interceptor } from '../model/models';
import { ApiKeyAuth } from '../model/models';
export declare enum PolicySvcApiApiKeys {
    BearerAuth = 0
}
export declare class PolicySvcApi {
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
    setApiKey(key: PolicySvcApiApiKeys, value: string): void;
    addInterceptor(interceptor: Interceptor): void;
    /**
     * Check records a resource access and returns if the access is allowed.
     * @summary Check
     * @param body Check Request
     */
    check(body: PolicySvcCheckRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: PolicySvcCheckResponse;
    }>;
    /**
     * Allows user to upsert a new policy instance based on a template.
     * @summary Upsert an Instance
     * @param instanceId Instance ID
     * @param body Upsert Instance Request
     */
    upsertInstance(instanceId: string, body: PolicySvcUpsertInstanceRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
}
