/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.33
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class UserSvcGrant {
    'id'?: string;
    'permission': string;
    /**
    * Role IDs that have been granted the specified permission.  Originally, grants were designed for slugs to facilitate service-to-service calls. Due to their convenience—especially with CLI and infrastructure-as-code support—they were later extended to roles.  Alternatively, permissions can be assigned to roles using UserSvc.SaveGrants. Grants currently offer a more streamlined approach, though this may evolve over time.
    */
    'roles'?: Array<string>;
    /**
    * Slugs that have been granted the specified permission.
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
        }    ];

    static getAttributeTypeMap() {
        return UserSvcGrant.attributeTypeMap;
    }
}

