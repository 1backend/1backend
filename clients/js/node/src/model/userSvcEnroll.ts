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

export class UserSvcEnroll {
    /**
    * App of the enroll. Use `*` to match all apps, such as when bootstrapping in services.
    */
    'app'?: string;
    /**
    * ContactId is the the recipient of the enroll. If the user is already registered, the role is assigned immediately; otherwise, it is applied upon registration.
    */
    'contactId'?: string;
    'createdAt': string;
    /**
    * CreatedBy contains the ID of the user who created the Enroll.
    */
    'createdBy'?: string;
    'deletedAt'?: string;
    'id': string;
    /**
    * Role specifies the role to be assigned to the ContactId. Callers can only assign roles they own, identified by their service slug (e.g., if \"my-service\" creates an enroll, the role must be \"my-service:admin\"). Dynamic organization roles can also be assigned (e.g., \"user-svc:org:{%orgId}:admin\" or \"user-svc:org:{%orgId}:user\"), but in this case, the caller must be an admin of the target organization.
    */
    'role': string;
    'updatedAt': string;
    /**
    * UserId is the recipient of the enroll. If the user is already registered, the role is assigned immediately; otherwise, it is applied upon registration.
    */
    'userId'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "app",
            "baseName": "app",
            "type": "string"
        },
        {
            "name": "contactId",
            "baseName": "contactId",
            "type": "string"
        },
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "string"
        },
        {
            "name": "createdBy",
            "baseName": "createdBy",
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
            "name": "role",
            "baseName": "role",
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
        }    ];

    static getAttributeTypeMap() {
        return UserSvcEnroll.attributeTypeMap;
    }
}

