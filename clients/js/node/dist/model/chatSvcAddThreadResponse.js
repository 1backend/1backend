/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ChatSvcAddThreadResponse {
    static getAttributeTypeMap() {
        return ChatSvcAddThreadResponse.attributeTypeMap;
    }
}
ChatSvcAddThreadResponse.discriminator = undefined;
ChatSvcAddThreadResponse.attributeTypeMap = [
    {
        "name": "thread",
        "baseName": "thread",
        "type": "ChatSvcThread"
    }
];
