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
export class FileSvcUpload {
    static getAttributeTypeMap() {
        return FileSvcUpload.attributeTypeMap;
    }
}
FileSvcUpload.discriminator = undefined;
FileSvcUpload.attributeTypeMap = [
    {
        "name": "createdAt",
        "baseName": "createdAt",
        "type": "string"
    },
    {
        "name": "fileId",
        "baseName": "fileId",
        "type": "string"
    },
    {
        "name": "fileName",
        "baseName": "fileName",
        "type": "string"
    },
    {
        "name": "filePath",
        "baseName": "filePath",
        "type": "string"
    },
    {
        "name": "fileSize",
        "baseName": "fileSize",
        "type": "number"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "nodeId",
        "baseName": "nodeId",
        "type": "string"
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
