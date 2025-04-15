/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { DatastoreQuery } from './DatastoreQuery';
/**
 *
 * @export
 * @interface DataSvcQueryRequest
 */
export interface DataSvcQueryRequest {
    /**
     *
     * @type {DatastoreQuery}
     * @memberof DataSvcQueryRequest
     */
    query?: DatastoreQuery;
    /**
     *
     * @type {Array<string>}
     * @memberof DataSvcQueryRequest
     */
    readers?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof DataSvcQueryRequest
     */
    table?: string;
}
/**
 * Check if a given object implements the DataSvcQueryRequest interface.
 */
export declare function instanceOfDataSvcQueryRequest(value: object): value is DataSvcQueryRequest;
export declare function DataSvcQueryRequestFromJSON(json: any): DataSvcQueryRequest;
export declare function DataSvcQueryRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcQueryRequest;
export declare function DataSvcQueryRequestToJSON(json: any): DataSvcQueryRequest;
export declare function DataSvcQueryRequestToJSONTyped(value?: DataSvcQueryRequest | null, ignoreDiscriminator?: boolean): any;
