/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { ChatSvcMessage } from './chatSvcMessage';

export class ChatSvcGetMessagesResponse {
    'messages'?: Array<ChatSvcMessage>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "messages",
            "baseName": "messages",
            "type": "Array<ChatSvcMessage>"
        }    ];

    static getAttributeTypeMap() {
        return ChatSvcGetMessagesResponse.attributeTypeMap;
    }
}

