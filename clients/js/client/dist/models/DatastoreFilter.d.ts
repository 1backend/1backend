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
import type { DatastoreOp } from './DatastoreOp';
/**
 *
 * @export
 * @interface DatastoreFilter
 */
export interface DatastoreFilter {
    /**
     *
     * @type {Array<string>}
     * @memberof DatastoreFilter
     */
    fields?: Array<string>;
    /**
     * JSONValues is a JSON marshalled array of values.
     * It's JSON marhalled due to the limitations of the
     * Swaggo -> OpenAPI 2.0 -> OpenAPI Go generator toolchain.
     * @type {string}
     * @memberof DatastoreFilter
     */
    jsonValues?: string;
    /**
     *
     * @type {DatastoreOp}
     * @memberof DatastoreFilter
     */
    op?: DatastoreOp;
    /**
     * SubFilters is used for operations like OR where multiple filters are combined.
     * @type {Array<DatastoreFilter>}
     * @memberof DatastoreFilter
     */
    subFilters?: Array<DatastoreFilter>;
}
/**
 * Check if a given object implements the DatastoreFilter interface.
 */
export declare function instanceOfDatastoreFilter(value: object): value is DatastoreFilter;
export declare function DatastoreFilterFromJSON(json: any): DatastoreFilter;
export declare function DatastoreFilterFromJSONTyped(json: any, ignoreDiscriminator: boolean): DatastoreFilter;
export declare function DatastoreFilterToJSON(json: any): DatastoreFilter;
export declare function DatastoreFilterToJSONTyped(value?: DatastoreFilter | null, ignoreDiscriminator?: boolean): any;
