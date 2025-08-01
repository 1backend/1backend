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
export declare class DataSvcCreateObjectFields {
    /**
    * Authors is a list of user ID and organization ID who created the object. If an organization ID is not provided, the currently active organization will be queried from the User Svc.
    */
    'authors'?: Array<string>;
    'data': {
        [key: string]: any;
    };
    /**
    * Deleters is a list of user IDs and role IDs that can delete the object. `_self` can be used to refer to the caller user\'s userId and `_org` can be used to refer to the user\'s currently active organization (if exists).
    */
    'deleters'?: Array<string>;
    'id'?: string;
    /**
    * Readers is a list of user IDs and role IDs that can read the object. `_self` can be used to refer to the caller user\'s userId and `_org` can be used to refer to the user\'s currently active organization (if exists).
    */
    'readers'?: Array<string>;
    'table': string;
    /**
    * Writers is a list of user IDs and role IDs that can write the object. `_self` can be used to refer to the caller user\'s userId and `_org` can be used to refer to the user\'s currently active organization (if exists).
    */
    'writers'?: Array<string>;
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
