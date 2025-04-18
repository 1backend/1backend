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
import { ChatSvcThreadFromJSON, ChatSvcThreadToJSON, } from './ChatSvcThread';
/**
 * Check if a given object implements the ChatSvcAddThreadResponse interface.
 */
export function instanceOfChatSvcAddThreadResponse(value) {
    return true;
}
export function ChatSvcAddThreadResponseFromJSON(json) {
    return ChatSvcAddThreadResponseFromJSONTyped(json, false);
}
export function ChatSvcAddThreadResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'thread': json['thread'] == null ? undefined : ChatSvcThreadFromJSON(json['thread']),
    };
}
export function ChatSvcAddThreadResponseToJSON(json) {
    return ChatSvcAddThreadResponseToJSONTyped(json, false);
}
export function ChatSvcAddThreadResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'thread': ChatSvcThreadToJSON(value['thread']),
    };
}
