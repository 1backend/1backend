/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.32
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { DatastoreFilter } from './DatastoreFilter';
import type { DatastoreOrderBy } from './DatastoreOrderBy';
/**
 *
 * @export
 * @interface DatastoreQuery
 */
export interface DatastoreQuery {
    /**
     * Count true means return the count of the dataset filtered by Filters
     * without after or limit.
     * @type {boolean}
     * @memberof DatastoreQuery
     */
    count?: boolean;
    /**
     * Filters are filtering options of a query. It is advised to use
     * It's advised to use helper functions in your respective client library such as filter constructors (`all`, `equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
     * @type {Array<DatastoreFilter>}
     * @memberof DatastoreQuery
     */
    filters?: Array<DatastoreFilter>;
    /**
     * JSONAfter is used for cursor-based pagination, which is more
     * effective in scalable and distributed environments compared
     * to offset-based pagination.
     *
     * JSONAfter is a JSON-encoded string due to limitations in Swaggo (e.g., []interface{} gets converted to []map[string]interface{}).
     * There is no way to specify a type that results in an any/interface{} type in the `go -> openapi -> go` generation process.
     * As a result, JSONAfter is a JSON-marshalled string representing an array, e.g., `[42]`.
     * @type {string}
     * @memberof DatastoreQuery
     */
    jsonAfter?: string;
    /**
     * Limit the number of records in the result set.
     * @type {number}
     * @memberof DatastoreQuery
     */
    limit?: number;
    /**
     * OrderBys order the result set.
     * @type {Array<DatastoreOrderBy>}
     * @memberof DatastoreQuery
     */
    orderBys?: Array<DatastoreOrderBy>;
}
/**
 * Check if a given object implements the DatastoreQuery interface.
 */
export declare function instanceOfDatastoreQuery(value: object): value is DatastoreQuery;
export declare function DatastoreQueryFromJSON(json: any): DatastoreQuery;
export declare function DatastoreQueryFromJSONTyped(json: any, ignoreDiscriminator: boolean): DatastoreQuery;
export declare function DatastoreQueryToJSON(json: any): DatastoreQuery;
export declare function DatastoreQueryToJSONTyped(value?: DatastoreQuery | null, ignoreDiscriminator?: boolean): any;
