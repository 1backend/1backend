/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.7.6
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class UserSvcOrganization {
    'app'?: string;
    'createdAt': string;
    'deletedAt'?: string;
    'id': string;
    /**
    * Full name of the organization
    */
    'name': string;
    /**
    * URL-friendly unique (inside the Singularon platform) identifier for the `organization`.
    */
    'slug': string;
    'thumbnailFileId'?: string;
    'updatedAt': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "app",
            "baseName": "app",
            "type": "string"
        },
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
            "name": "name",
            "baseName": "name",
            "type": "string"
        },
        {
            "name": "slug",
            "baseName": "slug",
            "type": "string"
        },
        {
            "name": "thumbnailFileId",
            "baseName": "thumbnailFileId",
            "type": "string"
        },
        {
            "name": "updatedAt",
            "baseName": "updatedAt",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcOrganization.attributeTypeMap;
    }
}

