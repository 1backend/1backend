/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { FileSvcDownloadFromJSON, FileSvcDownloadToJSON, } from './FileSvcDownload';
/**
 * Check if a given object implements the FileSvcDownloadsResponse interface.
 */
export function instanceOfFileSvcDownloadsResponse(value) {
    return true;
}
export function FileSvcDownloadsResponseFromJSON(json) {
    return FileSvcDownloadsResponseFromJSONTyped(json, false);
}
export function FileSvcDownloadsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'downloads': json['downloads'] == null ? undefined : (json['downloads'].map(FileSvcDownloadFromJSON)),
    };
}
export function FileSvcDownloadsResponseToJSON(json) {
    return FileSvcDownloadsResponseToJSONTyped(json, false);
}
export function FileSvcDownloadsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'downloads': value['downloads'] == null ? undefined : (value['downloads'].map(FileSvcDownloadToJSON)),
    };
}
