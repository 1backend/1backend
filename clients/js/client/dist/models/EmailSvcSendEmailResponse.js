/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the EmailSvcSendEmailResponse interface.
 */
export function instanceOfEmailSvcSendEmailResponse(value) {
    return true;
}
export function EmailSvcSendEmailResponseFromJSON(json) {
    return EmailSvcSendEmailResponseFromJSONTyped(json, false);
}
export function EmailSvcSendEmailResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'emailId': json['emailId'] == null ? undefined : json['emailId'],
        'status': json['status'] == null ? undefined : json['status'],
    };
}
export function EmailSvcSendEmailResponseToJSON(json) {
    return EmailSvcSendEmailResponseToJSONTyped(json, false);
}
export function EmailSvcSendEmailResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'emailId': value['emailId'],
        'status': value['status'],
    };
}
