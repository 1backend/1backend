/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
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
 * @interface EmailSvcErrorResponse
 */
export interface EmailSvcErrorResponse {
    /**
     *
     * @type {string}
     * @memberof EmailSvcErrorResponse
     */
    error?: string;
}
/**
 * Check if a given object implements the EmailSvcErrorResponse interface.
 */
export declare function instanceOfEmailSvcErrorResponse(value: object): value is EmailSvcErrorResponse;
export declare function EmailSvcErrorResponseFromJSON(json: any): EmailSvcErrorResponse;
export declare function EmailSvcErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmailSvcErrorResponse;
export declare function EmailSvcErrorResponseToJSON(json: any): EmailSvcErrorResponse;
export declare function EmailSvcErrorResponseToJSONTyped(value?: EmailSvcErrorResponse | null, ignoreDiscriminator?: boolean): any;
