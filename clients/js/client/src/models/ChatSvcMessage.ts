/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
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
 * @interface ChatSvcMessage
 */
export interface ChatSvcMessage {
    /**
     * 
     * @type {string}
     * @memberof ChatSvcMessage
     */
    createdAt?: string;
    /**
     * FileIds defines the file attachments the message has.
     * @type {Array<string>}
     * @memberof ChatSvcMessage
     */
    fileIds?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof ChatSvcMessage
     */
    id: string;
    /**
     * 
     * @type {{ [key: string]: any; }}
     * @memberof ChatSvcMessage
     */
    meta?: { [key: string]: any; };
    /**
     * Text content of the message eg. "Hi, what's up?"
     * @type {string}
     * @memberof ChatSvcMessage
     */
    text?: string;
    /**
     * ThreadId of the message.
     * @type {string}
     * @memberof ChatSvcMessage
     */
    threadId: string;
    /**
     * 
     * @type {string}
     * @memberof ChatSvcMessage
     */
    updatedAt?: string;
    /**
     * UserId is the id of the user who wrote the message.
     * For AI messages this field is empty.
     * @type {string}
     * @memberof ChatSvcMessage
     */
    userId?: string;
}

/**
 * Check if a given object implements the ChatSvcMessage interface.
 */
export function instanceOfChatSvcMessage(value: object): value is ChatSvcMessage {
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('threadId' in value) || value['threadId'] === undefined) return false;
    return true;
}

export function ChatSvcMessageFromJSON(json: any): ChatSvcMessage {
    return ChatSvcMessageFromJSONTyped(json, false);
}

export function ChatSvcMessageFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcMessage {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'fileIds': json['fileIds'] == null ? undefined : json['fileIds'],
        'id': json['id'],
        'meta': json['meta'] == null ? undefined : json['meta'],
        'text': json['text'] == null ? undefined : json['text'],
        'threadId': json['threadId'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}

export function ChatSvcMessageToJSON(json: any): ChatSvcMessage {
    return ChatSvcMessageToJSONTyped(json, false);
}

export function ChatSvcMessageToJSONTyped(value?: ChatSvcMessage | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'createdAt': value['createdAt'],
        'fileIds': value['fileIds'],
        'id': value['id'],
        'meta': value['meta'],
        'text': value['text'],
        'threadId': value['threadId'],
        'updatedAt': value['updatedAt'],
        'userId': value['userId'],
    };
}

