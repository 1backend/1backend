/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ContainerSvcRunContainerRequest {
    static getAttributeTypeMap() {
        return ContainerSvcRunContainerRequest.attributeTypeMap;
    }
}
ContainerSvcRunContainerRequest.discriminator = undefined;
ContainerSvcRunContainerRequest.attributeTypeMap = [
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
        "name": "ports",
        "baseName": "ports",
        "type": "Array<ContainerSvcPortMapping>"
    }
];
