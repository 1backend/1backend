/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface EmailSvcSendEmailResponse
 */
export interface EmailSvcSendEmailResponse {
    /**
     * Unique identifier for the sent email
     * @type {string}
     * @memberof EmailSvcSendEmailResponse
     */
    emailId?: string;
    /**
     * Status of the email send operation ("sent", "queued", etc.)
     * @type {string}
     * @memberof EmailSvcSendEmailResponse
     */
    status?: string;
}
/**
 * Check if a given object implements the EmailSvcSendEmailResponse interface.
 */
export declare function instanceOfEmailSvcSendEmailResponse(value: object): value is EmailSvcSendEmailResponse;
export declare function EmailSvcSendEmailResponseFromJSON(json: any): EmailSvcSendEmailResponse;
export declare function EmailSvcSendEmailResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmailSvcSendEmailResponse;
export declare function EmailSvcSendEmailResponseToJSON(json: any): EmailSvcSendEmailResponse;
export declare function EmailSvcSendEmailResponseToJSONTyped(value?: EmailSvcSendEmailResponse | null, ignoreDiscriminator?: boolean): any;
