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
export class PromptSvcStreamChunk {
    static getAttributeTypeMap() {
        return PromptSvcStreamChunk.attributeTypeMap;
    }
}
PromptSvcStreamChunk.discriminator = undefined;
PromptSvcStreamChunk.attributeTypeMap = [
    {
        "name": "messageId",
        "baseName": "messageId",
        "type": "string"
    },
    {
        "name": "text",
        "baseName": "text",
        "type": "string"
    },
    {
        "name": "type",
        "baseName": "type",
        "type": "PromptSvcStreamChunkType"
    }
];
