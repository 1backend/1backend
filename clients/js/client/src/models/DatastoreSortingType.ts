/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * A language-agnostic microservices framework for building AI applications.
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
} as const;
export type DatastoreSortingType = typeof DatastoreSortingType[keyof typeof DatastoreSortingType];


export function instanceOfDatastoreSortingType(value: any): boolean {
    for (const key in DatastoreSortingType) {
        if (Object.prototype.hasOwnProperty.call(DatastoreSortingType, key)) {
            if (DatastoreSortingType[key as keyof typeof DatastoreSortingType] === value) {
                return true;
            }
        }
    }
    return false;
}

export function DatastoreSortingTypeFromJSON(json: any): DatastoreSortingType {
    return DatastoreSortingTypeFromJSONTyped(json, false);
}

export function DatastoreSortingTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): DatastoreSortingType {
    return json as DatastoreSortingType;
}

export function DatastoreSortingTypeToJSON(value?: DatastoreSortingType | null): any {
    return value as any;
}

export function DatastoreSortingTypeToJSONTyped(value: any, ignoreDiscriminator: boolean): DatastoreSortingType {
    return value as DatastoreSortingType;
}

