/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class ModelSvcModel {
    static getAttributeTypeMap() {
        return ModelSvcModel.attributeTypeMap;
    }
}
ModelSvcModel.discriminator = undefined;
ModelSvcModel.attributeTypeMap = [
    {
        "name": "assets",
        "baseName": "assets",
        "type": "Array<ModelSvcAsset>"
    },
    {
        "name": "bits",
        "baseName": "bits",
        "type": "number"
    },
    {
        "name": "description",
        "baseName": "description",
        "type": "string"
    },
    {
        "name": "extension",
        "baseName": "extension",
        "type": "string"
    },
    {
        "name": "flavour",
        "baseName": "flavour",
        "type": "string"
    },
    {
        "name": "fullName",
        "baseName": "fullName",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "maxBits",
        "baseName": "maxBits",
        "type": "number"
    },
    {
        "name": "maxRam",
        "baseName": "maxRam",
        "type": "number"
    },
    {
        "name": "mirrors",
        "baseName": "mirrors",
        "type": "Array<string>"
    },
    {
        "name": "name",
        "baseName": "name",
        "type": "string"
    },
    {
        "name": "parameters",
        "baseName": "parameters",
        "type": "string"
    },
    {
        "name": "platformId",
        "baseName": "platformId",
        "type": "string"
    },
    {
        "name": "promptTemplate",
        "baseName": "promptTemplate",
        "type": "string"
    },
    {
        "name": "quality",
        "baseName": "quality",
        "type": "string"
    },
    {
        "name": "quantComment",
        "baseName": "quantComment",
        "type": "string"
    },
    {
        "name": "size",
        "baseName": "size",
        "type": "number"
    },
    {
        "name": "tags",
        "baseName": "tags",
        "type": "Array<string>"
    },
    {
        "name": "uncensored",
        "baseName": "uncensored",
        "type": "boolean"
    },
    {
        "name": "version",
        "baseName": "version",
        "type": "string"
    }
];
