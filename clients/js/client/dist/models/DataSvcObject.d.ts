/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.31
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 *
 * @export
 * @interface DataSvcObject
 */
export interface DataSvcObject {
    /**
     * Authors is a list of user ID and organization ID who created the object.
     * The authors field tracks which users or organizations created an entry, helping to prevent spam.
     * If an organization ID is not provided, the currently active organization will
     * be queried from the User Svc.
     * @type {Array<string>}
     * @memberof DataSvcObject
     */
    authors?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof DataSvcObject
     */
    createdAt?: string;
    /**
     *
     * @type {{ [key: string]: any; }}
     * @memberof DataSvcObject
     */
    data: {
        [key: string]: any;
    };
    /**
     * Deleters is a list of user IDs and role IDs that can delete the object.
     * `_self` can be used to refer to the caller user's userId and
     * `_org` can be used to refer to the user's currently active organization (if exists).
     * @type {Array<string>}
     * @memberof DataSvcObject
     */
    deleters?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof DataSvcObject
     */
    id?: string;
    /**
     * Readers is a list of user IDs and role IDs that can read the object.
     * `_self` can be used to refer to the caller user's userId and
     * `_org` can be used to refer to the user's currently active organization (if exists).
     * @type {Array<string>}
     * @memberof DataSvcObject
     */
    readers?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof DataSvcObject
     */
    table: string;
    /**
     *
     * @type {string}
     * @memberof DataSvcObject
     */
    updatedAt?: string;
    /**
     * Writers is a list of user IDs and role IDs that can write the object.
     * `_self` can be used to refer to the caller user's userId and
     * `_org` can be used to refer to the user's currently active organization (if exists).
     * @type {Array<string>}
     * @memberof DataSvcObject
     */
    writers?: Array<string>;
}
/**
 * Check if a given object implements the DataSvcObject interface.
 */
export declare function instanceOfDataSvcObject(value: object): value is DataSvcObject;
export declare function DataSvcObjectFromJSON(json: any): DataSvcObject;
export declare function DataSvcObjectFromJSONTyped(json: any, ignoreDiscriminator: boolean): DataSvcObject;
export declare function DataSvcObjectToJSON(json: any): DataSvcObject;
export declare function DataSvcObjectToJSONTyped(value?: DataSvcObject | null, ignoreDiscriminator?: boolean): any;
