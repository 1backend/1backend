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
/**
 * 
 * @export
 * @interface EmailSvcFile
 */
export interface EmailSvcFile {
    /**
     * Base64-encoded content of the file
     * @type {string}
     * @memberof EmailSvcFile
     */
    content: string;
    /**
     * MIME type of the file (e.g., "application/pdf")
     * @type {string}
     * @memberof EmailSvcFile
     */
    contentType: string;
    /**
     * Name of the attached file
     * @type {string}
     * @memberof EmailSvcFile
     */
    filename: string;
}

/**
 * Check if a given object implements the EmailSvcFile interface.
 */
export function instanceOfEmailSvcFile(value: object): value is EmailSvcFile {
    if (!('content' in value) || value['content'] === undefined) return false;
    if (!('contentType' in value) || value['contentType'] === undefined) return false;
    if (!('filename' in value) || value['filename'] === undefined) return false;
    return true;
}

export function EmailSvcFileFromJSON(json: any): EmailSvcFile {
    return EmailSvcFileFromJSONTyped(json, false);
}

export function EmailSvcFileFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmailSvcFile {
    if (json == null) {
        return json;
    }
    return {
        
        'content': json['content'],
        'contentType': json['contentType'],
        'filename': json['filename'],
    };
}

export function EmailSvcFileToJSON(json: any): EmailSvcFile {
    return EmailSvcFileToJSONTyped(json, false);
}

export function EmailSvcFileToJSONTyped(value?: EmailSvcFile | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'content': value['content'],
        'contentType': value['contentType'],
        'filename': value['filename'],
    };
}

