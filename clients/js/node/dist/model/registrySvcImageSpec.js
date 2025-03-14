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
export class RegistrySvcImageSpec {
    static getAttributeTypeMap() {
        return RegistrySvcImageSpec.attributeTypeMap;
    }
}
RegistrySvcImageSpec.discriminator = undefined;
RegistrySvcImageSpec.attributeTypeMap = [
    {
        "name": "internalPorts",
        "baseName": "internalPorts",
        "type": "Array<number>"
    },
    {
        "name": "name",
        "baseName": "name",
        "type": "string"
    }
];
