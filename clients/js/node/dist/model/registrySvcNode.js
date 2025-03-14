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
export class RegistrySvcNode {
    static getAttributeTypeMap() {
        return RegistrySvcNode.attributeTypeMap;
    }
}
RegistrySvcNode.discriminator = undefined;
RegistrySvcNode.attributeTypeMap = [
    {
        "name": "availabilityZone",
        "baseName": "availabilityZone",
        "type": "string"
    },
    {
        "name": "gpus",
        "baseName": "gpus",
        "type": "Array<RegistrySvcGPU>"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "lastHeartbeat",
        "baseName": "lastHeartbeat",
        "type": "string"
    },
    {
        "name": "region",
        "baseName": "region",
        "type": "string"
    },
    {
        "name": "url",
        "baseName": "url",
        "type": "string"
    },
    {
        "name": "usage",
        "baseName": "usage",
        "type": "RegistrySvcResourceUsage"
    }
];
