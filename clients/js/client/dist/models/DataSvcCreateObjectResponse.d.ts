/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
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
 * @interface DataSvcCreateObjectResponse
 */
export interface DataSvcCreateObjectResponse {
    /**
     *
     * @type {DataSvcObject}
     * @memberof DataSvcCreateObjectResponse
     */
    object?: DataSvcObject;
}
/**
 * Check if a given object implements the DataSvcCreateObjectResponse interface.
 */
export declare function instanceOfDataSvcCreateObjectResponse(value: object): value is DataSvcCreateObjectResponse;
export declare function DataSvcCreateObjectResponseFromJSON(json: any): DataSvcCreateObjectResponse;
export declare function DataSvcCreateObjectResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcCreateObjectResponse;
export declare function DataSvcCreateObjectResponseToJSON(json: any): DataSvcCreateObjectResponse;
export declare function DataSvcCreateObjectResponseToJSONTyped(value?: DataSvcCreateObjectResponse | null, ignoreDiscriminator?: boolean): any;
