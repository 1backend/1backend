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
import { UserSvcPermit } from './userSvcPermit';

export class UserSvcListPermitsResponse {
    'permits'?: Array<UserSvcPermit>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "permits",
            "baseName": "permits",
            "type": "Array<UserSvcPermit>"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcListPermitsResponse.attributeTypeMap;
    }
}

