/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { ChatSvcMessage } from './ChatSvcMessage';
import {
    ChatSvcMessageFromJSON,
    ChatSvcMessageFromJSONTyped,
    ChatSvcMessageToJSON,
    ChatSvcMessageToJSONTyped,
} from './ChatSvcMessage';

/**
 * 
 * @export
 * @interface ChatSvcGetMessagesResponse
 */
export interface ChatSvcGetMessagesResponse {
    /**
     * 
     * @type {Array<ChatSvcMessage>}
     * @memberof ChatSvcGetMessagesResponse
     */
    messages?: Array<ChatSvcMessage>;
}

/**
 * Check if a given object implements the ChatSvcGetMessagesResponse interface.
 */
export function instanceOfChatSvcGetMessagesResponse(value: object): value is ChatSvcGetMessagesResponse {
    return true;
}

export function ChatSvcGetMessagesResponseFromJSON(json: any): ChatSvcGetMessagesResponse {
    return ChatSvcGetMessagesResponseFromJSONTyped(json, false);
}

export function ChatSvcGetMessagesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcGetMessagesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'messages': json['messages'] == null ? undefined : ((json['messages'] as Array<any>).map(ChatSvcMessageFromJSON)),
    };
}

export function ChatSvcGetMessagesResponseToJSON(json: any): ChatSvcGetMessagesResponse {
    return ChatSvcGetMessagesResponseToJSONTyped(json, false);
}

export function ChatSvcGetMessagesResponseToJSONTyped(value?: ChatSvcGetMessagesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'messages': value['messages'] == null ? undefined : ((value['messages'] as Array<any>).map(ChatSvcMessageToJSON)),
    };
}

