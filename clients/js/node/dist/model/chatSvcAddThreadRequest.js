/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ChatSvcAddThreadRequest {
    static getAttributeTypeMap() {
        return ChatSvcAddThreadRequest.attributeTypeMap;
    }
}
ChatSvcAddThreadRequest.discriminator = undefined;
ChatSvcAddThreadRequest.attributeTypeMap = [
    {
        "name": "thread",
        "baseName": "thread",
        "type": "ChatSvcThread"
    }
];
