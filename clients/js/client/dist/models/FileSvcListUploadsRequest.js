/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the FileSvcListUploadsRequest interface.
 */
export function instanceOfFileSvcListUploadsRequest(value) {
    return true;
}
export function FileSvcListUploadsRequestFromJSON(json) {
    return FileSvcListUploadsRequestFromJSONTyped(json, false);
}
export function FileSvcListUploadsRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'after': json['after'] == null ? undefined : json['after'],
        'limit': json['limit'] == null ? undefined : json['limit'],
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}
export function FileSvcListUploadsRequestToJSON(json) {
    return FileSvcListUploadsRequestToJSONTyped(json, false);
}
export function FileSvcListUploadsRequestToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'after': value['after'],
        'limit': value['limit'],
        'userId': value['userId'],
    };
}
