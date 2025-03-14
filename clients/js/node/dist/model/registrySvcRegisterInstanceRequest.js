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
export class RegistrySvcRegisterInstanceRequest {
    static getAttributeTypeMap() {
        return RegistrySvcRegisterInstanceRequest.attributeTypeMap;
    }
}
RegistrySvcRegisterInstanceRequest.discriminator = undefined;
RegistrySvcRegisterInstanceRequest.attributeTypeMap = [
    {
        "name": "deploymentId",
        "baseName": "deploymentId",
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
        "name": "url",
        "baseName": "url",
        "type": "string"
    }
];
