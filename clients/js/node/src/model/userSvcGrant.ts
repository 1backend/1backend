/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.30
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class UserSvcGrant {
    'id'?: string;
    'permissionId'?: string;
    /**
    * Slugs who are granted the PermissionId
    */
    'slugs'?: Array<string>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "permissionId",
            "baseName": "permissionId",
            "type": "string"
        },
        {
            "name": "slugs",
            "baseName": "slugs",
            "type": "Array<string>"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcGrant.attributeTypeMap;
    }
}

