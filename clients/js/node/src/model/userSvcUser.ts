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

import { RequestFile } from './models';
import { UserSvcContact } from './userSvcContact';

export class UserSvcUser {
    /**
    * Contacts are used for login and identification purposes.
    */
    'contacts'?: Array<UserSvcContact>;
    'createdAt'?: string;
    'deletedAt'?: string;
    'id'?: string;
    /**
    * Full name of the organization.
    */
    'name'?: string;
    'passwordHash'?: string;
    /**
    * URL-friendly unique (inside the Singularon platform) identifier for the `user`.
    */
    'slug'?: string;
    'updatedAt'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "contacts",
            "baseName": "contacts",
            "type": "Array<UserSvcContact>"
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
            "name": "passwordHash",
            "baseName": "passwordHash",
            "type": "string"
        },
        {
            "name": "slug",
            "baseName": "slug",
            "type": "string"
        },
        {
            "name": "updatedAt",
            "baseName": "updatedAt",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcUser.attributeTypeMap;
    }
}

