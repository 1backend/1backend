/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface DataSvcErrorResponse
 */
export interface DataSvcErrorResponse {
    /**
     *
     * @type {string}
     * @memberof DataSvcErrorResponse
     */
    error?: string;
}
/**
 * Check if a given object implements the DataSvcErrorResponse interface.
 */
export declare function instanceOfDataSvcErrorResponse(value: object): value is DataSvcErrorResponse;
export declare function DataSvcErrorResponseFromJSON(json: any): DataSvcErrorResponse;
export declare function DataSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcErrorResponse;
export declare function DataSvcErrorResponseToJSON(json: any): DataSvcErrorResponse;
export declare function DataSvcErrorResponseToJSONTyped(value?: DataSvcErrorResponse | null, ignoreDiscriminator?: boolean): any;
