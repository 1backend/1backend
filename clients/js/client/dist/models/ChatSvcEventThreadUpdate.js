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
/**
 * Check if a given object implements the ChatSvcEventThreadUpdate interface.
 */
export function instanceOfChatSvcEventThreadUpdate(value) {
    return true;
}
export function ChatSvcEventThreadUpdateFromJSON(json) {
    return ChatSvcEventThreadUpdateFromJSONTyped(json, false);
}
export function ChatSvcEventThreadUpdateFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'threadId': json['threadId'] == null ? undefined : json['threadId'],
    };
}
export function ChatSvcEventThreadUpdateToJSON(json) {
    return ChatSvcEventThreadUpdateToJSONTyped(json, false);
}
export function ChatSvcEventThreadUpdateToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'threadId': value['threadId'],
    };
}
