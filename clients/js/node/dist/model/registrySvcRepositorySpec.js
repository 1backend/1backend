/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class RegistrySvcRepositorySpec {
    static getAttributeTypeMap() {
        return RegistrySvcRepositorySpec.attributeTypeMap;
    }
}
RegistrySvcRepositorySpec.discriminator = undefined;
RegistrySvcRepositorySpec.attributeTypeMap = [
    {
        "name": "buildContext",
        "baseName": "buildContext",
        "type": "string"
    },
    {
        "name": "containerFile",
        "baseName": "containerFile",
        "type": "string"
    },
    {
        "name": "internalPorts",
        "baseName": "internalPorts",
        "type": "Array<number>"
    },
    {
        "name": "url",
        "baseName": "url",
        "type": "string"
    },
    {
        "name": "version",
        "baseName": "version",
        "type": "string"
    }
];
