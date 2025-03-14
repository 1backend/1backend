/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class RegistrySvcProcess {
    static getAttributeTypeMap() {
        return RegistrySvcProcess.attributeTypeMap;
    }
}
RegistrySvcProcess.discriminator = undefined;
RegistrySvcProcess.attributeTypeMap = [
    {
        "name": "memoryUsage",
        "baseName": "memoryUsage",
        "type": "number"
    },
    {
        "name": "pid",
        "baseName": "pid",
        "type": "number"
    },
    {
        "name": "processName",
        "baseName": "processName",
        "type": "string"
    }
];
