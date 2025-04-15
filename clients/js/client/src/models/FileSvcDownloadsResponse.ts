/* tslint:disable */
/* eslint-disable */
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

import { mapValues } from '../runtime';
import type { FileSvcDownload } from './FileSvcDownload';
import {
    FileSvcDownloadFromJSON,
    FileSvcDownloadFromJSONTyped,
    FileSvcDownloadToJSON,
    FileSvcDownloadToJSONTyped,
} from './FileSvcDownload';

/**
 * 
 * @export
 * @interface FileSvcDownloadsResponse
 */
export interface FileSvcDownloadsResponse {
    /**
     * 
     * @type {Array<FileSvcDownload>}
     * @memberof FileSvcDownloadsResponse
     */
    downloads?: Array<FileSvcDownload>;
}

/**
 * Check if a given object implements the FileSvcDownloadsResponse interface.
 */
export function instanceOfFileSvcDownloadsResponse(value: object): value is FileSvcDownloadsResponse {
    return true;
}

export function FileSvcDownloadsResponseFromJSON(json: any): FileSvcDownloadsResponse {
    return FileSvcDownloadsResponseFromJSONTyped(json, false);
}

export function FileSvcDownloadsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcDownloadsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'downloads': json['downloads'] == null ? undefined : ((json['downloads'] as Array<any>).map(FileSvcDownloadFromJSON)),
    };
}

export function FileSvcDownloadsResponseToJSON(json: any): FileSvcDownloadsResponse {
    return FileSvcDownloadsResponseToJSONTyped(json, false);
}

export function FileSvcDownloadsResponseToJSONTyped(value?: FileSvcDownloadsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'downloads': value['downloads'] == null ? undefined : ((value['downloads'] as Array<any>).map(FileSvcDownloadToJSON)),
    };
}

