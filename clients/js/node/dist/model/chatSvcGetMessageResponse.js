/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ChatSvcGetMessageResponse {
    static getAttributeTypeMap() {
        return ChatSvcGetMessageResponse.attributeTypeMap;
    }
}
ChatSvcGetMessageResponse.discriminator = undefined;
ChatSvcGetMessageResponse.attributeTypeMap = [
    {
        "name": "exists",
        "baseName": "exists",
        "type": "boolean"
    },
    {
        "name": "message",
        "baseName": "message",
        "type": "ChatSvcMessage"
    }
];
