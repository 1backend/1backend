/**
 * 1Backend
 * AI-native microservices platform.
 *
 * The version of the OpenAPI document: 0.3.0-rc.35
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
export declare class UserSvcUserRecord {
    'contactIds'?: Array<string>;
    'createdAt'?: string;
    'id': string;
    /**
    * Full name of the user.
    */
    'name'?: string;
    'roles'?: Array<string>;
    /**
    * URL-friendly unique (inside the 1Backend platform) identifier for the `user`.
    */
    'slug': string;
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
