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
 *
 * @export
 */
export declare const DatastoreSortingType: {
    readonly SortingTypeDefault: "";
    readonly SortingTypeNumeric: "numeric";
    readonly SortingTypeText: "text";
    readonly SortingTypeDate: "date";
};
export type DatastoreSortingType = typeof DatastoreSortingType[keyof typeof DatastoreSortingType];
export declare function instanceOfDatastoreSortingType(value: any): boolean;
export declare function DatastoreSortingTypeFromJSON(json: any): DatastoreSortingType;
export declare function DatastoreSortingTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): DatastoreSortingType;
export declare function DatastoreSortingTypeToJSON(value?: DatastoreSortingType | null): any;
export declare function DatastoreSortingTypeToJSONTyped(value: any, ignoreDiscriminator: boolean): DatastoreSortingType;
