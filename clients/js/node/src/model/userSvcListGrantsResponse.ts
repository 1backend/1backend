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
import { UserSvcGrant } from './userSvcGrant';

export class UserSvcListGrantsResponse {
    'grants'?: Array<UserSvcGrant>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "grants",
            "baseName": "grants",
            "type": "Array<UserSvcGrant>"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcListGrantsResponse.attributeTypeMap;
    }
}

