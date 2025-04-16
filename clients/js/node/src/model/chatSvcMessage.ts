/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class ChatSvcMessage {
    'createdAt'?: string;
    /**
    * FileIds defines the file attachments the message has.
    */
    'fileIds'?: Array<string>;
    'id': string;
    'meta'?: { [key: string]: any; };
    /**
    * Text content of the message eg. \"Hi, what\'s up?\"
    */
    'text'?: string;
    /**
    * ThreadId of the message.
    */
    'threadId': string;
    'updatedAt'?: string;
    /**
    * UserId is the id of the user who wrote the message. For AI messages this field is empty.
    */
    'userId'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "string"
        },
        {
            "name": "fileIds",
            "baseName": "fileIds",
            "type": "Array<string>"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "meta",
            "baseName": "meta",
            "type": "{ [key: string]: any; }"
        },
        {
            "name": "text",
            "baseName": "text",
            "type": "string"
        },
        {
            "name": "threadId",
            "baseName": "threadId",
            "type": "string"
        },
        {
            "name": "updatedAt",
            "baseName": "updatedAt",
            "type": "string"
        },
        {
            "name": "userId",
            "baseName": "userId",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return ChatSvcMessage.attributeTypeMap;
    }
}

