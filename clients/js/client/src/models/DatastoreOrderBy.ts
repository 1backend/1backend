/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { DatastoreSortingType } from './DatastoreSortingType';
import {
    DatastoreSortingTypeFromJSON,
    DatastoreSortingTypeFromJSONTyped,
    DatastoreSortingTypeToJSON,
    DatastoreSortingTypeToJSONTyped,
} from './DatastoreSortingType';

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
export function instanceOfDatastoreOrderBy(value: object): value is DatastoreOrderBy {
    return true;
}

export function DatastoreOrderByFromJSON(json: any): DatastoreOrderBy {
    return DatastoreOrderByFromJSONTyped(json, false);
}

export function DatastoreOrderByFromJSONTyped(json: any, ignoreDiscriminator: boolean): DatastoreOrderBy {
    if (json == null) {
        return json;
    }
    return {
        
        'desc': json['desc'] == null ? undefined : json['desc'],
        'field': json['field'] == null ? undefined : json['field'],
        'randomize': json['randomize'] == null ? undefined : json['randomize'],
        'sortingType': json['sortingType'] == null ? undefined : DatastoreSortingTypeFromJSON(json['sortingType']),
    };
}

export function DatastoreOrderByToJSON(json: any): DatastoreOrderBy {
    return DatastoreOrderByToJSONTyped(json, false);
}

export function DatastoreOrderByToJSONTyped(value?: DatastoreOrderBy | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'desc': value['desc'],
        'field': value['field'],
        'randomize': value['randomize'],
        'sortingType': DatastoreSortingTypeToJSON(value['sortingType']),
    };
}

