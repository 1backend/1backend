/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
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
export function instanceOfEmailSvcSendEmailResponse(value: object): value is EmailSvcSendEmailResponse {
    return true;
}

export function EmailSvcSendEmailResponseFromJSON(json: any): EmailSvcSendEmailResponse {
    return EmailSvcSendEmailResponseFromJSONTyped(json, false);
}

export function EmailSvcSendEmailResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmailSvcSendEmailResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'emailId': json['emailId'] == null ? undefined : json['emailId'],
        'status': json['status'] == null ? undefined : json['status'],
    };
}

export function EmailSvcSendEmailResponseToJSON(json: any): EmailSvcSendEmailResponse {
    return EmailSvcSendEmailResponseToJSONTyped(json, false);
}

export function EmailSvcSendEmailResponseToJSONTyped(value?: EmailSvcSendEmailResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'emailId': value['emailId'],
        'status': value['status'],
    };
}

