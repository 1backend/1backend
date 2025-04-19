/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { FileSvcUploadFromJSON, FileSvcUploadToJSON, } from './FileSvcUpload';
/**
 * Check if a given object implements the FileSvcUploadFileResponse interface.
 */
export function instanceOfFileSvcUploadFileResponse(value) {
    return true;
}
export function FileSvcUploadFileResponseFromJSON(json) {
    return FileSvcUploadFileResponseFromJSONTyped(json, false);
}
export function FileSvcUploadFileResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'upload': json['upload'] == null ? undefined : FileSvcUploadFromJSON(json['upload']),
    };
}
export function FileSvcUploadFileResponseToJSON(json) {
    return FileSvcUploadFileResponseToJSONTyped(json, false);
}
export function FileSvcUploadFileResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'upload': FileSvcUploadToJSON(value['upload']),
    };
}
