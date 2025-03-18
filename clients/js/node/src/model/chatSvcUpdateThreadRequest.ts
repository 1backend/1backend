/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { ChatSvcThread } from './chatSvcThread';

export class ChatSvcUpdateThreadRequest {
    'thread'?: ChatSvcThread;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "thread",
            "baseName": "thread",
            "type": "ChatSvcThread"
        }    ];

    static getAttributeTypeMap() {
        return ChatSvcUpdateThreadRequest.attributeTypeMap;
    }
}

