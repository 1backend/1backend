/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ChatSvcMessage interface.
 */
export function instanceOfChatSvcMessage(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('threadId' in value) || value['threadId'] === undefined)
        return false;
    return true;
}
export function ChatSvcMessageFromJSON(json) {
    return ChatSvcMessageFromJSONTyped(json, false);
}
export function ChatSvcMessageFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'fileIds': json['fileIds'] == null ? undefined : json['fileIds'],
        'id': json['id'],
        'meta': json['meta'] == null ? undefined : json['meta'],
        'text': json['text'] == null ? undefined : json['text'],
        'threadId': json['threadId'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}
export function ChatSvcMessageToJSON(json) {
    return ChatSvcMessageToJSONTyped(json, false);
}
export function ChatSvcMessageToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'createdAt': value['createdAt'],
        'fileIds': value['fileIds'],
        'id': value['id'],
        'meta': value['meta'],
        'text': value['text'],
        'threadId': value['threadId'],
        'updatedAt': value['updatedAt'],
        'userId': value['userId'],
    };
}
