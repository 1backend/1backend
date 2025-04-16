/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.38
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class UserSvcPermit {
    'createdAt': string;
    'deletedAt'?: string;
    'id': string;
    'permission': string;
    /**
    * Role IDs that have been permited the specified permission.  Originally, permits were designed for slugs to facilitate service-to-service calls. Due to their convenience—especially with CLI and infrastructure-as-code support—they were later extended to roles.
    */
    'roles'?: Array<string>;
    /**
    * Slugs that have been permited the specified permission.
    */
    'slugs'?: Array<string>;
    'updatedAt': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "string"
        },
        {
            "name": "deletedAt",
            "baseName": "deletedAt",
            "type": "string"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "permission",
            "baseName": "permission",
            "type": "string"
        },
        {
            "name": "roles",
            "baseName": "roles",
            "type": "Array<string>"
        },
        {
            "name": "slugs",
            "baseName": "slugs",
            "type": "Array<string>"
        },
        {
            "name": "updatedAt",
            "baseName": "updatedAt",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcPermit.attributeTypeMap;
    }
}

