/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class PolicySvcInstance {
    static getAttributeTypeMap() {
        return PolicySvcInstance.attributeTypeMap;
    }
}
PolicySvcInstance.discriminator = undefined;
PolicySvcInstance.attributeTypeMap = [
    {
        "name": "endpoint",
        "baseName": "endpoint",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "parameters",
        "baseName": "parameters",
        "type": "PolicySvcParameters"
    },
    {
        "name": "templateId",
        "baseName": "templateId",
        "type": "PolicySvcTemplateId"
    }
];
