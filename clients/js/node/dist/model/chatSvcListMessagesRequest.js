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
export class ChatSvcListMessagesRequest {
    static getAttributeTypeMap() {
        return ChatSvcListMessagesRequest.attributeTypeMap;
    }
}
ChatSvcListMessagesRequest.discriminator = undefined;
ChatSvcListMessagesRequest.attributeTypeMap = [
    {
        "name": "ids",
        "baseName": "ids",
        "type": "Array<string>"
    },
    {
        "name": "threadId",
        "baseName": "threadId",
        "type": "string"
    }
];
