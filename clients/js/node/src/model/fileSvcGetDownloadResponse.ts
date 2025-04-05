/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { FileSvcDownload } from './fileSvcDownload';

export class FileSvcGetDownloadResponse {
    'download'?: FileSvcDownload;
    'exists': boolean;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "download",
            "baseName": "download",
            "type": "FileSvcDownload"
        },
        {
            "name": "exists",
            "baseName": "exists",
            "type": "boolean"
        }    ];

    static getAttributeTypeMap() {
        return FileSvcGetDownloadResponse.attributeTypeMap;
    }
}

