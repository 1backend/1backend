/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.29
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { UserSvcContact } from './userSvcContact';
export declare class UserSvcUser {
    /**
    * Contacts are used for login and identification purposes.
    */
    'contacts'?: Array<UserSvcContact>;
    'createdAt'?: string;
    'deletedAt'?: string;
    'id'?: string;
    /**
    * Full name of the organization.
    */
    'name'?: string;
    'passwordHash'?: string;
    /**
    * URL-friendly unique (inside the Singularon platform) identifier for the `user`.
    */
    'slug'?: string;
    'updatedAt'?: string;
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
