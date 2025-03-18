/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { DatastoreFilter } from './DatastoreFilter';
/**
 *
 * @export
 * @interface DataSvcDeleteObjectRequest
 */
export interface DataSvcDeleteObjectRequest {
    /**
     *
     * @type {Array<DatastoreFilter>}
     * @memberof DataSvcDeleteObjectRequest
     */
    filters?: Array<DatastoreFilter>;
    /**
     *
     * @type {string}
     * @memberof DataSvcDeleteObjectRequest
     */
    table?: string;
}
/**
 * Check if a given object implements the DataSvcDeleteObjectRequest interface.
 */
export declare function instanceOfDataSvcDeleteObjectRequest(value: object): value is DataSvcDeleteObjectRequest;
export declare function DataSvcDeleteObjectRequestFromJSON(json: any): DataSvcDeleteObjectRequest;
export declare function DataSvcDeleteObjectRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcDeleteObjectRequest;
export declare function DataSvcDeleteObjectRequestToJSON(json: any): DataSvcDeleteObjectRequest;
export declare function DataSvcDeleteObjectRequestToJSONTyped(value?: DataSvcDeleteObjectRequest | null, ignoreDiscriminator?: boolean): any;
