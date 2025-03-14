/* tslint:disable */
/* eslint-disable */
/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
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
 * Check if a given object implements the ChatSvcAddThreadRequest interface.
 */
export function instanceOfChatSvcAddThreadRequest(value) {
    return true;
}
export function ChatSvcAddThreadRequestFromJSON(json) {
    return ChatSvcAddThreadRequestFromJSONTyped(json, false);
}
export function ChatSvcAddThreadRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'thread': json['thread'] == null ? undefined : ChatSvcThreadFromJSON(json['thread']),
    };
}
export function ChatSvcAddThreadRequestToJSON(json) {
    return ChatSvcAddThreadRequestToJSONTyped(json, false);
}
export function ChatSvcAddThreadRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'thread': ChatSvcThreadToJSON(value['thread']),
    };
}
