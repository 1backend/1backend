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
/**
 *
 * @export
 * @interface PolicySvcErrorResponse
 */
export interface PolicySvcErrorResponse {
    /**
     *
     * @type {string}
     * @memberof PolicySvcErrorResponse
     */
    error?: string;
}
/**
 * Check if a given object implements the PolicySvcErrorResponse interface.
 */
export declare function instanceOfPolicySvcErrorResponse(value: object): value is PolicySvcErrorResponse;
export declare function PolicySvcErrorResponseFromJSON(json: any): PolicySvcErrorResponse;
export declare function PolicySvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicySvcErrorResponse;
export declare function PolicySvcErrorResponseToJSON(json: any): PolicySvcErrorResponse;
export declare function PolicySvcErrorResponseToJSONTyped(value?: PolicySvcErrorResponse | null, ignoreDiscriminator?: boolean): any;
