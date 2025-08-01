/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { DataSvcObject } from './DataSvcObject';
/**
 *
 * @export
 * @interface DataSvcUpsertObjectResponse
 */
export interface DataSvcUpsertObjectResponse {
    /**
     *
     * @type {DataSvcObject}
     * @memberof DataSvcUpsertObjectResponse
     */
    object?: DataSvcObject;
}
/**
 * Check if a given object implements the DataSvcUpsertObjectResponse interface.
 */
export declare function instanceOfDataSvcUpsertObjectResponse(value: object): value is DataSvcUpsertObjectResponse;
export declare function DataSvcUpsertObjectResponseFromJSON(json: any): DataSvcUpsertObjectResponse;
export declare function DataSvcUpsertObjectResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcUpsertObjectResponse;
export declare function DataSvcUpsertObjectResponseToJSON(json: any): DataSvcUpsertObjectResponse;
export declare function DataSvcUpsertObjectResponseToJSONTyped(value?: DataSvcUpsertObjectResponse | null, ignoreDiscriminator?: boolean): any;
