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
export class ContainerSvcResources {
    static getAttributeTypeMap() {
        return ContainerSvcResources.attributeTypeMap;
    }
}
ContainerSvcResources.discriminator = undefined;
ContainerSvcResources.attributeTypeMap = [
    {
        "name": "cpu",
        "baseName": "cpu",
        "type": "number"
    },
    {
        "name": "diskMB",
        "baseName": "diskMB",
        "type": "number"
    },
    {
        "name": "memoryMB",
        "baseName": "memoryMB",
        "type": "number"
    }
];
