/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.34
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { UserSvcUserRecord } from './userSvcUserRecord';

export class UserSvcListUsersResponse {
    'after'?: string;
    'count'?: number;
    'users'?: Array<UserSvcUserRecord>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "after",
            "baseName": "after",
            "type": "string"
        },
        {
            "name": "count",
            "baseName": "count",
            "type": "number"
        },
        {
            "name": "users",
            "baseName": "users",
            "type": "Array<UserSvcUserRecord>"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcListUsersResponse.attributeTypeMap;
    }
}

