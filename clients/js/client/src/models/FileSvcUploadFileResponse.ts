/* tslint:disable */
/* eslint-disable */
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

import { mapValues } from '../runtime';
import type { FileSvcUpload } from './FileSvcUpload';
import {
    FileSvcUploadFromJSON,
    FileSvcUploadFromJSONTyped,
    FileSvcUploadToJSON,
    FileSvcUploadToJSONTyped,
} from './FileSvcUpload';

/**
 * 
 * @export
 * @interface FileSvcUploadFileResponse
 */
export interface FileSvcUploadFileResponse {
    /**
     * 
     * @type {FileSvcUpload}
     * @memberof FileSvcUploadFileResponse
     */
    upload?: FileSvcUpload;
}

/**
 * Check if a given object implements the FileSvcUploadFileResponse interface.
 */
export function instanceOfFileSvcUploadFileResponse(value: object): value is FileSvcUploadFileResponse {
    return true;
}

export function FileSvcUploadFileResponseFromJSON(json: any): FileSvcUploadFileResponse {
    return FileSvcUploadFileResponseFromJSONTyped(json, false);
}

export function FileSvcUploadFileResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcUploadFileResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'upload': json['upload'] == null ? undefined : FileSvcUploadFromJSON(json['upload']),
    };
}

export function FileSvcUploadFileResponseToJSON(json: any): FileSvcUploadFileResponse {
    return FileSvcUploadFileResponseToJSONTyped(json, false);
}

export function FileSvcUploadFileResponseToJSONTyped(value?: FileSvcUploadFileResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'upload': FileSvcUploadToJSON(value['upload']),
    };
}

