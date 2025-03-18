/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ContainerSvcListContainersRequest {
    static getAttributeTypeMap() {
        return ContainerSvcListContainersRequest.attributeTypeMap;
    }
}
ContainerSvcListContainersRequest.discriminator = undefined;
ContainerSvcListContainersRequest.attributeTypeMap = [
    {
        "name": "containerId",
        "baseName": "containerId",
        "type": "string"
    },
    {
        "name": "limit",
        "baseName": "limit",
        "type": "number"
    },
    {
        "name": "nodeId",
        "baseName": "nodeId",
        "type": "string"
    }
];
