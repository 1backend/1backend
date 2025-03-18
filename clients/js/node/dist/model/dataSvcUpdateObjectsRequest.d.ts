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
import { DataSvcObject } from './dataSvcObject';
import { DatastoreFilter } from './datastoreFilter';
export declare class DataSvcUpdateObjectsRequest {
    /**
    * Filters to determine which objects will be updated. Only objects matching all filters will be modified.
    */
    'filters'?: Array<DatastoreFilter>;
    /**
    * The object containing the fields to update in matching objects.
    */
    'object'?: DataSvcObject;
    'table'?: string;
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
