/* tslint:disable */
/* eslint-disable */
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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface FileSvcListUploadsRequest
 */
export interface FileSvcListUploadsRequest {
    /**
     * After time value
     * @type {string}
     * @memberof FileSvcListUploadsRequest
     */
    after?: string;
    /**
     * 
     * @type {number}
     * @memberof FileSvcListUploadsRequest
     */
    limit?: number;
    /**
     * 
     * @type {string}
     * @memberof FileSvcListUploadsRequest
     */
    userId?: string;
}

/**
 * Check if a given object implements the FileSvcListUploadsRequest interface.
 */
export function instanceOfFileSvcListUploadsRequest(value: object): value is FileSvcListUploadsRequest {
    return true;
}

export function FileSvcListUploadsRequestFromJSON(json: any): FileSvcListUploadsRequest {
    return FileSvcListUploadsRequestFromJSONTyped(json, false);
}

export function FileSvcListUploadsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileSvcListUploadsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'after': json['after'] == null ? undefined : json['after'],
        'limit': json['limit'] == null ? undefined : json['limit'],
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}

export function FileSvcListUploadsRequestToJSON(json: any): FileSvcListUploadsRequest {
    return FileSvcListUploadsRequestToJSONTyped(json, false);
}

export function FileSvcListUploadsRequestToJSONTyped(value?: FileSvcListUploadsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'after': value['after'],
        'limit': value['limit'],
        'userId': value['userId'],
    };
}

