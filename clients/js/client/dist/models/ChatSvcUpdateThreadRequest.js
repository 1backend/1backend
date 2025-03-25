/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { ChatSvcThreadFromJSON, ChatSvcThreadToJSON, } from './ChatSvcThread';
/**
 * Check if a given object implements the ChatSvcUpdateThreadRequest interface.
 */
export function instanceOfChatSvcUpdateThreadRequest(value) {
    return true;
}
export function ChatSvcUpdateThreadRequestFromJSON(json) {
    return ChatSvcUpdateThreadRequestFromJSONTyped(json, false);
}
export function ChatSvcUpdateThreadRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'thread': json['thread'] == null ? undefined : ChatSvcThreadFromJSON(json['thread']),
    };
}
export function ChatSvcUpdateThreadRequestToJSON(json) {
    return ChatSvcUpdateThreadRequestToJSONTyped(json, false);
}
export function ChatSvcUpdateThreadRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'thread': ChatSvcThreadToJSON(value['thread']),
    };
}
