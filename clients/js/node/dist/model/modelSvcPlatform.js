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
export class ModelSvcPlatform {
    static getAttributeTypeMap() {
        return ModelSvcPlatform.attributeTypeMap;
    }
}
ModelSvcPlatform.discriminator = undefined;
ModelSvcPlatform.attributeTypeMap = [
    {
        "name": "architectures",
        "baseName": "architectures",
        "type": "ModelSvcArchitectures"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "name",
        "baseName": "name",
        "type": "string"
    },
    {
        "name": "types",
        "baseName": "types",
        "type": "Array<PromptSvcPromptType>"
    },
    {
        "name": "version",
        "baseName": "version",
        "type": "number"
    }
];
