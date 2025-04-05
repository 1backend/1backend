/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class PolicySvcParameters {
    static getAttributeTypeMap() {
        return PolicySvcParameters.attributeTypeMap;
    }
}
PolicySvcParameters.discriminator = undefined;
PolicySvcParameters.attributeTypeMap = [
    {
        "name": "blocklist",
        "baseName": "blocklist",
        "type": "PolicySvcBlocklistParameters"
    },
    {
        "name": "rateLimit",
        "baseName": "rateLimit",
        "type": "PolicySvcRateLimitParameters"
    }
];
