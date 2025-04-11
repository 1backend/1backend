/* tslint:disable */
/* eslint-disable */
/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface DataSvcCreateObjectFields
 */
export interface DataSvcCreateObjectFields {
    /**
     * Authors is a list of user ID and organization ID who created the object.
     * If an organization ID is not provided, the currently active organization will
     * be queried from the User Svc.
     * @type {Array<string>}
     * @memberof DataSvcCreateObjectFields
     */
    authors?: Array<string>;
    /**
     * 
     * @type {{ [key: string]: any; }}
     * @memberof DataSvcCreateObjectFields
     */
    data: { [key: string]: any; };
    /**
     * Deleters is a list of user IDs and role IDs that can delete the object.
     * `_self` can be used to refer to the caller user's userId and
     * `_org` can be used to refer to the user's currently active organization (if exists).
     * @type {Array<string>}
     * @memberof DataSvcCreateObjectFields
     */
    deleters?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof DataSvcCreateObjectFields
     */
    id?: string;
    /**
     * Readers is a list of user IDs and role IDs that can read the object.
     * `_self` can be used to refer to the caller user's userId and
     * `_org` can be used to refer to the user's currently active organization (if exists).
     * @type {Array<string>}
     * @memberof DataSvcCreateObjectFields
     */
    readers?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof DataSvcCreateObjectFields
     */
    table: string;
    /**
     * Writers is a list of user IDs and role IDs that can write the object.
     * `_self` can be used to refer to the caller user's userId and
     * `_org` can be used to refer to the user's currently active organization (if exists).
     * @type {Array<string>}
     * @memberof DataSvcCreateObjectFields
     */
    writers?: Array<string>;
}

/**
 * Check if a given object implements the DataSvcCreateObjectFields interface.
 */
export function instanceOfDataSvcCreateObjectFields(value: object): value is DataSvcCreateObjectFields {
    if (!('data' in value) || value['data'] === undefined) return false;
    if (!('table' in value) || value['table'] === undefined) return false;
    return true;
}

export function DataSvcCreateObjectFieldsFromJSON(json: any): DataSvcCreateObjectFields {
    return DataSvcCreateObjectFieldsFromJSONTyped(json, false);
}

export function DataSvcCreateObjectFieldsFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcCreateObjectFields {
    if (json == null) {
        return json;
    }
    return {
        
        'authors': json['authors'] == null ? undefined : json['authors'],
        'data': json['data'],
        'deleters': json['deleters'] == null ? undefined : json['deleters'],
        'id': json['id'] == null ? undefined : json['id'],
        'readers': json['readers'] == null ? undefined : json['readers'],
        'table': json['table'],
        'writers': json['writers'] == null ? undefined : json['writers'],
    };
}

export function DataSvcCreateObjectFieldsToJSON(json: any): DataSvcCreateObjectFields {
    return DataSvcCreateObjectFieldsToJSONTyped(json, false);
}

export function DataSvcCreateObjectFieldsToJSONTyped(value?: DataSvcCreateObjectFields | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'authors': value['authors'],
        'data': value['data'],
        'deleters': value['deleters'],
        'id': value['id'],
        'readers': value['readers'],
        'table': value['table'],
        'writers': value['writers'],
    };
}

