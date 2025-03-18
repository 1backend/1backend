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
import { FileSvcDownloadFromJSON, FileSvcDownloadToJSON, } from './FileSvcDownload';
/**
 * Check if a given object implements the FileSvcGetDownloadResponse interface.
 */
export function instanceOfFileSvcGetDownloadResponse(value) {
    if (!('_exists' in value) || value['_exists'] === undefined)
        return false;
    return true;
}
export function FileSvcGetDownloadResponseFromJSON(json) {
    return FileSvcGetDownloadResponseFromJSONTyped(json, false);
}
export function FileSvcGetDownloadResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'download': json['download'] == null ? undefined : FileSvcDownloadFromJSON(json['download']),
        '_exists': json['exists'],
    };
}
export function FileSvcGetDownloadResponseToJSON(json) {
    return FileSvcGetDownloadResponseToJSONTyped(json, false);
}
export function FileSvcGetDownloadResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'download': FileSvcDownloadToJSON(value['download']),
        'exists': value['_exists'],
    };
}
