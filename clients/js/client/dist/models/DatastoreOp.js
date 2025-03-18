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
 *
 * @export
 */
export const DatastoreOp = {
    OpEquals: 'equals',
    OpContainsSubstring: 'containsSubstring',
    OpStartsWith: 'startsWith',
    OpIntersects: 'intersects',
    OpIsInList: 'isInList'
};
export function instanceOfDatastoreOp(value) {
    for (const key in DatastoreOp) {
        if (Object.prototype.hasOwnProperty.call(DatastoreOp, key)) {
            if (DatastoreOp[key] === value) {
                return true;
            }
        }
    }
    return false;
}
export function DatastoreOpFromJSON(json) {
    return DatastoreOpFromJSONTyped(json, false);
}
export function DatastoreOpFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
export function DatastoreOpToJSON(value) {
    return value;
}
export function DatastoreOpToJSONTyped(value, ignoreDiscriminator) {
    return value;
}
