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
import { DatastoreOpFromJSON, DatastoreOpToJSON, } from './DatastoreOp';
/**
 * Check if a given object implements the DatastoreFilter interface.
 */
export function instanceOfDatastoreFilter(value) {
    return true;
}
export function DatastoreFilterFromJSON(json) {
    return DatastoreFilterFromJSONTyped(json, false);
}
export function DatastoreFilterFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'fields': json['fields'] == null ? undefined : json['fields'],
        'jsonValues': json['jsonValues'] == null ? undefined : json['jsonValues'],
        'op': json['op'] == null ? undefined : DatastoreOpFromJSON(json['op']),
    };
}
export function DatastoreFilterToJSON(json) {
    return DatastoreFilterToJSONTyped(json, false);
}
export function DatastoreFilterToJSONTyped(value, ignoreDiscriminator = false) {
    if (value == null) {
        return value;
    }
    return {
        'fields': value['fields'],
        'jsonValues': value['jsonValues'],
        'op': DatastoreOpToJSON(value['op']),
    };
}
