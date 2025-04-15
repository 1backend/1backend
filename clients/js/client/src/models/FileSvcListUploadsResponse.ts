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
 * @interface FileSvcListUploadsResponse
 */
export interface FileSvcListUploadsResponse {
    /**
     * 
     * @type {Array<FileSvcUpload>}
     * @memberof FileSvcListUploadsResponse
     */
    uploads?: Array<FileSvcUpload>;
}

/**
 * Check if a given object implements the FileSvcListUploadsResponse interface.
 */
export function instanceOfFileSvcListUploadsResponse(value: object): value is FileSvcListUploadsResponse {
    return true;
}

export function FileSvcListUploadsResponseFromJSON(json: any): FileSvcListUploadsResponse {
    return FileSvcListUploadsResponseFromJSONTyped(json, false);
}

export function FileSvcListUploadsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcListUploadsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'uploads': json['uploads'] == null ? undefined : ((json['uploads'] as Array<any>).map(FileSvcUploadFromJSON)),
    };
}

export function FileSvcListUploadsResponseToJSON(json: any): FileSvcListUploadsResponse {
    return FileSvcListUploadsResponseToJSONTyped(json, false);
}

export function FileSvcListUploadsResponseToJSONTyped(value?: FileSvcListUploadsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'uploads': value['uploads'] == null ? undefined : ((value['uploads'] as Array<any>).map(FileSvcUploadToJSON)),
    };
}

