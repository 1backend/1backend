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
export class ChatSvcGetMessagesResponse {
    static getAttributeTypeMap() {
        return ChatSvcGetMessagesResponse.attributeTypeMap;
    }
}
ChatSvcGetMessagesResponse.discriminator = undefined;
ChatSvcGetMessagesResponse.attributeTypeMap = [
    {
        "name": "messages",
        "baseName": "messages",
        "type": "Array<ChatSvcMessage>"
    }
];
