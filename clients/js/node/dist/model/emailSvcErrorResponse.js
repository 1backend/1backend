/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class EmailSvcErrorResponse {
    static getAttributeTypeMap() {
        return EmailSvcErrorResponse.attributeTypeMap;
    }
}
EmailSvcErrorResponse.discriminator = undefined;
EmailSvcErrorResponse.attributeTypeMap = [
    {
        "name": "error",
        "baseName": "error",
        "type": "string"
    }
];
