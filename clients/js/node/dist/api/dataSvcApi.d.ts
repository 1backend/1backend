/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import http from 'http';
import { DataSvcCreateObjectRequest } from '../model/dataSvcCreateObjectRequest';
import { DataSvcCreateObjectResponse } from '../model/dataSvcCreateObjectResponse';
import { DataSvcDeleteObjectRequest } from '../model/dataSvcDeleteObjectRequest';
import { DataSvcQueryRequest } from '../model/dataSvcQueryRequest';
import { DataSvcQueryResponse } from '../model/dataSvcQueryResponse';
import { DataSvcUpdateObjectsRequest } from '../model/dataSvcUpdateObjectsRequest';
import { DataSvcUpsertObjectRequest } from '../model/dataSvcUpsertObjectRequest';
import { DataSvcUpsertObjectResponse } from '../model/dataSvcUpsertObjectResponse';
import { Authentication, Interceptor } from '../model/models';
import { ApiKeyAuth } from '../model/models';
export declare enum DataSvcApiApiKeys {
    BearerAuth = 0
}
export declare class DataSvcApi {
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
    setApiKey(key: DataSvcApiApiKeys, value: string): void;
    addInterceptor(interceptor: Interceptor): void;
    /**
     * Creates a new object with the provided details. Requires authorization and user authentication.
     * @summary Create a Generic Object
     * @param body Create request payload
     */
    createObject(body: DataSvcCreateObjectRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: DataSvcCreateObjectResponse;
    }>;
    /**
     * Deletes all objects matchin the provided filters.
     * @summary Delete Objects
     * @param body Delete request payload
     */
    deleteObjects(body: DataSvcDeleteObjectRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
    /**
     * Retrieves objects from a specified table based on search criteria. Requires authorization and user authentication.   Use helper functions in your respective client library such as condition constructors (`equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
     * @summary Query Objects
     * @param body Query Request
     */
    queryObjects(body?: DataSvcQueryRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: DataSvcQueryResponse;
    }>;
    /**
     * Update fields of objects that match the given filters using the provided object. Any fields not included in the incoming object will remain unchanged.
     * @summary Update Objects
     * @param body Update request payload
     */
    updateObjects(body: DataSvcUpdateObjectsRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: object;
    }>;
    /**
     * Creates a new dynamic object or updates an existing one based on the provided data. Requires authorization and user authentication.
     * @summary Upsert a Generic Object
     * @param objectId Object ID
     * @param body Upsert request payload
     */
    upsertObject(objectId: string, body: DataSvcUpsertObjectRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: DataSvcUpsertObjectResponse;
    }>;
    /**
     * Upserts objects by ids.
     * @summary Upsert Objects
     * @param body Upsert request payload
     */
    upsertObjects(body: DataSvcUpsertObjectRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: DataSvcUpsertObjectResponse;
    }>;
}
