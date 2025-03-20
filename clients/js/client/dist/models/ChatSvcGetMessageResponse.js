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
import { ChatSvcMessageFromJSON, ChatSvcMessageToJSON, } from './ChatSvcMessage';
/**
 * Check if a given object implements the ChatSvcGetMessageResponse interface.
 */
export function instanceOfChatSvcGetMessageResponse(value) {
    return true;
}
export function ChatSvcGetMessageResponseFromJSON(json) {
    return ChatSvcGetMessageResponseFromJSONTyped(json, false);
}
export function ChatSvcGetMessageResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        '_exists': json['exists'] == null ? undefined : json['exists'],
        'message': json['message'] == null ? undefined : ChatSvcMessageFromJSON(json['message']),
    };
}
export function ChatSvcGetMessageResponseToJSON(json) {
    return ChatSvcGetMessageResponseToJSONTyped(json, false);
}
export function ChatSvcGetMessageResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'exists': value['_exists'],
        'message': ChatSvcMessageToJSON(value['message']),
    };
}
