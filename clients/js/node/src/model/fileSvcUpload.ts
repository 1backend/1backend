/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class FileSvcUpload {
    'createdAt'?: string;
    /**
    * Logical file ID spanning all replicas
    */
    'fileId'?: string;
    /**
    * Filename is the original name of the file
    */
    'fileName'?: string;
    /**
    * FilePath is the full node local path of the file
    */
    'filePath'?: string;
    'fileSize': number;
    /**
    * Unique ID for this replica
    */
    'id'?: string;
    /**
    * ID of the node storing this replica
    */
    'nodeId'?: string;
    'updatedAt'?: string;
    'userId'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
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
        }    ];

    static getAttributeTypeMap() {
        return FileSvcUpload.attributeTypeMap;
    }
}

