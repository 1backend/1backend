/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.39
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcListUsersOrderByField } from './userSvcListUsersOrderByField';
export declare class UserSvcListUsersRequest {
    /**
    * AfterTime is a time in RFC3339 format. It is used to paginate the results when the `orderByField` is set to `createdAt` or `updatedAt`. The results will be returned after this time.
    */
    'afterTime'?: string;
    /**
    * ContactId is the id of the contact the user is associated with. Will return a user list with one element if set.
    */
    'contactId'?: string;
    /**
    * Count is a flag that indicates if the count of the users should be returned.
    */
    'count'?: boolean;
    /**
    * Ids of the users to list.
    */
    'ids'?: Array<string>;
    'limit'?: number;
    'offset'?: number;
    'orderByDesc'?: boolean;
    'orderByField'?: UserSvcListUsersOrderByField;
    static discriminator: string | undefined;
    static attributeTypeMap: Array<{
        name: string;
        baseName: string;
        type: string;
    }>;
    static getAttributeTypeMap(): {
        name: string;
        baseName: string;
        type: string;
    }[];
}
export declare namespace UserSvcListUsersRequest {
}
