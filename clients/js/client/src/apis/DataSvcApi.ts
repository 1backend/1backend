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
  DataSvcCreateObjectRequest,
  DataSvcCreateObjectResponse,
  DataSvcDeleteObjectRequest,
  DataSvcErrorResponse,
  DataSvcQueryRequest,
  DataSvcQueryResponse,
  DataSvcUpdateObjectsRequest,
  DataSvcUpsertObjectRequest,
  DataSvcUpsertObjectResponse,
} from '../models/index';
import {
    DataSvcCreateObjectRequestFromJSON,
    DataSvcCreateObjectRequestToJSON,
    DataSvcCreateObjectResponseFromJSON,
    DataSvcCreateObjectResponseToJSON,
    DataSvcDeleteObjectRequestFromJSON,
    DataSvcDeleteObjectRequestToJSON,
    DataSvcErrorResponseFromJSON,
    DataSvcErrorResponseToJSON,
    DataSvcQueryRequestFromJSON,
    DataSvcQueryRequestToJSON,
    DataSvcQueryResponseFromJSON,
    DataSvcQueryResponseToJSON,
    DataSvcUpdateObjectsRequestFromJSON,
    DataSvcUpdateObjectsRequestToJSON,
    DataSvcUpsertObjectRequestFromJSON,
    DataSvcUpsertObjectRequestToJSON,
    DataSvcUpsertObjectResponseFromJSON,
    DataSvcUpsertObjectResponseToJSON,
} from '../models/index';

export interface CreateObjectRequest {
    body: DataSvcCreateObjectRequest;
}

export interface DeleteObjectsRequest {
    body: DataSvcDeleteObjectRequest;
}

export interface QueryObjectsRequest {
    body?: DataSvcQueryRequest;
}

export interface UpdateObjectsRequest {
    body: DataSvcUpdateObjectsRequest;
}

export interface UpsertObjectRequest {
    objectId: string;
    body: DataSvcUpsertObjectRequest;
}

export interface UpsertObjectsRequest {
    body: DataSvcUpsertObjectRequest;
}

/**
 * 
 */
export class DataSvcApi extends runtime.BaseAPI {

    /**
     * Creates a new object with the provided details. Requires authorization and user authentication.
     * Create a Generic Object
     */
    async createObjectRaw(requestParameters: CreateObjectRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DataSvcCreateObjectResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling createObject().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/data-svc/object`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: DataSvcCreateObjectRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DataSvcCreateObjectResponseFromJSON(jsonValue));
    }

    /**
     * Creates a new object with the provided details. Requires authorization and user authentication.
     * Create a Generic Object
     */
    async createObject(requestParameters: CreateObjectRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DataSvcCreateObjectResponse> {
        const response = await this.createObjectRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Deletes all objects matchin the provided filters.
     * Delete Objects
     */
    async deleteObjectsRaw(requestParameters: DeleteObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling deleteObjects().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/data-svc/objects/delete`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: DataSvcDeleteObjectRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Deletes all objects matchin the provided filters.
     * Delete Objects
     */
    async deleteObjects(requestParameters: DeleteObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.deleteObjectsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieves objects from a specified table based on search criteria. Requires authorization and user authentication.   Use helper functions in your respective client library such as condition constructors (`equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
     * Query Objects
     */
    async queryObjectsRaw(requestParameters: QueryObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DataSvcQueryResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/data-svc/objects`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: DataSvcQueryRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DataSvcQueryResponseFromJSON(jsonValue));
    }

    /**
     * Retrieves objects from a specified table based on search criteria. Requires authorization and user authentication.   Use helper functions in your respective client library such as condition constructors (`equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
     * Query Objects
     */
    async queryObjects(requestParameters: QueryObjectsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DataSvcQueryResponse> {
        const response = await this.queryObjectsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Update fields of objects that match the given filters using the provided object. Any fields not included in the incoming object will remain unchanged.
     * Update Objects
     */
    async updateObjectsRaw(requestParameters: UpdateObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling updateObjects().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/data-svc/objects/update`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: DataSvcUpdateObjectsRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Update fields of objects that match the given filters using the provided object. Any fields not included in the incoming object will remain unchanged.
     * Update Objects
     */
    async updateObjects(requestParameters: UpdateObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object> {
        const response = await this.updateObjectsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Creates a new dynamic object or updates an existing one based on the provided data. Requires authorization and user authentication.
     * Upsert a Generic Object
     */
    async upsertObjectRaw(requestParameters: UpsertObjectRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DataSvcUpsertObjectResponse>> {
        if (requestParameters['objectId'] == null) {
            throw new runtime.RequiredError(
                'objectId',
                'Required parameter "objectId" was null or undefined when calling upsertObject().'
            );
        }

        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling upsertObject().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/data-svc/object/{objectId}`.replace(`{${"objectId"}}`, encodeURIComponent(String(requestParameters['objectId']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: DataSvcUpsertObjectRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DataSvcUpsertObjectResponseFromJSON(jsonValue));
    }

    /**
     * Creates a new dynamic object or updates an existing one based on the provided data. Requires authorization and user authentication.
     * Upsert a Generic Object
     */
    async upsertObject(requestParameters: UpsertObjectRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DataSvcUpsertObjectResponse> {
        const response = await this.upsertObjectRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Upserts objects by ids.
     * Upsert Objects
     */
    async upsertObjectsRaw(requestParameters: UpsertObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DataSvcUpsertObjectResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling upsertObjects().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/data-svc/objects/upsert`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: DataSvcUpsertObjectRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DataSvcUpsertObjectResponseFromJSON(jsonValue));
    }

    /**
     * Upserts objects by ids.
     * Upsert Objects
     */
    async upsertObjects(requestParameters: UpsertObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DataSvcUpsertObjectResponse> {
        const response = await this.upsertObjectsRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
