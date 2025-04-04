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
import * as runtime from '../runtime';
import type { EmailSvcSendEmailRequest, EmailSvcSendEmailResponse } from '../models/index';
export interface SendEmailRequest {
    body: EmailSvcSendEmailRequest;
}
/**
 *
 */
export declare class EmailSvcApi extends runtime.BaseAPI {
    /**
     * Send an email with attachments.
     * Send an Email
     */
    sendEmailRaw(requestParameters: SendEmailRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<EmailSvcSendEmailResponse>>;
    /**
     * Send an email with attachments.
     * Send an Email
     */
    sendEmail(requestParameters: SendEmailRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<EmailSvcSendEmailResponse>;
}
