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
import type { DatastoreFilter } from './DatastoreFilter';
import type { DataSvcObject } from './DataSvcObject';
/**
 *
 * @export
 * @interface DataSvcUpdateObjectsRequest
 */
export interface DataSvcUpdateObjectsRequest {
    /**
     * Filters to determine which objects will be updated.
     * Only objects matching all filters will be modified.
     * @type {Array<DatastoreFilter>}
     * @memberof DataSvcUpdateObjectsRequest
     */
    filters?: Array<DatastoreFilter>;
    /**
     * The object containing the fields to update in matching objects.
     * @type {DataSvcObject}
     * @memberof DataSvcUpdateObjectsRequest
     */
    object?: DataSvcObject;
    /**
     *
     * @type {string}
     * @memberof DataSvcUpdateObjectsRequest
     */
    table?: string;
}
/**
 * Check if a given object implements the DataSvcUpdateObjectsRequest interface.
 */
export declare function instanceOfDataSvcUpdateObjectsRequest(value: object): value is DataSvcUpdateObjectsRequest;
export declare function DataSvcUpdateObjectsRequestFromJSON(json: any): DataSvcUpdateObjectsRequest;
export declare function DataSvcUpdateObjectsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcUpdateObjectsRequest;
export declare function DataSvcUpdateObjectsRequestToJSON(json: any): DataSvcUpdateObjectsRequest;
export declare function DataSvcUpdateObjectsRequestToJSONTyped(value?: DataSvcUpdateObjectsRequest | null, ignoreDiscriminator?: boolean): any;
