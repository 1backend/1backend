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
export class EmailSvcSendEmailRequest {
    static getAttributeTypeMap() {
        return EmailSvcSendEmailRequest.attributeTypeMap;
    }
}
EmailSvcSendEmailRequest.discriminator = undefined;
EmailSvcSendEmailRequest.attributeTypeMap = [
    {
        "name": "attachments",
        "baseName": "attachments",
        "type": "Array<EmailSvcAttachment>"
    },
    {
        "name": "bcc",
        "baseName": "bcc",
        "type": "Array<string>"
    },
    {
        "name": "body",
        "baseName": "body",
        "type": "string"
    },
    {
        "name": "cc",
        "baseName": "cc",
        "type": "Array<string>"
    },
    {
        "name": "contentType",
        "baseName": "contentType",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "subject",
        "baseName": "subject",
        "type": "string"
    },
    {
        "name": "to",
        "baseName": "to",
        "type": "Array<string>"
    }
];
