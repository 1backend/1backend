/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class DeploySvcTargetRegion {
    static getAttributeTypeMap() {
        return DeploySvcTargetRegion.attributeTypeMap;
    }
}
DeploySvcTargetRegion.discriminator = undefined;
DeploySvcTargetRegion.attributeTypeMap = [
    {
        "name": "cluster",
        "baseName": "cluster",
        "type": "string"
    },
    {
        "name": "zone",
        "baseName": "zone",
        "type": "string"
    }
];
