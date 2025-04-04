/* tslint:disable */
/* eslint-disable */
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
 * @interface ChatSvcGetThreadResponse
 */
export interface ChatSvcGetThreadResponse {
    /**
     * 
     * @type {boolean}
     * @memberof ChatSvcGetThreadResponse
     */
    _exists?: boolean;
    /**
     * 
     * @type {ChatSvcThread}
     * @memberof ChatSvcGetThreadResponse
     */
    thread?: ChatSvcThread;
}

/**
 * Check if a given object implements the ChatSvcGetThreadResponse interface.
 */
export function instanceOfChatSvcGetThreadResponse(value: object): value is ChatSvcGetThreadResponse {
    return true;
}

export function ChatSvcGetThreadResponseFromJSON(json: any): ChatSvcGetThreadResponse {
    return ChatSvcGetThreadResponseFromJSONTyped(json, false);
}

export function ChatSvcGetThreadResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcGetThreadResponse {
    if (json == null) {
        return json;
    }
    return {
        
        '_exists': json['exists'] == null ? undefined : json['exists'],
        'thread': json['thread'] == null ? undefined : ChatSvcThreadFromJSON(json['thread']),
    };
}

export function ChatSvcGetThreadResponseToJSON(json: any): ChatSvcGetThreadResponse {
    return ChatSvcGetThreadResponseToJSONTyped(json, false);
}

export function ChatSvcGetThreadResponseToJSONTyped(value?: ChatSvcGetThreadResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'exists': value['_exists'],
        'thread': ChatSvcThreadToJSON(value['thread']),
    };
}

