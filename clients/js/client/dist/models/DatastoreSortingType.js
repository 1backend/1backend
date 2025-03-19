/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A unified backend development platform for microservices-based AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 */
export const DatastoreSortingType = {
    SortingTypeDefault: '',
    SortingTypeNumeric: 'numeric',
    SortingTypeText: 'text',
    SortingTypeDate: 'date'
};
export function instanceOfDatastoreSortingType(value) {
    for (const key in DatastoreSortingType) {
        if (Object.prototype.hasOwnProperty.call(DatastoreSortingType, key)) {
            if (DatastoreSortingType[key] === value) {
                return true;
            }
        }
    }
    return false;
}
export function DatastoreSortingTypeFromJSON(json) {
    return DatastoreSortingTypeFromJSONTyped(json, false);
}
export function DatastoreSortingTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
export function DatastoreSortingTypeToJSON(value) {
    return value;
}
export function DatastoreSortingTypeToJSONTyped(value, ignoreDiscriminator) {
    return value;
}
