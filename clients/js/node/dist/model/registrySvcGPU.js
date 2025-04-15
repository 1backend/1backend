/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export class RegistrySvcGPU {
    static getAttributeTypeMap() {
        return RegistrySvcGPU.attributeTypeMap;
    }
}
RegistrySvcGPU.discriminator = undefined;
RegistrySvcGPU.attributeTypeMap = [
    {
        "name": "busId",
        "baseName": "busId",
        "type": "string"
    },
    {
        "name": "computeMode",
        "baseName": "computeMode",
        "type": "string"
    },
    {
        "name": "cudaVersion",
        "baseName": "cudaVersion",
        "type": "string"
    },
    {
        "name": "gpuUtilization",
        "baseName": "gpuUtilization",
        "type": "number"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "intraNodeId",
        "baseName": "intraNodeId",
        "type": "number"
    },
    {
        "name": "memoryTotal",
        "baseName": "memoryTotal",
        "type": "number"
    },
    {
        "name": "memoryUsage",
        "baseName": "memoryUsage",
        "type": "number"
    },
    {
        "name": "name",
        "baseName": "name",
        "type": "string"
    },
    {
        "name": "performanceState",
        "baseName": "performanceState",
        "type": "string"
    },
    {
        "name": "powerCap",
        "baseName": "powerCap",
        "type": "number"
    },
    {
        "name": "powerUsage",
        "baseName": "powerUsage",
        "type": "number"
    },
    {
        "name": "processDetails",
        "baseName": "processDetails",
        "type": "Array<RegistrySvcProcess>"
    },
    {
        "name": "temperature",
        "baseName": "temperature",
        "type": "number"
    }
];
