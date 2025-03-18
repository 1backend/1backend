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
export class PromptSvcPrompt {
    static getAttributeTypeMap() {
        return PromptSvcPrompt.attributeTypeMap;
    }
}
PromptSvcPrompt.discriminator = undefined;
PromptSvcPrompt.attributeTypeMap = [
    {
        "name": "createdAt",
        "baseName": "createdAt",
        "type": "string"
    },
    {
        "name": "engineParameters",
        "baseName": "engineParameters",
        "type": "PromptSvcEngineParameters"
    },
    {
        "name": "error",
        "baseName": "error",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "lastRun",
        "baseName": "lastRun",
        "type": "string"
    },
    {
        "name": "maxRetries",
        "baseName": "maxRetries",
        "type": "number"
    },
    {
        "name": "modelId",
        "baseName": "modelId",
        "type": "string"
    },
    {
        "name": "parameters",
        "baseName": "parameters",
        "type": "PromptSvcParameters"
    },
    {
        "name": "prompt",
        "baseName": "prompt",
        "type": "string"
    },
    {
        "name": "requestMessageId",
        "baseName": "requestMessageId",
        "type": "string"
    },
    {
        "name": "responseMessageId",
        "baseName": "responseMessageId",
        "type": "string"
    },
    {
        "name": "runCount",
        "baseName": "runCount",
        "type": "number"
    },
    {
        "name": "status",
        "baseName": "status",
        "type": "PromptSvcPromptStatus"
    },
    {
        "name": "sync",
        "baseName": "sync",
        "type": "boolean"
    },
    {
        "name": "threadId",
        "baseName": "threadId",
        "type": "string"
    },
    {
        "name": "type",
        "baseName": "type",
        "type": "PromptSvcPromptType"
    },
    {
        "name": "updatedAt",
        "baseName": "updatedAt",
        "type": "string"
    },
    {
        "name": "userId",
        "baseName": "userId",
        "type": "string"
    }
];
