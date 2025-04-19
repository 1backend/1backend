/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { UserSvcContact } from './userSvcContact';
import { UserSvcOrganization } from './userSvcOrganization';
import { UserSvcUser } from './userSvcUser';

export class UserSvcReadSelfResponse {
    /**
    * Active organization of the caller user, if it has any.
    */
    'activeOrganizationId'?: string;
    /**
    * Contacts of the caller user.
    */
    'contacts'?: Array<UserSvcContact>;
    /**
    * Organizations of the caller user.
    */
    'organizations'?: Array<UserSvcOrganization>;
    /**
    * Roles the token has that made this request.
    */
    'roles'?: Array<string>;
    /**
    * The user who made the request.
    */
    'user': UserSvcUser;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "activeOrganizationId",
            "baseName": "activeOrganizationId",
            "type": "string"
        },
        {
            "name": "contacts",
            "baseName": "contacts",
            "type": "Array<UserSvcContact>"
        },
        {
            "name": "organizations",
            "baseName": "organizations",
            "type": "Array<UserSvcOrganization>"
        },
        {
            "name": "roles",
            "baseName": "roles",
            "type": "Array<string>"
        },
        {
            "name": "user",
            "baseName": "user",
            "type": "UserSvcUser"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcReadSelfResponse.attributeTypeMap;
    }
}

