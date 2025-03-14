/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
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
 * @interface DataSvcCreateObjectRequest
 */
export interface DataSvcCreateObjectRequest {
    /**
     *
     * @type {DataSvcCreateObjectFields}
     * @memberof DataSvcCreateObjectRequest
     */
    object?: DataSvcCreateObjectFields;
}
/**
 * Check if a given object implements the DataSvcCreateObjectRequest interface.
 */
export declare function instanceOfDataSvcCreateObjectRequest(value: object): value is DataSvcCreateObjectRequest;
export declare function DataSvcCreateObjectRequestFromJSON(json: any): DataSvcCreateObjectRequest;
export declare function DataSvcCreateObjectRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcCreateObjectRequest;
export declare function DataSvcCreateObjectRequestToJSON(json: any): DataSvcCreateObjectRequest;
export declare function DataSvcCreateObjectRequestToJSONTyped(value?: DataSvcCreateObjectRequest | null, ignoreDiscriminator?: boolean): any;
