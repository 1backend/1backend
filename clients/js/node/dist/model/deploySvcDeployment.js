/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class DeploySvcDeployment {
    static getAttributeTypeMap() {
        return DeploySvcDeployment.attributeTypeMap;
    }
}
DeploySvcDeployment.discriminator = undefined;
DeploySvcDeployment.attributeTypeMap = [
    {
        "name": "autoScaling",
        "baseName": "autoScaling",
        "type": "DeploySvcAutoScalingConfig"
    },
    {
        "name": "definitionId",
        "baseName": "definitionId",
        "type": "string"
    },
    {
        "name": "description",
        "baseName": "description",
        "type": "string"
    },
    {
        "name": "details",
        "baseName": "details",
        "type": "string"
    },
    {
        "name": "envars",
        "baseName": "envars",
        "type": "{ [key: string]: string; }"
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
        "name": "replicas",
        "baseName": "replicas",
        "type": "number"
    },
    {
        "name": "resources",
        "baseName": "resources",
        "type": "DeploySvcResourceLimits"
    },
    {
        "name": "status",
        "baseName": "status",
        "type": "DeploySvcDeploymentStatus"
    },
    {
        "name": "strategy",
        "baseName": "strategy",
        "type": "DeploySvcDeploymentStrategy"
    },
    {
        "name": "targetRegions",
        "baseName": "targetRegions",
        "type": "Array<DeploySvcTargetRegion>"
    }
];
