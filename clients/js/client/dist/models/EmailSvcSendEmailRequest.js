/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { EmailSvcAttachmentFromJSON, EmailSvcAttachmentToJSON, } from './EmailSvcAttachment';
/**
 * Check if a given object implements the EmailSvcSendEmailRequest interface.
 */
export function instanceOfEmailSvcSendEmailRequest(value) {
    if (!('body' in value) || value['body'] === undefined)
        return false;
    if (!('contentType' in value) || value['contentType'] === undefined)
        return false;
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('subject' in value) || value['subject'] === undefined)
        return false;
    if (!('to' in value) || value['to'] === undefined)
        return false;
    return true;
}
export function EmailSvcSendEmailRequestFromJSON(json) {
    return EmailSvcSendEmailRequestFromJSONTyped(json, false);
}
export function EmailSvcSendEmailRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'attachments': json['attachments'] == null ? undefined : (json['attachments'].map(EmailSvcAttachmentFromJSON)),
        'bcc': json['bcc'] == null ? undefined : json['bcc'],
        'body': json['body'],
        'cc': json['cc'] == null ? undefined : json['cc'],
        'contentType': json['contentType'],
        'id': json['id'],
        'subject': json['subject'],
        'to': json['to'],
    };
}
export function EmailSvcSendEmailRequestToJSON(json) {
    return EmailSvcSendEmailRequestToJSONTyped(json, false);
}
export function EmailSvcSendEmailRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'attachments': value['attachments'] == null ? undefined : (value['attachments'].map(EmailSvcAttachmentToJSON)),
        'bcc': value['bcc'],
        'body': value['body'],
        'cc': value['cc'],
        'contentType': value['contentType'],
        'id': value['id'],
        'subject': value['subject'],
        'to': value['to'],
    };
}
