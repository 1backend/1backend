/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the EmailSvcAttachment interface.
 */
export function instanceOfEmailSvcAttachment(value) {
    if (!('contentType' in value) || value['contentType'] === undefined)
        return false;
    if (!('filename' in value) || value['filename'] === undefined)
        return false;
    return true;
}
export function EmailSvcAttachmentFromJSON(json) {
    return EmailSvcAttachmentFromJSONTyped(json, false);
}
export function EmailSvcAttachmentFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'content': json['content'] == null ? undefined : json['content'],
        'contentType': json['contentType'],
        'fileId': json['fileId'] == null ? undefined : json['fileId'],
        'filename': json['filename'],
    };
}
export function EmailSvcAttachmentToJSON(json) {
    return EmailSvcAttachmentToJSONTyped(json, false);
}
export function EmailSvcAttachmentToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'content': value['content'],
        'contentType': value['contentType'],
        'fileId': value['fileId'],
        'filename': value['filename'],
    };
}
