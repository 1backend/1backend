/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
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
