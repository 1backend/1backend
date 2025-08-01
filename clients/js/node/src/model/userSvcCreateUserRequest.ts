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
import { UserSvcContact } from './userSvcContact';
import { UserSvcUserInput } from './userSvcUserInput';

export class UserSvcCreateUserRequest {
    'app'?: string;
    'contacts'?: Array<UserSvcContact>;
    'password'?: string;
    'roleIds'?: Array<string>;
    'user'?: UserSvcUserInput;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "app",
            "baseName": "app",
            "type": "string"
        },
        {
            "name": "contacts",
            "baseName": "contacts",
            "type": "Array<UserSvcContact>"
        },
        {
            "name": "password",
            "baseName": "password",
            "type": "string"
        },
        {
            "name": "roleIds",
            "baseName": "roleIds",
            "type": "Array<string>"
        },
        {
            "name": "user",
            "baseName": "user",
            "type": "UserSvcUserInput"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcCreateUserRequest.attributeTypeMap;
    }
}

