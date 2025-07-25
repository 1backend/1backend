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
import { ChatSvcThreadFromJSON, ChatSvcThreadToJSON, } from './ChatSvcThread';
/**
 * Check if a given object implements the ChatSvcListThreadsResponse interface.
 */
export function instanceOfChatSvcListThreadsResponse(value) {
    return true;
}
export function ChatSvcListThreadsResponseFromJSON(json) {
    return ChatSvcListThreadsResponseFromJSONTyped(json, false);
}
export function ChatSvcListThreadsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'threads': json['threads'] == null ? undefined : (json['threads'].map(ChatSvcThreadFromJSON)),
    };
}
export function ChatSvcListThreadsResponseToJSON(json) {
    return ChatSvcListThreadsResponseToJSONTyped(json, false);
}
export function ChatSvcListThreadsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'threads': value['threads'] == null ? undefined : (value['threads'].map(ChatSvcThreadToJSON)),
    };
}
