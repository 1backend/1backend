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
export class ContainerSvcAsset {
    static getAttributeTypeMap() {
        return ContainerSvcAsset.attributeTypeMap;
    }
}
ContainerSvcAsset.discriminator = undefined;
ContainerSvcAsset.attributeTypeMap = [
    {
        "name": "envVarKey",
        "baseName": "envVarKey",
        "type": "string"
    },
    {
        "name": "url",
        "baseName": "url",
        "type": "string"
    }
];
