/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ChatSvcGetThreadsResponse {
    static getAttributeTypeMap() {
        return ChatSvcGetThreadsResponse.attributeTypeMap;
    }
}
ChatSvcGetThreadsResponse.discriminator = undefined;
ChatSvcGetThreadsResponse.attributeTypeMap = [
    {
        "name": "threads",
        "baseName": "threads",
        "type": "Array<ChatSvcThread>"
    }
];
