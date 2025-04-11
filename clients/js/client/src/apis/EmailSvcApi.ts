/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  EmailSvcErrorResponse,
  EmailSvcSendEmailRequest,
  EmailSvcSendEmailResponse,
} from '../models/index';
import {
    EmailSvcErrorResponseFromJSON,
    EmailSvcErrorResponseToJSON,
    EmailSvcSendEmailRequestFromJSON,
    EmailSvcSendEmailRequestToJSON,
    EmailSvcSendEmailResponseFromJSON,
    EmailSvcSendEmailResponseToJSON,
} from '../models/index';

export interface SendEmailRequest {
    body: EmailSvcSendEmailRequest;
}

/**
 * 
 */
export class EmailSvcApi extends runtime.BaseAPI {

    /**
     * Send an email with attachments.
     * Send an Email
     */
    async sendEmailRaw(requestParameters: SendEmailRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<EmailSvcSendEmailResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling sendEmail().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/email-svc/email`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EmailSvcSendEmailRequestToJSON(requestParameters['body']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => EmailSvcSendEmailResponseFromJSON(jsonValue));
    }

    /**
     * Send an email with attachments.
     * Send an Email
     */
    async sendEmail(requestParameters: SendEmailRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<EmailSvcSendEmailResponse> {
        const response = await this.sendEmailRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
