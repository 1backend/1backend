/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { ChatSvcThread } from './ChatSvcThread';
import {
    ChatSvcThreadFromJSON,
    ChatSvcThreadFromJSONTyped,
    ChatSvcThreadToJSON,
    ChatSvcThreadToJSONTyped,
} from './ChatSvcThread';

/**
 * 
 * @export
 * @interface ChatSvcListThreadsResponse
 */
export interface ChatSvcListThreadsResponse {
    /**
     * 
     * @type {Array<ChatSvcThread>}
     * @memberof ChatSvcListThreadsResponse
     */
    threads?: Array<ChatSvcThread>;
}

/**
 * Check if a given object implements the ChatSvcListThreadsResponse interface.
 */
export function instanceOfChatSvcListThreadsResponse(value: object): value is ChatSvcListThreadsResponse {
    return true;
}

export function ChatSvcListThreadsResponseFromJSON(json: any): ChatSvcListThreadsResponse {
    return ChatSvcListThreadsResponseFromJSONTyped(json, false);
}

export function ChatSvcListThreadsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcListThreadsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'threads': json['threads'] == null ? undefined : ((json['threads'] as Array<any>).map(ChatSvcThreadFromJSON)),
    };
}

export function ChatSvcListThreadsResponseToJSON(json: any): ChatSvcListThreadsResponse {
    return ChatSvcListThreadsResponseToJSONTyped(json, false);
}

export function ChatSvcListThreadsResponseToJSONTyped(value?: ChatSvcListThreadsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'threads': value['threads'] == null ? undefined : ((value['threads'] as Array<any>).map(ChatSvcThreadToJSON)),
    };
}

