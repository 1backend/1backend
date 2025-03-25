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
/**
 *
 * @export
 * @interface FirehoseSvcErrorResponse
 */
export interface FirehoseSvcErrorResponse {
    /**
     *
     * @type {string}
     * @memberof FirehoseSvcErrorResponse
     */
    error?: string;
}
/**
 * Check if a given object implements the FirehoseSvcErrorResponse interface.
 */
export declare function instanceOfFirehoseSvcErrorResponse(value: object): value is FirehoseSvcErrorResponse;
export declare function FirehoseSvcErrorResponseFromJSON(json: any): FirehoseSvcErrorResponse;
export declare function FirehoseSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FirehoseSvcErrorResponse;
export declare function FirehoseSvcErrorResponseToJSON(json: any): FirehoseSvcErrorResponse;
export declare function FirehoseSvcErrorResponseToJSONTyped(value?: FirehoseSvcErrorResponse | null, ignoreDiscriminator?: boolean): any;
