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
export class FirehoseSvcEventPublishRequest {
    static getAttributeTypeMap() {
        return FirehoseSvcEventPublishRequest.attributeTypeMap;
    }
}
FirehoseSvcEventPublishRequest.discriminator = undefined;
FirehoseSvcEventPublishRequest.attributeTypeMap = [
    {
        "name": "event",
        "baseName": "event",
        "type": "FirehoseSvcEvent"
    }
];
