/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
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
 * @interface ChatSvcThread
 */
export interface ChatSvcThread {
    /**
     * 
     * @type {string}
     * @memberof ChatSvcThread
     */
    createdAt?: string;
    /**
     * 
     * @type {string}
     * @memberof ChatSvcThread
     */
    id: string;
    /**
     * Title of the thread.
     * @type {string}
     * @memberof ChatSvcThread
     */
    title?: string;
    /**
     * TopicIds defines which topics the thread belongs to.
     * Topics can roughly be thought of as tags for threads.
     * @type {Array<string>}
     * @memberof ChatSvcThread
     */
    topicIds?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof ChatSvcThread
     */
    updatedAt?: string;
    /**
     * UserIds the ids of the users who can see this thread.
     * @type {Array<string>}
     * @memberof ChatSvcThread
     */
    userIds?: Array<string>;
}

/**
 * Check if a given object implements the ChatSvcThread interface.
 */
export function instanceOfChatSvcThread(value: object): value is ChatSvcThread {
    if (!('id' in value) || value['id'] === undefined) return false;
    return true;
}

export function ChatSvcThreadFromJSON(json: any): ChatSvcThread {
    return ChatSvcThreadFromJSONTyped(json, false);
}

export function ChatSvcThreadFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcThread {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'id': json['id'],
        'title': json['title'] == null ? undefined : json['title'],
        'topicIds': json['topicIds'] == null ? undefined : json['topicIds'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
        'userIds': json['userIds'] == null ? undefined : json['userIds'],
    };
}

export function ChatSvcThreadToJSON(json: any): ChatSvcThread {
    return ChatSvcThreadToJSONTyped(json, false);
}

export function ChatSvcThreadToJSONTyped(value?: ChatSvcThread | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'createdAt': value['createdAt'],
        'id': value['id'],
        'title': value['title'],
        'topicIds': value['topicIds'],
        'updatedAt': value['updatedAt'],
        'userIds': value['userIds'],
    };
}

