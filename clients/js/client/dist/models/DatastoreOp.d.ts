/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
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
export declare const DatastoreOp: {
    readonly OpOr: "or";
    readonly OpEquals: "equals";
    readonly OpContainsSubstring: "containsSubstring";
    readonly OpStartsWith: "startsWith";
    readonly OpIntersects: "intersects";
    readonly OpIsInList: "isInList";
};
export type DatastoreOp = typeof DatastoreOp[keyof typeof DatastoreOp];
export declare function instanceOfDatastoreOp(value: any): boolean;
export declare function DatastoreOpFromJSON(json: any): DatastoreOp;
export declare function DatastoreOpFromJSONTyped(json: any, ignoreDiscriminator: boolean): DatastoreOp;
export declare function DatastoreOpToJSON(value?: DatastoreOp | null): any;
export declare function DatastoreOpToJSONTyped(value: any, ignoreDiscriminator: boolean): DatastoreOp;
