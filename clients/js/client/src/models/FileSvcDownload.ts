/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
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
    createdAt: string;
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
    updatedAt: string;
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
export function instanceOfFileSvcDownload(value: object): value is FileSvcDownload {
    if (!('createdAt' in value) || value['createdAt'] === undefined) return false;
    if (!('updatedAt' in value) || value['updatedAt'] === undefined) return false;
    return true;
}

export function FileSvcDownloadFromJSON(json: any): FileSvcDownload {
    return FileSvcDownloadFromJSONTyped(json, false);
}

export function FileSvcDownloadFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcDownload {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['createdAt'],
        'downloadedBytes': json['downloadedBytes'] == null ? undefined : json['downloadedBytes'],
        'error': json['error'] == null ? undefined : json['error'],
        'fileName': json['fileName'] == null ? undefined : json['fileName'],
        'filePath': json['filePath'] == null ? undefined : json['filePath'],
        'fileSize': json['fileSize'] == null ? undefined : json['fileSize'],
        'id': json['id'] == null ? undefined : json['id'],
        'progress': json['progress'] == null ? undefined : json['progress'],
        'status': json['status'] == null ? undefined : json['status'],
        'updatedAt': json['updatedAt'],
        'url': json['url'] == null ? undefined : json['url'],
    };
}

export function FileSvcDownloadToJSON(json: any): FileSvcDownload {
    return FileSvcDownloadToJSONTyped(json, false);
}

export function FileSvcDownloadToJSONTyped(value?: FileSvcDownload | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'createdAt': value['createdAt'],
        'downloadedBytes': value['downloadedBytes'],
        'error': value['error'],
        'fileName': value['fileName'],
        'filePath': value['filePath'],
        'fileSize': value['fileSize'],
        'id': value['id'],
        'progress': value['progress'],
        'status': value['status'],
        'updatedAt': value['updatedAt'],
        'url': value['url'],
    };
}

