/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
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
} as const;
export type DatastoreOp = typeof DatastoreOp[keyof typeof DatastoreOp];


export function instanceOfDatastoreOp(value: any): boolean {
    for (const key in DatastoreOp) {
        if (Object.prototype.hasOwnProperty.call(DatastoreOp, key)) {
            if (DatastoreOp[key as keyof typeof DatastoreOp] === value) {
                return true;
            }
        }
    }
    return false;
}

export function DatastoreOpFromJSON(json: any): DatastoreOp {
    return DatastoreOpFromJSONTyped(json, false);
}

export function DatastoreOpFromJSONTyped(json: any, ignoreDiscriminator: boolean): DatastoreOp {
    return json as DatastoreOp;
}

export function DatastoreOpToJSON(value?: DatastoreOp | null): any {
    return value as any;
}

export function DatastoreOpToJSONTyped(value: any, ignoreDiscriminator: boolean): DatastoreOp {
    return value as DatastoreOp;
}

