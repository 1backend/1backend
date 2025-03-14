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
 * Check if a given object implements the ChatSvcGetThreadResponse interface.
 */
export function instanceOfChatSvcGetThreadResponse(value) {
    return true;
}
export function ChatSvcGetThreadResponseFromJSON(json) {
    return ChatSvcGetThreadResponseFromJSONTyped(json, false);
}
export function ChatSvcGetThreadResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        '_exists': json['exists'] == null ? undefined : json['exists'],
        'thread': json['thread'] == null ? undefined : ChatSvcThreadFromJSON(json['thread']),
    };
}
export function ChatSvcGetThreadResponseToJSON(json) {
    return ChatSvcGetThreadResponseToJSONTyped(json, false);
}
export function ChatSvcGetThreadResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'exists': value['_exists'],
        'thread': ChatSvcThreadToJSON(value['thread']),
    };
}
