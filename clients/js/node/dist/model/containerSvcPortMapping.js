/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ContainerSvcPortMapping {
    static getAttributeTypeMap() {
        return ContainerSvcPortMapping.attributeTypeMap;
    }
}
ContainerSvcPortMapping.discriminator = undefined;
ContainerSvcPortMapping.attributeTypeMap = [
    {
        "name": "host",
        "baseName": "host",
        "type": "number"
    },
    {
        "name": "internal",
        "baseName": "internal",
        "type": "number"
    }
];
