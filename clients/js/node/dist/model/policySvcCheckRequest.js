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
export class PolicySvcCheckRequest {
    static getAttributeTypeMap() {
        return PolicySvcCheckRequest.attributeTypeMap;
    }
}
PolicySvcCheckRequest.discriminator = undefined;
PolicySvcCheckRequest.attributeTypeMap = [
    {
        "name": "endpoint",
        "baseName": "endpoint",
        "type": "string"
    },
    {
        "name": "ip",
        "baseName": "ip",
        "type": "string"
    },
    {
        "name": "method",
        "baseName": "method",
        "type": "string"
    },
    {
        "name": "userId",
        "baseName": "userId",
        "type": "string"
    }
];
