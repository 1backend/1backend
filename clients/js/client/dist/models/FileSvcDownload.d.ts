/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface FileSvcDownload
 */
export interface FileSvcDownload {
    /**
     *
     * @type {string}
     * @memberof FileSvcDownload
     */
    createdAt?: string;
    /**
     * DownloadedBytes exists to show the download progress in terms of the number of bytes already downloaded.
     * @type {number}
     * @memberof FileSvcDownload
     */
    downloadedBytes?: number;
    /**
     *
     * @type {string}
     * @memberof FileSvcDownload
     */
    error?: string;
    /**
     *
     * @type {string}
     * @memberof FileSvcDownload
     */
    fileName?: string;
    /**
     *
     * @type {string}
     * @memberof FileSvcDownload
     */
    filePath?: string;
    /**
     * FileSize is the full final downloaded file size.
     * @type {number}
     * @memberof FileSvcDownload
     */
    fileSize?: number;
    /**
     *
     * @type {string}
     * @memberof FileSvcDownload
     */
    id?: string;
    /**
     *
     * @type {number}
     * @memberof FileSvcDownload
     */
    progress?: number;
    /**
     *
     * @type {string}
     * @memberof FileSvcDownload
     */
    status?: string;
    /**
     *
     * @type {string}
     * @memberof FileSvcDownload
     */
    updatedAt?: string;
    /**
     *
     * @type {string}
     * @memberof FileSvcDownload
     */
    url?: string;
}
/**
 * Check if a given object implements the FileSvcDownload interface.
 */
export declare function instanceOfFileSvcDownload(value: object): value is FileSvcDownload;
export declare function FileSvcDownloadFromJSON(json: any): FileSvcDownload;
export declare function FileSvcDownloadFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcDownload;
export declare function FileSvcDownloadToJSON(json: any): FileSvcDownload;
export declare function FileSvcDownloadToJSONTyped(value?: FileSvcDownload | null, ignoreDiscriminator?: boolean): any;
