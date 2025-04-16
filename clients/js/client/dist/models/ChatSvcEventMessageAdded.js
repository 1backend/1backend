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
/**
 * Check if a given object implements the ChatSvcEventMessageAdded interface.
 */
export function instanceOfChatSvcEventMessageAdded(value) {
    return true;
}
export function ChatSvcEventMessageAddedFromJSON(json) {
    return ChatSvcEventMessageAddedFromJSONTyped(json, false);
}
export function ChatSvcEventMessageAddedFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'threadId': json['threadId'] == null ? undefined : json['threadId'],
    };
}
export function ChatSvcEventMessageAddedToJSON(json) {
    return ChatSvcEventMessageAddedToJSONTyped(json, false);
}
export function ChatSvcEventMessageAddedToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'threadId': value['threadId'],
    };
}
