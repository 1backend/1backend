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
export class ContainerSvcContainer {
    static getAttributeTypeMap() {
        return ContainerSvcContainer.attributeTypeMap;
    }
}
ContainerSvcContainer.discriminator = undefined;
ContainerSvcContainer.attributeTypeMap = [
    {
        "name": "assets",
        "baseName": "assets",
        "type": "Array<ContainerSvcAsset>"
    },
    {
        "name": "capabilities",
        "baseName": "capabilities",
        "type": "ContainerSvcCapabilities"
    },
    {
        "name": "envs",
        "baseName": "envs",
        "type": "Array<ContainerSvcEnvVar>"
    },
    {
        "name": "hash",
        "baseName": "hash",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "image",
        "baseName": "image",
        "type": "string"
    },
    {
        "name": "keeps",
        "baseName": "keeps",
        "type": "Array<ContainerSvcKeep>"
    },
    {
        "name": "labels",
        "baseName": "labels",
        "type": "Array<ContainerSvcLabel>"
    },
    {
        "name": "names",
        "baseName": "names",
        "type": "Array<string>"
    },
    {
        "name": "network",
        "baseName": "network",
        "type": "ContainerSvcNetwork"
    },
    {
        "name": "nodeId",
        "baseName": "nodeId",
        "type": "string"
    },
    {
        "name": "ports",
        "baseName": "ports",
        "type": "Array<ContainerSvcPortMapping>"
    },
    {
        "name": "resources",
        "baseName": "resources",
        "type": "ContainerSvcResources"
    },
    {
        "name": "runtime",
        "baseName": "runtime",
        "type": "string"
    },
    {
        "name": "status",
        "baseName": "status",
        "type": "string"
    },
    {
        "name": "volumes",
        "baseName": "volumes",
        "type": "Array<ContainerSvcVolume>"
    }
];
