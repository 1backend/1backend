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

import { RequestFile } from './models';

export class DataSvcObject {
    /**
    * Authors is a list of user ID and organization ID who created the object. The authors field tracks which users or organizations created an entry, helping to prevent spam. If an organization ID is not provided, the currently active organization will be queried from the User Svc.
    */
    'authors'?: Array<string>;
    'createdAt'?: string;
    'data': { [key: string]: any; };
    /**
    * Deleters is a list of user IDs and role IDs that can delete the object. `_self` can be used to refer to the caller user\'s userId and `_org` can be used to refer to the user\'s currently active organization (if exists).
    */
    'deleters'?: Array<string>;
    'id'?: string;
    /**
    * Readers is a list of user IDs and role IDs that can read the object. `_self` can be used to refer to the caller user\'s userId and `_org` can be used to refer to the user\'s currently active organization (if exists).
    */
    'readers'?: Array<string>;
    'table': string;
    'updatedAt'?: string;
    /**
    * Writers is a list of user IDs and role IDs that can write the object. `_self` can be used to refer to the caller user\'s userId and `_org` can be used to refer to the user\'s currently active organization (if exists).
    */
    'writers'?: Array<string>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "authors",
            "baseName": "authors",
            "type": "Array<string>"
        },
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "string"
        },
        {
            "name": "data",
            "baseName": "data",
            "type": "{ [key: string]: any; }"
        },
        {
            "name": "deleters",
            "baseName": "deleters",
            "type": "Array<string>"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "readers",
            "baseName": "readers",
            "type": "Array<string>"
        },
        {
            "name": "table",
            "baseName": "table",
            "type": "string"
        },
        {
            "name": "updatedAt",
            "baseName": "updatedAt",
            "type": "string"
        },
        {
            "name": "writers",
            "baseName": "writers",
            "type": "Array<string>"
        }    ];

    static getAttributeTypeMap() {
        return DataSvcObject.attributeTypeMap;
    }
}

