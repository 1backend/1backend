/**
 * 1Backend
 * A common backend for your AI applications—microservices-based and built to scale.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { UserSvcPermission } from './userSvcPermission';

export class UserSvcSavePermissionsResponse {
    'permissions'?: Array<UserSvcPermission>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "permissions",
            "baseName": "permissions",
            "type": "Array<UserSvcPermission>"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcSavePermissionsResponse.attributeTypeMap;
    }
}

