/**
 * 1Backend
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
import { UserSvcRole } from './userSvcRole';

export class UserSvcGetRolesResponse {
    'roles'?: Array<UserSvcRole>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "roles",
            "baseName": "roles",
            "type": "Array<UserSvcRole>"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcGetRolesResponse.attributeTypeMap;
    }
}

