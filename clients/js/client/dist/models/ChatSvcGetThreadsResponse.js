/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { ChatSvcThreadFromJSON, ChatSvcThreadToJSON, } from './ChatSvcThread';
/**
 * Check if a given object implements the ChatSvcGetThreadsResponse interface.
 */
export function instanceOfChatSvcGetThreadsResponse(value) {
    return true;
}
export function ChatSvcGetThreadsResponseFromJSON(json) {
    return ChatSvcGetThreadsResponseFromJSONTyped(json, false);
}
export function ChatSvcGetThreadsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'threads': json['threads'] == null ? undefined : (json['threads'].map(ChatSvcThreadFromJSON)),
    };
}
export function ChatSvcGetThreadsResponseToJSON(json) {
    return ChatSvcGetThreadsResponseToJSONTyped(json, false);
}
export function ChatSvcGetThreadsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'threads': value['threads'] == null ? undefined : (value['threads'].map(ChatSvcThreadToJSON)),
    };
}
