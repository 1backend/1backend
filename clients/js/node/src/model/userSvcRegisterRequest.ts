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
import { UserSvcContact } from './userSvcContact';

export class UserSvcRegisterRequest {
    'contact'?: UserSvcContact;
    'name'?: string;
    'password'?: string;
    /**
    * Slug is a URL-friendly unique (inside the 1Backend platform) identifier for the `user`. Required due to its central role in the platform. If your project has no use for a slug, just derive it from the email or similar.
    */
    'slug': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "contact",
            "baseName": "contact",
            "type": "UserSvcContact"
        },
        {
            "name": "name",
            "baseName": "name",
            "type": "string"
        },
        {
            "name": "password",
            "baseName": "password",
            "type": "string"
        },
        {
            "name": "slug",
            "baseName": "slug",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcRegisterRequest.attributeTypeMap;
    }
}

