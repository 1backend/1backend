/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class PromptSvcListPromptsResponse {
    static getAttributeTypeMap() {
        return PromptSvcListPromptsResponse.attributeTypeMap;
    }
}
PromptSvcListPromptsResponse.discriminator = undefined;
PromptSvcListPromptsResponse.attributeTypeMap = [
    {
        "name": "after",
        "baseName": "after",
        "type": "object"
    },
    {
        "name": "count",
        "baseName": "count",
        "type": "number"
    },
    {
        "name": "prompts",
        "baseName": "prompts",
        "type": "Array<PromptSvcPrompt>"
    }
];
