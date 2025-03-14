/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ChatSvcThread {
    static getAttributeTypeMap() {
        return ChatSvcThread.attributeTypeMap;
    }
}
ChatSvcThread.discriminator = undefined;
ChatSvcThread.attributeTypeMap = [
    {
        "name": "createdAt",
        "baseName": "createdAt",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "title",
        "baseName": "title",
        "type": "string"
    },
    {
        "name": "topicIds",
        "baseName": "topicIds",
        "type": "Array<string>"
    },
    {
        "name": "updatedAt",
        "baseName": "updatedAt",
        "type": "string"
    },
    {
        "name": "userIds",
        "baseName": "userIds",
        "type": "Array<string>"
    }
];
