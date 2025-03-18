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
 * Check if a given object implements the FileSvcUpload interface.
 */
export function instanceOfFileSvcUpload(value) {
    if (!('fileSize' in value) || value['fileSize'] === undefined)
        return false;
    return true;
}
export function FileSvcUploadFromJSON(json) {
    return FileSvcUploadFromJSONTyped(json, false);
}
export function FileSvcUploadFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'fileId': json['fileId'] == null ? undefined : json['fileId'],
        'fileName': json['fileName'] == null ? undefined : json['fileName'],
        'filePath': json['filePath'] == null ? undefined : json['filePath'],
        'fileSize': json['fileSize'],
        'id': json['id'] == null ? undefined : json['id'],
        'nodeId': json['nodeId'] == null ? undefined : json['nodeId'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}
export function FileSvcUploadToJSON(json) {
    return FileSvcUploadToJSONTyped(json, false);
}
export function FileSvcUploadToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'createdAt': value['createdAt'],
        'fileId': value['fileId'],
        'fileName': value['fileName'],
        'filePath': value['filePath'],
        'fileSize': value['fileSize'],
        'id': value['id'],
        'nodeId': value['nodeId'],
        'updatedAt': value['updatedAt'],
        'userId': value['userId'],
    };
}
