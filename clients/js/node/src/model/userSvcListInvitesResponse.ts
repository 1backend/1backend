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
import { UserSvcInvite } from './userSvcInvite';

export class UserSvcListInvitesResponse {
    'invites': Array<UserSvcInvite>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "invites",
            "baseName": "invites",
            "type": "Array<UserSvcInvite>"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcListInvitesResponse.attributeTypeMap;
    }
}

