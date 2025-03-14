/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class EmailSvcSendEmailResponse {
    static getAttributeTypeMap() {
        return EmailSvcSendEmailResponse.attributeTypeMap;
    }
}
EmailSvcSendEmailResponse.discriminator = undefined;
EmailSvcSendEmailResponse.attributeTypeMap = [
    {
        "name": "emailId",
        "baseName": "emailId",
        "type": "string"
    },
    {
        "name": "status",
        "baseName": "status",
        "type": "string"
    }
];
