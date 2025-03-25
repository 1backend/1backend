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
export class ContainerSvcLog {
    static getAttributeTypeMap() {
        return ContainerSvcLog.attributeTypeMap;
    }
}
ContainerSvcLog.discriminator = undefined;
ContainerSvcLog.attributeTypeMap = [
    {
        "name": "containerId",
        "baseName": "containerId",
        "type": "string"
    },
    {
        "name": "content",
        "baseName": "content",
        "type": "string"
    },
    {
        "name": "createdAt",
        "baseName": "createdAt",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "nodeId",
        "baseName": "nodeId",
        "type": "string"
    }
];
