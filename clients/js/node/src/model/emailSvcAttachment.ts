/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class EmailSvcAttachment {
    /**
    * Base64-encoded file content. Use this for small files. Required for inline attachments (i.e., those not using File Svc, see FileId).
    */
    'content'?: string;
    /**
    * MIME type of the file (e.g., \"application/pdf\", \"image/png\") Required for inline attachments (i.e., those not using File Svc, see FileId).
    */
    'contentType': string;
    /**
    * A File Svc file ID. Requires the file to be uploaded separately. Recommended for mid to large-sized files. If this field is specified, all other fields are optional.
    */
    'fileId'?: string;
    /**
    * File name for the attachment. Required for inline attachments (i.e., those not using File Svc, see FileId).
    */
    'filename': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "content",
            "baseName": "content",
            "type": "string"
        },
        {
            "name": "contentType",
            "baseName": "contentType",
            "type": "string"
        },
        {
            "name": "fileId",
            "baseName": "fileId",
            "type": "string"
        },
        {
            "name": "filename",
            "baseName": "filename",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return EmailSvcAttachment.attributeTypeMap;
    }
}

