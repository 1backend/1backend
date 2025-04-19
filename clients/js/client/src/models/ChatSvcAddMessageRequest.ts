/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
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
 * @interface ChatSvcAddMessageRequest
 */
export interface ChatSvcAddMessageRequest {
    /**
     * 
     * @type {ChatSvcMessage}
     * @memberof ChatSvcAddMessageRequest
     */
    message?: ChatSvcMessage;
}

/**
 * Check if a given object implements the ChatSvcAddMessageRequest interface.
 */
export function instanceOfChatSvcAddMessageRequest(value: object): value is ChatSvcAddMessageRequest {
    return true;
}

export function ChatSvcAddMessageRequestFromJSON(json: any): ChatSvcAddMessageRequest {
    return ChatSvcAddMessageRequestFromJSONTyped(json, false);
}

export function ChatSvcAddMessageRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcAddMessageRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'message': json['message'] == null ? undefined : ChatSvcMessageFromJSON(json['message']),
    };
}

export function ChatSvcAddMessageRequestToJSON(json: any): ChatSvcAddMessageRequest {
    return ChatSvcAddMessageRequestToJSONTyped(json, false);
}

export function ChatSvcAddMessageRequestToJSONTyped(value?: ChatSvcAddMessageRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'message': ChatSvcMessageToJSON(value['message']),
    };
}

