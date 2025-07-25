/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class RegistrySvcInstance {
    static getAttributeTypeMap() {
        return RegistrySvcInstance.attributeTypeMap;
    }
}
RegistrySvcInstance.discriminator = undefined;
RegistrySvcInstance.attributeTypeMap = [
    {
        "name": "deploymentId",
        "baseName": "deploymentId",
        "type": "string"
    },
    {
        "name": "details",
        "baseName": "details",
        "type": "string"
    },
    {
        "name": "host",
        "baseName": "host",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "ip",
        "baseName": "ip",
        "type": "string"
    },
    {
        "name": "lastHeartbeat",
        "baseName": "lastHeartbeat",
        "type": "string"
    },
    {
        "name": "nodeUrl",
        "baseName": "nodeUrl",
        "type": "string"
    },
    {
        "name": "path",
        "baseName": "path",
        "type": "string"
    },
    {
        "name": "port",
        "baseName": "port",
        "type": "number"
    },
    {
        "name": "scheme",
        "baseName": "scheme",
        "type": "string"
    },
    {
        "name": "slug",
        "baseName": "slug",
        "type": "string"
    },
    {
        "name": "status",
        "baseName": "status",
        "type": "RegistrySvcInstanceStatus"
    },
    {
        "name": "tags",
        "baseName": "tags",
        "type": "Array<string>"
    },
    {
        "name": "url",
        "baseName": "url",
        "type": "string"
    }
];
