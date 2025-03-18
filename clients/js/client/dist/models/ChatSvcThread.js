/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the ChatSvcThread interface.
 */
export function instanceOfChatSvcThread(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    return true;
}
export function ChatSvcThreadFromJSON(json) {
    return ChatSvcThreadFromJSONTyped(json, false);
}
export function ChatSvcThreadFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'id': json['id'],
        'title': json['title'] == null ? undefined : json['title'],
        'topicIds': json['topicIds'] == null ? undefined : json['topicIds'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
        'userIds': json['userIds'] == null ? undefined : json['userIds'],
    };
}
export function ChatSvcThreadToJSON(json) {
    return ChatSvcThreadToJSONTyped(json, false);
}
export function ChatSvcThreadToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'createdAt': value['createdAt'],
        'id': value['id'],
        'title': value['title'],
        'topicIds': value['topicIds'],
        'updatedAt': value['updatedAt'],
        'userIds': value['userIds'],
    };
}
