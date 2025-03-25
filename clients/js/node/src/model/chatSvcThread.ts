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

import { RequestFile } from './models';

export class ChatSvcThread {
    'createdAt'?: string;
    'id': string;
    /**
    * Title of the thread.
    */
    'title'?: string;
    /**
    * TopicIds defines which topics the thread belongs to. Topics can roughly be thought of as tags for threads.
    */
    'topicIds'?: Array<string>;
    'updatedAt'?: string;
    /**
    * UserIds the ids of the users who can see this thread.
    */
    'userIds'?: Array<string>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "string"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "title",
            "baseName": "title",
            "type": "string"
        },
        {
            "name": "topicIds",
            "baseName": "topicIds",
            "type": "Array<string>"
        },
        {
            "name": "updatedAt",
            "baseName": "updatedAt",
            "type": "string"
        },
        {
            "name": "userIds",
            "baseName": "userIds",
            "type": "Array<string>"
        }    ];

    static getAttributeTypeMap() {
        return ChatSvcThread.attributeTypeMap;
    }
}

