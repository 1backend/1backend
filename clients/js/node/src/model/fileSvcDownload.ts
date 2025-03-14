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

export class FileSvcDownload {
    'createdAt'?: string;
    /**
    * DownloadedBytes exists to show the download progress in terms of the number of bytes already downloaded.
    */
    'downloadedBytes'?: number;
    'error'?: string;
    'fileName'?: string;
    'filePath'?: string;
    /**
    * FileSize is the full final downloaded file size.
    */
    'fileSize'?: number;
    'id'?: string;
    'progress'?: number;
    'status'?: string;
    'updatedAt'?: string;
    'url'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "string"
        },
        {
            "name": "downloadedBytes",
            "baseName": "downloadedBytes",
            "type": "number"
        },
        {
            "name": "error",
            "baseName": "error",
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
            "name": "progress",
            "baseName": "progress",
            "type": "number"
        },
        {
            "name": "status",
            "baseName": "status",
            "type": "string"
        },
        {
            "name": "updatedAt",
            "baseName": "updatedAt",
            "type": "string"
        },
        {
            "name": "url",
            "baseName": "url",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return FileSvcDownload.attributeTypeMap;
    }
}

