/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.37
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
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
    data: {
        [key: string]: any;
    };
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
export declare function instanceOfDataSvcCreateObjectFields(value: object): value is DataSvcCreateObjectFields;
export declare function DataSvcCreateObjectFieldsFromJSON(json: any): DataSvcCreateObjectFields;
export declare function DataSvcCreateObjectFieldsFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcCreateObjectFields;
export declare function DataSvcCreateObjectFieldsToJSON(json: any): DataSvcCreateObjectFields;
export declare function DataSvcCreateObjectFieldsToJSONTyped(value?: DataSvcCreateObjectFields | null, ignoreDiscriminator?: boolean): any;
