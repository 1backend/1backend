/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { DataSvcCreateObjectFields } from './DataSvcCreateObjectFields';
/**
 *
 * @export
 * @interface DataSvcUpsertObjectRequest
 */
export interface DataSvcUpsertObjectRequest {
    /**
     *
     * @type {DataSvcCreateObjectFields}
     * @memberof DataSvcUpsertObjectRequest
     */
    object?: DataSvcCreateObjectFields;
}
/**
 * Check if a given object implements the DataSvcUpsertObjectRequest interface.
 */
export declare function instanceOfDataSvcUpsertObjectRequest(value: object): value is DataSvcUpsertObjectRequest;
export declare function DataSvcUpsertObjectRequestFromJSON(json: any): DataSvcUpsertObjectRequest;
export declare function DataSvcUpsertObjectRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcUpsertObjectRequest;
export declare function DataSvcUpsertObjectRequestToJSON(json: any): DataSvcUpsertObjectRequest;
export declare function DataSvcUpsertObjectRequestToJSONTyped(value?: DataSvcUpsertObjectRequest | null, ignoreDiscriminator?: boolean): any;
