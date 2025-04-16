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
import { FileSvcUploadFromJSON, FileSvcUploadToJSON, } from './FileSvcUpload';
/**
 * Check if a given object implements the FileSvcListUploadsResponse interface.
 */
export function instanceOfFileSvcListUploadsResponse(value) {
    return true;
}
export function FileSvcListUploadsResponseFromJSON(json) {
    return FileSvcListUploadsResponseFromJSONTyped(json, false);
}
export function FileSvcListUploadsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'uploads': json['uploads'] == null ? undefined : (json['uploads'].map(FileSvcUploadFromJSON)),
    };
}
export function FileSvcListUploadsResponseToJSON(json) {
    return FileSvcListUploadsResponseToJSONTyped(json, false);
}
export function FileSvcListUploadsResponseToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'uploads': value['uploads'] == null ? undefined : (value['uploads'].map(FileSvcUploadToJSON)),
    };
}
