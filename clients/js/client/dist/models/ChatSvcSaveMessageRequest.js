/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ChatSvcSaveMessageRequest interface.
 */
export function instanceOfChatSvcSaveMessageRequest(value) {
    return true;
}
export function ChatSvcSaveMessageRequestFromJSON(json) {
    return ChatSvcSaveMessageRequestFromJSONTyped(json, false);
}
export function ChatSvcSaveMessageRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'fileIds': json['fileIds'] == null ? undefined : json['fileIds'],
        'id': json['id'] == null ? undefined : json['id'],
        'meta': json['meta'] == null ? undefined : json['meta'],
        'text': json['text'] == null ? undefined : json['text'],
        'threadId': json['threadId'] == null ? undefined : json['threadId'],
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}
export function ChatSvcSaveMessageRequestToJSON(json) {
    return ChatSvcSaveMessageRequestToJSONTyped(json, false);
}
export function ChatSvcSaveMessageRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'fileIds': value['fileIds'],
        'id': value['id'],
        'meta': value['meta'],
        'text': value['text'],
        'threadId': value['threadId'],
        'userId': value['userId'],
    };
}
