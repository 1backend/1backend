/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
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
 * @interface ChatSvcAddThreadResponse
 */
export interface ChatSvcAddThreadResponse {
    /**
     * 
     * @type {ChatSvcThread}
     * @memberof ChatSvcAddThreadResponse
     */
    thread?: ChatSvcThread;
}

/**
 * Check if a given object implements the ChatSvcAddThreadResponse interface.
 */
export function instanceOfChatSvcAddThreadResponse(value: object): value is ChatSvcAddThreadResponse {
    return true;
}

export function ChatSvcAddThreadResponseFromJSON(json: any): ChatSvcAddThreadResponse {
    return ChatSvcAddThreadResponseFromJSONTyped(json, false);
}

export function ChatSvcAddThreadResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcAddThreadResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'thread': json['thread'] == null ? undefined : ChatSvcThreadFromJSON(json['thread']),
    };
}

export function ChatSvcAddThreadResponseToJSON(json: any): ChatSvcAddThreadResponse {
    return ChatSvcAddThreadResponseToJSONTyped(json, false);
}

export function ChatSvcAddThreadResponseToJSONTyped(value?: ChatSvcAddThreadResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'thread': ChatSvcThreadToJSON(value['thread']),
    };
}

