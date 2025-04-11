/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
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
 * @interface ChatSvcEventMessageAdded
 */
export interface ChatSvcEventMessageAdded {
    /**
     * 
     * @type {string}
     * @memberof ChatSvcEventMessageAdded
     */
    threadId?: string;
}

/**
 * Check if a given object implements the ChatSvcEventMessageAdded interface.
 */
export function instanceOfChatSvcEventMessageAdded(value: object): value is ChatSvcEventMessageAdded {
    return true;
}

export function ChatSvcEventMessageAddedFromJSON(json: any): ChatSvcEventMessageAdded {
    return ChatSvcEventMessageAddedFromJSONTyped(json, false);
}

export function ChatSvcEventMessageAddedFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcEventMessageAdded {
    if (json == null) {
        return json;
    }
    return {
        
        'threadId': json['threadId'] == null ? undefined : json['threadId'],
    };
}

export function ChatSvcEventMessageAddedToJSON(json: any): ChatSvcEventMessageAdded {
    return ChatSvcEventMessageAddedToJSONTyped(json, false);
}

export function ChatSvcEventMessageAddedToJSONTyped(value?: ChatSvcEventMessageAdded | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'threadId': value['threadId'],
    };
}

