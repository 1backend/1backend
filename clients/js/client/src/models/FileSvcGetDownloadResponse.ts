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
 * @interface FileSvcGetDownloadResponse
 */
export interface FileSvcGetDownloadResponse {
    /**
     * 
     * @type {FileSvcDownload}
     * @memberof FileSvcGetDownloadResponse
     */
    download?: FileSvcDownload;
    /**
     * 
     * @type {boolean}
     * @memberof FileSvcGetDownloadResponse
     */
    _exists: boolean;
}

/**
 * Check if a given object implements the FileSvcGetDownloadResponse interface.
 */
export function instanceOfFileSvcGetDownloadResponse(value: object): value is FileSvcGetDownloadResponse {
    if (!('_exists' in value) || value['_exists'] === undefined) return false;
    return true;
}

export function FileSvcGetDownloadResponseFromJSON(json: any): FileSvcGetDownloadResponse {
    return FileSvcGetDownloadResponseFromJSONTyped(json, false);
}

export function FileSvcGetDownloadResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcGetDownloadResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'download': json['download'] == null ? undefined : FileSvcDownloadFromJSON(json['download']),
        '_exists': json['exists'],
    };
}

export function FileSvcGetDownloadResponseToJSON(json: any): FileSvcGetDownloadResponse {
    return FileSvcGetDownloadResponseToJSONTyped(json, false);
}

export function FileSvcGetDownloadResponseToJSONTyped(value?: FileSvcGetDownloadResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'download': FileSvcDownloadToJSON(value['download']),
        'exists': value['_exists'],
    };
}

