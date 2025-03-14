/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
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
 * @interface FileSvcDownloadFileRequest
 */
export interface FileSvcDownloadFileRequest {
    /**
     * 
     * @type {string}
     * @memberof FileSvcDownloadFileRequest
     */
    folderPath?: string;
    /**
     * 
     * @type {string}
     * @memberof FileSvcDownloadFileRequest
     */
    url: string;
}

/**
 * Check if a given object implements the FileSvcDownloadFileRequest interface.
 */
export function instanceOfFileSvcDownloadFileRequest(value: object): value is FileSvcDownloadFileRequest {
    if (!('url' in value) || value['url'] === undefined) return false;
    return true;
}

export function FileSvcDownloadFileRequestFromJSON(json: any): FileSvcDownloadFileRequest {
    return FileSvcDownloadFileRequestFromJSONTyped(json, false);
}

export function FileSvcDownloadFileRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcDownloadFileRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'folderPath': json['folderPath'] == null ? undefined : json['folderPath'],
        'url': json['url'],
    };
}

export function FileSvcDownloadFileRequestToJSON(json: any): FileSvcDownloadFileRequest {
    return FileSvcDownloadFileRequestToJSONTyped(json, false);
}

export function FileSvcDownloadFileRequestToJSONTyped(value?: FileSvcDownloadFileRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'folderPath': value['folderPath'],
        'url': value['url'],
    };
}

