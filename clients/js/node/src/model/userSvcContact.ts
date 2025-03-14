/**
 * OpenOrch
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

export class UserSvcContact {
    'createdAt'?: string;
    'deletedAt'?: string;
    /**
    * The unique identifier, which can be a URL.  Example values: \"joe12\" (openorch username), \"twitter.com/thejoe\" (twitter url), \"joe@joesdomain.com\" (email)
    */
    'id'?: string;
    /**
    * If this is the primary contact method
    */
    'isPrimary'?: boolean;
    /**
    * Platform of the contact (e.g., \"email\", \"phone\", \"twitter\")
    */
    'platform'?: string;
    'updatedAt'?: string;
    'userId'?: string;
    /**
    * Value is the platform local unique identifier. Ie. while the `id` of a Twitter contact is `twitter.com/thejoe`, the value will be only `thejoe`. For email and phones the `id` and the `value` will be the same. This field mostly exists for display purposes.  Example values: \"joe12\" (openorch username), \"thejoe\" (twitter username), \"joe@joesdomain.com\" (email)
    */
    'value'?: string;
    /**
    * Whether the contact is verified
    */
    'verified'?: boolean;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "string"
        },
        {
            "name": "deletedAt",
            "baseName": "deletedAt",
            "type": "string"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "isPrimary",
            "baseName": "isPrimary",
            "type": "boolean"
        },
        {
            "name": "platform",
            "baseName": "platform",
            "type": "string"
        },
        {
            "name": "updatedAt",
            "baseName": "updatedAt",
            "type": "string"
        },
        {
            "name": "userId",
            "baseName": "userId",
            "type": "string"
        },
        {
            "name": "value",
            "baseName": "value",
            "type": "string"
        },
        {
            "name": "verified",
            "baseName": "verified",
            "type": "boolean"
        }    ];

    static getAttributeTypeMap() {
        return UserSvcContact.attributeTypeMap;
    }
}

