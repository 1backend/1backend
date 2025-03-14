/**
 * OpenOrch
 * A language-agnostic microservices framework for building AI applications.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { DatastoreFilter } from './datastoreFilter';
import { DatastoreOrderBy } from './datastoreOrderBy';
export declare class DatastoreQuery {
    /**
    * Count true means return the count of the dataset filtered by Filters without after or limit.
    */
    'count'?: boolean;
    /**
    * Filters are filtering options of a query. It is advised to use It\'s advised to use helper functions in your respective client library such as filter constructors (`all`, `equal`, `contains`, `startsWith`) and field selectors (`field`, `fields`, `id`) for easier access.
    */
    'filters'?: Array<DatastoreFilter>;
    /**
    * JSONAfter is used for cursor-based pagination, which is more effective in scalable and distributed environments compared to offset-based pagination.  JSONAfter is a JSON-encoded string due to limitations in Swaggo (e.g., []interface{} gets converted to []map[string]interface{}). There is no way to specify a type that results in an any/interface{} type in the `go -> openapi -> go` generation process. As a result, JSONAfter is a JSON-marshalled string representing an array, e.g., `[42]`.
    */
    'jsonAfter'?: string;
    /**
    * Limit the number of records in the result set.
    */
    'limit'?: number;
    /**
    * OrderBys order the result set.
    */
    'orderBys'?: Array<DatastoreOrderBy>;
    static discriminator: string | undefined;
    static attributeTypeMap: Array<{
        name: string;
        baseName: string;
        type: string;
    }>;
    static getAttributeTypeMap(): {
        name: string;
        baseName: string;
        type: string;
    }[];
}
