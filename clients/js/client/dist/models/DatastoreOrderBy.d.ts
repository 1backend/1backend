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
import type { DatastoreSortingType } from './DatastoreSortingType';
/**
 *
 * @export
 * @interface DatastoreOrderBy
 */
export interface DatastoreOrderBy {
    /**
     * Desc indicates whether the sorting should be in descending order.
     * @type {boolean}
     * @memberof DatastoreOrderBy
     */
    desc?: boolean;
    /**
     * The field by which to order the results
     * @type {string}
     * @memberof DatastoreOrderBy
     */
    field?: string;
    /**
     * Randomize indicates that the results should be randomized instead of ordered by the `field` and `desc` criteria
     * @type {boolean}
     * @memberof DatastoreOrderBy
     */
    randomize?: boolean;
    /**
     * Defines the type of sorting to apply (numeric, text, date, etc.)
     * @type {DatastoreSortingType}
     * @memberof DatastoreOrderBy
     */
    sortingType?: DatastoreSortingType;
}
/**
 * Check if a given object implements the DatastoreOrderBy interface.
 */
export declare function instanceOfDatastoreOrderBy(value: object): value is DatastoreOrderBy;
export declare function DatastoreOrderByFromJSON(json: any): DatastoreOrderBy;
export declare function DatastoreOrderByFromJSONTyped(json: any, ignoreDiscriminator: boolean): DatastoreOrderBy;
export declare function DatastoreOrderByToJSON(json: any): DatastoreOrderBy;
export declare function DatastoreOrderByToJSONTyped(value?: DatastoreOrderBy | null, ignoreDiscriminator?: boolean): any;
