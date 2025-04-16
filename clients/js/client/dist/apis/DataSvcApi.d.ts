/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import * as runtime from '../runtime';
import type { DataSvcCreateObjectRequest, DataSvcCreateObjectResponse, DataSvcDeleteObjectRequest, DataSvcQueryRequest, DataSvcQueryResponse, DataSvcUpdateObjectsRequest, DataSvcUpsertObjectRequest, DataSvcUpsertObjectResponse } from '../models/index';
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
export declare class DataSvcApi extends runtime.BaseAPI {
    /**
     * Creates a new object with the provided details. Requires authorization and user authentication.
     * Create a Generic Object
     */
    createObjectRaw(requestParameters: CreateObjectRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DataSvcCreateObjectResponse>>;
    /**
     * Creates a new object with the provided details. Requires authorization and user authentication.
     * Create a Generic Object
     */
    createObject(requestParameters: CreateObjectRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DataSvcCreateObjectResponse>;
    /**
     * Deletes all objects matchin the provided filters.
     * Delete Objects
     */
    deleteObjectsRaw(requestParameters: DeleteObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Deletes all objects matchin the provided filters.
     * Delete Objects
     */
    deleteObjects(requestParameters: DeleteObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Retrieves objects from a specified table based on search criteria. Requires authorization and user authentication.   Use helper functions in your respective client library such as condition constructors (`equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
     * Query Objects
     */
    queryObjectsRaw(requestParameters: QueryObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DataSvcQueryResponse>>;
    /**
     * Retrieves objects from a specified table based on search criteria. Requires authorization and user authentication.   Use helper functions in your respective client library such as condition constructors (`equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
     * Query Objects
     */
    queryObjects(requestParameters?: QueryObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DataSvcQueryResponse>;
    /**
     * Update fields of objects that match the given filters using the provided object. Any fields not included in the incoming object will remain unchanged.
     * Update Objects
     */
    updateObjectsRaw(requestParameters: UpdateObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Update fields of objects that match the given filters using the provided object. Any fields not included in the incoming object will remain unchanged.
     * Update Objects
     */
    updateObjects(requestParameters: UpdateObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Creates a new dynamic object or updates an existing one based on the provided data. Requires authorization and user authentication.
     * Upsert a Generic Object
     */
    upsertObjectRaw(requestParameters: UpsertObjectRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DataSvcUpsertObjectResponse>>;
    /**
     * Creates a new dynamic object or updates an existing one based on the provided data. Requires authorization and user authentication.
     * Upsert a Generic Object
     */
    upsertObject(requestParameters: UpsertObjectRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DataSvcUpsertObjectResponse>;
    /**
     * Upserts objects by ids.
     * Upsert Objects
     */
    upsertObjectsRaw(requestParameters: UpsertObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DataSvcUpsertObjectResponse>>;
    /**
     * Upserts objects by ids.
     * Upsert Objects
     */
    upsertObjects(requestParameters: UpsertObjectsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DataSvcUpsertObjectResponse>;
}
