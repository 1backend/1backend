/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { ChatSvcMessageFromJSON, ChatSvcMessageToJSON, } from './ChatSvcMessage';
/**
 * Check if a given object implements the ChatSvcGetMessagesResponse interface.
 */
export function instanceOfChatSvcGetMessagesResponse(value) {
    return true;
}
export function ChatSvcGetMessagesResponseFromJSON(json) {
    return ChatSvcGetMessagesResponseFromJSONTyped(json, false);
}
export function ChatSvcGetMessagesResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'messages': json['messages'] == null ? undefined : (json['messages'].map(ChatSvcMessageFromJSON)),
    };
}
export function ChatSvcGetMessagesResponseToJSON(json) {
    return ChatSvcGetMessagesResponseToJSONTyped(json, false);
}
export function ChatSvcGetMessagesResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'messages': value['messages'] == null ? undefined : (value['messages'].map(ChatSvcMessageToJSON)),
    };
}
