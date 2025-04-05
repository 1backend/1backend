/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
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
 * @interface ChatSvcEventThreadAdded
 */
export interface ChatSvcEventThreadAdded {
    /**
     * 
     * @type {string}
     * @memberof ChatSvcEventThreadAdded
     */
    threadId?: string;
}

/**
 * Check if a given object implements the ChatSvcEventThreadAdded interface.
 */
export function instanceOfChatSvcEventThreadAdded(value: object): value is ChatSvcEventThreadAdded {
    return true;
}

export function ChatSvcEventThreadAddedFromJSON(json: any): ChatSvcEventThreadAdded {
    return ChatSvcEventThreadAddedFromJSONTyped(json, false);
}

export function ChatSvcEventThreadAddedFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatSvcEventThreadAdded {
    if (json == null) {
        return json;
    }
    return {
        
        'threadId': json['threadId'] == null ? undefined : json['threadId'],
    };
}

export function ChatSvcEventThreadAddedToJSON(json: any): ChatSvcEventThreadAdded {
    return ChatSvcEventThreadAddedToJSONTyped(json, false);
}

export function ChatSvcEventThreadAddedToJSONTyped(value?: ChatSvcEventThreadAdded | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'threadId': value['threadId'],
    };
}

