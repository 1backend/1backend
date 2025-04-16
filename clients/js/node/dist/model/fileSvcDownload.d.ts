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
export declare class FileSvcDownload {
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
    static discriminator: string | undefined;
    static attributeTypeMap: Array<{
        name: string;
        baseName: string;
        type: string;
    }>;
    static getAttributeTypeMap(): {
        name: string;
        baseName: string;
        type: string;
    }[];
}
