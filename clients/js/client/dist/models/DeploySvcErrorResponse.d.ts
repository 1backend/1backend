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
/**
 *
 * @export
 * @interface DeploySvcErrorResponse
 */
export interface DeploySvcErrorResponse {
    /**
     *
     * @type {string}
     * @memberof DeploySvcErrorResponse
     */
    error?: string;
}
/**
 * Check if a given object implements the DeploySvcErrorResponse interface.
 */
export declare function instanceOfDeploySvcErrorResponse(value: object): value is DeploySvcErrorResponse;
export declare function DeploySvcErrorResponseFromJSON(json: any): DeploySvcErrorResponse;
export declare function DeploySvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcErrorResponse;
export declare function DeploySvcErrorResponseToJSON(json: any): DeploySvcErrorResponse;
export declare function DeploySvcErrorResponseToJSONTyped(value?: DeploySvcErrorResponse | null, ignoreDiscriminator?: boolean): any;
